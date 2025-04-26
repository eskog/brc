package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

//files are:
//1b.txt
//1m.txt
//10k.txt

const inputfile = "/Users/skogen/Projects/1brc/data/1b.txt"

type measurepoint struct {
	min, max, sum float32
	count         int
}

func newMeasurepoint(value float32) *measurepoint {
	return &measurepoint{
		min:   value,
		max:   value,
		sum:   value,
		count: 1,
	}
}

func insertData(mp *measurepoint, value float32) *measurepoint {
	if mp.min > value {
		mp.min = value
	}
	if mp.max < value {
		mp.max = value
	}
	mp.sum += value
	mp.count += 1
	return mp
}

func parse(mp map[string]*measurepoint, stationname string, value float32) {
	station, exists := mp[stationname]
	if !exists {
		mp[stationname] = newMeasurepoint(value)
		return
	}
	mp[stationname] = insertData(station, value)
}

func main() {
	starttime := time.Now()
	defer func() {
		fmt.Printf("Execution time: %s", time.Since(starttime))
	}()

	measurePoints := make(map[string]*measurepoint)
	file, err := os.Open(inputfile)
	if err != nil {
		log.Fatalln("Unable to open file")
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Printf("Unable to read line: %s", err.Error())
		}
		vals := strings.Split(line, ";")
		parsedFloat, err := strconv.ParseFloat(strings.TrimSuffix(vals[1], "\n"), 32)
		if err != nil {
			log.Printf("could not parse string: %s to float: %s", vals[1], err.Error())
		}

		parse(measurePoints, vals[0], float32(parsedFloat))

	}
	for stationname, mp := range measurePoints {
		fmt.Printf("%s: %.1f/%.1f/%.1f\n", stationname, mp.min, mp.sum/float32(mp.count), mp.max)
	}
}
