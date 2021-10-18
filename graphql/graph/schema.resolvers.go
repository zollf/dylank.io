package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"api/graph/generated"
	"api/graph/model"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func (r *queryResolver) Projects(ctx context.Context, tags []*string) ([]*model.Project, error) {
	filename := "graph/projects.json"
	jsonFile, err := os.Open(filename)
	fmt.Println(err)
	if err != nil {
		panic(fmt.Errorf(fmt.Sprintf("Error: could not import %q", filename)))
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var projects []*model.Project

	json.Unmarshal(byteValue, &projects)

	return projects, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
