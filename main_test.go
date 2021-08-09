package main

import (
	"fmt"
	"testing"
)

var o = Operator{True, False}

func TestType(t *testing.T) {
	type TestCase struct {
		Name     string
		Function TypeFunction
		Expected Type
	}

	testCases := []*TestCase{
		{
			Name:     "should F function return false",
			Function: o.F,
			Expected: False,
		},
		{
			Name:     "should T function return true",
			Function: o.T,
			Expected: True,
		},
	}

	for _, tc := range testCases {
		result := tc.Function(o.a, o.b)
		fmt.Printf("%s\n", tc.Name)
		if result != tc.Expected {
			t.Errorf("result %v != expected %v", result, tc.Expected)
		}

	}
}

func TestNotLogicalOperator(t *testing.T) {
	type TestCase struct {
		Name     string
		Function func(p TypeFunction) Type
		Arg      TypeFunction
		Expected Type
	}

	testCases := []*TestCase{
		{
			Name: "should Not returns false when given True",

			Function: o.Not,
			Arg:      o.T,
			Expected: False,
		},
		{
			Name:     "should Not returns true when given False",
			Function: o.Not,
			Arg:      o.F,
			Expected: True,
		},
	}

	for _, tc := range testCases {
		result := tc.Function(tc.Arg)
		fmt.Printf("%s\n", tc.Name)
		if result != tc.Expected {
			t.Errorf("result %v != expected %v", result, tc.Expected)
		}
	}
}

func TestLogicalOperator(t *testing.T) {
	type TestCase struct {
		Name     string
		Function func(p, q TypeFunction) Type
		Arg0     TypeFunction
		Arg1     TypeFunction
		Expected Type
	}

	testCases := []*TestCase{
		{
			Name:     "should And return true when given True and True",
			Function: o.And,
			Arg0:     o.T,
			Arg1:     o.T,
			Expected: True,
		},
		{
			Name:     "should Or return false when given True or False",
			Function: o.Or,
			Arg0:     o.T,
			Arg1:     o.F,
			Expected: True,
		},
	}

	for _, tc := range testCases {
		result := tc.Function(tc.Arg0, tc.Arg1)
		fmt.Printf("%s\n", tc.Name)
		if result != tc.Expected {
			t.Errorf("result %v != expected %v", result, tc.Expected)
		}
	}
}

func TestConditionalOperator(t *testing.T) {
	type TestCase struct {
		Name     string
		Function func(p, a, b TypeFunction) Type
		Arg0     TypeFunction
		Arg1     TypeFunction
		Arg2     TypeFunction
		Expected Type
	}

	testCases := []*TestCase{
		{
			Name:     "if True then True else False",
			Function: o.If,
			Arg0:     o.T,
			Arg1:     o.T,
			Arg2:     o.F,
			Expected: True,
		},
		{
			Name:     "if True then False else True",
			Function: o.If,
			Arg0:     o.T,
			Arg1:     o.F,
			Arg2:     o.T,
			Expected: False,
		},
	}

	for _, tc := range testCases {
		result := tc.Function(tc.Arg0, tc.Arg1, tc.Arg2)
		fmt.Printf("%s\n", tc.Name)
		if result != tc.Expected {
			t.Errorf("result %v != expected %v", result, tc.Expected)
		}
	}
}

func TestNumerals(t *testing.T) {
	type TestCase struct {
		Name     string
		Function func(f func(i int) int, x int) int
		Arg0     int
		Arg1     int
		Expected int
	}

	acc := func(i int) int { return i + 1 }

	testCases := []*TestCase{
		{
			Name:     "function should return 0",
			Function: _0,
			Arg1:     0,
			Expected: 0,
		},
		{
			Name:     "function should return 1",
			Function: _1,
			Arg1:     0,
			Expected: 1,
		},
		{
			Name:     "function should return 2",
			Function: _2,
			Arg1:     0,
			Expected: 2,
		},
		{
			Name:     "function should return 3",
			Function: _3,
			Arg1:     0,
			Expected: 3,
		},
		{
			Name:     "function should return 4",
			Function: _4,
			Arg1:     0,
			Expected: 4,
		},
		{
			Name:     "function should return 5",
			Function: _5,
			Arg1:     0,
			Expected: 5,
		},
		{
			Name:     "function should return 6",
			Function: _6,
			Arg1:     0,
			Expected: 6,
		},
		{
			Name:     "function should return 7",
			Function: _7,
			Arg1:     0,
			Expected: 7,
		},
		{
			Name:     "function should return 8",
			Function: _8,
			Arg1:     0,
			Expected: 8,
		},
		{
			Name:     "function should return 9",
			Function: _9,
			Arg1:     0,
			Expected: 9,
		},
	}

	for _, tc := range testCases {
		result := tc.Function(acc, tc.Arg1)
		fmt.Printf("%s\n", tc.Name)
		if result != tc.Expected {
			t.Errorf("result %v != expected %v", result, tc.Expected)
		}
	}
}
