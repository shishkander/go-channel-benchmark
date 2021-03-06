package main

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func benchmarkChannel(b *testing.B, P, S, T int) {
	var c chan byte
	if S == 0 {
		c = make(chan byte)
	} else {
		c = make(chan byte, S)
	}
	wg := sync.WaitGroup{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		wg.Add(P + 1)
		go func() {
			for j := 0; j < T*P; j++ {
				c <- 1
			}
			wg.Done()
		}()
		for k := 0; k < P; k++ {
			go func() {
				for j := 0; j < T; j++ {
					<-c
					time.Sleep(time.Microsecond)
				}
				wg.Done()
			}()
		}
		b.SetBytes(int64(T) * int64(P))
		wg.Wait()
	}
	b.StopTimer()
	close(c)
}

func makeBenchmarkFunc(p, s, t int) func(b *testing.B) {
	return func(b *testing.B) {
		benchmarkChannel(b, p, s, t)
	}
}

func main() {
	Ss := []int{0, 10, 50, 100, 500, 1000, 1500, 5000}
	Ts := []int{0, 10, 50, 100, 150, 250, 400, 800}
	Ps := []int{1, 4, 8, 16}
	//MaxProcs := []int{4, 3, 2, 1}
	MaxProcs := []int{4, 2, 1}

	tabTable := "        " // 8 spaces.
	printDivider := func() {
		fmt.Print(tabTable + "|")
		for i := 1; i < 13+len(Ts)*8-1; i++ {
			fmt.Print("=")
		}
		fmt.Println("|")
	}
	printHeader := func(p int) {
		printDivider()
		fmt.Printf(tabTable+"| %9s |", "buf\\tasks")
		for _, t := range Ts {
			if t == 0 {
				fmt.Printf(" %5s |", "none")
			} else {
				fmt.Printf(" %5d |", p*t)
			}
		}
		fmt.Println()
		printDivider()
	}
	for _, c := range MaxProcs {
		runtime.GOMAXPROCS(c)
		for _, p := range Ps {
			fmt.Printf("\n### GOMAXPROCS = %02d, Producers = %02d, Consumers = %02d\n\n", c, 1, p)
			printHeader(p)
			for _, s := range Ss {
				if s == 0 {
					fmt.Printf(tabTable+"| %5s     |", "unbuf")
				} else {
					fmt.Printf(tabTable+"| %5d     |", s)
				}
				for _, t := range Ts {
					r := testing.Benchmark(makeBenchmarkFunc(p, s, t))
					sp := r.T.Seconds() / float64(r.N)
					if t > 0 {
						sp = sp / float64(t) / float64(p)
					}
					fmt.Printf(" %5.1f |", sp*1e6)
				}
				fmt.Println()
			}
			printDivider()
		}
	}
}
