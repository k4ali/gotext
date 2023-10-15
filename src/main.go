package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
	"unicode"
)

type UserInput struct {
	Input  string
	Length int
}

type StringCharacter struct {
	Value      string
	ValueInt32 rune
	Valid      bool
	IsDigit    bool
}

var _numbers int = 9
var _alphabet int = 26

func try_user_input_request() {
	inputRequest := func() UserInput {
		var data UserInput

		fmt.Println("Type something here: ")
		scan := bufio.NewScanner(os.Stdin)
		scan.Scan()

		data.Input = scan.Text()
		data.Length = len(data.Input)
		return data
	}

	isInputCorrect := func(length int) bool {
		return (length > 0)
	}

	var data UserInput = inputRequest()
	if !isInputCorrect(data.Length) {
		fmt.Println("Error, wrong text form, try again..\n")

		try_user_input_request()
		return
	}

	start_text_animation(data)
}

func start_text_animation(data UserInput) {
	randomLetter := func(uppercaseState bool) int32 {
		letter := rune('a' + rand.Intn(_alphabet))
		if uppercaseState {
			letter = rune('A' + rand.Intn(_alphabet))
		}

		return letter
	}

	isCharacterValid := func(char byte) StringCharacter {
		_char := rune(char)
		return StringCharacter{string(char), _char, (unicode.IsLetter(_char) == true || unicode.IsDigit(_char) == true), unicode.IsDigit(_char) == true}
	}

	var completed string = ""
	for i := 0; i < data.Length; i++ {
		characterData := isCharacterValid(data.Input[i])

		if characterData.Valid {
			if !characterData.IsDigit {
				for j := 0; j < _alphabet; j++ {
					time.Sleep(2)
					fmt.Println(completed + string(randomLetter(unicode.IsUpper(characterData.ValueInt32))))
				}
			} else {
				for j := 0; j < _numbers; j++ {
					time.Sleep(2)
					fmt.Println(completed + strconv.Itoa(rand.Intn(_numbers)))
				}
			}
		}

		completed = completed + characterData.Value
	}

	fmt.Println(completed + "\n\n")
	completed = ""

	try_user_input_request()
}

func main() {
	try_user_input_request()
}
