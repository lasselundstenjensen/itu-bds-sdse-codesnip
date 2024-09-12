package main

import (
	"fmt"
	"time"

	"itu.dk/course/coursetypes"
)

func main() {
	// Instantiate/create a new instructor
	lasse := coursetypes.Instructor{
		FirstName: "Lasse",
		LastName:  "Jensen"}

	// Instantiate/create a new instructor
	paolo := coursetypes.NewInstructor("Paolo", "Tell")
	fmt.Println(paolo)

	// One way to instantiate a new empty Course, then set the properties one by one afterwards
	course := coursetypes.Course{}
	course.Name = "Go course"
	course.Instructor = lasse
	course.Date = time.Now()

	fmt.Println(lasse)
	fmt.Println(course)

	// One way to instantiate a new Workshop "object" in one go
	workshop := coursetypes.Worshop{
		Course:     coursetypes.Course{Name: "Workshop on Go", Date: time.Now()},
		SignupDate: time.Now(),
		Name:       "Workshop",
	}

	fmt.Println(workshop)
	// Accessing different levels of the embedding
	fmt.Println(workshop.Name)
	fmt.Println(workshop.Course.Name)

	fmt.Println()

	// These two calls should have the same output of "Hello()"
	lasse.Hello()
	course.Instructor.Hello()

	// If we change the name..
	lasse.FirstName = "Paolo"
	// ..what will the output of this be?
	course.Instructor.Hello()
	// Note: It will be "Lasse", because we are not setting the value on
	// the "course.Instructor", but on instance of instructor named "lasse"

	// Using interface "Signable" to put two different structs in the same list
	var courses [2]coursetypes.Signable
	courses[0] = course
	courses[1] = workshop
	// and loop over them
	for _, c := range courses {
		fmt.Println(c)
	}
}
