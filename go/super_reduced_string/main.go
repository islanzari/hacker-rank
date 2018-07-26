package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// Complete the superReducedString function below.
func superReducedString(s string) string {

	var count int32

	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s); i++ {

			if s[i] == s[j] {
				count++

				if count == 2 {
					s = s[:i-1] + s[i+1:]

				}

			}
		}
	}

	return s
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	//stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	//checkError(err)

	//defer stdout.Close()

	writer := bufio.NewWriterSize(os.Stdout, 16*1024*1024)

	s := readLine(reader)

	result := superReducedString(s)

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
