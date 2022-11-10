package cmd

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/chzyer/readline"
	"github.com/fatih/color"
	"github.com/wlredeye/jsonlines"
	"golang.org/x/exp/slices"
)

type Project struct {
	Name string
	Id   string
}

type Application struct {
	Name     string
	Id       string
	Status   string
	Deployed string
}

type NoBellStdout struct{}

func getProjects() []Project {
	buf, err := os.ReadFile(dbProjects)
	if err != nil {
		log.Fatal(err)
	}
	projects := []Project{}
	err = jsonlines.Decode(strings.NewReader(string(buf)), &projects)
	if err != nil {
		log.Fatal(err)
	}
	return projects
}

func addProject(p Project) {
	projects := getProjects()
	projects = append(projects, p)
	var buf bytes.Buffer
	err := jsonlines.Encode(&buf, &projects)
	if err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile(dbProjects, buf.Bytes(), 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func getCurrentProject() Project {
	data, err := os.ReadFile(cfgFile)
	if err != nil {
		cyan := color.New(color.FgCyan).SprintFunc()
		fmt.Printf("Log in by running %s\n", cyan("foo login"))
		os.Exit(0)
	}
	projects := getProjects()
	idx := slices.IndexFunc(projects, func(p Project) bool { return p.Id == string(data) })
	return projects[idx]
}

func setCurrentProject(id string) {
	buf := []byte(id)
	if err := os.WriteFile(cfgFile, buf, 0644); err != nil {
		log.Fatal(err)
	}
}

func getApplications() []Application {
	buf, err := os.ReadFile(dbApplications)
	if err != nil {
		log.Fatal(err)
	}
	apps := []Application{}
	err = jsonlines.Decode(strings.NewReader(string(buf)), &apps)
	if err != nil {
		log.Fatal(err)
	}
	return apps
}

// Disable terminal bell https://github.com/manifoldco/promptui/issues/49#issuecomment-1012640880

var noBellStdout = &NoBellStdout{}

func (n *NoBellStdout) Write(p []byte) (int, error) {
	if len(p) == 1 && p[0] == readline.CharBell {
		return 0, nil
	}
	return readline.Stdout.Write(p)
}

func (n *NoBellStdout) Close() error {
	return readline.Stdout.Close()
}
