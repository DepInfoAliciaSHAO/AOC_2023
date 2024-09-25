package main

import (
	"AOC2023/utils"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Rule
/*
Encodes a rule.

ratingCategory is an int, c.f. the XMAS map to know which rating it refers to.

compare is 1 if the rules contains a >, -1 if it contains a <.

threshold is the threshold of the rule.

nextWorkflow is the next workflow's ID string representation. It could be a workflow or A and R.
*/
type Rule struct {
	ratingCategory int
	compare        int
	threshold      int
	nextWorkflow   string
}

//Workflow
/*
A workflow is represented by a list of rules and its default state i.e.
the ID of a workflow, A or R.
*/
type Workflow struct {
	rules        []Rule
	defaultState string
}

//Part
/*
Alias encoding the ratings of each part.

For example the rating at index 0 is the rating in category x.
*/
type Part = []int

//Workflows
/*
Alias for a map of workflow where their ID is in string.
*/
type Workflows = map[string]Workflow

//SPLITTER
/*
Map to easily have access to the compare attribute of Rule.
*/
var SPLITTER = map[string]int{
	">": 1,
	"<": -1,
}

//XMAS
/*
Map to easily have access to the index representing a category of rating.
*/
var XMAS = map[string]int{"x": 0, "m": 1, "a": 2, "s": 3}

/*
Returns the rating of a part in the category used by a given rule.
*/
func getRating(p Part, r Rule) int {
	return p[r.ratingCategory]
}

/*
Applies a rule to a list of ratings of a part.

Returns true if the part respects the rules, false if it does not.
*/
func apply(p Part, r Rule) bool {
	switch r.compare {
	case -1:
		return getRating(p, r) < r.threshold
	case 1:
		return getRating(p, r) > r.threshold
	default:
		return false
	}
}

/*
Applies the succession of rules of a given workflow to a part.

Returns the next workflow to be applied to the part.

A, R are special workflows.
*/
func workflow(p Part, w Workflow) string {
	for _, rule := range w.rules {
		if apply(p, rule) {
			return rule.nextWorkflow
		}
	}
	return w.defaultState
}

/*
Finds out whether a part is accepted by the succession of workflows it goes through.

ws is the map of all the workflows.

Returns true if the part is accepted, false if it is not.
*/
func findState(p Part, ws Workflows) bool {
	//First workflow is in.
	var currentWorkflow = "in"
	//workflow function is implemented so that A and R are also returned.
	for currentWorkflow != "A" && currentWorkflow != "R" {
		//Continuing to apply workflows until a final state is reached.
		currentWorkflow = workflow(p, ws[currentWorkflow])
	}
	return currentWorkflow == "A"
}

/*
From an array of string representations of rules, returns the associated Rule array.
*/
func splitRules(stringRules []string) []Rule {
	var rules = make([]Rule, 0)
	for _, str := range stringRules {
		//Choosing splitter according to the nature of the rule.
		var splitter string
		if strings.Contains(str, "<") {
			splitter = "<"
		} else {
			splitter = ">"
		}
		//components = [category, threshold:nextWorkflow]
		var components = strings.Split(str, splitter)
		var category = XMAS[components[0]]
		//tokens = [threshold, nextWorkflow]
		var tokens = strings.Split(components[1], ":")
		var threshold, _ = strconv.Atoi(tokens[0])
		var next = tokens[1]
		var rule = Rule{category, SPLITTER[splitter], threshold, next}
		rules = append(rules, rule)
	}
	return rules
}

/*
From the string representation of the value of a rating, returns the value of the rating.
*/
func getCategoryNumber(ratingStr string) int {
	var n, _ = strconv.Atoi(strings.Split(ratingStr, "=")[1])
	return n
}

/*
From an array of the string representation of the ratings, returns the Part ratings of a part.
*/
func splitParts(partsStr []string) []int {
	var parts = make([]int, 4)
	parts[0] = getCategoryNumber(partsStr[0])
	parts[1] = getCategoryNumber(partsStr[1])
	parts[2] = getCategoryNumber(partsStr[2])
	parts[3] = getCategoryNumber(partsStr[3])
	return parts
}

/*
From the line in the input representing the workflow,
returns the workflow and its string ID.
*/
func buildWorkflow(line string) (*Workflow, string) {
	//ID
	var token1 = strings.Split(line, "{")
	var id = token1[0]
	var token2 = strings.TrimSuffix(token1[1], "}")
	var token3 = strings.Split(token2, ",")
	//Default State is the last element.
	var defaultState = token3[len(token3)-1]
	//Getting the rules.
	var rulesStr = token3[0 : len(token3)-1]
	var rules = splitRules(rulesStr)
	var workflow = Workflow{rules, defaultState}
	return &workflow, id
}

/*
From the input given as an array of its line, returns the workflows and parts described in it.
*/
func buildPartsAndWorkflows(input []string) (Workflows, []Part) {
	var readWorkflows = true
	var parts = make([]Part, 0)
	var ws = make(map[string]Workflow)
	for _, line := range input {
		if len(line) == 0 {
			readWorkflows = false
		} else if readWorkflows {
			var workflow, id = buildWorkflow(line)
			ws[id] = *workflow
		} else {
			var ratingsStr = strings.Split(strings.TrimPrefix(strings.TrimSuffix(line, "}"), "{"), ",")
			var ratingsNum = splitParts(ratingsStr)
			parts = append(parts, Part{ratingsNum[0], ratingsNum[1], ratingsNum[2], ratingsNum[3]})
		}
	}
	return ws, parts
}

/*
Solves part one.

For each part, we find out if it's accepted or not, if it is the sum of its ratings is added to the total rating.
*/
func partOne(input string) int {
	var res = 0
	var ws, parts = buildPartsAndWorkflows(utils.LineByLine(input))

	for _, p := range parts {
		if findState(p, ws) {
			res += utils.Sum(p)
		}
	}
	return res
}

//PartCombination
/*
Used to represent a combination of part ratings.
*/
type RatingCombination struct {
	intervalX utils.Interval
	intervalM utils.Interval
	intervalA utils.Interval
	intervalS utils.Interval
}

/*
Finds out which category rating interval is affected by a given rule.

Returns the interval affected by this rule.
*/
func (pc RatingCombination) getPartInterval(r Rule) utils.Interval {
	switch r.ratingCategory {
	case 0:
		return pc.intervalX
	case 1:
		return pc.intervalM
	case 2:
		return pc.intervalA
	case 3:
		return pc.intervalS
	default:
		panic("No parts.")
	}
}

/*
Computes the different combinations
resulting from the combination of the category ratings of this RatingCombination instance.

For an interval I, you can pick up to I.Offset() ratings, thus the formula.
*/
func (pc RatingCombination) combinations() int {
	return pc.intervalX.Offset() * pc.intervalM.Offset() * pc.intervalA.Offset() * pc.intervalS.Offset()
}

/*
Returns a copy of this RatingCombination instance.
*/
func (pc RatingCombination) copy() RatingCombination {
	return RatingCombination{pc.intervalX, pc.intervalM, pc.intervalA, pc.intervalS}
}

/*
Splits an interval according to a threshold of a given rule.

The first interval will be part of the rating combination to be processed through another workflow.
The second interval will be part of the one which will continue to be to processed to the current workflow.
*/
func split(interval utils.Interval, threshold int, switchOrder bool) (utils.Interval, utils.Interval) {
	var before, after = interval.Split(threshold)
	if switchOrder {
		return after, before
	} else {
		return before, after
	}
}

/*
From a rating combination, returns the result of the application of a rule.

The first rating combination is to be processed on the workflow the rule is pointing at.
The second rating combination is to be processed by the next rule in the workflow.
The last boolean indicates whether the rating combination was affected by the rule.
If it was not, the first rating combination is empty.

In practice, this method is not called if the rating combination doesn't need to be cut in two.
*/
func (pc RatingCombination) cut(r Rule) (RatingCombination, RatingCombination, bool) {
	//The rating combination will be cut into if the threshold of the rule is in the category interval.
	var cut = pc.getPartInterval(r).Contains(r.threshold)
	if !cut {
		return RatingCombination{}, pc.copy(), false
	} else {
		var pcChangeNext = pc.copy()
		var pcContinue = pc.copy()
		var plusEndFirstInterval int
		//If the rule is a > rule, the result of the interval split has to be switched.
		var switchOrder = r.compare == 1
		if !switchOrder {
			//If the rule is a < rule, the end for the first interval is minus one the threshold.
			plusEndFirstInterval = -1
		}
		//Changing the interval of the category on which the rule operates.
		switch r.ratingCategory {
		case 0:
			var next, continue_ = split(pc.intervalX, r.threshold+plusEndFirstInterval, switchOrder)
			pcChangeNext.intervalX = next
			pcContinue.intervalX = continue_
		case 1:
			var next, continue_ = split(pc.intervalM, r.threshold+plusEndFirstInterval, switchOrder)
			pcChangeNext.intervalM = next
			pcContinue.intervalM = continue_
		case 2:
			var next, continue_ = split(pc.intervalA, r.threshold+plusEndFirstInterval, switchOrder)
			pcChangeNext.intervalA = next
			pcContinue.intervalA = continue_
		case 3:
			var next, continue_ = split(pc.intervalS, r.threshold+plusEndFirstInterval, switchOrder)
			pcChangeNext.intervalS = next
			pcContinue.intervalS = continue_
		}
		return pcChangeNext, pcContinue, cut
	}
}

/*
Computes recursively the total number of combinations accepted by all the workflows.

ws is the map of all the workflows
rc RatingCombination is the starting rating combination.
currentWorkflow is the string id of the currentWorkflow, could also be A or R.
*/
func totalCombinations(ws Workflows, currentWorkflow string, rc RatingCombination) int {
	switch currentWorkflow {
	case "A": //Terminal case one, the combinations are accepted.
		return rc.combinations()
	case "R": //Terminal case two, the combinations are not accepted.
		return 0
	default:
		//Recursive case, the combination is split by the workflow to different workflows.

		//Combinations is the number of combinations accepted by the other workflows the rules point at.
		var combinations = 0
		//remaining parts are updated through the traversal of a workflow, they eventually end up
		//in the default state.
		var remainingParts = rc
		//Iterations over the rules
		for _, rule := range ws[currentWorkflow].rules {
			var toBeNexted, toBeContinued, cut = remainingParts.cut(rule)
			if cut {
				//If the combination was cut, recursive call to the results on the cut and the next workflow.
				combinations += totalCombinations(ws, rule.nextWorkflow, toBeNexted)
			}
			//Remaining parts that have to continue through the current workflow are updated.
			remainingParts = toBeContinued
		}
		// The result are the sum of combinations accepted "en route" in the workflow
		// and the remaining combinations accepted by the workflow's default state.
		return combinations + totalCombinations(ws, ws[currentWorkflow].defaultState, remainingParts)
	}
}

/*
Parses the input to build the map of all the workflows.
*/
func buildWorkflows(input []string) Workflows {
	var ws = make(map[string]Workflow)
	for _, line := range input {
		if len(line) == 0 {
			break
		} else {
			var workflow, id = buildWorkflow(line)
			ws[id] = *workflow
		}
	}
	return ws
}

func partTwo(input string) int {
	var ws = buildWorkflows(utils.LineByLine(input))
	var startCombination = RatingCombination{
		intervalX: utils.Interval{Start: 1, End: 4000},
		intervalM: utils.Interval{Start: 1, End: 4000},
		intervalA: utils.Interval{Start: 1, End: 4000},
		intervalS: utils.Interval{Start: 1, End: 4000},
	}
	return totalCombinations(ws, "in", startCombination)
}
func main() {
	var start = time.Now()
	fmt.Println(partOne("19/input.txt"))
	fmt.Println(time.Since(start))
	start = time.Now()
	fmt.Println(partTwo("19/input.txt"))
	fmt.Println(time.Since(start))
}
