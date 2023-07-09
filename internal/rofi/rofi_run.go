package rofi

import (
	"bufio"
	"log"
	"os/exec"
	"strings"
)

type RofiItem struct {
	Label string
	Value string
}

type ByLabel []RofiItem

func (items ByLabel) Len() int           { return len(items) }
func (items ByLabel) Swap(i, j int)      { items[i], items[j] = items[j], items[i] }
func (items ByLabel) Less(i, j int) bool { return strings.Compare(items[i].Label, items[j].Label) < 0 }

type RofiMenuOptions struct {
	Title    string
	Items    []RofiItem
	Theme    string
	OnSelect func(label string, value string)
}

func LaunchMenu(options RofiMenuOptions) {
	lines := ""
	for idx, item := range options.Items {
		lines += item.Label
		if idx < len(options.Items)-1 {
			lines += "\n"
		}
	}

	cmd1 := exec.Command("echo", lines)

	cmd2 := exec.Command(
		"rofi",
		"-dmenu",
		"-i",
		"-theme", options.Theme,
		"-p", options.Title,
	)

	cmd2.Stdin, _ = cmd1.StdoutPipe()

	cmd2StdoutPipe, _ := cmd2.StdoutPipe()
	scanner := bufio.NewScanner(cmd2StdoutPipe)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Start both commands
	if err := cmd1.Start(); err != nil {
		log.Fatal(err)
	}
	if err := cmd2.Start(); err != nil {
		log.Fatal(err)
	}

	if err := cmd1.Wait(); err != nil {
		log.Fatal(err)
	}

	for scanner.Scan() {
		rofiOutput := scanner.Text()
		
		selectedItem := findSelectedItem(rofiOutput, &options.Items)

		options.OnSelect(
			selectedItem.Label,
			selectedItem.Value,
		)
	}

	if err := cmd2.Wait(); err != nil {
		log.Fatal(err)
	}
}

func findSelectedItem(choiche string, items *[]RofiItem) *RofiItem {
	for _, item := range *items {
		if item.Label == choiche {
			return &item
		}
	}

	return nil
}
