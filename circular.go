package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

/*
 * Complete the 'getTime' function below.
 *
 * The function is expected to return a LONG_INTEGER.
 * The function accepts STRING s as parameter.
 */

func getTime(s string) int64 {
	// Write your code here
	var sum int64
	sum = 0
	var start int64
	for _, c := range s {
		end := int64(c) - 65
		res := start - end
		if res < 0 {
			res = res + 26
		}

		if res > 13 {
			res = 26 - res
		}
		fmt.Println(res)
		sum = sum + res
		start = end
	}
	return sum
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	os.Setenv("OUTPUT_PATH", "1")
	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	s := readLine(reader)

	result := getTime(s)

	fmt.Fprintf(writer, "%d\n", result)
	fmt.Println("Total", result)

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
