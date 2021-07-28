package util

import (
	"polymorphism/model"
)

var skills = []model.Knowledge{model.NewBackend(), model.NewFrontend(), model.NewGeneric()}

func CheckCandidateProfile(c *model.Candidate) {
	for _, v := range skills {
		v.Evaluation(c)
	}
}
