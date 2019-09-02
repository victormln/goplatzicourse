package main

import "fmt"

// Course struct
type Course struct {
	Name   string
	Slug   string
	Skills []string
}

// Subscribe method
func (p Course) Subscribe(name string) {
	fmt.Printf("The person %s has subscribed to %s course\n", name, p.Name)
}

// Career struct
type Career struct {
	Course
}

// Subscribe method
func (p Career) Subscribe(name string) {
	fmt.Printf("The person %s has subscribed to %s career\n", name, p.Name)
}

// Organization interface
type Organization interface {
	Subscribe(name string)
}

func callSubscribe(p Organization) {
	p.Subscribe("Victor")
}

// InterfaceTest method
func InterfaceTest() {
	myCourse := Course{Name: "Go", Slug: "go", Skills: []string{"1", "2"}}
	myCareer := new(Career)
	myCareer.Name = "Go 2"
	myCareer.Slug = "go2"
	myCareer.Skills = []string{"3", "4"}
	callSubscribe(myCourse)
	callSubscribe(myCareer)
}

func deferTest() {
	fmt.Println("The function InterfaceTest ended")
}

func main() {
	defer deferTest()
	InterfaceTest()
}
