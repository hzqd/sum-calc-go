package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sum-calc/ext"
)

func readLn() string {
	r, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	return strings.TrimRight(r, "\n")
}

func main() {
	for {
		str := readLn()
		if len(str) == 0 {
			break
		}
		ss := strings.FieldsFunc(str, func(c rune) bool {
			return c != '.' && c != '-' && c < '0' || c > '9'
		})
		floats := ext.FilterMap(ss, func(s string) (float64, bool) {
			num, err := strconv.ParseFloat(s, 64)
			if err != nil {
				fmt.Println(err)
			}
			return num, err == nil
		})
		result, _ := ext.Reduce(floats, func(acc float64, x float64) float64 {
			return acc + x
		})
		fmt.Println(result)
	}
}
