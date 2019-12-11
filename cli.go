package main

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/manifoldco/promptui"
	"github.com/shaletjohnson/christmas-swap/giftassignment"
)

const (
	makeGroup       = "Add the name of your group"
	viewAssignments = "View current Assignments"
	finish          = "Done adding names"
	addPeople       = "Add a new person to your group"
)

func main() {
	//add for loop for it to return to main func when it ends prompts to add people
	fmt.Println()

	//create variable to call newgroup and store values to append

	for {
		prompt := promptui.Select{
			Label: "Welcome to the sibling gift exchange assignor!",
			Items: []string{
				makeGroup,
				viewAssignments,
			},
		}

		_, result, err := prompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}
		fmt.Println(result)

		switch result {
		case makeGroup:
			err := newGroupPrompt()
			if err != nil {
				fmt.Printf("Prompt failed %v\n", err)
				return
			}

		case viewAssignments:
			err := viewAssignmentsPrompt(giftassignment.GetGroup())
			if err != nil {
				fmt.Printf("Prompt failed %v\n", err)
				return
			}
		}
	}
}

func newGroupPrompt() error {
	namePrompt := promptui.Prompt{
		Label: "Name of your Group",
	}

	name, err := namePrompt.Run()
	if err != nil {
		return err
	}

	budget, err := promptFloat("Set Budget")
	if err != nil {
		return err
	}

	newGroup := &giftassignment.Group{
		Name:   name,
		Budget: budget,
	}

	giftassignment.SetGroup(newGroup)

	fmt.Printf("You have added the Group %vwith a budget of $%v", name, budget)

	err = peoplePrompt(newGroup)
	if err != nil {
		return err
	}

	newGroup.Assign()

	return nil
}

func viewAssignmentsPrompt(group *giftassignment.Group) error {

	availableAssignments := group.ListAssignments()
	fmt.Println(availableAssignments)

	if len(availableAssignments) == 0 {
		fmt.Println("No assignments have been made yet!")
		return nil
	}
	return nil
}

func peoplePrompt(newGroup *giftassignment.Group) error {
	//giftassignment.AddPerson()
	for {
		fmt.Println()

		prompt := promptui.Select{
			Label: "Select Action",
			Items: []string{
				addPeople,
				finish,
			},
		}
		_, result, err := prompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return nil
		}

		switch result {
		case addPeople:
			namePrompt := promptui.Prompt{
				Label: "Add a persons name",
			}
			name, err := namePrompt.Run()
			if err != nil {
				return err
			}
			newGroup.AddPerson(name)

		case finish:

			return nil

		}

	}

	return nil
}

// func doneAddingGroup() {
// 	return main()
// }

func promptFloat(label string) (float64, error) {
	validate := func(input string) error {
		_, err := strconv.ParseFloat(input, 64)
		if err != nil {
			return errors.New("Invalid number")
		}
		return nil
	}

	prompt := promptui.Prompt{
		Validate: validate,
		Label:    label,
	}
	inputStr, err := prompt.Run()
	if err != nil {
		return 0, err
	}
	input, err := strconv.ParseFloat(inputStr, 64)
	if err != nil {
		return 0, err
	}

	return input, nil
}
