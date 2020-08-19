package main

// Student struct
type Student struct {
	firstName, lastName string
	grade               string
	age                 int
	totalMark           int
}

func (s *Student) getName() string {
	return s.firstName + s.lastName
}

func (s *Student) setName(fname, lname string) {
	s.firstName = fname
	s.lastName = lname
}

func (s *Student) getGrade() string {
	return s.grade
}

func (s *Student) setGrade(g string) {
	s.grade = g
}

func (s *Student) getAge() int {
	return s.age
}

func (s *Student) setAge(a int) {
	s.age = a
}

func (s *Student) getTotalMark() int {
	return s.totalMark
}

func (s *Student) setTotalMark(mark int) {
	s.totalMark = mark
}
