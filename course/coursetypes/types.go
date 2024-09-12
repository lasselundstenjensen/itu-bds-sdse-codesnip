package coursetypes

import (
	"fmt"
	"time"
)

// Declare interface "Signable": this is the contract to fulfil, which means
// a struct that wants to implement this interface must attach a function with this
// signature (i.e. create and attach a method to the struct)
type Signable interface {
	SignUp() string
}

type Instructor struct {
	FirstName string
	LastName  string
}

// Implement the "Stringer" interface to modify how 'fmt.Println()' prints the "Instructor"
// structure to console
func (i Instructor) Hello() {
	fmt.Printf("Hello, I'm %s %s\n", i.FirstName, i.LastName)
}

// Factory design pattern to construct a new instance of "Instructor"
func NewInstructor(firstname string, lastname string) Instructor {
	return Instructor{FirstName: firstname, LastName: lastname}
}

type Course struct {
	Name       string
	Instructor Instructor
	Date       time.Time
}

// Implement the "Signable" interface for struct "Course"
func (c Course) SignUp() string {
	return fmt.Sprintf("Signed up for course, %s", c.Name)
}

type Worshop struct {
	Course     // Embed struct "Course" into struct "Workshop", i.e. struct embedding
	Name       string
	SignupDate time.Time
}

// Implement the "Signable" interface for struct "Workshop"
func (w Worshop) SignUp() string {
	return fmt.Sprintf("Signed up for workshop, %s", w.Name)
}

// Implement the "Stringer" interface for struct "Course"
func (c Course) String() string {
	return fmt.Sprintf("Course name: %s", c.Name)
}
