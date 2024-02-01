package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

type City struct {
	City       string
	Country    string
	Population int
}

func genRows(r io.Reader) chan City {
	out := make(chan City)
	go func() {
		reader := csv.NewReader(r)
		// read the first line
		_, err := reader.Read()
		if err != nil {
			log.Fatal(err)
		}

		for {
			row, err := reader.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatal(err)
			}
			if row[9] == "" {

			}
			populationInt, err := strconv.Atoi(row[9])
			if err != nil {
				populationInt = 0
			}
			out <- City{
				City:       row[1],
				Country:    row[5],
				Population: populationInt,
			}
		}
		close(out)
	}()
	return out
}

func upperCityRows(rows <-chan City) <-chan City {
	out := make(chan City)
	go func() {
		for c := range rows {
			out <- City{City: strings.ToUpper(c.City), Population: c.Population}
		}
		close(out)
	}()
	return out
}

func filterOutPopulation(cities <-chan City) <-chan City {
	out := make(chan City)
	go func() {
		for c := range cities {
			if c.Population > 40000 {
				out <- c
			}
		}
		close(out)
	}()
	return out
}

func fanIn(workers ...<-chan City) <-chan City {
	wg := &sync.WaitGroup{}
	wg.Add(len(workers))
	out := make(chan City)

	for _, w := range workers {
		go func(c <-chan City) {
			for r := range c {
				out <- r
			}
			wg.Done()
		}(w)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func main() {
	file, err := os.Open("cities/worldcities.csv")
	if err != nil {
		log.Fatal(err)
	}
	start := time.Now()

	defer file.Close()

	rows := genRows(file)

	ur1 := upperCityRows(filterOutPopulation(rows))
	ur2 := upperCityRows(filterOutPopulation(rows))
	ur3 := upperCityRows(filterOutPopulation(rows))
	ur4 := upperCityRows(filterOutPopulation(rows))

	cities := fanIn(ur1, ur2, ur3, ur4)
	for c := range cities {
		fmt.Println(c.City, c.Population)
	}

	duration := time.Now().Sub(start)
	fmt.Println("total time: ", duration)
}
