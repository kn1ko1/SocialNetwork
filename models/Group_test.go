package models

import (
	"fmt"
	"math/rand"
	"testing"
)

var (
	sutTitle = []string{"Family Time", "Our Group", "Cinema Trip"}
)

func TestGroupValidateInvalidFieldExpectError(t *testing.T) {
	var groups []*Group
	for i := 0; i < tableRunCount; i++ {
		groups = append(groups, generateInvalidGroup())
	}
	for _, g := range groups {
		name := fmt.Sprintf("%+v", *g)
		t.Run(name, func(t *testing.T) {
			err := g.Validate()
			if err == nil {
				t.Error("expect error for invalid group field")
			}
		})
	}
}

func TestGroupValidateMissingFieldExpectError(t *testing.T) {
	var groups []*Group
	for i := 0; i < tableRunCount; i++ {
		groups = append(groups, generateMissingFieldGroup())
	}
	for _, g := range groups {
		name := fmt.Sprintf("%+v", *g)
		t.Run(name, func(t *testing.T) {
			err := g.Validate()
			if err == nil {
				t.Error("expect error for missing group field")
			}
		})
	}
}

func TestGroupValidateValidExpectNil(t *testing.T) {
	var groups []*Group
	for i := 0; i < tableRunCount; i++ {
		groups = append(groups, generateValidGroup())
	}
	for _, g := range groups {
		name := fmt.Sprintf("%+v", *g)
		t.Run(name, func(t *testing.T) {
			err := g.Validate()
			if err != nil {
				t.Error("expect nil for valid group")
			}
		})
	}
}

func generateValidGroup() *Group {
	ctime := rand.Int63n(1000) + 1
	idxDescriptions := rand.Intn(len(sutDescriptions))
	idxTitle := rand.Intn(len(sutTitle))
	g := &Group{
		CreatedAt:   ctime,
		CreatorID:   rand.Intn(1000) + 1,
		Description: sutDescriptions[idxDescriptions],
		Title:       sutTitle[idxTitle],
		UpdatedAt:   ctime,
	}
	return g
}

func generateMissingFieldGroup() *Group {
	g := generateValidGroup()
	missingField := rand.Intn(5)
	switch missingField {
	case 0:
		g.CreatedAt = 0
	case 1:
		g.CreatorID = 0
	case 2:
		g.Description = ""
	case 3:
		g.Title = ""
	case 4:
		g.UpdatedAt = 0
	}
	return g
}

func generateInvalidGroup() *Group {
	g := generateValidGroup()
	invalidField := rand.Intn(5)
	switch invalidField {
	case 0:
		g.CreatedAt = -g.CreatedAt
	case 1:
		g.CreatorID = -g.CreatorID
	case 2:
		g.Description = ""
	case 3:
		g.Title = ""
	case 4:
		g.UpdatedAt = -g.UpdatedAt
	}
	return g
}
