package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"api/graph/generated"
	"api/graph/model"
	"context"
	"sort"
)

func (r *queryResolver) Projects(ctx context.Context, tags []*string) ([]*model.Project, error) {
	return ParseProjects(), nil
}

func (r *queryResolver) Tags(ctx context.Context) ([]*model.TagInterface, error) {
	var tagInterface []*model.TagInterface
	var tagMap = make(map[string]*model.TagInterface)
	for _, project := range ParseProjects() {
		for _, tag := range project.Tags {
			if val, ok := tagMap[tag.Slug]; ok {
				val.Total = val.Total + 1
			} else {
				tagMap[tag.Slug] = &model.TagInterface{
					Tag:   tag,
					Total: 1,
				}
			}
		}
	}

	for _, element := range tagMap {
		tagInterface = append(tagInterface, element)
	}

	sort.SliceStable(tagInterface, func(i int, j int) bool {
		return tagInterface[i].Total > tagInterface[j].Total
	})

	return tagInterface, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
