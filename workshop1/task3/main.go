package main

import "fmt"

type FlyingBehavior interface {
	Fly() string
}

type WingFlying struct{}

type JetFlying struct{}

type NoFlying struct{}

func (flying WingFlying) Fly() string {
	return "flies by flapping its wings"
}

func (flying JetFlying) Fly() string {
	return "flies by using jet engie at its back"
}

func (flying NoFlying) Fly() string {
	return "cannot fly"
}

type QuackBehavior func() string

func NormalQuack() string {
	return "quack!"
}

func RoboQuack() string {
	return "kill all humans!"
}

func NoQuack() string {
	return ""
}

type Duck struct {
	name   string
	flying FlyingBehavior
	quack  QuackBehavior
}

type IDuck interface {
	Name() string
	Fly() string
	Quack() string
}

func (duck Duck) Fly() string {
	return duck.flying.Fly()
}

func (duck Duck) Quack() string {
	return duck.quack()
}

func (duck Duck) Name() string {
	return duck.name
}

func NewDuck(name string, flying FlyingBehavior, quack func() string) IDuck {
	return Duck{name, flying, quack}
}

func NewNormalDuck() IDuck {
	return NewDuck("Normal duck", WingFlying{}, NormalQuack)
}

func NewRoboDuck() IDuck {
	return NewDuck("Robo-duck", JetFlying{}, RoboQuack)
}

func NewDeadDuck() IDuck {
	return NewDuck("Dead duck", NoFlying{}, NoQuack)
}

func PrintDuck(duck IDuck) {
	fmt.Printf("Duck by the name of %q %v and tells you %q\n", duck.Name(), duck.Fly(), duck.Quack())
}

func PlayWithDuck(duck IDuck) {
	fmt.Println(duck.Fly())
	fmt.Println(duck.Fly())
	fmt.Println(duck.Fly())
	fmt.Printf("Says: %q\n", duck.Quack())
	fmt.Println(duck.Fly())
	fmt.Printf("Says: %q\n", duck.Quack())
	fmt.Printf("Says: %q\n", duck.Quack())
	fmt.Printf("Says: %q\n", duck.Quack())
}

func main() {
	ducks := []IDuck{
		NewNormalDuck(),
		NewRoboDuck(),
		NewDeadDuck(),
	}
	for _, duck := range ducks {
		PrintDuck(duck)
	}
	for _, duck := range ducks {
		PlayWithDuck(duck)
	}
}
