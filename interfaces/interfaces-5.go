package main

import "fmt"

type Seq interface {
	Next() (x float64, ok bool)
}

type SeqFunc func() (x float64, ok bool)

func (sf SeqFunc) Next() (x float64, ok bool) {
	return sf()
}

func Mk(vals ...float64) Seq {
	i := 0
	return SeqFunc(func() (x float64, ok bool) {
		if i >= len(vals) {
			return 0, false
		}
		x = vals[i]
		i++
		return x, true
	})
}

func Naturals() Seq {
	n := 0.0
	return SeqFunc(func() (x float64, ok bool) {
		n++
		return n, true
	})
}

func Take(n int, s Seq) Seq {
	return SeqFunc(func() (x float64, ok bool) {
		if n <= 0 {
			return 0, false
		}
		n--
		return s.Next()
	})
}

func Collect(s Seq) []float64 {
	var vals []float64
	for {
		x, ok := s.Next()
		if !ok {
			break
		}
		vals = append(vals, x)
	}
	return vals
}

func main() {
	fmt.Println(Collect(Take(10, Naturals())))
}
