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
var start time.Time

func startQuiz(p questions, c chan answers, cTime <-chan time.Time) {
	jawabanUser := make(answers, len(p))

	for i := range p {
		select {
		case <- cTime :
			fmt.Println("Kriiingg, waktu habis!", time.Since(start))
			c <- jawabanUser
			break

		default:
			var ans int
			fmt.Printf("%v = ", p[i])
			_, err := fmt.Scan(&ans)
			if err != nil {
				panic(err)
			}
			jawabanUser[i] = ans
		}
	}
	c <- jawabanUser
}

func main() {
	duration := 3 * time.Second
	start = time.Now()
	timer := time.NewTimer(duration)

	pertanyaan, kunciJawaban := readCSV()
	jawabanUser := make(chan answers)

	fmt.Printf("Quiz dimulai! waktu anda %v\n", duration)

	go startQuiz(pertanyaan, jawabanUser, timer.C)

	ju := <- jawabanUser
	fmt.Printf("Anda berhasil menjawab %v dari %v pertanyaan", getScore(ju, kunciJawaban), len(kunciJawaban))
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