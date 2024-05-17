package creategameform

import (
	"errors"
	"strconv"
	"unicode"

	"github.com/charmbracelet/huh"
	"github.com/maria-mz/bash-battle/tui/constants"
)

const (
	ROUNDS_KEY     = "rounds"
	ROUND_MINS_KEY = "roundMinutes"
	CONFIRM_KEY    = "done"
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
	n, err := strconv.Atoi(s)
	if err != nil {
		return false, err
	}
	return n > 0, nil
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
				Key(ROUNDS_KEY).
				Title("Number of rounds").
				Placeholder("Enter a number").
				Validate(validateNumericInput).
				CharLimit(2),
			huh.NewInput().
				Key(ROUND_MINS_KEY).
				Title("Round duration (minutes)").
				Placeholder("Enter a number").
				Validate(validateNumericInput).
				CharLimit(2),
			huh.NewConfirm().
				Key(CONFIRM_KEY).
				Title("Create game?").
				Description(
					"Selecting 'Yes' should create a new game on the server. \n"+
						"You should shortly receive a code that others can use "+
						"to join the game.",
				).
				Affirmative("Yes").
				Negative("No"),
		).
			WithShowHelp(false),
	)

	form.WithTheme(getFormTheme())

	return form
}

func getRounds(form *huh.Form) string {
	return form.GetString(ROUNDS_KEY)
}

func getRoundMinutes(form *huh.Form) string {
	return form.GetString(ROUND_MINS_KEY)
}

func wantsToCreateGame(form *huh.Form) bool {
	return form.GetBool(CONFIRM_KEY)
}
