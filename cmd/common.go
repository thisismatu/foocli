/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bytes"
	"log"
	"os"
	"strings"

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
		log.Fatal(err)
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
