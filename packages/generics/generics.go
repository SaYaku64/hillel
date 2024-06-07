package generics

import "fmt"

// simplest usage
func Print[T any](value T) {
	fmt.Println(value)
}

////////////////////////

// generic type struct
type Container[T any] struct {
	value T
}

// method for generic type
func (c Container[T]) GetValue() T {
	return c.value
}

func GenericTypeExample() {
	intContainer := Container[int]{value: 42}
	fmt.Println(intContainer.GetValue()) // 42

	stringContainer := Container[string]{value: "Hello"}
	fmt.Println(stringContainer.GetValue()) // Hello
}

////////////////////////

// generic constraints (обмеження)
type Number interface {
	int | float64 | uint
}

func AddNum[T Number](a, b T) T {
	return a + b
}

////////////////////////

type Adder[T any] interface {
	Add(a, b T) T
}

type IntAdder struct{}

func (IntAdder) Add(a, b int) int {
	return a + b
}

func Add[T any](adder Adder[T], a, b T) T {
	return adder.Add(a, b)
}
