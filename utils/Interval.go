package utils

import "strconv"

type Interval struct {
	Start int
	End   int
}

func emptyInterval() Interval {
	return Interval{-1, -1}
}
func (I Interval) Offset() int {
	return I.End - I.Start + 1
}

func (I Interval) isEmpty() bool {
	return I.Start == -1 && I.End == -1
}

func (I Interval) Contains(element int) bool {
	return I.Start <= element && element <= I.End
}

func intersect(I Interval, J Interval) bool {
	var I_ = Interval{}
	var J_ = Interval{}
	if I.Start < J.Start {
		I_ = I
		J_ = J
	} else {
		I_ = J
		J_ = I
	}
	return J_.Start <= I_.End
}

func (I Interval) isIncludedIn(J Interval) bool {
	return J.Start <= I.Start && I.End <= J.End
}

func (I Interval) equals(J Interval) bool {
	return (I.Start == J.Start) && (I.End == J.End)
}

func (I Interval) Split(splitter int) (Interval, Interval) {
	if I.Contains(splitter) {
		return Interval{I.Start, splitter}, Interval{splitter + 1, I.End}
	} else {
		return emptyInterval(), emptyInterval()
	}
}
func (I Interval) toString() string {
	var start = strconv.Itoa(I.Start)
	var end = strconv.Itoa(I.End)
	return "Start: " + start + ", End: " + end
}
