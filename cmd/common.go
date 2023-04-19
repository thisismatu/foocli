package cmd

import (
	"bytes"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/chzyer/readline"
	"github.com/fatih/color"
	"github.com/google/uuid"
	"github.com/juju/ansiterm"
	"github.com/kaduartur/go-cli-spinner/pkg/spinner"
	"github.com/wlredeye/jsonlines"
	"golang.org/x/exp/slices"
)

type Project struct {
	Name        string
	Id          string
	PaymentPlan string
	UserCount   int
}

type Model struct {
	Name        string
	Id          string
	ProjectId   string
	Language    string
	Status      string
	IsAdaptable bool
	BaseModel   string
}

type Deployment struct {
	Date     time.Time
	Url      string
	Status   string
	Duration string
}

type NoBellStdout struct{}

func logError(err error) {
	fmt.Printf("%s %v\n", color.RedString("Error:"), err)
	os.Exit(0)
}

func logWarning(msg string) {
	fmt.Printf("%s %s\n", color.YellowString("Warning:"), msg)
}

func logSuccess(msg string) {
	fmt.Printf("%s %s\n", color.GreenString("Success!"), msg)
}

func faint(str string) string {
	return color.New(color.Faint).Sprint(str)
}

func fmtCmd(cmd string) string {
	backtick := faint("`")
	command := color.CyanString(cmd)
	return backtick + command + backtick
}

func fmtExample(desc string, cmd string, isLast bool) string {
	nl := "\n\n"
	if isLast {
		nl = "\n"
	}
	return fmt.Sprintf("  %s\n  %s %s %s", desc, faint("$"), color.CyanString(cmd), nl)
}

func getProjects() []Project {
	buf, err := os.ReadFile(dbProjects)
	if err != nil {
		logError(err)
	}
	projects := []Project{}
	err = jsonlines.Decode(strings.NewReader(string(buf)), &projects)
	if err != nil {
		logError(err)
	}
	return projects
}

func getCurrentProject() Project {
	data, err := os.ReadFile(cfgFile)
	if err != nil || len(data) == 0 {
		logWarning("you are not logged in")
		fmt.Printf("To log in run %s\n", fmtCmd("foo login"))
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
		logError(err)
	}
}

func addProject(p Project) {
	projects := getProjects()
	projects = append(projects, p)
	var buf bytes.Buffer
	err := jsonlines.Encode(&buf, &projects)
	if err != nil {
		logError(err)
	}
	err = os.WriteFile(dbProjects, buf.Bytes(), 0644)
	if err != nil {
		logError(err)
	}
}

func getAdaptedModels(pId string) []Model {
	buf, err := os.ReadFile(dbModels)
	if err != nil {
		logError(err)
	}
	models := []Model{}
	err = jsonlines.Decode(strings.NewReader(string(buf)), &models)
	if err != nil {
		logError(err)
	}
	filteredModels := []Model{}
	for i := range models {
		if models[i].ProjectId == pId {
			filteredModels = append(filteredModels, models[i])
		}
	}
	return filteredModels
}

func getBaseModels() []Model {
	buf, err := os.ReadFile(dbModels)
	if err != nil {
		logError(err)
	}
	models := []Model{}
	err = jsonlines.Decode(strings.NewReader(string(buf)), &models)
	if err != nil {
		logError(err)
	}
	filteredModels := []Model{}
	for i := range models {
		if models[i].ProjectId == "all" {
			filteredModels = append(filteredModels, models[i])
		}
	}
	return filteredModels
}

func getAllModels(pId string) []Model {
	baseModels := getBaseModels()
	adaptedModels := getAdaptedModels(pId)
	models := append(baseModels, adaptedModels...)
	return models
}

func getModel(mId string) (Model, error) {
	buf, err := os.ReadFile(dbModels)
	if err != nil {
		logError(err)
	}
	models := []Model{}
	err = jsonlines.Decode(strings.NewReader(string(buf)), &models)
	if err != nil {
		logError(err)
	}
	for i := range models {
		if models[i].Id == mId {
			return models[i], nil
		}
	}
	errMsg := fmt.Sprintf("model '%s' does not exist", mId)
	return Model{}, errors.New(errMsg)
}

func addModel(m Model) {
	models := getAllModels(m.ProjectId)
	models = append(models, m)
	var buf bytes.Buffer
	err := jsonlines.Encode(&buf, &models)
	if err != nil {
		logError(err)
	}
	err = os.WriteFile(dbModels, buf.Bytes(), 0644)
	if err != nil {
		logError(err)
	}
}

func removeModel(m Model) {
	models := getAllModels(m.ProjectId)
	for i := len(models) - 1; i >= 0; i-- {
		if models[i].Id == m.Id {
			models = append(models[:i], models[i+1:]...)
		}
	}
	var buf bytes.Buffer
	err := jsonlines.Encode(&buf, &models)
	if err != nil {
		logError(err)
	}
	err = os.WriteFile(dbModels, buf.Bytes(), 0644)
	if err != nil {
		logError(err)
	}
}

func printModelInfo(m Model) {
	writer := ansiterm.NewTabWriter(os.Stdout, 0, 8, 2, '\t', 0)
	sc := color.New(statusColor(m.Status)).SprintFunc()

	fmt.Println()
	fmt.Fprintf(writer, "  %s\t%s\n", faint("Name"), m.Name)
	fmt.Fprintf(writer, "  %s\t%s\n", faint("Model ID"), m.Id)
	fmt.Fprintf(writer, "  %s\t%s\n", faint("Language"), m.Language)
	if m.ProjectId == "all" {
		fmt.Fprintf(writer, "  %s\t%s\n", faint("Description"), "Model description goes here. It should briefly describe the model characteristics.")
	} else {
		fmt.Fprintf(writer, "  %s\t%s\n", faint("Base model"), m.BaseModel)
	}
	fmt.Fprintf(writer, "  %s\t%s\n", faint("Deployed"), time.Now().Local().String())
	fmt.Fprintf(writer, "  %s\t%s %s\n", faint("Status"), sc("●"), m.Status)
	writer.Flush()
	fmt.Println()
}

func getDeployments(mId string) []Deployment {
	deployments := []Deployment{}
	for x := 0; x < 10; x++ {
		hash := uuid.New().String()[0:8]
		url := fmt.Sprintf("https://dashboard.foo.com/model/%s/%s", mId, hash)
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

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}
