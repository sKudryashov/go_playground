package methods

import (
	"fmt"
	"reflect"
)

type GreetingsManager struct {
	Country string
	Text string
	GMT int
}

type Greetings []GreetingsManager

type GreetingAble interface  {
	AddSimpleGreeting()
	SaySomething()
}

type GreetingShow interface {
	GreetingAble
	ShowGreetings()
}

type temporary interface {
	Temporary() bool
}

// IsTemporary returns true if err is temporary.
func IsTemporary(err error) bool {
	te, ok := err.(temporary)
	return ok && te.Temporary()
}

type Write interface  {
	Write(p []byte) (n int, err error)
}

func (manager GreetingsManager) Write(p []byte) (n int, err error) {
	fmt.Println("Writer interface realization: ", p, " and memory address ", &p)
	n = len(p)
	err = nil
	return
}

func (manager *GreetingsManager) SaySomething() {
	fmt.Println("SaySomething call :) from Greetings manager")
}

func (greeting *Greetings) PushGreeting(manager GreetingsManager) {
	//make()
	//append(Greetings, manager)
}

//func (greeting *Greetings) PullGreeting() *Greetings {
//	return Greetings
//}

func (manager *GreetingsManager) AddSimpleGreeting () {
	manager.Country = "country"
	manager.Text = "greeting"
	manager.GMT = 5

	fmt.Println("")
	fmt.Println("GreetingsManager:", manager, " and type ", reflect.TypeOf(manager))
}

func (greetings *Greetings) ShowGreetings() {
	fmt.Println("Show greetings memory address: ", &greetings, " and type ", reflect.TypeOf(greetings),
		" and value ", *greetings)

	// when i dereferenced greetins it worked just perfectly
	for index, value := range *greetings {
		fmt.Println("Methods:ShowGreetings -> show greetings: ", index, value)
	}
}

func (greetings *Greetings) SpinGreetings()  {
	fmt.Println("Show greetings memory address: ", &greetings, " and type ", reflect.TypeOf(greetings),
		" and value ", greetings)
	for index, value := range *greetings {
		fmt.Println("Methods:SpinGreetings -> show greetings: ", index, value)
	}
}