package main

import (
	"fmt"
	"os"
	"strings"
	//"io"
	"reflect"

	"github.com/sKudryashov/go_playground/branching"
	"github.com/sKudryashov/go_playground/concurrency"
	"github.com/sKudryashov/go_playground/flow"
	"github.com/sKudryashov/go_playground/interfaces"
	"github.com/sKudryashov/go_playground/loops"
	"github.com/sKudryashov/go_playground/maps"
	"github.com/sKudryashov/go_playground/methods"
	"github.com/sKudryashov/go_playground/variables"
)

type UserGreeting struct {
	username string
	message  string
}

type GreetingFactory func(message string) (str1, str2 string)

type JapanGreeting struct {
	message string
	auth    string
}

type T struct {
	name string
}

func (t T) M() {
	fmt.Println(t.name)
}

type I interface {
	M()
}

type TT struct {
	name string
}

func (t *TT) M(string) {}
func (t TT) N(float64) {}

type TTT struct {
	name string
}

func (t TTT) M() string {
	return t.name
}

type U struct {
	*TTT
}

func main() {
	//todo: probably need to be moved to methods
	//t1 := T{name: "Michał Int Test"}
	//t2 := T{name: "Tom Int TEST"}
	//// very interesting way to call methods
	////m := I.M
	////m(t1)
	////m(t2)
	////
	////t := TT{name: "Michał"}
	////m = t.M
	////n := t.N
	////m("foo") // it's completely legal
	//n(1.1)

	//var _ io.Writer = (*UserGreeting)(nil) check if interface satisfies this type

	////////////////////////
	flow.InitFlow()
	module := "function basics"
	author := os.Getenv("USER")
	var newMod, newAuth string
	var engMsg = UserGreeting{strings.ToUpper(author), "Hi new user " + author}
	var ukrMsg = UserGreeting{message: "Vitaemo " + author, username: strings.ToUpper(author)}

	//5
	bestFinish := beastLeagueFinishes(13, 10, 13, 17, 14, 16, 22, 34, 5)
	//Tor
	shortestNameOfTheTeam := shortestNameOfTheTeam("Manchester", "Liverpool", "Arsenal", "Thunderland", "Tor")

	newMod, newAuth = converter(module, author)
	newModuleN, newAuthorN := namedConverter(module, author)

	//let's try Ukrainian
	invitateCurrentSystemUser(ukrMsg, printUkr)
	//and Englisj
	_, message := invitateCurrentSystemUser(engMsg, printEng)

	var japanFactory GreetingFactory = getJapanGreetingFactory(author)
	// kudryashov-san, Yakamotu coageri password Nakamichi
	japanAuthor, japanMessage := japanFactory(author)

	fmt.Println(newMod, newAuth)
	fmt.Println("BestValue is:", bestFinish)
	fmt.Println("Shortest name of the team is:", shortestNameOfTheTeam)
	// data from named function
	fmt.Println("Named function data: ", newAuthorN, newModuleN)
	fmt.Println("Just a one message left:", message)
	fmt.Println("Japan type of system greeting: ", japanAuthor, japanMessage)
	fmt.Println(branching.MakeCustomGreet("japan", branching.NATION_JAPANESE, branching.NATION_UKRAINIAN))

	branching.GetTypeOfTheVar(japanFactory)
	branching.GetTypeOfTheVar(2)

	loops.MultipleGreeting()

	japanGreetingAsSlice := []branching.JapanGreeting{
		{"Tetsuya", "Komuro"},
		{"Satoshi", "Tomie"},
		{"Takish", "Shita"},
	}

	loops.MultipleMapGreeting(japanGreetingAsSlice)
	loops.WhileMapGreeting()
	loops.ForGreeting()
	loops.MultipleGreeting()

	fmt.Println(maps.MapPrefixes("Canadian"))

	maps.CountryCodes()
	maps.MapPrefixes("British")
	maps.MapGreetingPrefixes("Canadian")

	variables.TryRefs()

	fmt.Println(interfaces.GetGreetingsForKazakhstan())
	putinFansGreeting := interfaces.GetGreetingsForPutinFans()

	fmt.Println("Type of variable -> putinFansGreeting: ", reflect.TypeOf(putinFansGreeting),
		putinFansGreeting)
	putinFansGreetingRef := putinFansGreeting

	putinFansGreeting.IQ = 80

	fmt.Println("Type of variable -> putinFansGreetingRef: ", reflect.TypeOf(putinFansGreetingRef),
		putinFansGreetingRef)
	fmt.Println("Memory address: ", &putinFansGreeting)

	fmt.Println(putinFansGreetingRef)
	fmt.Println(interfaces.GetGreetingsForPutinFans())

	greetingsManager := methods.GreetingsManager{
		Country: "BRA",
		Text:    "Estrado di biando",
		GMT:     6,
	}

	greetings := methods.Greetings{
		{"USA", "Hi folks", 9},
		{"CAN", "Je vu i", 10},
		{"MEX", "Estrado di biando", 7},
		greetingsManager,
	}

	greetingsManager.AddSimpleGreeting()
	makeGreetingFromMethPackage(&greetingsManager)

	greetings.ShowGreetings()
	greetings.SpinGreetings()

	//io.Writer implementation
	fmt.Fprint(greetingsManager, "Good news, get buffer")
	//
	//	methods.GreetingsManager.AddSimpleGreeting()
	//	methods.GreetingsManager.AddSimpleGreeting()
	concurrency.Launch()
}

func makeGreetingFromMethPackage(m methods.GreetingAble) {
	m.AddSimpleGreeting()
	i := 1
	fmt.Println("i", i)
}

//Closure works as usual there
func getJapanGreetingFactory(author string) GreetingFactory {

	var passwordWord string = "Nakamichi"

	return func(message string) (str1, str2 string) {
		var one string = "Yakamotu coageri"
		var two string = author + "-san, password " + passwordWord
		return one, two
	}
}

func invitateCurrentSystemUser(userGreeting UserGreeting, do func(UserGreeting) (str1, str2 string)) (str1, str2 string) {
	user, message := do(userGreeting)
	return user, message
}

func printUkr(data UserGreeting) (str1, str2 string) {
	fmt.Println("Koristuvach:", data.username, " ta povidomlennya - ", data.message)
	return strings.ToUpper(data.username), strings.ToUpper(data.message)
}

func printEng(data UserGreeting) (str1, str2 string) {
	fmt.Println("The user is:", data.username, "and the message for him is - ", data.message)
	return strings.ToLower(data.username), strings.ToLower(data.message)
}

func namedConverter(module, author string) (newModuleNamed, newAuthorNamed string) {
	newAuthorNamed = strings.ToLower(author)
	newModuleNamed = strings.ToUpper(module)
	return
}

func shortestNameOfTheTeam(names ...string) string {
	var shortest string = names[0]
	for _, i := range names {
		if len(names[0]) > len(i) {
			shortest = i
		}
	}
	return shortest
}

func beastLeagueFinishes(finishes ...int) int {
	var best int
	best = finishes[0]
	for _, i := range finishes {
		if i < best {
			best = i
		}
	}
	return best
}

func converter(s1, s2 string) (str1, str2 string) {
	s1 = strings.ToLower(s1)
	s2 = strings.ToUpper(s2)
	return s1, s2
}
