package foopkg

import "fmt"

func Greet(somebody string) string {
	return fmt.Sprintf("Greetings, %s", somebody)
}
