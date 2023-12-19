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

func copyRules(current []WorkflowRule, new *WorkflowRule, before ...WorkflowRule) []WorkflowRule {
	c := make([]WorkflowRule, len(current))
	copy(c, current)
	c = append(c, before...)

	if new != nil {
		c = append(c, *new)
	}

	return c
}

func (s *System) findAcceptedRules(allConditions *[][]WorkflowRule, currentCondition []WorkflowRule, workflow Workflow) {
	for ruleIndex, rule := range workflow.Rules {
		if rule.Target == "R" {
			continue
		}

		if rule.Target == "A" {
			if ruleIndex == 0 {
				*allConditions = append(
					*allConditions,
					copyRules(currentCondition, &rule),
				)
			} else {
				*allConditions = append(
					*allConditions,
					copyRules(currentCondition, &rule, negateRules(workflow.Rules[0:ruleIndex])...),
				)
			}
			continue
		}

		if ruleIndex == 0 {
			if rule.Condition == "" {
				s.findAcceptedRules(allConditions, currentCondition, s.FindWorkflowByName(rule.Target))
			} else {
				s.findAcceptedRules(
					allConditions,
					copyRules(currentCondition, &rule),
					s.FindWorkflowByName(rule.Target),
				)
			}
		} else {
			if rule.Condition == "" {
				s.findAcceptedRules(
					allConditions,
					copyRules(currentCondition, nil, negateRules(workflow.Rules[0:ruleIndex])...),
					s.FindWorkflowByName(rule.Target),
				)
			} else {
				s.findAcceptedRules(
					allConditions,
					copyRules(currentCondition, &rule, negateRules(workflow.Rules[0:ruleIndex])...),
					s.FindWorkflowByName(rule.Target),
				)
			}
		}
	}
}

func negateRules(rules []WorkflowRule) []WorkflowRule {
	negated := []WorkflowRule{}

	for _, r := range rules {
		if r.Condition != "" {
			negated = append(negated, r.Negate())
		}
	}

	return negated
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

func (r *WorkflowRule) Negate() WorkflowRule {
	switch r.Condition {
	case ">":
		return WorkflowRule{
			Variable:  r.Variable,
			Condition: "<",
			Value:     r.Value + 1,
		}
	case "<":
		return WorkflowRule{
			Variable:  r.Variable,
			Condition: ">",
			Value:     r.Value - 1,
		}
	}

	return *r
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

type PartRange struct {
	X Range
	M Range
	A Range
	S Range
}

func (pr *PartRange) Apply(rule WorkflowRule) {
	switch rule.Condition {
	case "<":
		pr.AdjustRight(rule)
	case ">":
		pr.AdjustLeft(rule)
	}
}

func (pr *PartRange) AdjustRight(rule WorkflowRule) {
	switch rule.Variable {
	case "s":
		pr.S.RightInclusive = rule.Value - 1
	case "x":
		pr.X.RightInclusive = rule.Value - 1
	case "a":
		pr.A.RightInclusive = rule.Value - 1
	case "m":
		pr.M.RightInclusive = rule.Value - 1
	}
}

func (pr *PartRange) AdjustLeft(rule WorkflowRule) {
	switch rule.Variable {
	case "s":
		pr.S.LeftInclusive = rule.Value + 1
	case "x":
		pr.X.LeftInclusive = rule.Value + 1
	case "a":
		pr.A.LeftInclusive = rule.Value + 1
	case "m":
		pr.M.LeftInclusive = rule.Value + 1
	}
}

func (pr *PartRange) Combinations() int64 {
	return pr.A.Combinations() * pr.X.Combinations() * pr.M.Combinations() * pr.S.Combinations()
}

type Range struct {
	LeftInclusive  int
	RightInclusive int
}

func (r *Range) Combinations() int64 {
	return int64(r.RightInclusive) - int64(r.LeftInclusive) + 1
}

func A(input string) int {
	d, err := os.ReadFile(input)
	if err != nil {
		panic(err)
	}

	system := parseSystem(string(d))
	return system.SumAccepted()
}

func B(input string) int64 {
	d, err := os.ReadFile(input)
	if err != nil {
		panic(err)
	}

	system := parseSystem(string(d))

	allConditions := [][]WorkflowRule{}
	system.findAcceptedRules(&allConditions, []WorkflowRule{}, system.FindWorkflowByName("in"))

	total := int64(0)
	for _, c := range allConditions {
		pr := &PartRange{
			X: Range{LeftInclusive: 1, RightInclusive: 4000},
			S: Range{LeftInclusive: 1, RightInclusive: 4000},
			M: Range{LeftInclusive: 1, RightInclusive: 4000},
			A: Range{LeftInclusive: 1, RightInclusive: 4000},
		}

		for _, r := range c {
			pr.Apply(r)
		}

		total += pr.Combinations()
	}

	return total
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
