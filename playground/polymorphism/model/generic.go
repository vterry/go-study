package model

import (
	"fmt"
)

type Generic struct{}

func (g Generic) SendEmail(email string) {
	fmt.Println("Sorry but were not in that time!!!")
}

func (g *Generic) Evaluation(c *Candidate) {
	if len(c.GetKnowledge()) == 0 {
		g.SendEmail(c.GetEmail())
	}
}

func NewGeneric() *Generic {
	return &Generic{}
}
