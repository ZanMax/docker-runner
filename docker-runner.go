package main

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"moul.io/banner"
	"os"
	"os/exec"
)

func main() {
	clear()
	fmt.Println(banner.Inline("docker runner"))
	var options []string
	options = append(options, "MySql", "Postgres")
	var qs = []*survey.Question{
		{
			Name: "server",
			Prompt: &survey.Select{
				Message: "Choose server:",
				Options: options,
				Default: options[0],
			},
		},
	}

	err := survey.Ask(qs, &options, survey.WithPageSize(len(options)))
	checkError(err)

}

func clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
