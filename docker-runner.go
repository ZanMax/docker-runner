package main

import (
	"encoding/json"
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"moul.io/banner"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
)

type ContainersConfigs struct {
	Containers []struct {
		Name     string `json:"name"`
		Command  string `json:"command"`
		DataPath string `json:"data_path"`
		Notes    string `json:"notes"`
	} `json:"containers"`
}

func main() {
	appDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	checkError(err)
	configPath := path.Join(appDir, "config.json")
	data, err := os.ReadFile(configPath)
	checkError(err)

	var config ContainersConfigs

	err = json.Unmarshal(data, &config)
	checkError(err)

	dockerConfig := map[string]string{}
	var options []string

	for i := 0; i < len(config.Containers); i++ {
		options = append(options, config.Containers[i].Name)
		dockerConfig[config.Containers[i].Name] = config.Containers[i].Command
	}
	recipes := getDirsList()
	if len(recipes) > 0 {
		options = append(options, "Additional")
	}
	options = append(options, "PRUNE")
	options = append(options, "Exit")

	command("clear")

	fmt.Println(banner.Inline("docker runner"))

	var qs = []*survey.Question{
		{
			Name: "Docker",
			Prompt: &survey.Select{
				Message: "Choose:",
				Options: options,
				Default: options[0],
			},
		},
	}

	answers := struct {
		Docker string `survey:"Docker"`
	}{}

	for {
		err := survey.Ask(qs, &answers, survey.WithPageSize(len(options)))
		checkError(err)

		if dockerConfig[answers.Docker] == "Exit" {
			os.Exit(0)
		} else if dockerConfig[answers.Docker] == "Additional" {
			var additionalOptions []string
			additionalOptions = append(additionalOptions, getDirsList()...)
			var qsAdditional = []*survey.Question{
				{
					Name: "Recipe",
					Prompt: &survey.Select{
						Message: "Choose:",
						Options: additionalOptions,
						Default: additionalOptions[0],
					},
				},
			}
			answersAdditional := struct {
				Docker string `survey:"Docker"`
			}{}
			err = survey.Ask(qsAdditional, &answersAdditional, survey.WithPageSize(len(options)))
			checkError(err)
		} else {
			command("clear")
			fmt.Println(answers.Docker, " starting ... ")
			command(dockerConfig[answers.Docker])
			command("docker ps")
		}
	}
}

func command(command string) {
	cmd := exec.Command("/bin/sh", "-c", command)
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	checkError(err)
}

func getDirsList() []string {
	file, err := os.Open("recipes")
	checkError(err)
	defer file.Close()
	dirs, err := file.Readdirnames(0)
	checkError(err)
	var dirsList []string
	for _, dir := range dirs {
		if !strings.HasPrefix(dir, ".") {
			dirsList = append(dirsList, dir)
		}
	}
	return dirsList
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
