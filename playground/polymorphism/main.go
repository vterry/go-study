package main

import (
	"fmt"
	"polymorphism/model"
	"polymorphism/util"
)

func main() {
	c := new(model.Candidate)

	c.SetName("Expected BackDev")
	c.SetEmail("c@sample.com")
	c.SetSkills(map[string]int{
		"Java":   10,
		"Golang": 10,
	})

	c1 := new(model.Candidate)

	c1.SetName("Expected Front")
	c1.SetEmail("c1@sample.com")
	c1.SetSkills(map[string]int{
		"Javascript": 10,
		"Vue":        10,
	})

	c2 := new(model.Candidate)

	c2.SetName("Expected All")
	c2.SetEmail("full@sample.com")
	c2.SetSkills(map[string]int{
		"Java":       10,
		"Golang":     10,
		"Javascript": 10,
		"Vue":        10,
	})

	c3 := new(model.Candidate)

	c3.SetName("Expected Nothing")
	c3.SetEmail("nothing@sample.com")
	c3.SetSkills(map[string]int{
		"Java":   0,
		"Golang": 0,
	})

	util.CheckCandidateProfile(c)
	util.CheckCandidateProfile(c1)
	util.CheckCandidateProfile(c2)
	util.CheckCandidateProfile(c3)

	fmt.Printf("Candidate c: %+v\n", c)
	fmt.Printf("Candidate c1: %+v\n", c1)
	fmt.Printf("Candidate c2: %+v\n", c2)
	fmt.Printf("Candidate c3: %+v\n", c3)
}
