package main

import (
	"AOC2023/utils"
	"fmt"
	"slices"
	"strings"
	"time"
)

type Pulse = int

type Module interface {
	isModule()
}

func compute(modules Modules, moduleID string, p Pulse, emitter string) (Pulse, Module) {
	switch module := modules[moduleID].(type) {
	case FlipFlop:
		return module.compute(p)
	case Conjunction:
		return module.compute(p, emitter)
	default:
		panic("Unknown module.")
	}
}

func getId(m Module) string {
	switch module := m.(type) {
	case FlipFlop:
		return module.ID
	case Conjunction:
		return module.ID
	default:
		panic("Unknown module.")
	}
}

type FlipFlop struct {
	ID    string
	state int
}

func (_ FlipFlop) isModule() {}

func (ff FlipFlop) copy() FlipFlop {
	return FlipFlop{ff.ID, ff.state}
}

func (ff FlipFlop) flip() FlipFlop {
	return FlipFlop{ff.ID, 1 - ff.state}
}

func (ff FlipFlop) isOn() bool {
	return ff.state == 1
}
func (ff FlipFlop) compute(p Pulse) (Pulse, Module) {
	if p == 1 {
		return -1, ff.copy()
	} else {
		ff = ff.flip()
		if ff.isOn() {
			return 1, ff.copy()
		} else {
			return 0, ff.copy()
		}
	}
}

func (ff FlipFlop) isInDefaultState() bool {
	return ff.state == 0
}

type Conjunction struct {
	ID            string
	ConnectionsIn map[string]Pulse
}

func (_ Conjunction) isModule() {}

func (c Conjunction) copy() Conjunction {
	return Conjunction{c.ID, c.ConnectionsIn}
}

func (c Conjunction) allInputsHP() bool {
	for emitter := range c.ConnectionsIn {
		if c.ConnectionsIn[emitter] == 0 {
			return false
		}
	}
	return true
}

func (c Conjunction) allInputsLP() bool {
	for emitter := range c.ConnectionsIn {
		if c.ConnectionsIn[emitter] == 1 {
			return false
		}
	}
	return true
}

func (c Conjunction) changeEmitterState(emitter string, p Pulse) Conjunction {
	var newMap = c.ConnectionsIn
	newMap[emitter] = p
	return Conjunction{c.ID, newMap}
}

func (c Conjunction) compute(p Pulse, emitter string) (Pulse, Module) {
	var newC = c.changeEmitterState(emitter, p)
	if newC.allInputsHP() {
		return 0, newC.copy()
	} else {
		return 1, newC.copy()
	}
}

func (c Conjunction) isInDefaultState() bool {
	return c.allInputsLP()
}

type Broadcaster struct {
	broadcastTo []string
}

func (_ Broadcaster) isModule() {}

type Compute struct {
	moduleID string
	pulse    Pulse
	emitter  string
}

func (b Broadcaster) broadcast(modules Modules) utils.FIFO[Compute] {
	var toBeComputed = utils.FIFO[Compute]{Queue: make([]Compute, 0)}
	var lowPulse = 0
	for _, module := range b.broadcastTo {
		var compute = Compute{getId(modules[module]), lowPulse, "broadcaster"}
		toBeComputed = toBeComputed.Enqueue(compute)
	}
	return toBeComputed
}

type Network = map[string][]string
type Modules = map[string]Module

func copyModules(modules Modules) Modules {
	var newModules = make(map[string]Module)
	for module := range modules {
		newModules[module] = modules[module]
	}
	return newModules
}

func pushButton(network Network, modules Modules, b Broadcaster) (int, int, Modules) {
	var numberLowPulses = 0
	var numberHighPulses = 0
	var LeftToBeComputed = b.broadcast(modules)
	numberLowPulses += len(b.broadcastTo) + 1
	for LeftToBeComputed.Len() != 0 {
		var computed = Compute{}
		LeftToBeComputed, computed = LeftToBeComputed.Dequeue()
		//ModuleID to be computed
		var newPulse, newModule = compute(modules, computed.moduleID, computed.pulse, computed.emitter)
		modules[getId(newModule)] = newModule
		if newPulse > -1 {
			var connections = network[computed.moduleID]
			for _, connection := range connections {
				var _, inNetwork = modules[connection]
				if inNetwork {
					LeftToBeComputed = LeftToBeComputed.Enqueue(Compute{getId(modules[connection]), newPulse, getId(newModule)})
				}
				if newPulse == 0 {
					numberLowPulses += 1
				} else {
					numberHighPulses += 1
				}
			}
		}

	}
	return numberLowPulses, numberHighPulses, modules
}

func isInDefaultState(network Network, modules Modules) bool {
	for module := range network {
		switch module := modules[module].(type) {
		case FlipFlop:
			if !module.isInDefaultState() {
				return false
			}
		case Conjunction:
			if !module.isInDefaultState() {
				return false
			}
		case Broadcaster:
		default:
			panic("Unknown module.")
		}
	}
	return true
}

func split(line string) []string {
	return strings.Split(line, " -> ")
}

func splitConnections(connections string) []string {
	return strings.Split(connections, ", ")
}

func connectionsToConjunctions(flipflops map[string][]string, conjunctions map[string][]string, broadcasted []string) map[string]map[string]int {
	var connectionsIn = make(map[string]map[string]int)
	for flipflop := range flipflops {
		for _, connection := range flipflops[flipflop] {
			var _, isConjunction = conjunctions[connection]
			if isConjunction {
				if connectionsIn[connection] == nil {
					connectionsIn[connection] = make(map[string]int)
				}
				connectionsIn[connection][flipflop] = 0
			}
		}
	}
	for conjunction := range conjunctions {
		for _, connection := range conjunctions[conjunction] {
			var _, isConjunction = conjunctions[connection]
			if isConjunction {
				if connectionsIn[connection] == nil {
					connectionsIn[connection] = make(map[string]int)
				}
				connectionsIn[connection][conjunction] = 0
			}
		}
	}
	for _, connection := range broadcasted {
		var _, isConjunction = conjunctions[connection]
		if isConjunction {
			if connectionsIn[connection] == nil {
				connectionsIn[connection] = make(map[string]int)
			}
			connectionsIn[connection]["broadcaster"] = 0
		}
	}
	return connectionsIn
}

func convertMaps(flipflops map[string][]string, conjunctions map[string][]string, broadcasted []string) (Network, Modules, Broadcaster) {
	var connectionsToConjunctions = connectionsToConjunctions(flipflops, conjunctions, broadcasted)
	var transpose = make(map[string]Module)
	var network = make(map[string][]string)
	for flipflop := range flipflops {
		transpose[flipflop] = FlipFlop{flipflop, 0}
		network[flipflop] = flipflops[flipflop]
	}
	for conjunction := range conjunctions {
		transpose[conjunction] = Conjunction{conjunction, connectionsToConjunctions[conjunction]}
		network[conjunction] = conjunctions[conjunction]
	}

	transpose["broadcast"] = Broadcaster{broadcasted}

	return network, transpose, Broadcaster{broadcasted}
}

func getNetwork(input []string) (Network, Modules, Broadcaster) {
	var flipflops = make(map[string][]string)
	var conjunctions = make(map[string][]string)
	var broadcasted []string
	for _, line := range input {
		var tokens = split(line)
		switch tokens[0][0] {
		case '&':
			conjunctions[strings.TrimPrefix(tokens[0], "&")] = splitConnections(tokens[1])
		case '%':
			flipflops[strings.TrimPrefix(tokens[0], "%")] = splitConnections(tokens[1])
		case 'b':
			broadcasted = splitConnections(tokens[1])
		}
	}
	return convertMaps(flipflops, conjunctions, broadcasted)
}

func partOne(input string) int {
	var k = 1
	var nHighPulse, nLowPulse int
	var resLowPulse []int = make([]int, 0)
	var resHighPulse = make([]int, 0)
	var network, modules, broadcaster = getNetwork(utils.LineByLine(input))
	nLowPulse, nHighPulse, modules = pushButton(network, modules, broadcaster)
	resLowPulse = append(resLowPulse, nLowPulse)
	resHighPulse = append(resHighPulse, nHighPulse)
	for !isInDefaultState(network, modules) && k < 1000 {
		nLowPulse, nHighPulse, modules = pushButton(network, modules, broadcaster)
		resLowPulse = append(resLowPulse, nLowPulse)
		resHighPulse = append(resHighPulse, nHighPulse)
		k += 1
	}
	var cycles = 1000 / k
	var remainder = 1000 % k
	return cycles*(utils.Sum(resLowPulse)*cycles*utils.Sum(resHighPulse)) +
		utils.Sum(resLowPulse[0:remainder])*utils.Sum(resHighPulse[0:remainder])
}

// rx's unique antecedent is a conjunction called sq
func getSQAntecedents(network Network) map[string]int {
	var antecedents = make(map[string]int)
	for module := range network {
		if slices.Contains(network[module], "sq") {
			antecedents[module] = 0
		}
	}
	return antecedents
}

func pushButtonWatch(network Network, modules Modules, b Broadcaster, watched []string) (map[string]bool, Modules) {
	var res = make(map[string]bool)
	var LeftToBeComputed = b.broadcast(modules)
	for LeftToBeComputed.Len() != 0 {
		var computed = Compute{}
		LeftToBeComputed, computed = LeftToBeComputed.Dequeue()
		var newPulse, newModule = compute(modules, computed.moduleID, computed.pulse, computed.emitter)
		if slices.Contains(watched, getId(newModule)) {
			if newPulse == 0 {
				res[getId(newModule)] = true
			}
		}
		modules[getId(newModule)] = newModule

		if newPulse > -1 {
			var connections = network[computed.moduleID]
			for _, connection := range connections {
				var _, inNetwork = modules[connection]
				if inNetwork {
					LeftToBeComputed = LeftToBeComputed.Enqueue(Compute{getId(modules[connection]), newPulse, getId(newModule)})
				}
			}
		}
	}

	return res, modules
}

func partTwo(input string) int {
	var network, modules, broadcaster = getNetwork(utils.LineByLine(input))
	var antecedents = map[string]int{"vk": 0, "dn": 0, "kb": 0, "vm": 0}
	var foundAll = false
	var k = 0
	for !foundAll {
		k += 1
		var pushResults map[string]bool
		pushResults, modules = pushButtonWatch(network, modules, broadcaster, []string{"vk", "dn", "kb", "vm"})
		fmt.Printf("\rPush: %d.", k)

		for module := range pushResults {
			if pushResults[module] && antecedents[module] == 0 {
				antecedents[module] = k
			}
		}

		var findAll = true
		for module := range antecedents {
			if antecedents[module] == 0 {
				findAll = false
			}
		}
		foundAll = findAll
	}
	var press = make([]int, 0)
	for module := range antecedents {
		press = append(press, antecedents[module])
	}
	fmt.Println("")
	return utils.PPCM(press)
}

func main() {
	var start = time.Now()
	fmt.Println(partOne("20/input.txt"))
	fmt.Println(time.Since(start))
	start = time.Now()
	fmt.Println(partTwo("20/input.txt"))
	fmt.Println(time.Since(start))
}
