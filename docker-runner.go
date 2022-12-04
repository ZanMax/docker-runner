package main

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"moul.io/banner"
	"os"
	"os/exec"
)

func main() {
	dockerConfig := map[string]string{
		"MySql":      "sudo docker run --rm -d -p 3306:3306 --name=mysql -v ${DATA_FOLDER}mysql_data:/var/lib/mysql --env=\"MYSQL_ROOT_PASSWORD=docker\" mysql mysqld --default-authentication-plugin=mysql_native_password",
		"Postgres":   "docker ps",
		"Mongo":      "docker ps",
		"RabbitMQ":   "docker ps",
		"Redis":      "docker ps",
		"Memcached":  "docker ps",
		"Additional": "Additional",
		"PRUNE":      "docker system prune -a",
		"Exit":       "Exit",
	}

	command("clear")

	fmt.Println(banner.Inline("docker runner"))
	var options []string
	options = append(options, "MySql", "Postgres", "Mongo", "RabbitMQ", "Redis", "Memcached", "Additional", "PRUNE", "Exit")
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
			fmt.Println("Additional")
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

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
