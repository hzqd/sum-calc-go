package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sum-calc/ext"
)

func main() {
	for {
		reader := bufio.NewReader(os.Stdin)
		read, _ := reader.ReadString('\n')
		str := strings.TrimRight(read, "\n")
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
		result := ext.FoldDefault(floats, func(acc float64, x float64) float64 {
			return acc + x
		})
		fmt.Println(result)
	}
}
