// call your code from another module
// package greetings

// import(
// 	"fmt"
// )

// func Hello(name string) string {
// 	msg := fmt.Sprintf("Hi %v. Welcome !", name)
// 	return msg
// }

// return a rondom greetings

package greetings

import(
	"errors"
	"fmt"
	"math/rand"
	"time"
	
)

func Hello(name string) ( string, error) {
	if name == ""{
		return "", errors.New("empty name")
	}

	// msg := fmt.Sprintf("Hi %v. Welcome !", name)
	msg := fmt.Sprintf(randomFormat(), name)
	return msg, nil
}

// init sets initial values for variables used in the function.
func init() {
    rand.Seed(time.Now().UnixNano())
}

// randomFormat returns one of a set of greeting messages. The returned
// message is selected at random.
func randomFormat() string {
    // A slice of message formats.
    formats := []string{
        "Hi, %v. Welcome!",
        "Great to see you, %v!",
        "Hail, %v! Well met!",
    }

    // Return a randomly selected message format by specifying
    // a random index for the slice of formats.
    return formats[rand.Intn(len(formats))]
}