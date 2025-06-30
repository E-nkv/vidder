package cli

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/AlecAivazis/survey/v2"
)

// PromptSelect displays a select prompt and returns the selected string.

type LabelValue struct {
	Label string
	Value string
}

func PromptSelectWithValues(message string, LabVals []LabelValue) (string, error) {
	// Extract display labels (keys) for the prompt
	labels := make([]string, 0, len(LabVals))
	for _, LabVal := range LabVals {
		labels = append(labels, LabVal.Label)
	}

	var selectedLabel string
	prompt := &survey.Select{
		Message: message,
		Options: labels,
	}
	err := survey.AskOne(prompt, &selectedLabel)
	if err != nil {
		return "", err
	}

	// Map the selected label back to the actual value

	for _, LabVal := range LabVals {
		if LabVal.Label == selectedLabel {
			return LabVal.Value, nil
		}
	}

	return "", fmt.Errorf("couldnt find value")
}
func PromptString(message string) string {
	var input string

	prompt := &survey.Input{
		Message: message,
	}
	err := survey.AskOne(prompt, &input)
	if err != nil {
		log.Fatal(err)
	}
	return input
}

// PromptBool displays a yes/no prompt and returns the boolean answer.
func PromptBool(message string) (bool, error) {
	var answer bool
	prompt := &survey.Confirm{
		Message: message,
	}
	err := survey.AskOne(prompt, &answer)
	return answer, err
}

// PromptIntRange prompts for an integer within a range.
func PromptIntRange(message string, min, max int) int {
	for {
		var input string
		prompt := &survey.Input{
			Message: message,
		}
		err := survey.AskOne(prompt, &input)
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)
		val, err := strconv.Atoi(input)
		if err == nil && val >= min && val <= max {
			return val
		}
		fmt.Printf("😡☠️ Please enter a number between %d and %d\n", min, max)
	}
}
