package main

import (
	"fmt"
	"math/rand"
	"time"
)

func CreateNumber() string {
	rand.Seed(time.Now().Unix())
	arr := make([]int, 0)
	n := 0
	for i := 0; i < 4; {
		n = rand.Intn(10)
		add := true
		for _, v := range arr {
			if v == n {
				add = false
				break
			}
		}
		if add {
			arr = append(arr, n)
			i++
		}
	}
	s := fmt.Sprintf("%v%v%v%v", arr[0], arr[1], arr[2], arr[3])
	return s
}
func CheckNumber(answer, input string) (int, int, string) {
	cow := 0
	bull := 0
	if len(input) != 4 {
		return 0, 0, "требуется ввести 4х значное число!\n"
	}
	if answer == input {
		return 0, 4, ""
	}
	list := make([]rune, 0)
	for _, v := range input {
		list = append(list, v)
	}
	for outList, outValue := range list {
		for inList, inValue := range list {
			if outList == inList {
				continue
			}
			if outValue == inValue {
				return 0, 0, "Числа не должны повторяться!\n"
			}
		}
	}
	for i, inputValue := range input {
		if input[i] == answer[i] {
			bull++
		} else {
			for _, answerValue := range answer {
				if inputValue == answerValue {
					cow++
				}
			}
		}
	}
	return cow, bull, ""
}
func NewBullAndCow() Game {
	g := Game{number: CreateNumber(), state: "none", count: 1}
	return g
}
