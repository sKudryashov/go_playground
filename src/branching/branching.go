package branching
import "fmt"

type JapanGreeting struct {
	Message string
	Auth string
}
const NATION_JAPANESE string = "japan"
const NATION_UKRAINIAN string = "ukrainian"

func MakeCustomGreet(nation string, nations ... string) string {
	var response string
	if prefix := getPrefix("American"); nation == NATION_JAPANESE {
		response = prefix + " Cunayiki atomoro " + nations[0] + getSuffix(nation)
	} else if nation == NATION_UKRAINIAN {
		response = "Vitaemo tut" + nations[1] + getSuffix(nation)
	}
	return response
}

func getSuffix(nation string) (suffix string) {
	switch  {
		case nation == NATION_UKRAINIAN:
			suffix = " Bajaemo garnogo podoroju"
		case nation == NATION_JAPANESE:
			suffix = " Atari cirigamo"
		case nation == "British", nation == "American":
			suffix = " wish you a nice trip"
			fallthrough
		case nation == "Ca" + nation:
			suffix = " the same as for british and americans"
	}
	return suffix
}

func GetTypeOfTheVar(x interface{})  {
	switch t := x.(type) {
	case int: fmt.Println("Given type is int", t)
	case string: fmt.Println("Given type is string")
	case JapanGreeting: fmt.Println("Given type is Japan Greeting")
	default:
		fmt.Println("Unknown")
	}
}

func getPrefix (nation string) (prefix string) {
	switch nation {
		case NATION_UKRAINIAN: prefix = " Pan "
		case NATION_JAPANESE: prefix = " San "
	    case "British", "American":
			prefix = " Mr "
			fallthrough
		case "Canadian": prefix = " Dear sir "

		default: prefix = " doctor "
	}
	return prefix
}