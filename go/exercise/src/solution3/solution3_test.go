package solution3

import (
	"testing"
	"time"
	"os"
	"github.com/edsrzf/mmap-go"
	"encoding/csv"
	"bytes"
	"io"
	"strconv"
	"sync"
	"fmt"
	"math/rand"
	"runtime"
)

func FileToDictMmap(file string) (map[int]int, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	mmap, err := mmap.Map(f, mmap.RDONLY, 0)
	if err != nil {
		return nil, err
	}
	defer mmap.Unmap()

	ret := csv.NewReader(bytes.NewReader(mmap))

	var dict map[int]int

	dict = make(map[int]int, 100*10000)

	for {
		record, err := ret.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		idx, _ := strconv.Atoi(record[0])
		dict[idx]++
	}

	return dict, nil
}

//多goroutine mmap
func TestTestSolution3(t *testing.T) {
	N := runtime.NumCPU()
	t1 := time.Now()

	//构建hash表
	dict, err := FileToDictMmap("../t2.csv")

	if err != nil {
		t.Error(err)
	}

	f1, err := os.Open("../t1.csv")
	if err != nil {
		t.Error(err)
	}

	defer f1.Close()

	//映射为mmap
	mmap, err := mmap.Map(f1, mmap.RDONLY, 0)
	if err != nil {
		t.Error(err)
	}

	ret := csv.NewReader(bytes.NewReader(mmap))
	//构建作为分发信息的chan
	chans := make([]chan int, N)
	for i := 0; i < N; i++ {
		chans[i] = make(chan int, 1024*1024)
	}

	go func() {
		for {
			record, err := ret.Read()

			if err == io.EOF {
				break
			}

			if err != nil {
				t.Error(err)
			}

			idx, _ := strconv.Atoi(record[0])
			r_idx := rand.Intn(N)
			chans[r_idx] <- idx
		}

		for i := 0; i < N; i++ {
			close(chans[i])
		}
	}()

	var (
		count      int
		countSlice []int
	)

	//分开汇总
	countSlice = make([]int, N)

	var wg sync.WaitGroup
	wg.Add(N)

	//启动多个goroutine进行比较
	for i := 0; i < N; i++ {
		go func(idx int, c chan int) {
			defer wg.Done()
			for x := range c {
				if t2V, ok := dict[x]; ok {
					countSlice[idx] = countSlice[idx] + t2V
				}
			}
		}(i, chans[i])
	}

	wg.Wait()

	for i := 0; i < N; i++ {
		count += countSlice[i]
	}

	fmt.Println(count)
	fmt.Println(time.Since(t1))

}
func BenchmarkSolution3(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		const N = 2
		t1 := time.Now()
		dict, err := FileToDictMmap("../t2.csv")

		if err != nil {
			b.Error(err)
		}

		f1, err := os.Open("../t1.csv")
		if err != nil {
			b.Error(err)
		}

		defer f1.Close()

		mmap, err := mmap.Map(f1, mmap.RDONLY, 0)
		if err != nil {
			b.Error(err)
		}

		ret := csv.NewReader(bytes.NewReader(mmap))
		chans := make([]chan int, N)
		for i := 0; i < N; i++ {
			chans[i] = make(chan int, 1024*1024)
		}

		go func() {
			for {
				record, err := ret.Read()

				if err == io.EOF {
					break
				}

				if err != nil {
					b.Error(err)
				}

				idx, _ := strconv.Atoi(record[0])
				r_idx := rand.Intn(N)
				chans[r_idx] <- idx
			}

			for i := 0; i < N; i++ {
				close(chans[i])
			}
		}()

		var (
			count      int
			countSlice []int
		)

		countSlice = make([]int, N)

		var wg sync.WaitGroup
		wg.Add(N)

		for i := 0; i < N; i++ {
			go func(idx int, c chan int) {
				defer wg.Done()
				for x := range c {
					if t2V, ok := dict[x]; ok {
						countSlice[idx] = countSlice[idx] + t2V
					}
				}
			}(i, chans[i])
		}

		wg.Wait()

		for i := 0; i < N; i++ {
			count += countSlice[i]
		}

		fmt.Println(count)
		fmt.Println(time.Since(t1))

	}
}
