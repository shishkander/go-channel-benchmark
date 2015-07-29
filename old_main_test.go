package main

import (
	"runtime"
	"sync"
	"testing"
)

func BenchmarkChanUNBx001(b *testing.B) { benchmarkChannelOld(b, 0, 1) }

func BenchmarkChan001x001(b *testing.B) { benchmarkChannelOld(b, 1, 1) }
func BenchmarkChan010x001(b *testing.B) { benchmarkChannelOld(b, 10, 1) }
func BenchmarkChan100x001(b *testing.B) { benchmarkChannelOld(b, 100, 1) }
func BenchmarkChan150x001(b *testing.B) { benchmarkChannelOld(b, 150, 1) }
func BenchmarkChan01Kx001(b *testing.B) { benchmarkChannelOld(b, 1000, 1) }
func BenchmarkChan10Kx001(b *testing.B) { benchmarkChannelOld(b, 10000, 1) }

func BenchmarkChanUNBx010(b *testing.B) { benchmarkChannelOld(b, 0, 10) }
func BenchmarkChan001x010(b *testing.B) { benchmarkChannelOld(b, 1, 10) }
func BenchmarkChan010x010(b *testing.B) { benchmarkChannelOld(b, 10, 10) }
func BenchmarkChan100x010(b *testing.B) { benchmarkChannelOld(b, 100, 10) }
func BenchmarkChan150x010(b *testing.B) { benchmarkChannelOld(b, 150, 10) }
func BenchmarkChan01Kx010(b *testing.B) { benchmarkChannelOld(b, 1000, 10) }
func BenchmarkChan10Kx010(b *testing.B) { benchmarkChannelOld(b, 10000, 10) }

func BenchmarkChanUNBx100(b *testing.B) { benchmarkChannelOld(b, 0, 100) }
func BenchmarkChan001x100(b *testing.B) { benchmarkChannelOld(b, 1, 100) }
func BenchmarkChan010x100(b *testing.B) { benchmarkChannelOld(b, 10, 100) }
func BenchmarkChan100x100(b *testing.B) { benchmarkChannelOld(b, 100, 100) }
func BenchmarkChan150x100(b *testing.B) { benchmarkChannelOld(b, 150, 100) }
func BenchmarkChan01Kx100(b *testing.B) { benchmarkChannelOld(b, 1000, 100) }
func BenchmarkChan10Kx100(b *testing.B) { benchmarkChannelOld(b, 10000, 100) }

func BenchmarkChanUNBx150(b *testing.B) { benchmarkChannelOld(b, 0, 150) }
func BenchmarkChan001x150(b *testing.B) { benchmarkChannelOld(b, 1, 150) }
func BenchmarkChan010x150(b *testing.B) { benchmarkChannelOld(b, 10, 150) }
func BenchmarkChan100x150(b *testing.B) { benchmarkChannelOld(b, 100, 150) }
func BenchmarkChan150x150(b *testing.B) { benchmarkChannelOld(b, 150, 150) }
func BenchmarkChan01Kx150(b *testing.B) { benchmarkChannelOld(b, 1000, 150) }
func BenchmarkChan10Kx150(b *testing.B) { benchmarkChannelOld(b, 10000, 150) }

func BenchmarkChanUNBx01K(b *testing.B) { benchmarkChannelOld(b, 0, 1000) }
func BenchmarkChan001x01K(b *testing.B) { benchmarkChannelOld(b, 1, 1000) }
func BenchmarkChan010x01K(b *testing.B) { benchmarkChannelOld(b, 10, 1000) }
func BenchmarkChan100x01K(b *testing.B) { benchmarkChannelOld(b, 100, 1000) }
func BenchmarkChan150x01K(b *testing.B) { benchmarkChannelOld(b, 150, 1000) }
func BenchmarkChan01Kx01K(b *testing.B) { benchmarkChannelOld(b, 1000, 1000) }
func BenchmarkChan10Kx01K(b *testing.B) { benchmarkChannelOld(b, 10000, 1000) }

func benchmarkChannelOld(b *testing.B, S, T int) {
	var c chan byte
	if S == 0 {
		c = make(chan byte)
	} else {
		c = make(chan byte, S)
	}
	wg := sync.WaitGroup{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		wg.Add(runtime.NumCPU() + 1)
		go func() {
			for j := 0; j < T*runtime.NumCPU(); j++ {
				c <- 1
			}
			wg.Done()
		}()
		for k := 0; k < runtime.NumCPU(); k++ {
			go func() {
				for j := 0; j < T; j++ {
					<-c
					//time.Sleep(time.Microsecond)
				}
				wg.Done()
			}()
		}
		b.SetBytes(int64(T) * int64(runtime.NumCPU()))
		wg.Wait()
	}
	b.StopTimer()
	close(c)
}
