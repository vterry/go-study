package model

type Candidate struct {
	name      string
	skills    map[string]int
	knowledge []Knowledge
	email     string
}

func (c *Candidate) SetName(n string) {
	c.name = n
}

func (c *Candidate) SetSkills(m map[string]int) {
	c.skills = m
}

func (c *Candidate) SetEmail(email string) {
	c.email = email
}

func (c *Candidate) SetKnowledge(p Knowledge) {
	c.knowledge = append(c.knowledge, p)
}

func (c *Candidate) GetKnowledge() []Knowledge {
	return c.knowledge
}

func (c *Candidate) GetEmail() string {
	return c.email
}
