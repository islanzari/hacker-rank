package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the catAndMouse function below.
func catAndMouse(x int32, y int32, z int32) string {

	var fin string
	for i := 0; i < 100; i++ {
		if x > z {
			x--

		}
		if y > z {
			y--

		}
		if y < z {
			y++

		}
		if x < z {
			x++

		}
		{

			if x == z {
				if y != z {
					return "Cat A"
				}
			}
			if y == z {
				if x != z {
					return "Cat B"
				}

			}
			if y == z && x == z {

				return "Mouse C"

			}
		}

	}

	return fin
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	//  stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	// checkError(err)

	//defer stdout.Close()

	writer := bufio.NewWriterSize(os.Stdout, 1024*1024)

	qTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	q := int32(qTemp)

	for qItr := 0; qItr < int(q); qItr++ {
		xyz := strings.Split(readLine(reader), " ")

		xTemp, err := strconv.ParseInt(xyz[0], 10, 64)
		checkError(err)
		x := int32(xTemp)

		yTemp, err := strconv.ParseInt(xyz[1], 10, 64)
		checkError(err)
		y := int32(yTemp)

		zTemp, err := strconv.ParseInt(xyz[2], 10, 64)
		checkError(err)
		z := int32(zTemp)

		result := catAndMouse(x, y, z)

		fmt.Fprintf(writer, "%s\n", result)
	}

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
