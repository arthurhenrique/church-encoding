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
