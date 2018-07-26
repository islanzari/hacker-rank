package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func repeating(s string, n int64) string {

	//lenghtStr := len(s)
	for i := int64(0); i < n; i++ {
		if int64(len(s)) < n {
			s += s
		} else {
			s = s[:n]
		}
	}
	//	fmt.Println(str)
	return s
}

// Complete the repeatedString function below.
func repeatedString(s string, n int64) int64 {
	rep := repeating(s, n)
	var countA int64
	if len(rep) == 1 {
		return n
	}
	for i := 0; i < len(rep); i++ {
		if rep[i] == 'a' {
			countA++
		}
	}
	return countA
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	//stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	//checkError(err)

	//defer stdout.Close()

	writer := bufio.NewWriterSize(os.Stdout, 1024*1024)

	s := readLine(reader)

	n, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)

	result := repeatedString(s, n)

	fmt.Fprintf(writer, "%d\n", result)

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
