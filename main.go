package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type measurepoint struct {
	name          string
	min, max, sum float32
	count         int
}

func newMeasurepoint(stationname string, value float32) measurepoint {
	return measurepoint{
		name:  stationname,
		min:   value,
		max:   value,
		sum:   value,
		count: 1,
	}
}

func insertData(mp measurepoint, stationname string, value float32) measurepoint {
	if mp.name != stationname {
		return mp
	}
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

func parse(mp *[]measurepoint, stationname string, value float32) {
	for i := 0; i < len(*mp); i++ {
		if (*mp)[i].name == stationname {
			(*mp)[i] = insertData((*mp)[i], stationname, value)
			return
		}
	}
	*mp = append(*mp, newMeasurepoint(stationname, value))
}

func main() {

	measurePoints := new([]measurepoint)
	file, err := os.Open("input-short.txt")
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
	for _, mp := range *measurePoints {
		fmt.Printf("%s: %.1f/%.1f/%.1f\n", mp.name, mp.min, mp.sum/float32(mp.count), mp.max)
	}
}
