package model

type Knowledge interface {
	SendEmail(s string)
	Evaluation(c *Candidate)
}
