package solution1

import (
	"encoding/csv"
	"io"
	"io/ioutil"
	"strconv"
	"strings"
	"os"
	"fmt"
	"bufio"
	"testing"
	"time"
	_ "net/http/pprof"
)

func FileToDict(file string) map[int]int {
	data, err := ioutil.ReadFile(file)

	if err != nil {
		panic(err)
	}

	var dict map[int]int

	dict = make(map[int]int)

	ret := csv.NewReader(strings.NewReader(string(data[:])))

	for {
		record, err := ret.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			panic(err)
		}

		idx, _ := strconv.Atoi(record[0])
		//fmt.Println(idx)
		dict[idx]++
	}

	return dict
}
func FileToDictBuffIo(file string) (map[int]int, error) {
	f, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var dict map[int]int

	dict = make(map[int]int,100*10000)

	ret := csv.NewReader(bufio.NewReader(f))

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
		//fmt.Println(idx)
		dict[idx]++
	}

	return dict, nil

}



//单goroutine buffio
func TestTestSolution1(t *testing.T) {
	fmt.Println("I am coming back go")
	t1 := time.Now()

	dict1, err := FileToDictBuffIo("../t1.csv")

	if err != nil {
		panic(err)
	}

	dict2, err := FileToDictBuffIo("../t2.csv")

	if err != nil {
		panic(err)
	}

	var (
		count int
	)

	for t1_k, t1_v := range dict1 {
		if t2_v, ok := dict2[t1_k]; ok {
			count = count + t1_v*t2_v
		}
	}

	fmt.Println(count)
	fmt.Println(time.Since(t1))
}
//
////单goroutine mmap
//func TestTestSolution6(t *testing.T) {
//	fmt.Println("I am coming back go")
//	t1 := time.Now()
//
//	dict1, err := FileToDictMmap("t1.csv")
//	if err != nil {
//		panic(err)
//	}
//
//	dict2, err := FileToDictMmap("t2.csv")
//
//	if err != nil {
//		panic(err)
//	}
//
//	var (
//		count int
//	)
//
//	for t1_k, t1_v := range dict1 {
//		if t2_v, ok := dict2[t1_k]; ok {
//			count = count + t1_v*t2_v
//		}
//	}
//
//	fmt.Println(count)
//	fmt.Println(time.Since(t1))
//}
//
////多goroutine mmap
//func TestTestSolution4(t *testing.T) {
//	const N = 4
//	t1 := time.Now()
//	dict, err := FileToDictMmap("t2.csv")
//
//	if err != nil {
//		t.Error(err)
//	}
//
//	f1, err := os.Open("t1.csv")
//	if err != nil {
//		t.Error(err)
//	}
//
//	defer f1.Close()
//
//	mmap, err := mmap.Map(f1, mmap.RDONLY, 0)
//	if err != nil {
//		t.Error(err)
//	}
//
//	//ret := csv.NewReader(strings.NewReader(string(mmap)))
//	ret := csv.NewReader(bytes.NewReader(mmap))
//	chans := make([]chan int, N)
//	for i := 0; i < N; i++ {
//		chans[i] = make(chan int, 1024)
//	}
//
//	go func() {
//		for {
//			record, err := ret.Read()
//
//			if err == io.EOF {
//				break
//			}
//
//			if err != nil {
//				t.Error(err)
//			}
//
//			idx, _ := strconv.Atoi(record[0])
//			r_idx := rand.Intn(N)
//			chans[r_idx] <- idx
//		}
//
//		for i := 0; i < N; i++ {
//			close(chans[i])
//		}
//	}()
//
//	var (
//		count      int
//		countSlice []int
//	)
//
//	countSlice = make([]int, N)
//
//	var wg sync.WaitGroup
//	wg.Add(N)
//
//	for i := 0; i < N; i++ {
//		go func(idx int, c chan int) {
//			defer wg.Done()
//			for x := range c {
//				if t2V, ok := dict[x]; ok {
//					countSlice[idx] = countSlice[idx] + t2V
//				}
//			}
//		}(i, chans[i])
//	}
//
//	wg.Wait()
//
//	for i := 0; i < N; i++ {
//		count += countSlice[i]
//	}
//
//	fmt.Println(count)
//
//	fmt.Println(time.Since(t1))
//
//}

func BenchmarkSolution1(b *testing.B){
	b.ReportAllocs()
	fmt.Println("I am coming back go")
	t1 := time.Now()

	dict1, err := FileToDictBuffIo("../t1.csv")

	if err != nil {
		b.Error(err)
	}

	dict2, err := FileToDictBuffIo("../t2.csv")

	if err != nil {
		b.Error(err)
	}

	var (
		count int
	)

	for t1_k, t1_v := range dict1 {
		if t2_v, ok := dict2[t1_k]; ok {
			count = count + t1_v*t2_v
		}
	}

	fmt.Println(count)
	fmt.Println(time.Since(t1))
}