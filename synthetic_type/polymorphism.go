package main

import "fmt"

type Starcraft interface {
	Attack()
}

type Protoss struct {
	name string
}

func (p Protoss) Attack() {
	fmt.Println(p.name, " 공격")
}

type Zerg struct {
	name string
}

func (z Zerg) Attack() {
	fmt.Println(z.name, " 공격")
}

type Terran struct {
	name string
}

func (t Terran) Attack() {
	fmt.Println(t.name, " 공격")
}

func main() {
	var star Starcraft
	star = Protoss{"질럿"}
	star.Attack()

	star = Zerg{"히드라"}
	star.Attack()

	star = Terran{"마린"}
	star.Attack()
}
