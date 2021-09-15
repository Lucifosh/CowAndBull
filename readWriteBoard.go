package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func readBoard() []Board {
	file, err := os.Open("board.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	ArrBoard := make([]Board, 0)

	ByteText, err := ioutil.ReadFile("board.txt")
	if err != nil {
		fmt.Println(err)
	}
	text := string(ByteText)
	fmt.Printf("\n\n***\nчтение файла: %v\n***\n\n", text)
	if len(text) == 0 {
		return ArrBoard
	}
	data := strings.Split(text, " ")

	length := len(data)
	board := Board{}
	for i := 0; i < length; i++ {
		n, err := strconv.Atoi(data[i])
		if err != nil {
			fmt.Println(err)
		}
		if i%2 == 0 || i == 0 {
			board.ID = n
		} else {
			board.score = n
			ArrBoard = append(ArrBoard, board)
		}
	}
	return ArrBoard
}
func swap(b []Board, idx1, idx2 int) []Board {
	nb := b
	temp := nb[idx1]
	nb[idx1] = nb[idx2]
	nb[idx2] = temp
	return nb
}
func soartBoard(b []Board) []Board {
	nb := b
	for {
		flag := false
		for i := 1; i < len(nb); i++ {
			if nb[i].score > nb[i-1].score {
				nb = swap(nb, i, i-1)
				break
			}
			flag = true
		}
		if flag {
			break
		}
	}
	return nb
}
func WriteBoard(id, score int) {
	b := readBoard()
	fmt.Printf("\n\n***\n%v\n***\n\n", b)
	idx := -1
	for i, v := range b {
		if v.ID == id {
			idx = i
		}
	}
	if idx == -1 {
		//add new user
		newBoard := Board{ID: id, score: score}
		b = append(b, newBoard)
	} else {
		//добавление очков к существующему
		b[idx].score += score
	}
	fmt.Println("после добавления очков ", b)
	b = soartBoard(b)
	fmt.Println("после сортирования ", b)

	os.Truncate("board.txt", 0)

	fmt.Println("файл пуст")
	file, err := os.Create("board.txt")
	if err != nil {
		fmt.Println(err)
	}
	length := len(b)
	for i := 0; i < length; i++ {
		s := fmt.Sprintf("%v %v ", b[i].ID, b[i].score)
		file.WriteString(s)
	}
	defer file.Close()
}
func SendBoard(id int) string {
	b := readBoard()
	s := ""
	for idx, v := range b {
		if v.ID == id {
			temp := fmt.Sprintf("%v.Вы      %v\n", idx+1, v.score)
			s += temp
		} else if idx < 10 {
			temp := fmt.Sprintf("%v.Игрок %v\n", idx+1, v.score)
			s += temp
		} else if idx == 10 {
			s += "...\n"
		}
	}
	return s
}
func getScoreFromID(id int) int {
	b := readBoard()
	for _, v := range b {
		if v.ID == id {
			return v.score
		}
	}
	return 0
}
