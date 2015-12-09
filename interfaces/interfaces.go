package interfaces
//For interface instance look at methods.go

import "fmt"

type CommonUserGreetingInterface struct {
	User string
	Name string
	Surname string
	Age int8
	Greeting string
}

type AsianUserGreetingInterface struct {
	User string
	Name string
	Surname string
	Country string
	Greeting string
}

type ExUSSRGreetingInterface struct {
	User string
	Name string
	Greeting string
	Surname string
	Country string
	IQ int8
	IsVatnique bool
}

func GetGreetingsForPutinFans() *ExUSSRGreetingInterface {
	var message *ExUSSRGreetingInterface = new(ExUSSRGreetingInterface)
	message.User = "Low"
	message.Name = "Ivan"
	message.Greeting = getMessageForVatnique()
	message.Country = "ru"
	message.IQ = 40
	message.IsVatnique = true
	fmt.Println("Memory address: ", &message)

	return message
}

func GetGreetingsForKazakhstan() ExUSSRGreetingInterface {
	message := ExUSSRGreetingInterface{
		User: "Standard",
		Name: "Shyktyn",
		Greeting: "Gim vakre gumm mettr",
		Country: "kz",
		IQ: 60,
		IsVatnique:false,
	}

	return message
}