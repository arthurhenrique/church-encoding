package main

import (
	"fmt"
)

type Type bool

type TypeFunction func(Type, Type) Type

type Operator struct {
	a Type
	b Type
}

const True = true
const False = false

func (o Operator) T(a, b Type) Type {
	return a
}

func (o Operator) F(a, b Type) Type {
	return b
}

func (o Operator) Not(p TypeFunction) Type {
	return p(o.F(o.a, o.b), o.T(o.a, o.b))
}

func (o Operator) And(p, q TypeFunction) Type {
	return p(q(o.a, o.b), p(o.a, o.b))
}

func (o Operator) Or(p, q TypeFunction) Type {
	return p(p(o.a, o.b), q(o.a, o.b))
}

func (o Operator) If(p, a, b TypeFunction) Type {
	return p(a(o.a, o.b), b(o.a, o.b))
}

func _0(f func(int) int, x int) int {
	return x
}

func _1(f func(int) int, x int) int {
	return f(x)
}

func _2(f func(int) int, x int) int {
	return f(f(x))
}

func _3(f func(int) int, x int) int {
	return f(f(f(x)))
}
func _4(f func(int) int, x int) int {
	return f(f(f(f(x))))
}

func _5(f func(int) int, x int) int {
	return f((f(f(f(x)))))
}
func _6(f func(int) int, x int) int {
	return f(f(f(f(f(f(x))))))
}

func _7(f func(int) int, x int) int {
	return f(f(f(f(f(f(f(x)))))))
}

func _8(f func(int) int, x int) int {
	return f(f(f(f(f(f(f(f(x))))))))
}

func _9(f func(int) int, x int) int {
	return f(f(f(f(f(f(f(f(f(x)))))))))
}

func main() {

	o := Operator{True, False}

	// types
	fmt.Printf("types\n")
	fmt.Printf("%v\n", o.T(o.a, o.b))
	fmt.Printf("%v\n", o.F(o.a, o.b))

	//operators
	fmt.Printf("\noperators\n")
	fmt.Printf("Not True = %v\n", o.Not(o.T))
	fmt.Printf("True and True = %v\n", o.And(o.T, o.T))
	fmt.Printf("True or False = %v\n", o.Or(o.T, o.F))
	fmt.Printf("If False Then True Else False = %v\n", o.If(o.T, o.T, o.F))

	// numerals
	fmt.Printf("\nnumerals\n")
	acc := func(i int) int { return i + 1 }
	fmt.Printf("%v\n", _0(acc, 0))
	fmt.Printf("%v\n", _1(acc, 0))
	fmt.Printf("%v\n", _2(acc, 0))
	fmt.Printf("%v\n", _3(acc, 0))
	fmt.Printf("%v\n", _4(acc, 0))
	fmt.Printf("%v\n", _5(acc, 0))
	fmt.Printf("%v\n", _6(acc, 0))
	fmt.Printf("%v\n", _7(acc, 0))
	fmt.Printf("%v\n", _8(acc, 0))
	fmt.Printf("%v\n", _9(acc, 0))

	// todo y-combinators
}
