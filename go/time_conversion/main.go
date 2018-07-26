package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the timeConversion function below.
 */
func timeConversion(s string) string {
	/*
	 * Write your code here.
	 */

	//k, err := strconv.Atoi(s)
	ost := s
	str := strings.Split(s, ":")   // rozdziela
	k, err := strconv.Atoi(str[0]) // string to int
	if err != nil {
		fmt.Println(err)
		return ""
	}
	last := string(s[8]) + string(s[9])

	if string(last) == "PM" {

		if k != 12 {
			hour := k + 12
			hours := strconv.Itoa(hour)
			s = hours + ":" + str[1] + ":" + str[2]

		}
		ost = strings.TrimSuffix(s, "PM")
	}

	if last == "AM" {

		if k == 12 {
			str[0] = "00"
			s = str[0] + ":" + str[1] + ":" + str[2]

		}
		ost = strings.TrimSuffix(s, "AM")
	}

	strings.TrimSuffix(s, "AM")
	return ost
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	outputFile, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer outputFile.Close()

	writer := bufio.NewWriterSize(outputFile, 1024*1024)

	s := readLine(reader)

	result := timeConversion(s)

	fmt.Fprintf(writer, "%s\n", result)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
