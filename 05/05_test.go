package main

import "testing"

func TestPartOne(t *testing.T) {
	var actual = lowestLocation("test.txt")
	var expected = 35
	if actual != expected {
		t.Errorf("Lowest Location = %d; want %d", actual, expected)
	}
}

func TestOffset(t *testing.T) {
	var I = Interval{5, 9}
	var actual = I.offset()
	var expected = 5
	if actual != expected {
		t.Errorf("Offset = %d; want %d", actual, expected)
	}
}

func TestStrictlyContains(t *testing.T) {
	var I = Interval{6, 12}
	var element = 8
	var expected = I.contains(element)
	if true != expected {
		t.Errorf("Error in contains")
	}
}

func TestContains(t *testing.T) {
	var I = Interval{6, 12}
	var element = 12
	var actual = I.contains(element)
	if !actual {
		t.Errorf("Error in contains.")
	}
}

func TestIntersectSingleton(t *testing.T) {
	var I = Interval{6, 12}
	var J = Interval{5, 6}
	var actual = intersect(I, J)
	if !actual {
		t.Errorf("Error in intersect.")
	}
}

func TestIntersectOverlap(t *testing.T) {
	var I = Interval{6, 12}
	var J = Interval{5, 8}
	var actual = intersect(I, J)
	if !actual {
		t.Errorf("Error in intersect.")
	}
}

func TestIntersectNotActuallyIntersected(t *testing.T) {
	var I = Interval{6, 12}
	var J = Interval{13, 78}
	var actual = intersect(I, J)
	if actual {
		t.Errorf("Error in intersect.")
	}
}

func TestIsIncludedIn(t *testing.T) {
	var I = Interval{14, 14}
	var J = Interval{13, 78}
	var actual = I.isIncludedIn(J)
	if !actual {
		t.Errorf("Error in intersect.")
	}
}

func TestImageNotIntersected(t *testing.T) {
	var I = Interval{13, 14}
	var J = Interval{15, 78}
	var K, L, M = I.image(J)
	var actual = K.isEmpty() && L.isEmpty() && M.isEmpty()
	if !actual {
		t.Errorf("One interval is non empty in Image when I inter J is null.")
	}
}

func TestImageJIncludedInI(t *testing.T) {
	var J = Interval{17, 20}
	var I = Interval{15, 78}
	var K, L, M = I.image(J)
	var Ke, Le, Me = Interval{15, 16}, Interval{17, 20}, Interval{21, 78}
	if !K.equals(Ke) {
		t.Errorf("Problem in LHS.")
	}
	if !L.equals(Le) {
		t.Errorf("Problem in middle.")
	}
	if !M.equals(Me) {
		t.Errorf("Problem in RHS.")
	}
}

func TestImageLeftOverlap(t *testing.T) {
	var I = Interval{17, 25}
	var J = Interval{12, 21}
	var K, L, M = I.image(J)
	var Ke, Le, Me = emptyInterval(), Interval{17, 21}, Interval{22, 25}
	if !K.equals(Ke) {
		t.Errorf("Problem in LHS. Actual: " + K.toString() + ", expected: " + Ke.toString())
	}
	if !L.equals(Le) {
		t.Errorf("Problem in middle. Actual: " + L.toString() + ", expected: " + Le.toString())
	}
	if !M.equals(Me) {
		t.Errorf("Problem in RHS. Actual: " + M.toString() + ", expected: " + Me.toString())
	}
}

func TestImageRightOverlap(t *testing.T) {
	var I = Interval{17, 25}
	var J = Interval{21, 30}
	var K, L, M = I.image(J)
	var Ke, Le, Me = Interval{17, 20}, Interval{21, 25}, emptyInterval()
	if !K.equals(Ke) {
		t.Errorf("Problem in LHS. Actual: " + K.toString() + ", expected: " + Ke.toString())
	}
	if !L.equals(Le) {
		t.Errorf("Problem in middle. Actual: " + L.toString() + ", expected: " + Le.toString())
	}
	if !M.equals(Me) {
		t.Errorf("Problem in RHS. Actual: " + M.toString() + ", expected: " + Me.toString())
	}
}

func TestPartTwo(t *testing.T) {
	var actual = lowestLocation2("test.txt")
	var expected = 46
	if actual != expected {
		t.Errorf("Lowest Location = %d; want %d", actual, expected)
	}
}
