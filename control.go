package play

import (
	"fmt"
)

type Person struct {
	name string
	age  int
}

func NewPerson(n string, a int) *Person {
	p := new(Person)
	p.name = n
	p.age = a
	return p
}

func CheckAge(p *Person, threshold int) string {
	if p.age > threshold {
		return "Old"
	} else {
		return "Young"
	}
}

func RunControl(name string, age int, threshold int) string {
	person := NewPerson(name, age)
	return CheckAge(person, threshold)
}

func RunControl2(name string, age int, numOfPeople int) string {
	people := make([]*Person, numOfPeople)
	for i := 0; i < numOfPeople; i++ {
		people[i] = NewPerson(fmt.Sprintf("%s_%d", name, i), age+i)
	}

	sum := 0
	for i := 0; i < numOfPeople; i++ {
		sum += people[i].age
	}

	average := sum / numOfPeople
	return fmt.Sprintf("%d", average)
}
