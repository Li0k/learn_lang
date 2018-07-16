package fail

import (
	"fmt"
	"io"
	"strconv"
	"os"
	"sync"
	"encoding/csv"
	"strings"
	"testing"
	"time"
)

type chunk struct {
	bufsize int
	offset  int64
}

func TestSolution2(t testing.T) {
	t1 := time.Now()

	var (
		count int
	)

	dict1 := FileToDictGoroutine("t1.csv")
	dict2 := FileToDictGoroutine("t2.csv")

	for t1_k, t1_v := range dict1 {
		if t2_v, ok := dict2[t1_k]; ok {
			count = count + t1_v*t2_v
		}
	}

	fmt.Println(count)
	fmt.Println(time.Since(t1))
}

func FileToDictGoroutine(f string) (map[int]int) {
	var dict map[int]int
	dict = make(map[int]int)

	const BufferSize = 4096
	file, err := os.Open("t1.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fileinfo, err := file.Stat()
	if err != nil {
		panic(err)
	}

	filesize := int(fileinfo.Size())
	// Number of go routines we need to spawn.
	concurrency := filesize / BufferSize
	// buffer sizes that each of the go routine below should use. ReadAt
	// returns an error if the buffer size is larger than the bytes returned
	// from the file.
	chunksizes := make([]chunk, concurrency)

	// All buffer sizes are the same in the normal case. Offsets depend on the
	// index. Second go routine should start at 100, for example, given our
	// buffer size of 100.
	for i := 0; i < concurrency; i++ {
		chunksizes[i].bufsize = BufferSize
		chunksizes[i].offset = int64(BufferSize * i)
	}

	// check for any left over bytes. Add the residual number of bytes as the
	// the last chunk size.
	if remainder := filesize % BufferSize; remainder != 0 {
		c := chunk{bufsize: remainder, offset: int64(concurrency * BufferSize)}
		concurrency++
		chunksizes = append(chunksizes, c)
	}

	var wg sync.WaitGroup
	wg.Add(concurrency)

	var lock sync.RWMutex

	for i := 0; i < concurrency; i++ {
		go func(chunksizes []chunk, i int) {
			defer wg.Done()

			chunk := chunksizes[i]
			buffer := make([]byte, chunk.bufsize)
			bytesread, err := file.ReadAt(buffer, chunk.offset)

			if err != nil {
				fmt.Println(err)
				return
			}

			r := csv.NewReader(strings.NewReader(string(buffer[:bytesread])))
			//r.Comma = '\t'
			r.FieldsPerRecord = -1

			lock.Lock()
			for {
				record, err := r.Read()

				if err == io.EOF {
					break
				}

				if err != nil {
					panic(err)
				}
				//lock.Lock()
				idx, _ := strconv.Atoi(record[0])
				//		//fmt.Println(idx)
				dict[idx]++
			}
			lock.Unlock()
		}(chunksizes, i)
	}

	wg.Wait()

	return dict
}
