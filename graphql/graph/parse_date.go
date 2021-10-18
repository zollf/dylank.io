package graph

import (
	"api/graph/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func ParseProjects() []*model.Project {
	var projects []*model.Project
	filename := "data/projects.json"
	jsonFile, err := os.Open(filename)

	if err != nil {
		panic(fmt.Errorf(fmt.Sprintf("Error: could not import %q", filename)))
	}

	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &projects)

	return projects
}
