package main

import (
	"io"
	"io/ioutil"
	"fmt"
	"encoding/csv"
	"strings"
	"strconv"
)

type questions []string
type answers []int

func main() {
	pertanyaan, kunciJawaban := readCSV()

	var jawabanUser answers

	for i := range pertanyaan {
		var ans int
		fmt.Printf("%v = ", pertanyaan[i])
		_, err := fmt.Scan(&ans)
		if err != nil {
			panic(err)
		}
		jawabanUser = append(jawabanUser, ans)
	}

	fmt.Printf("Anda berhasil menjawab %v dari %v pertanyaan", getScore(jawabanUser, kunciJawaban), len(kunciJawaban))
}

func getScore(jawabanUser answers, kunciJawaban answers) int {
	wrongAns := 0
	for i := range jawabanUser {
		if jawabanUser[i] != kunciJawaban[i] {
			wrongAns++
		}
	}

	return len(jawabanUser) - wrongAns
}

func readCSV() (questions, answers) {
	byteSlice, err := ioutil.ReadFile("./problems.csv")
	if err != nil {
		panic(err)
	}
	pString := string(byteSlice)

	r := csv.NewReader(strings.NewReader(pString))

	var q questions
	var kj answers

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		q = append(q, record[0])
		ans, err := strconv.Atoi(record[1])
		if err != nil {
			panic(err)
		}
		kj = append(kj, ans)
	}
	return q, kj
}