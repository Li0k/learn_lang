package solution2

import (
	"testing"
	"os"
	"github.com/edsrzf/mmap-go"
	"encoding/csv"
	"bytes"
	"io"
	"fmt"
	"strconv"
	"time"
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

	dict = make(map[int]int,100*10000)

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

//单goroutine mmap 异步
func TestTestSolution2(t *testing.T) {
	fmt.Println("I am coming back go")
	t1 := time.Now()

	//构建第一个hash表
	dict1, err := FileToDictMmap("../t1.csv")
	if err != nil {
		panic(err)
	}

	f2, err := os.Open("../t2.csv")
	if err != nil {
		t.Error(err)
	}
	//映射为mmap
	mmap, err := mmap.Map(f2, mmap.RDONLY, 0)
	if err != nil {
		t.Error(err)
	}

	ret := csv.NewReader(bytes.NewReader(mmap))
	var (
		count int
	)

	c := make(chan int, 1024*1024)

	//解析为record，分发
	go func() {
		defer close(c)
		for {
			record, err := ret.Read()

			if err == io.EOF {
				break
			}

			if err != nil {
				t.Error(err)
			}

			idx, _ := strconv.Atoi(record[0])

			c <- idx
		}
	}()

	//var wg sync.WaitGroup

	//wg.Add(1)

	//使用flag来同步
	flag := make(chan int,1)
	go func() {
		//defer  wg.Done()
		for e := range c {
			if t2V, ok := dict1[e]; ok {
				count += t2V
			}
		}

		flag <- 1
	}()

	//wg.Wait()
	<-flag

	fmt.Println(count)
	fmt.Println(time.Since(t1))
}

func BenchmarkSolution2(b *testing.B) {
	b.ReportAllocs()

	for i := 0;i < b.N ;i++ {
		fmt.Println("I am coming back go")
		t1 := time.Now()

		dict1, err := FileToDictMmap("../t1.csv")
		if err != nil {
			panic(err)
		}

		f2, err := os.Open("../t2.csv")
		if err != nil {
			b.Error(err)
		}
		mmap, err := mmap.Map(f2, mmap.RDONLY, 0)
		if err != nil {
			b.Error(err)
		}

		ret := csv.NewReader(bytes.NewReader(mmap))
		var (
			count int
		)

		c := make(chan int, 1024*1024)
		go func() {
			defer close(c)
			for {
				record, err := ret.Read()

				if err == io.EOF {
					break
				}

				if err != nil {
					b.Error(err)
				}

				idx, _ := strconv.Atoi(record[0])

				c <- idx
			}
		}()

		//var wg sync.WaitGroup

		//wg.Add(1)

		flag := make(chan int)
		go func() {
			//defer  wg.Done()
			for e := range c {
				if t2V, ok := dict1[e]; ok {
					count += t2V
				}
			}

			flag <- 1
		}()

		//wg.Wait()
		<-flag

		fmt.Println(count)
		fmt.Println(time.Since(t1))
	}
}
