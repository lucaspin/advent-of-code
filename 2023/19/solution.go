package pkg202319

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type System struct {
	Workflows []Workflow
	Parts     []Part
}

func (s *System) AddWorkflow(w Workflow) {
	s.Workflows = append(s.Workflows, w)
}

func (s *System) AddPart(p Part) {
	s.Parts = append(s.Parts, p)
}

func (s *System) FindWorkflowByName(workflowName string) Workflow {
	for _, w := range s.Workflows {
		if w.Name == workflowName {
			return w
		}
	}

	panic(fmt.Errorf("workflow %s not found", workflowName))
}

func (s *System) IsAccepted(part Part) bool {
	current := "in"

	for current != "R" && current != "A" {
		w := s.FindWorkflowByName(current)
		current = w.Apply(part)
	}

	return current == "A"
}

func (s *System) Accepted() []Part {
	accepted := []Part{}

	for _, p := range s.Parts {
		if s.IsAccepted(p) {
			accepted = append(accepted, p)
		}
	}

	return accepted
}

func (s *System) SumAccepted() int {
	t := 0
	for _, p := range s.Accepted() {
		t += p.X + p.M + p.A + p.S
	}

	return t
}

type Workflow struct {
	Name  string
	Rules []WorkflowRule
}

func (w *Workflow) Apply(part Part) string {
	for _, r := range w.Rules {
		if r.Apply(part) {
			return r.Target
		}
	}

	panic("should not happen")
}

func (w *Workflow) AddRule(r WorkflowRule) {
	w.Rules = append(w.Rules, r)
}

type WorkflowRule struct {
	Variable  string
	Condition string
	Value     int
	Target    string
}

func (r *WorkflowRule) Apply(part Part) bool {
	switch r.Condition {
	case ">":
		return part.Get(r.Variable) > r.Value
	case "<":
		return part.Get(r.Variable) < r.Value
	}

	return true
}

type Part struct {
	X int
	M int
	A int
	S int
}

func (p *Part) Get(name string) int {
	switch name {
	case "x":
		return p.X
	case "m":
		return p.M
	case "a":
		return p.A
	case "s":
		return p.S
	}

	panic(fmt.Errorf("variable %s does not exist", name))
}

func A(input string) int {
	d, err := os.ReadFile(input)
	if err != nil {
		panic(err)
	}

	system := parseSystem(string(d))
	return system.SumAccepted()
}

func parseSystem(input string) *System {
	s := &System{Workflows: []Workflow{}, Parts: []Part{}}
	workflowsAndParts := strings.Split(input, "\n\n")

	for _, w := range strings.Split(workflowsAndParts[0], "\n") {
		s.AddWorkflow(parseWorkflow(w))
	}

	for _, p := range strings.Split(workflowsAndParts[1], "\n") {
		s.AddPart(parsePart(p))
	}

	return s
}

func parseWorkflow(line string) Workflow {
	workflow := Workflow{Rules: []WorkflowRule{}}
	nameAndRules := strings.Split(line, "{")
	workflow.Name = nameAndRules[0]

	rules := strings.ReplaceAll(nameAndRules[1], "}", "")
	for _, r := range strings.Split(rules, ",") {
		workflow.AddRule(parseRule(r))
	}

	return workflow
}

func parseRule(r string) WorkflowRule {
	if !strings.Contains(r, ":") {
		return WorkflowRule{Target: r}
	}

	wr := WorkflowRule{}
	condAndTarget := strings.Split(r, ":")
	wr.Target = condAndTarget[1]

	condition := condAndTarget[0]

	if strings.Contains(condition, ">") {
		varAndValue := strings.Split(condition, ">")
		wr.Condition = ">"
		wr.Variable = varAndValue[0]
		v, _ := strconv.Atoi(varAndValue[1])
		wr.Value = v
		return wr
	}

	if strings.Contains(condition, "<") {
		varAndValue := strings.Split(condition, "<")
		wr.Condition = "<"
		wr.Variable = varAndValue[0]
		v, _ := strconv.Atoi(varAndValue[1])
		wr.Value = v
		return wr
	}

	panic("should not happen")
}

func parsePart(line string) Part {
	part := Part{}
	line = strings.ReplaceAll(line, "{", "")
	line = strings.ReplaceAll(line, "}", "")
	for _, p := range strings.Split(line, ",") {
		nameAndValue := strings.Split(p, "=")
		value, _ := strconv.Atoi(nameAndValue[1])
		switch nameAndValue[0] {
		case "x":
			part.X = value
		case "m":
			part.M = value
		case "a":
			part.A = value
		case "s":
			part.S = value
		default:
			panic("nope")
		}
	}

	return part
}
