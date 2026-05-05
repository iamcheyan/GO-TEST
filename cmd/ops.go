package cmd

import (
	"log"
	"strconv"
)

var ops = map[string]func(float64, float64) float64{
	"add": func(a, b float64) float64 { return a + b },
	"sub": func(a, b float64) float64 { return a - b },
	"mul": func(a, b float64) float64 { return a * b },
	"div": func(a, b float64) float64 { return a / b },
	"mod": func(a, b float64) float64 { return float64(int(a) % int(b)) },
}

func parseOrFatal(s string) float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		log.Fatalf("invalid number: %s", s)
	}
	return f
}