package main

import (
	"fmt"
	"image/color"
	"time"
)

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

type NormalDuck struct {
	Duck
	color color.Color
}

type RoboDuck struct {
	Duck
	id string
}

type DeadDuck struct {
	Duck
	timeOfDeath time.Time
}

type IDuck interface {
	Name() string
	Fly() string
	Quack() string
}

func (duck *Duck) Fly() string {
	return duck.flying.Fly()
}

func (duck *Duck) Quack() string {
	return duck.quack()
}

func (duck *Duck) Name() string {
	return duck.name
}

func NewNormalDuck(color color.Color) *NormalDuck {
	return &NormalDuck{Duck{"Normal duck", WingFlying{}, NormalQuack}, color}
}

func NewRoboDuck(id string) *RoboDuck {
	return &RoboDuck{Duck{"Robo-duck", JetFlying{}, RoboQuack}, id}
}

func NewDeadDuck(timeOfDeath time.Time) *DeadDuck {
	return &DeadDuck{Duck{"Dead duck", NoFlying{}, NoQuack}, timeOfDeath}
}

func Details(duck IDuck) string {
	switch d := duck.(type) {
	case *NormalDuck:
		return fmt.Sprint("has a color of ", d.color)
	case *RoboDuck:
		return fmt.Sprint("has an id of ", d.id)
	case *DeadDuck:
		return fmt.Sprint("diead at ", d.timeOfDeath)
	}
	return ""
}

func PrintDuck(duck IDuck) {
	fmt.Printf("Duck by the name of %q %v, tells you %q and %v\n", duck.Name(), duck.Fly(), duck.Quack(), Details(duck))
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
		NewNormalDuck(color.White),
		NewRoboDuck("42373216assdf5632476"),
		NewDeadDuck(time.Now()),
	}
	for _, duck := range ducks {
		PrintDuck(duck)
	}
	for _, duck := range ducks {
		PlayWithDuck(duck)
	}
}
