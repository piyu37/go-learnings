package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

/*
 * Complete the 'rateLimiter' function below.
 *
 * The function accepts following parameters:
 *  1. STRING inputFileName
 *  2. STRING outputFileName
 */

func rateLimiter(inputFileName string, outputFileName string) {
	doneTicker := make(chan bool)
	ticker := time.NewTicker(1 * time.Second)
	idMap := make(map[string]int)
	lock := sync.Mutex{}
	inputFile, err := os.Open(inputFileName)
	checkError(err)
	defer inputFile.Close()
	outputFile, err := os.Create(outputFileName)
	checkError(err)
	defer outputFile.Close()
	scanner := bufio.NewScanner(inputFile)
	go func() {
		for {
			select {
			case <-doneTicker:
				ticker.Stop()
				break
			case <-ticker.C:
				refreshMap(&idMap, &lock)
			}
		}
	}()

	scanner.Scan()

	for scanner.Scan() {
		line := scanner.Text()
		lineValues := strings.Split(line, ",")
		lock.Lock()
		if idMap[lineValues[0]] >= 5 {
			_, err := outputFile.Write([]byte("BLOCKED\n"))
			checkError(err)
		} else {
			_, err := outputFile.Write([]byte("PASSED\n"))
			checkError(err)
		}

		idMap[lineValues[0]] += 1
		lock.Unlock()
	}

	close(doneTicker)
}

func refreshMap(idMap *map[string]int, lock *sync.Mutex) {
	lock.Lock()
	*idMap = map[string]int{}
	lock.Unlock()
}

func rateLimiterMain() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int(nTemp)
	nStr := strconv.Itoa(n) + "\n"

	file, err := os.Create(inputFileName)
	checkError(err)
	_, err = os.Create(outputFileName)
	checkError(err)
	_, err = file.WriteString(nStr)
	checkError(err)
	for i := 0; i < n; i++ {
		inputString := readLine(reader)
		_, err = file.WriteString(inputString)
		checkError(err)
	}
	defer file.Close()
	rateLimiter(inputFileName, outputFileName)

	//validate the outputFileName
	rateLimiterOutput, errErr := ioutil.ReadFile(outputFileName)
	checkError(errErr)
	fmt.Println(string(rateLimiterOutput))
}

const inputFileName = "input"
const outputFileName = "output"

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return string(str) + "\n"
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
