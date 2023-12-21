package pkg202320

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type ModuleRegistry struct {
	Modules map[string]*Module
}

func (r *ModuleRegistry) AddModule(m *Module) {
	r.Modules[m.Name] = m
}

func (r *ModuleRegistry) PushButton() {
	r.Broadcast()
}

type queue struct {
	items []queueItem
}

func (q *queue) Push(item queueItem) {
	q.items = append(q.items, item)
}

func (q *queue) Pop() queueItem {
	head := q.items[0]
	q.items = q.items[1:]
	return head
}

func (q *queue) Len() int {
	return len(q.items)
}

type queueItem struct {
	To    string
	From  string
	Pulse PulseType
}

func (r *ModuleRegistry) Send(q queue, cycles map[string]int64, module string, presses int64) (int64, int64) {
	lowPulses := int64(0)
	highPulses := int64(0)

	for q.Len() > 0 {
		item := q.Pop()
		if item.Pulse == PulseTypeHigh {
			highPulses++
		} else {
			lowPulses++
		}

		if cycles != nil {
			if item.To == module && item.Pulse == PulseTypeHigh {
				cycles[item.From] = presses
			}
		}

		m := r.Modules[item.To]
		if m == nil {
			continue
		}

		out := m.Send(item.Pulse, item.From)
		if out != PulseTypeIgnore {
			for _, d := range m.Destinations {
				q.Push(queueItem{To: d, From: item.To, Pulse: out})
			}
		}
	}

	return lowPulses, highPulses
}

func (r *ModuleRegistry) Broadcast() (int64, int64) {
	queue := queue{items: []queueItem{}}
	queue.Push(queueItem{To: "broadcaster", From: "", Pulse: PulseTypeLow})
	return r.Send(queue, nil, "", int64(0))
}

func (r *ModuleRegistry) BroadcastAndTrack(cycles map[string]int64, module string, presses int64) (int64, int64) {
	queue := queue{items: []queueItem{}}
	queue.Push(queueItem{To: "broadcaster", From: "", Pulse: PulseTypeLow})
	return r.Send(queue, cycles, module, presses)
}

func in(list []string, item string) bool {
	for _, i := range list {
		if i == item {
			return true
		}
	}

	return false
}

func (r *ModuleRegistry) FindInputs(moduleName string) []string {
	inputs := []string{}
	for _, m := range r.Modules {
		if in(m.Destinations, moduleName) {
			inputs = append(inputs, m.Name)
		}
	}

	return inputs
}

func (r *ModuleRegistry) LoadMemories() {
	for k, m := range r.Modules {
		if m.Type == ModuleTypeConjunction {
			for _, i := range r.FindInputs(k) {
				m.Memory[i] = PulseTypeLow
			}
		}
	}
}

type ModuleType int

const (
	ModuleTypeFlipFlop ModuleType = iota
	ModuleTypeConjunction
	ModuleTypeBroadcaster
)

type ModuleState int

const (
	ModuleStateOff ModuleState = iota
	ModuleStateOn
)

type PulseType int

const (
	PulseTypeLow PulseType = iota
	PulseTypeHigh
	PulseTypeIgnore
)

func (p *PulseType) String() string {
	switch *p {
	case PulseTypeLow:
		return "low"
	case PulseTypeHigh:
		return "high"
	case PulseTypeIgnore:
		return "ignore"
	}

	return "???"
}

type Module struct {
	Name         string
	Type         ModuleType
	State        ModuleState
	Memory       map[string]PulseType
	Destinations []string
}

func (m *Module) Send(pulse PulseType, from string) PulseType {
	switch m.Type {
	case ModuleTypeBroadcaster:
		return pulse
	case ModuleTypeFlipFlop:
		if pulse == PulseTypeHigh {
			return PulseTypeIgnore
		}

		if m.ToggleState() == ModuleStateOn {
			return PulseTypeHigh
		}

		return PulseTypeLow
	case ModuleTypeConjunction:
		m.UpdateMemory(from, pulse)
		if m.AllHigh() {
			return PulseTypeLow
		}

		return PulseTypeHigh
	}

	panic("should not happen")
}

func (m *Module) UpdateMemory(from string, pulse PulseType) {
	m.Memory[from] = pulse
}

func (m *Module) AllHigh() bool {
	for _, p := range m.Memory {
		if p == PulseTypeLow {
			return false
		}
	}

	return true
}

func (m *Module) ToggleState() ModuleState {
	if m.State == ModuleStateOn {
		m.State = ModuleStateOff
	} else {
		m.State = ModuleStateOn
	}

	return m.State
}

func anyValueSatisfies(m map[string]int64, cond func(v int64) bool) bool {
	for _, v := range m {
		if cond(v) {
			return true
		}
	}

	return false
}

func (r *ModuleRegistry) FindCycles(module string) map[string]int64 {
	inputs := r.FindInputs(module)

	cycles := map[string]int64{}
	for _, input := range inputs {
		cycles[input] = 0
	}

	fmt.Printf("Inputs: %v\n", inputs)

	i := int64(0)
	for anyValueSatisfies(cycles, func(v int64) bool { return v == 0 }) {
		i++
		r.BroadcastAndTrack(cycles, module, i)
	}

	return cycles
}

func A(input string) int64 {
	f, err := os.Open(input)
	if err != nil {
		panic(err)
	}

	registry := parse(f)

	var low, high int64
	for i := 0; i < 1000; i++ {
		l, h := registry.Broadcast()
		fmt.Printf("low=%d high=%d\n", l, high)
		low += l
		high += h
	}

	fmt.Printf("low=%d, high=%d\n", low, high)
	return low * high
}

func B(input string) int64 {
	f, err := os.Open(input)
	if err != nil {
		panic(err)
	}

	registry := parse(f)

	// From the input, only one conjunction module goes into rx.
	// Into that conjunction module, only conjunction modules are served as inputs as well.
	// So we need to find the cycle lengths of all the inputs going into the module going into rx.
	inputs := registry.FindInputs("rx")
	if len(inputs) != 1 {
		panic("more than one input to rx")
	}

	cycles := registry.FindCycles(inputs[0])
	fmt.Printf("Cycles: %v\n", cycles)

	values := []int64{}
	for _, v := range cycles {
		values = append(values, v)
	}

	return lcm(values[0], values[1], values[2:]...)
}

func gcd(a, b int64) int64 {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}

	return a
}

func lcm(a, b int64, integers ...int64) int64 {
	result := a * b / gcd(a, b)
	for i := 0; i < len(integers); i++ {
		result = lcm(result, integers[i])
	}

	return result
}

func parse(f *os.File) *ModuleRegistry {
	registry := ModuleRegistry{Modules: map[string]*Module{}}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		registry.AddModule(parseModule(scanner.Text()))
	}

	registry.LoadMemories()
	return &registry
}

func parseModule(line string) *Module {
	parts := strings.Split(line, " -> ")
	return &Module{
		Name:         parseName(parts[0]),
		Type:         parseType(parts[0]),
		Destinations: strings.Split(parts[1], ", "),
		Memory:       map[string]PulseType{},
	}
}

func parseName(s string) string {
	switch s[0] {
	case '&', '%':
		return s[1:]
	}

	return s
}

func parseType(s string) ModuleType {
	switch s[0] {
	case '&':
		return ModuleTypeConjunction
	case '%':
		return ModuleTypeFlipFlop
	}

	return ModuleTypeBroadcaster
}
