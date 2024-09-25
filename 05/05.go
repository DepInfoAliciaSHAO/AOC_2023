package main

import (
	"AOC2023/utils"
	"fmt"
	"math"
	"strings"
	"time"
	"unicode"
)

////////////
//PART ONE//
////////////

/*
Updates the current locations of the seeds as maps are being read.
rangesToCheck stores all the different ranges of the property maps.
*/
func updateLocations(rangesToCheck [][]int, currentLocations []int) []int {
	//Iteration over the current locations.
	for i, location := range currentLocations {
		//For each location, let us check if it needs to be projected according to the property map.
		for _, propertyRange := range rangesToCheck {
			var source = propertyRange[1]
			var destination = propertyRange[0]
			var range_ = propertyRange[2]
			if location >= source && location < source+range_ {
				currentLocations[i] = destination + location - source
				//A seed can only be projected once.
				break
			}
		}
	}
	return currentLocations
}

/*
Returns the locations of the seeds after reading through the map.
*/
func getLocations(input []string) []int {
	var currentLocations = make([]int, 0)
	var readMap = false
	var nextSeedsString = strings.TrimSpace(strings.Split(input[0], ":")[1])
	currentLocations = utils.StringArrayToIntArray(strings.Split(nextSeedsString, " "))

	var rangesToCheck = make([][]int, 0)
	for i := 2; i < len(input); i++ {
		if len(input[i]) == 0 {
			//A map has been read fully, time to update the locations and reset the map parameters.
			readMap = false
			updateLocations(rangesToCheck, currentLocations)
			rangesToCheck = make([][]int, 0)
		} else if unicode.IsLetter(rune(input[i][0])) {
			//If the line encountered starts with a letter, a property map is starting on the next line.
			readMap = true
		} else if readMap {
			//While reading a map, ranges that are projected are stored.
			var propertyRange []int = utils.StringArrayToIntArray(strings.Split(input[i], " "))
			rangesToCheck = append(rangesToCheck, propertyRange)
		}
	}
	//The last map hasn't been used yet, let's update the locations one last time.
	updateLocations(rangesToCheck, currentLocations)
	return currentLocations
}

func lowestLocation(input string) int {
	return utils.MinimumArray(getLocations(utils.LineByLine(input)))
}

////////////
//PART TWO//
////////////

/*
Cuts an interval I by J. There is at most three interval as a result:
the left leftover part of I, J, and the right leftover part of J.

In case one of the interval doesn't exist (e.g. if I and J have the same left bound),
an empty interval is returned of the leftover part that doesn't exist.
*/
func image(I utils.Interval, J utils.Interval) (utils.Interval, utils.Interval, utils.Interval) {
	if !utils.Intersect(I, J) {
		// I: [----------------]
		// J:                               [----------]
		return utils.EmptyInterval(), utils.EmptyInterval(), utils.EmptyInterval()
	} else if J.IsIncludedIn(I) {
		// I: [---------------------]
		// J:         [----------]
		return utils.Interval{Start: I.Start, End: J.Start - 1},
			J,
			utils.Interval{Start: J.End + 1, End: I.End}
	} else if I.IsIncludedIn(J) {
		// I:    [--------------]
		// J: [--------------------------]
		return utils.EmptyInterval(),
			I,
			utils.EmptyInterval()
	} else if I.End <= J.End {
		// I: [-----------]
		// J:    [-------------]
		return utils.Interval{Start: I.Start, End: J.Start - 1},
			utils.Interval{Start: J.Start, End: I.End},
			utils.EmptyInterval()
	} else {
		// I:       [-----------]
		// J:   [-----------]
		return utils.EmptyInterval(),
			utils.Interval{Start: I.Start, End: J.End},
			utils.Interval{Start: J.End + 1, End: I.End}
	}
}

func offSetToEnd(start int, offset int) int {
	return start + offset - 1
}

/*
Projects a seed location interval to the next property map.
*/
func sourceToDestination(current utils.Interval, source int, destination int, offset int) (utils.Interval, utils.Interval, utils.Interval) {
	var sourceInterval = utils.Interval{Start: source, End: offSetToEnd(source, offset)}
	//Cutting the interval
	var leftLeftover, toBeProjected, rightLeftover = image(current, sourceInterval)
	//Projecting the interval that has to be projected if it exists.
	if !toBeProjected.IsEmpty() {
		var destinationInterval = utils.Interval{destination + toBeProjected.Start - source, destination + toBeProjected.End - source}
		return leftLeftover, destinationInterval, rightLeftover
	} else {
		return utils.EmptyInterval(), utils.EmptyInterval(), utils.EmptyInterval()
	}
}

func addAll(m1 map[utils.Interval]bool, m2 map[utils.Interval]bool) {
	for interval := range m2 {
		m1[interval] = true
	}
}

func getLocationRanges(input []string) map[utils.Interval]bool {
	var currentLocationRanges = make(map[utils.Interval]bool)
	var nextLocationRanges = make(map[utils.Interval]bool)
	var readMap = false

	//Initialization of intervals
	var nextSeedsString = strings.TrimSpace(strings.Split(input[0], ":")[1])
	var currentSeedRanges = utils.StringArrayToIntArray(strings.Split(nextSeedsString, " "))
	for i := 0; i < len(currentSeedRanges)/2; i++ {
		currentLocationRanges[utils.Interval{currentSeedRanges[2*i], offSetToEnd(currentSeedRanges[2*i], currentSeedRanges[2*i+1])}] = true
	}

	//Same reading process as in part one.
	for i := 2; i < len(input); i++ {
		if len(input[i]) == 0 {
			readMap = false
			addAll(currentLocationRanges, nextLocationRanges)
			nextLocationRanges = make(map[utils.Interval]bool)
		} else if unicode.IsLetter(rune(input[i][0])) {
			readMap = true
		} else if readMap {
			var propertyRange []int = utils.StringArrayToIntArray(strings.Split(input[i], " "))
			//Updating the intervals to be projected, for each property range.
			for interval := range currentLocationRanges {
				if currentLocationRanges[interval] {
					var K, L, M = sourceToDestination(interval, propertyRange[1], propertyRange[0], propertyRange[2])
					// If there hasn't been any projection, none of the next lines are read,
					// the location interval is kept.

					//If not, there are at least two intervals that are not empty and L is never empty.
					if !L.IsEmpty() {
						nextLocationRanges[L] = true
						//The original interval has been cut, it needs to be removed from the set of the current intervals that are considered.
						currentLocationRanges[interval] = false
					}
					if !K.IsEmpty() {
						currentLocationRanges[K] = true
					}
					if !M.IsEmpty() {
						currentLocationRanges[M] = true
					}
				}
			}
		}
	}
	addAll(currentLocationRanges, nextLocationRanges)
	return currentLocationRanges
}

func lowestLocationFromRanges(locations map[utils.Interval]bool) int {
	var mini = math.MaxInt
	for interval := range locations {
		if locations[interval] && interval.Start < mini {
			mini = interval.Start
		}
	}
	return mini
}

func lowestLocation2(input string) int {
	return lowestLocationFromRanges(getLocationRanges(utils.LineByLine(input)))
}
func main() {
	var start = time.Now()
	fmt.Println(lowestLocation("05/input.txt"))
	fmt.Println(time.Since(start))
	start = time.Now()
	fmt.Println(lowestLocation2("05/input.txt"))
	fmt.Println(time.Since(start))
}
