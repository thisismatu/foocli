package cmd

import (
	"bytes"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/chzyer/readline"
	"github.com/fatih/color"
	"github.com/juju/ansiterm"
	"github.com/kaduartur/go-cli-spinner/pkg/spinner"
	"github.com/wlredeye/jsonlines"
	"golang.org/x/exp/slices"
)

type Project struct {
	Name string
	Id   string
}

type Model struct {
	Name        string
	Id          string
	ProjectId   string
	Language    string
	Status      string
	IsAdaptable bool
}

type Deployment struct {
	Date     time.Time
	Url      string
	Status   string
	Duration string
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
	if err != nil || len(data) == 0 {
		fmt.Println("You are not logged in")
		fmt.Printf("To log in run %s\n", color.CyanString("foo login"))
		os.Exit(0)
	}
	currProjectId := string(data)
	projects := getProjects()
	idx := slices.IndexFunc(projects, func(p Project) bool { return p.Id == currProjectId })
	return projects[idx]
}

func setCurrentProject(id string) {
	buf := []byte(id)
	if err := os.WriteFile(cfgFile, buf, 0644); err != nil {
		log.Fatal(err)
	}
}

func getModel(mid string) Model {
	buf, err := os.ReadFile(dbModels)
	if err != nil {
		log.Fatal(err)
	}
	models := []Model{}
	err = jsonlines.Decode(strings.NewReader(string(buf)), &models)
	if err != nil {
		log.Fatal(err)
	}
	for i := range models {
		if models[i].Id == mid {
			return models[i]
		}
	}
	return models[0]
}

func getModels(pid string) []Model {
	buf, err := os.ReadFile(dbModels)
	if err != nil {
		log.Fatal(err)
	}
	models := []Model{}
	err = jsonlines.Decode(strings.NewReader(string(buf)), &models)
	if err != nil {
		log.Fatal(err)
	}
	filteredModels := []Model{}
	for i := range models {
		if models[i].ProjectId == pid || models[i].ProjectId == "all" {
			filteredModels = append(filteredModels, models[i])
		}
	}
	return filteredModels
}

func getBaseModels() []Model {
	buf, err := os.ReadFile(dbModels)
	if err != nil {
		log.Fatal(err)
	}
	models := []Model{}
	err = jsonlines.Decode(strings.NewReader(string(buf)), &models)
	if err != nil {
		log.Fatal(err)
	}
	filteredModels := []Model{}
	for i := range models {
		if models[i].ProjectId == "all" && models[i].IsAdaptable {
			filteredModels = append(filteredModels, models[i])
		}
	}
	return filteredModels
}

func addModel(pid string, m Model) {
	models := getModels(pid)
	models = append(models, m)
	var buf bytes.Buffer
	err := jsonlines.Encode(&buf, &models)
	if err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile(dbModels, buf.Bytes(), 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func getDeployments(mid string) []Deployment {
	deployments := []Deployment{}
	for x := 0; x < 10; x++ {
		url := fmt.Sprintf("https://dashboard.foo.com/model/%s/%d", mid, x)
		duration := fmt.Sprintf("%ds", rand.Intn(1000))
		deployments = append(deployments, Deployment{Date: time.Now(), Url: url, Status: "Ready", Duration: duration})
	}
	return deployments
}

func loading(s string, t time.Duration) {
	writer := ansiterm.NewWriter(os.Stdout)
	writer.SetStyle(ansiterm.Style(2))
	sp := spinner.New(s)
	sp.Output = writer
	sp.Start()
	time.Sleep(time.Second * t)
	sp.Stop()
	writer.Reset()
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

func statusColor(status string) color.Attribute {
	switch status {
	case "Ready":
		return color.FgGreen
	case "Training":
		return color.FgYellow
	case "Queued":
		return color.FgYellow
	case "Failed":
		return color.FgRed
	default:
		return color.Faint
	}
}
