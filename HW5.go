package main

import (
	"fmt"
)

type worker interface {
	callSkill()
}
func introduceOneself(w worker)  {
	w.callSkill()
}

type person struct {
	name string
	work string
}
type doctor struct {
	p person
	specialization string
}
type rescuer struct{
	p person
	isSwim bool
}
type blogger struct {
	p person
	SubscribersNumber int
}
func (d doctor) callSkill(){
	fmt.Println("Hi! I'm",d.p.name,". I work at the",d.p.work,"as a doctor. I treat people.")
}
func (r rescuer) callSkill(){
	fmt.Println("Hi! I'm",r.p.name,". I work at the",r.p.work,"as a rescuer.I save people.")
}
func (b blogger) callSkill(){
	fmt.Println("Hi! I'm",b.p.name,". I'm bloger. I save people.Like and subscribe!")
}

func typesElements(m map[int]worker) []string{
	var a[] string
	for _,v:=range m{
		a=append(a,fmt.Sprintf("%T", v))
		println(a,fmt.Sprintf("%T", v))
	}

	return a
}
func main() {
	m := make(map[int]worker)
	m[1]=doctor{
		p: person{
			name: "Marry",
			work: "Hospital"},
		specialization: "pediatrician"}
	m[2]=rescuer{
		p: person{
			name: "Volodymyr",
			work: "fire station"},
			isSwim: false}
	m[3]=blogger{
		p: person{
			name: "Nataliia",
			work: "Instagram"},
		SubscribersNumber:57600 }
	for _,n:=range m{
		introduceOneself(n)
	}
	typesElements(m)
}