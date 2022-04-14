package scigo

// Importing fmt package for the sake of printing
import (
	"fmt"
)

// Animal is the name we want but since we are
// to use it as an interface, we will change
// the name into Animaler.
type Animaler interface {

	// Note that we will
	// declare the methods to be used
	// later here in this
	// interface
	Eat()
	Move()
	Speak()
	Error()
}

// A struct holding a string variable: SuperAnimals
type SuperAnimals struct {
	locomotion string
}

// An embedded struct holding content
// from another struct and two
// other string variables
// named Animals
type Animals struct {
	SuperAnimals
	food  string
	sound string
}

// Now we are indirectly implementing
// the Animaler interface without any
// keywords.
// We are about to define each method
// declared in the Animaler interface.
func (x Animals) Eat() {
	// this method will access the variable
	// food in Animal class

	fmt.Println(x.food)
}

func (x Animals) Move() {
	// this method will access the variable
	// locomotion in Animal class
	fmt.Println(x.locomotion)
}

func (x Animals) Speak() {
	// this method will access the variable
	// sound in Animal class
	fmt.Println(x.sound)
}

func (x Animals) Error() {
	fmt.Println("Invalid query entered!")
}

// Finally reached main function where we can
// now test our "GO classes"
func main() {

	// Experiencing a dictionary / map in GO
	// For the animal name as a key,
	// that particular object is a value
	m := map[string]Animals{
		"cow":   Animals{SuperAnimals{"walk"}, "grass", "moo"},
		"Cow":   Animals{SuperAnimals{"walk"}, "grass", "moo"},
		"Bird":  Animals{SuperAnimals{"fly"}, "worms", "peep"},
		"bird":  Animals{SuperAnimals{"fly"}, "worms", "peep"},
		"Snake": Animals{SuperAnimals{"slither"}, "mice", "hsss"},
		"snake": Animals{SuperAnimals{"slither"}, "mice", "hsss"},
	}
	for i := 0; i < 3; i++ {
		fmt.Println("Enter animal name & query (eat / move / speak): ")
		fmt.Print(">")
		var animal, op string
		fmt.Scan(&animal)
		fmt.Print(">")
		fmt.Scan(&op)
		if op == "eat" {
			m[animal].Eat()
		} else if op == "move" {
			m[animal].Move()
		} else if op == "speak" {
			m[animal].Speak()
		} else {
			m[animal].Error()
		}
	}
}
