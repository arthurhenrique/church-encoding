package main

import (
	"fmt"
)

type Type bool
type TypeFunction func(Type, Type) Type

const True = true
const False = false

func T(a, b Type) Type {
	return a
}

func F(a, b Type) Type {
	return b
}

func Not(p TypeFunction, a, b Type) Type {
	return p(F(a, b), T(a, b))
}

func And(p TypeFunction, q TypeFunction, a, b Type) Type {
	return p(q(a, b), p(a, b))
}

func Or(p TypeFunction, q TypeFunction, a, b Type) Type {
	return p(p(a, b), q(a, b))
}

func If(p TypeFunction, q TypeFunction, a, b Type) Type {
	return p(p(a, b), q(a, b))
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
	//operators
	fmt.Printf("%v\n", T(True, False))
	fmt.Printf("%v\n", F(True, False))
	fmt.Printf("%v\n", Not(T, True, False))
	fmt.Printf("%v\n", And(T, T, True, False))
	fmt.Printf("%v\n", Or(T, T, True, False))
	fmt.Printf("%v\n", If(F, F, True, False))

	// numerals
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
}
