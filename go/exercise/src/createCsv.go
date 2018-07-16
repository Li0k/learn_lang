package main

import (
	"os"
	"encoding/csv"
	"math/rand"
	"time"
	"strconv"
)

func main() {
	CreateCsv("src/t1.csv")
	CreateCsv("src/t2.csv")
}

func CreateCsv(file string)  {
	f,err := os.Create(file)
	if err != nil {
		panic(err)
	}

	defer f.Close();

	//header := []string{"a","b"}
	w := csv.NewWriter(f)
	//w.Write(header)

	var data [][]string
	rand.Seed(time.Now().Unix())

	for i := 0;i < 5000000 ;i++  {
		a := rand.Intn(1000000)
		b := rand.Intn(1000000 - 1)

		tmp := []string{strconv.Itoa(a),strconv.Itoa(b)}
		data = append(data,tmp)
	}

	w.WriteAll(data)
}
