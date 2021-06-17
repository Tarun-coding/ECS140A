package triangle

import "testing"

func TestGetTriangleType(t *testing.T) {
	type Test struct {
		a, b, c  int
		expected triangleType
	}

	var tests = []Test{
		{1001, 5, 6, UnknownTriangle},
		{5, 2001, 6, UnknownTriangle},
		{5, 6, 3001, UnknownTriangle},
		{0, 5, 6, UnknownTriangle},
		{5, 0, 6, UnknownTriangle},
		{5, 6, 0, UnknownTriangle},
		{3, 4, 8, InvalidTriangle},
		{3, 4, 5, RightTriangle},
		{3, 4, 2, AcuteTriangle},
		{3, 4, 6, ObtuseTriangle},
		// TODO add more tests for 100% test coverage
	}

	for _, test := range tests {
		actual := getTriangleType(test.a, test.b, test.c)
		if actual != test.expected {
			t.Errorf("getTriangleType(%d, %d, %d)=%v; want %v", test.a, test.b, test.c, actual, test.expected)
		}
	}
}
