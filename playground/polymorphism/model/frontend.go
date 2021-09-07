package model

import "fmt"

type Frontend struct {
	skills []string
}

func NewFrontend() *Frontend {
	return &Frontend{[]string{"Javascript", "Vue", "PHP"}}
}

func (f Frontend) SendEmail(email string) {
	fmt.Println("You were evaluated as Frontend Developer!!!")
}

func (f *Frontend) Evaluation(c *Candidate) {
	skillLevel := 0

	for _, v := range f.skills {
		skillLevel += c.skills[v]
	}

	if skillLevel > 7 {
		f.SendEmail(c.GetEmail())
		c.SetKnowledge(f)
	} else {
		return
	}
}
