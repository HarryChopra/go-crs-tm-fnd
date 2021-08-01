package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

type City struct {
	Name       string
	Population int
	Status     string
}

func main() {
	rows1 := genRows("cities1.csv")
	rows2 := genRows("cities2.csv")
	rows3 := genRows("cities3.csv")

	// Fan-in: Consolidating multiple channels into one channel
	// Multiplexing each message

	// worker1 ----\
	// worker2 -----> mergeRows ----> rows
	// worker3 ----/

	rows := mergeRows(rows1, rows2, rows3)

	highlyPopulatedCities := filterPopulation(20000000)
	cities := formatCityName(highlyPopulatedCities(rows))

	// Fan-Out: Break up of One channel into multiple ones
	// by distributing each city/message to a func timeConsumingTask()

	//            /----worker1
	//   rows ---<-----worker2
	//            \----   ...
	//             \---workerN

	result := fanOut(cities)

	// Print the result
	for r := range result {
		fmt.Println(r)
	}
}

func genRows(fileName string) <-chan City {
	// open the csv file for reading
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal("error: file could not be opened\n", err)
	}

	out := make(chan City)
	go func() {
		defer f.Close()
		reader := csv.NewReader(f)
		// Skip the first row for csv headers
		if _, err := reader.Read(); err != nil {
			log.Fatal("error: csv header could not be read\n", err)
		}
		for {
			row, err := reader.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatal("error: csv row could not be read\n", err)
			}
			// Dataset may not be clean. Skip the row if population cant be parsed
			populationInt, err := strconv.Atoi(row[9])
			if err != nil {
				continue
			}
			out <- City{
				Name:       row[1],
				Population: populationInt,
			}
		}
		close(out)
	}()
	return out
}

func formatCityName(cities <-chan City) <-chan City {
	out := make(chan City)
	go func() {
		for c := range cities {
			c.Name = strings.ToLower(c.Name)
			out <- c
		}
		close(out)
	}()
	return out
}

func filterPopulation(min int) func(<-chan City) <-chan City {
	return func(cities <-chan City) <-chan City {
		out := make(chan City)
		go func() {
			for c := range cities {
				if c.Population >= min {
					out <- c
				}
			}
			close(out)
		}()
		return out
	}
}

func mergeRows(rows ...<-chan City) <-chan City {
	out := make(chan City)
	wg := &sync.WaitGroup{}
	wg.Add(len(rows))
	for _, cities := range rows {
		go func(c <-chan City) {
			for city := range c {
				out <- city
			}
			wg.Done()
		}(cities)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func fanOut(cities <-chan City) <-chan City {
	out := make(chan City)
	go func() {
		wg := &sync.WaitGroup{}
		for city := range cities {
			wg.Add(1)
			go func(c City) {
				out <- timeConsumingTask(c)
				wg.Done()
			}(city)
		}
		go func() {
			wg.Wait()
			close(out)
		}()
	}()
	return out
}

func timeConsumingTask(city City) City {
	time.Sleep(time.Duration(rand.Intn(400)+100) * time.Millisecond)
	city.Status = "Processed"
	return city
}
