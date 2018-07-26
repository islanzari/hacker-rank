package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func reverse(i int32) int32 {
	var modulo int32
	var revers int32
	for z := 0; z < 6; z++ {

		modulo = i % 10
		i = (i - modulo) / 10
		revers = revers*10 + modulo
		//fmt.Println(revers)
		if i <= 0 {
			break
		}
	}
	return revers
}

// Complete the beautifulDays function below.
func beautifulDays(i int32, j int32, k int32) int32 {
	var count int32

	for i < j {

		if (i-reverse(i))%k == 0 {
			count++
		}

		i++
	}
	return count
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	//stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	//checkError(err)

	// defer stdout.Close()

	writer := bufio.NewWriterSize(os.Stdout, 1024*1024)

	ijk := strings.Split(readLine(reader), " ")

	iTemp, err := strconv.ParseInt(ijk[0], 10, 64)
	checkError(err)
	i := int32(iTemp)

	jTemp, err := strconv.ParseInt(ijk[1], 10, 64)
	checkError(err)
	j := int32(jTemp)

	kTemp, err := strconv.ParseInt(ijk[2], 10, 64)
	checkError(err)
	k := int32(kTemp)

	result := beautifulDays(i, j, k)

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
