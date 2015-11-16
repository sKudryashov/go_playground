package loops

import (
	"fmt"
	"../branching"
)

func MultipleGreeting()  {
	count := 0
	for {
		if name := "Hi Tako!"; count >= 5 {
			fmt.Println(name)
			break
		}
		fmt.Println("Multiple Greeting Tako:")
		count++
	}
}

func MultipleMapGreeting(names []branching.JapanGreeting)  {
	for index, keyvalue := range names {
		fmt.Println("Multiple greeting is:", index, keyvalue)
	}
}

func WhileMapGreeting()  {
	count := 0
	for count <= 5 {
		fmt.Println("While we're waiting for greeting:", count)
		count++
	}
}

func ForGreeting()  {
	for i := 0; i <= 2; i++ {
		fmt.Println("For some greeting timeout")
	}
}
