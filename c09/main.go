package main

import (
	"encoding/json"
	"fmt"
)

type school struct {
	Name    string
	Classes map[int]*class
}

func newSchool(name string, classNum int) *school {
	return &school{
		Name:    name,
		Classes: make(map[int]*class, classNum),
	}
}

func (s *school) addClass(class *class) {
	s.Classes[class.ID] = class
}

func (s *school) addStudent(cID int, name string, age int, score score) bool {
	status, class := s.getClass(cID)
	if status {
		class.addStudent(name, age, score)
		return true
	}
	return false
}

func (s *school) getStudent(cID, ID int) (bool, *student) {
	status, class := s.getClass(cID)
	if status {
		student, ok := class.Students[ID]
		if ok {
			return true, student
		}
	}
	return false, nil
}

func (s *school) getClass(cID int) (bool, *class) {
	class, ok := s.Classes[cID]
	if ok {
		return true, class
	}
	return false, nil
}

type class struct {
	ID        int
	Name      string
	Students  map[int]*student
	studenter studenter
}

func newClass(id int, name string, stdNum int, studenter studenter) *class {
	return &class{
		ID:        id,
		Name:      name,
		Students:  make(map[int]*student, stdNum),
		studenter: studenter,
	}
}

func (c *class) addStudent(name string, age int, score score) {
	student := c.studenter(name, age, score)
	c.Students[student.ID] = student
}

type student struct {
	ID, Age int
	Name    string
	Score   score
}

type studenter func(name string, age int, score score) *student

func newStudenter() studenter {
	id := 0
	return func(name string, age int, score score) *student {
		id++
		return newStudent(id, name, age, score)
	}
}

func newStudent(id int, name string, age int, score score) *student {
	return &student{
		ID:    id,
		Age:   age,
		Name:  name,
		Score: score,
	}
}

type score struct {
	Chinese int
	Maths   int
	English int
}

func newScore(chinese int, maths int, english int) score {
	return score{
		Chinese: chinese,
		Maths:   maths,
		English: english,
	}
}

func altStudent(sch *school, cID int, ID int, arg interface{}) bool {
	status, student := sch.getStudent(cID, ID)
	if status {
		switch tmp := arg.(type) {
		case int:
			student.Age = tmp
		case string:
			student.Name = tmp
		case score:
			student.Score = tmp
		default:
			return false
		}
		return true
	}
	return false
}

func altClass(sch *school, cID int, name string) bool {
	status, class := sch.getClass(cID)
	if status {
		class.Name = name
		return true
	}
	return false
}

func initClass(cID int, name string, stdNum int) *class {
	ser := newStudenter()
	return newClass(cID, name, stdNum, ser)
}

func main() {
	sch := newSchool("hsgz", 13)
	sch.addClass(initClass(8, "gebb", 40))
	sch.addStudent(8, "TEST", 17, newScore(100, 100, 100))

	fmt.Println(sch.getClass(8))
	fmt.Println(sch.getStudent(8, 1))
	str, _ := json.Marshal(sch)
	fmt.Printf("%s", str)
}
