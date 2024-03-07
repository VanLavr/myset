package main

import (
	"fmt"
)

type MyConstraint interface {
	int | int64 | int32 | float32 | float64 | string | Human | bool
}

type MySet[E MyConstraint] map[E]struct{}

type Human struct {
	Name string
}

func NewSet[E MyConstraint](args ...E) MySet[E] {
	newSet := MySet[E]{}
	for _, arg := range args {
		newSet[arg] = struct{}{}
	}
	return newSet
}

func (s MySet[E]) Add(values ...E) {
	for _, value := range values {
		s[value] = struct{}{}
	}
}

func (s MySet[E]) Contains(value E) bool {
	_, ok := s[value]
	return ok
}

func (s MySet[E]) Members() []E {
	var result []E
	for member, _ := range s {
		result = append(result, member)
	}
	return result
}

func (s MySet[E]) Delete(member E) {
	_, ok := s[member]
	if !ok {
		return
	}
	delete(s, member)
}

func main() {
	testWithints := NewSet(1, 3, 5)
	testWithbools := NewSet(true, false, false)
	testWithstructs := NewSet(Human{Name: "Joe"}, Human{Name: "Ivan"})

	testWithints.Add(0, 3, 12, 1)
	testWithbools.Add(true)
	testWithstructs.Add(Human{Name: "Vasya"}, Human{Name: "Ivan"})
	fmt.Println(testWithints.Contains(1))
	fmt.Println(testWithstructs.Members(), testWithbools.Members(), testWithints.Members())
	testWithints.Delete(0)
	fmt.Println(testWithints.Members())
	testWithstructs.Delete(Human{Name: "Vasya"})
	fmt.Println(testWithstructs.Members())
}
