package templateFunc

import (
	. "controllers"
)

func init() {
	App.AddTemplateFunc("op", func(op string, xs ...float64) interface{} {
		l := len(xs)
		if l < 1 {
			return 0
		}

		var i interface{}
		switch op {
		case "+":
			sum := xs[0]
			for _, x := range xs[1:] {
				sum += x
			}

			i = sum
		case "-":
			sum := xs[0]
			for _, x := range xs[1:] {
				sum -= x
			}

			i = sum
		case "*":
			sum := xs[0]
			for _, x := range xs[1:] {
				sum *= x
			}

			i = sum
		case "/":
			sum := xs[0]
			for _, x := range xs[1:] {
				sum /= x
			}

			i = sum
		case "%":
			i = float64(int64(xs[0]) % int64(xs[1]))
		case "==":
			i = xs[0] == xs[1]
		case "!=":
			i = xs[0] != xs[1]
		case ">":
			i = xs[0] > xs[1]
		case ">=":
			i = xs[0] >= xs[1]
		case "<":
			i = xs[0] < xs[1]
		case "<=":
			i = xs[0] <= xs[1]
		}

		return i
	})
}
