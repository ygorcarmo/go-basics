package data

type Instructor struct {
	Id        int
	firstName string
	lastName  string
	score     int
}

func NewInstructor(first string, last string) Instructor {
	return Instructor{firstName: first, lastName: last}
}
