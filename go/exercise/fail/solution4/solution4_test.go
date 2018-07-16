package solution4

import (
	"os"
	"github.com/edsrzf/mmap-go"
	"encoding/csv"
	"bytes"
	"io"
	"fmt"
	"strconv"
	"testing"
	"time"
	"sync"
	"math/rand"
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

	dict = make(map[int]int)

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
func TestTestSolution4(t *testing.T) {
	const N = 1
	t1 := time.Now()
	dict, err := FileToDictMmap("../t1.csv")

	if err != nil {
		t.Error(err)
	}

	f1, err := os.Open("../t2.csv")
	if err != nil {
		t.Error(err)
	}

	defer f1.Close()

	mmap, err := mmap.Map(f1, mmap.RDONLY, 0)
	if err != nil {
		t.Error(err)
	}

	//ret := csv.NewReader(bytes.NewReader(mmap))
	chans := make([]chan int, N)
	for i := 0; i < N; i++ {
		chans[i] = make(chan int, 1024*2500)
	}

	//const disPatCount  = 4
	var wgDis sync.WaitGroup
	wgDis.Add(N)

	//var lockdis sync.RWMutex
	s := make([]int, 0)
	ret := csv.NewReader(bytes.NewReader(mmap))
	for {
		record, err := ret.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			t.Error(err)
		}

		idx, _ := strconv.Atoi(record[0])
		//r_idx := rand.Intn(N)
		//chans[r_idx] <- idx
		s = append(s, idx)
	}

	offset := len(s) / N
	//fmt.Println(len(s))

	for i := 0; i < N; i++ {
		go func(idx int) {
			defer wgDis.Done()

			for i :=offset * idx;i < offset*(idx + 1);i++ {
				r_idx := rand.Intn(N)
				chans[r_idx] <- s[i]
			}

		}(i)
	}

	//closer
	go func() {
		wgDis.Wait()

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

func TestSolution44(t *testing.T) {
	t1 := time.Now()

	f1, err := os.Open("../t4.csv")
	if err != nil {
		t.Error(err)
	}

	defer f1.Close()

	mmap, err := mmap.Map(f1, mmap.RDONLY, 0)
	if err != nil {
		t.Error(err)
	}
	const N = 4
	fmt.Println(len(mmap))
	fmt.Println(len(string(mmap)))

	//fmt.Println(len(mmap))
	ret := csv.NewReader(bytes.NewReader(mmap))
	var wgDis sync.WaitGroup
	wgDis.Add(N)

	s := make([]int, 0)
	for {
		//lockdis.Lock()
		record, err := ret.Read()
		//lockdis.Unlock()
		if err == io.EOF {
			break
		}

		if err != nil {
			t.Error(err)
		}

		idx, _ := strconv.Atoi(record[0])
		s = append(s, idx)
	}

	offset := len(s) / N
	fmt.Println(len(s))

	//var lockdis sync.RWMutex
	for i := 0; i < N; i++ {
		//fmt.Println(i)
		go func(idx int) {
			fmt.Println("start")
			defer wgDis.Done()
			for e := range s[offset*idx : offset*(idx+1)] {
				fmt.Println(e)
			}
		}(i)
	}
	wgDis.Wait()

	fmt.Println(time.Since(t1))

}
