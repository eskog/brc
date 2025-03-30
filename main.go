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

func main() {

	values := make(map[string][]float64)
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
		parsedFloat, err := strconv.ParseFloat(strings.TrimSuffix(vals[1], "\n"), 64)
		if err != nil {
			log.Printf("could not parse string: %s to float: %s", vals[1], err.Error())
		}
		values[vals[0]] = append(values[vals[0]], parsedFloat)

	}
	fmt.Println(values)
}
