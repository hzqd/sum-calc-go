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

func sumCalc[O ext.Option[float64], R ext.Result[float64, string]](str string) float64 {
	ss := strings.FieldsFunc(str, func(c rune) bool {
		return c != '.' && c != '-' && c < '0' || c > '9'
	})
	floats := ext.FilterMap[string, float64](ss, func(s string) O {
		var tmp O
		if num, err := strconv.ParseFloat(s, 64); err != nil {
			fmt.Println(err)
		} else {
			tmp = ext.NewSome[float64, O](num)
		}
		return tmp
	})
	option := ext.Reduce[float64, O, R](floats, func(acc float64, x float64) float64 {
		return acc + x
	})
	result := ext.OkOrElse[float64, string, O, R](option, func() string {
		return "Option has None value."
	})
	return ext.UnwrapOk[float64, string](result)
}

func main() {
	for {
		str := readLn()
		if len(str) == 0 {
			break
		}
		res := sumCalc[ext.Some[float64], ext.Ok[float64]](str)
		fmt.Println(res)
	}
}
