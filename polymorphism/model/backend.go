package model

import (
	"fmt"
)

type Backend struct {
	skills []string
}

func NewBackend() *Backend {
	return &Backend{[]string{"Java", "C#", "Golang"}}
}

func (b Backend) SendEmail(email string) {
	fmt.Println("You was evaluated as Backend Developer!!!")
}

func (b *Backend) Evaluation(c *Candidate) {
	skillLevel := 0.0

	for _, v := range b.skills {
		skillLevel += float64(c.skills[v])
	}

	if skillLevel > 7 {
		b.SendEmail(c.GetEmail())
		c.SetKnowledge(b)
	} else {
		return
	}
}
