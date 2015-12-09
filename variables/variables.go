package variables

import (
	"fmt"
	"runtime"
	"reflect"
	"os"
);

const (
	speedOfLight = 186000
	accelerationOfFreefalling = 2.3
)

func TryRefs() {
	name := os.Getenv("USER")
	const PI float64 = 3.14

	course := "Docker deep dive"
	module := 3.2

	//Hello GO! darwin string string float64
	ptr := &module
	fmt.Println("Hello GO!", runtime.GOOS, reflect.TypeOf(name), reflect.TypeOf(course),
		reflect.TypeOf(module), &module, module);

	//Let's try to dereference ptr:  0x18e058 3.2
	fmt.Println("Let's try to dereference ptr: ", ptr, *ptr)

	changeTheCourse(&course)
	changeStackForCourse := &course

	fmt.Println("variables:go -> changeStackForCourse: ", reflect.TypeOf(changeStackForCourse))

	//	the new course is: Go deep dive
	fmt.Println("the new course is:", course);

	//	And the course owner is kudryashov stdout is: &{0x8201ea180} 0x192c18
	fmt.Println("And the course owner is", name, "stdout is:", os.Stdout, &os.Stdout);

}

func changeTheCourse(data *string) string {
	*data = "Go deepest dive"
	return *data
}
