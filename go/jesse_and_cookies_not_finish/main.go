package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
 * Complete the cookies function below.
 */
func cookies(k int32, A []int32) int32 {
	var count int32
	s := make([]int, len(A))

	for i := 0; i < len(A); i++ {
		s[i] = int(A[i])
		//if A[i] < k {
		//    A[i] = 1*A[i] + 2*A[i+1]
		//s[i] = int(A[i])
		//sort.Ints(s)
	}

	for j := 0; j < len(s)-1; j++ {

		if s[j] < int(k) {

			s[j] = 1*s[j] + 2*s[j+1]
			count++
		}
		sort.Ints(s)
	}
	if k > int32(s[len(s)-1]) {
		count = -1
	}
	return count
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	// stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	//checkError(err)

	//defer stdout.Close()

	writer := bufio.NewWriterSize(os.Stdout, 1024*1024)

	nk := strings.Split(readLine(reader), " ")

	nTemp, err := strconv.ParseInt(nk[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	kTemp, err := strconv.ParseInt(nk[1], 10, 64)
	checkError(err)
	k := int32(kTemp)

	ATemp := strings.Split(readLine(reader), " ")

	var A []int32

	for AItr := 0; AItr < int(n); AItr++ {
		AItemTemp, err := strconv.ParseInt(ATemp[AItr], 10, 64)
		checkError(err)
		AItem := int32(AItemTemp)
		A = append(A, AItem)
	}

	result := cookies(k, A)

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
