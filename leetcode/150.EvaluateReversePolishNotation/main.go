package main

import (
	"log"
	"strconv"
)

func evalRPN(tokens []string) int {
	s := Stack{}

	for i := 0; i < len(tokens); i++ {
		if tokens[i] == "+" {
			s.Push(s.Pop() + s.Pop())

			continue
		}
		if tokens[i] == "*" {
			s.Push(s.Pop() * s.Pop())

			continue
		}
		if tokens[i] == "/" {
			b := s.Pop()
			a := s.Pop()
			s.Push(a / b)

			continue
		}
		if tokens[i] == "-" {
			b := s.Pop()
			a := s.Pop()
			s.Push(a - b)

			continue
		}
		integer, err := strconv.Atoi(tokens[i])
		if err != nil {
			log.Fatalln(err)
			return -1
		}
		s.Push(integer)
	}

	return s.Pop()
}

type Stack []int

func (s *Stack) Len() int {
	return len(*s)
}

func (s *Stack) Pop() int {
	if s.Len() == 0 {
		return -1
	}
	top := (*s)[s.Len()-1]
	*s = (*s)[:s.Len()-1]
	return top
}

func (s *Stack) Push(value int) {
	*s = append(*s, value)
}
