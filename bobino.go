package bobino

import (
	"errors"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cast"
	"strconv"
)

// AskForNumber prompts with a question expecting a number as an answer. The
// function returns only when a valid int value is entered
//
//   age := bobino.AskForNumber("What's your age?")
//   fmt.Printf("You're %d years old!", age)
func AskForNumber(question string) (answer int) {
	validateInteger := func(input string) error {
		_, err := strconv.ParseInt(input, 10, 64)
		if err != nil {
			return errors.New("Not a valid number")
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    question,
		Validate: validateInteger,
	}

	result, _ := prompt.Run()
	return cast.ToInt(result)
}

// AskForDecimalNumber prompts with a question expecting a decimal number as an answer. The
// function returns only when a valid float value is entered
//
//   savings := bobino.AskForDecimalNumber("How much did you save up this week?")
//   fmt.Printf("Wow, you saved %.2f!", savings)
func AskForDecimalNumber(question string) (answer float64) {
	validateDecimal := func(input string) error {
		_, err := strconv.ParseFloat(input, 64)
		if err != nil {
			return errors.New("Not a valid number")
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    question,
		Validate: validateDecimal,
	}

	result, _ := prompt.Run()
	return cast.ToFloat64(result)
}

// AskForInput prompts with a question expecting a string answer. The
// function returns only when a valid string value is entered
//
//   name := bobino.AskForInput("What's your name?")
//   fmt.Printf("Nice to meet you, %s!", name)
func AskForInput(question string) (answer string) {
	prompt := promptui.Prompt{
		Label: question,
	}

	answer, _ = prompt.Run()
	return answer
}

// AskWithChoice prompts with a question expecting an answer from a list of options. The
// function returns only when a valid option is selected
//
//   animal := bobino.AskWithChoice("What's your favorite animal?", "cat", "dog", "horse", "raven")
//   fmt.Printf("I love a nice %s too!", animal)
func AskWithChoice(question string, options ...string) (answer string) {
	prompt := promptui.Select{
		Label: question,
		Items: options,
	}

	_, answer, _ = prompt.Run()
	return answer
}
