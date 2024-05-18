package creategame

import (
	"errors"
	"unicode"

	"github.com/charmbracelet/huh"
	"github.com/maria-mz/bash-battle/tui/constants"
	"github.com/maria-mz/bash-battle/utils"
)

const (
	roundsKey    = "rounds"
	roundMinsKey = "roundMinutes"
	confirmKey   = "done"
)

func isInputEmpty(s string) bool {
	return len(s) == 0
}

func isInputNumeric(s string) bool {
	for _, char := range s {
		if !unicode.IsDigit(char) {
			return false
		}
	}
	return true
}

func isGreaterThanZero(s string) (bool, error) {
	v, err := utils.StringToInt(s)
	if err != nil {
		return false, err
	}
	return v > 0, nil
}

func validateNumericInput(s string) error {
	if isInputEmpty(s) {
		return errors.New("field cannot be empty")
	}

	if !isInputNumeric(s) {
		return errors.New("not a valid number")
	}

	isGreater, _ := isGreaterThanZero(s)

	if !isGreater {
		return errors.New("must be greater than zero")
	}
	return nil
}

func getFormTheme() *huh.Theme {
	t := huh.ThemeBase()

	t.Focused.Title = t.Focused.Title.Foreground(constants.BlueColor)
	t.Focused.ErrorMessage = t.Focused.ErrorMessage.Foreground(constants.RedColor)
	t.Focused.Description = t.Focused.Description.Foreground(constants.GrayColor)

	t.Blurred.ErrorMessage = t.Focused.ErrorMessage
	t.Blurred.Description = t.Focused.Description

	return t
}

func newForm() *huh.Form {
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Key(roundsKey).
				Title("Number of rounds").
				Placeholder("Enter a number").
				Validate(validateNumericInput).
				CharLimit(2),
			huh.NewInput().
				Key(roundMinsKey).
				Title("Round duration (minutes)").
				Placeholder("Enter a number").
				Validate(validateNumericInput).
				CharLimit(2),
			huh.NewConfirm().
				Key(confirmKey).
				Title("Create game?").
				Description(
					"Selecting 'Yes' should create a new game on the server. \n"+
						"You should shortly receive a code that others can use "+
						"to join the game.",
				).
				Affirmative("Yes").
				Negative("No"),
		).
			WithShowHelp(true),
	)

	form.WithTheme(getFormTheme())

	return form
}

func getRounds(form *huh.Form) string {
	return form.GetString(roundsKey)
}

func getRoundMinutes(form *huh.Form) string {
	return form.GetString(roundMinsKey)
}

func wantsToCreateGame(form *huh.Form) bool {
	return form.GetBool(confirmKey)
}
