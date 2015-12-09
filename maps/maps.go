package maps

import (
	"github.com/sKudryashov/go_playground/branching"
	"fmt"
)

type NorwegianGreeting struct {
	code     string
	fullname string
}

type Salutation struct {
	Name     string
	Greeting string
}

func MapPrefixes(nation string) (data string) {

	var prefix = make(map[string]string)

	prefix[branching.NATION_UKRAINIAN] = " Pan "
	prefix[branching.NATION_JAPANESE] = " San "
	prefix["British"] = " Mr "
	prefix["Canadian"] = " Dear sir "

	return prefix[nation]
}

func MapGreetingPrefixes(nation string) (data string) {

	var prefix = map[string]string{
		"JapanGreeting": nation,
		"ChineseGreeting": "Lee",
	}
	// let's check whether the key in map exists and then
	// delete it. This checking is not required though ...
	// delete works without it as well
	if value, exists := prefix["ChineseGreeting"]; exists {
		delete(prefix, value)
	}

	return prefix[nation]
}

func CountryCodes() (capacity, length int, codes interface{}) {
	var mainCodes = make([]string, 3, 4)

	// extending fixed sized slice
	secondaryCodes := [...]string{"Uganda", "Russia", "Botswana"}

	mainCodes = append(mainCodes, "US")
	mainCodes = append(mainCodes, "GBR")
	mainCodes = append(mainCodes, "AUS")
	mainCodes[0] = "Finland"
	mainCodes[1] = "Norway"
	mainCodes[2] = "Sweden"

	norwegianGreeting := []NorwegianGreeting{
		{"Uga", "Uganda"},
		{"Ru", "Russia"},
	}

	codePairs := []NorwegianGreeting{
		{"Bob", "Hello"},
		{"Joe", "Hi"},
		{"Mary", "What is up?"},
	}
	codePairs = append(codePairs, codePairs...)

	codePairs = codePairs[1:6]

	//this is slices deleting example
	codePairsDeleted := append(codePairs[:1], codePairs[3:]...)

	fmt.Println("Main country codes: ", mainCodes, " capacity ", cap(mainCodes), " and len ", len(mainCodes))
	fmt.Println("Secondary country codes: ", secondaryCodes, " len ", len(secondaryCodes))
	fmt.Println("norwegianGreeting:", norwegianGreeting, " len and capacity ", len(norwegianGreeting), cap(norwegianGreeting))
	fmt.Println("Code pairs:", codePairs, " len and capacity ", len(codePairs), cap(codePairs))
	fmt.Println("slice Deleted : ", codePairsDeleted, " len and capacity ", len(codePairsDeleted), cap(codePairsDeleted))

	return cap(mainCodes), len(mainCodes), mainCodes
}