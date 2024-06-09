package utils

type Stack struct {
	data []string
}

// добавляет элемент в вершину
func (s *Stack) Push(item string) {
	s.data = append(s.data, item)
}

// удаляет и возвращает элемент из вершины
// Если стек пуст, возвращает пустую строку.
func (s *Stack) Pop() string {
	if len(s.data) == 0 {
		return ""
	}
	index := len(s.data) - 1
	item := s.data[index]
	s.data = s.data[:index]
	return item
}

// возвращает элемент из вершины стека без удаления

func (s *Stack) Top() string {
	if len(s.data) == 0 {
		return ""
	}
	return s.data[len(s.data)-1]
}

// возвращает true на пустой стек
func (s *Stack) IsEmpty() bool {
	return len(s.data) == 0
}
