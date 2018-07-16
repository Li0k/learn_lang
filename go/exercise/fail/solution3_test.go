package fail

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"testing"
	"time"
)

const N = 4

func DeepCopy(value interface{}) interface{} {
	if valueMap, ok := value.(map[string]interface{}); ok {
		newMap := make(map[string]interface{})
		for k, v := range valueMap {
			newMap[k] = DeepCopy(v)
		}

		return newMap
	} else if valueSlice, ok := value.([]interface{}); ok {
		newSlice := make([]interface{}, len(valueSlice))
		for k, v := range valueSlice {
			newSlice[k] = DeepCopy(v)
		}

		return newSlice
	}

	return value
}

func TestSolution3(t *testing.T) {
	t1 := time.Now()

	//分割map1
	var partition []map[int]int
	partition = make([]map[int]int, N)
	for i := 0; i < N; i++ {
		partition[i] = make(map[int]int)
	}

	f, err := os.Open("t1.csv")
	if err != nil {
		panic(err)
	}

	ret := csv.NewReader(bufio.NewReader(f))

	for {
		record, err := ret.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			panic(err)
		}

		//产生随机数，分配到partition中不同的map
		r_idx := rand.Intn(N)

		idx, _ := strconv.Atoi(record[0])

		partition[r_idx][idx]++
	}

	var dict map[int]int

	dict = make(map[int]int)

	f2, err := os.Open("t2.csv")
	if err != nil {
		panic(err)
	}

	ret2 := csv.NewReader(bufio.NewReader(f2))

	for {
		//按行读取
		record, err := ret2.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			panic(err)
		}

		idx, _ := strconv.Atoi(record[0])

		//统计第一列值，出现的次数
		dict[idx]++
	}

	var (
		count      int
		countSlice []int //分开统计每个partition扫描后的累积和
	)

	countSlice = make([]int, N)

	var wg sync.WaitGroup
	//var lock sync.RWMutex
	wg.Add(N)

	for i := 0; i < N; i++ {
		var newMap map[int]int
		newMap = make(map[int]int)

		for k, v := range dict {
			newMap[k] = v
		}

		go func(idx int) {
			defer wg.Done()
			//lock.Lock()
			for t1K, t1V := range partition[idx] {
				if t2V, ok := newMap[t1K]; ok {
					//lock.Lock()
					countSlice[idx] = countSlice[idx] + t1V*t2V
					//lock.Unlock()
				}
			}
			//lock.Unlock()
		}(i)
	}

	wg.Wait()

	//累积
	for i := 0; i < N; i++ {
		count += countSlice[i]
	}

	fmt.Println(count)

	fmt.Println(time.Since(t1))
}

func ScanDict(d1, d2 map[int]int) (int) {
	for t1K, t1V := range d1 {
		if t2V, ok := d2[t1K]; ok {
			return t2V * t1V
		}
	}
	return 0
}

func TestSolution5(t *testing.T) {
	t1 := time.Now()

	//分割map1
	var partition []map[int]int
	partition = make([]map[int]int, N)
	for i := 0; i < N; i++ {
		partition[i] = make(map[int]int)
	}

	f, err := os.Open("t1.csv")
	if err != nil {
		panic(err)
	}

	ret := csv.NewReader(bufio.NewReader(f))

	for {
		record, err := ret.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			panic(err)
		}

		//产生随机数，分配到partition中不同的map
		r_idx := rand.Intn(N)

		idx, _ := strconv.Atoi(record[0])

		partition[r_idx][idx]++
	}

	var dict map[int]int

	dict = make(map[int]int)

	f2, err := os.Open("t2.csv")
	if err != nil {
		panic(err)
	}

	ret2 := csv.NewReader(bufio.NewReader(f2))

	for {
		//按行读取
		record, err := ret2.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			panic(err)
		}

		idx, _ := strconv.Atoi(record[0])

		//统计第一列值，出现的次数
		dict[idx]++
	}

	var (
		count int
		c     chan int
	)

	c = make(chan int)

	var wg sync.WaitGroup
	wg.Add(N)

	for i := 0; i < N; i++ {
		go func(idx int) {
			var ret int
			defer wg.Done()
			//查找dict中是否有相同值，记录在ret
			for t1K, t1V := range partition[idx] {
				if t2V, ok := dict[t1K]; ok {
					//lock.Lock()
					//count = count + t1V*t2V
					ret = t1V * t2V
					//lock.Unlock()
				}
			}
			c <- ret
		}(i)
	}

	//closer
	go func() {
		wg.Wait()
		close(c)
	}()

	//累积
	for e := range c {
		fmt.Println(e)
		count += e
	}

	fmt.Println(count)
	fmt.Println(time.Since(t1))
}
