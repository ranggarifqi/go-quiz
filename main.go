package main

import (
	"io"
	"io/ioutil"
	"fmt"
	"encoding/csv"
	"strings"
	"strconv"
	"time"
)

type questions []string
type answers []int

func startQuiz(p questions, c chan answers) {
	// var jawabanUser answers
	jawabanUser := make(answers, len(p))

	for i := range p {
		var ans int
		fmt.Printf("%v = ", p[i])
		_, err := fmt.Scan(&ans)
		if err != nil {
			panic(err)
		}
		jawabanUser[i] = ans
	}
	c <- jawabanUser
}


func main() {
	duration := 50 * time.Second

	pertanyaan, kunciJawaban := readCSV()
	jawabanUser := make(chan answers)

	go startQuiz(pertanyaan, jawabanUser)

	select {
	case m := <- jawabanUser:
		fmt.Printf("Anda berhasil menjawab %v dari %v pertanyaan", getScore(m, kunciJawaban), len(kunciJawaban))
	case <- time.After(duration):
		fmt.Printf("\nKriiingg, waktu habis")
	}
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