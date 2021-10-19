// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Project struct {
	ID          string  `json:"id" bson:"_id"`
	Slug        string  `json:"slug" bson:"slug"`
	Title       string  `json:"title" bson:"title"`
	Description string  `json:"description" bson:"description"`
	Image       string  `json:"image" bson:"image"`
	URL         *string `json:"url" bson:"url"`
	Git         *string `json:"git" bson:"git"`
	Tags        []*Tag  `json:"tags" bson:"tags"`
}

type Tag struct {
	ID    string `json:"id" bson:"_id"`
	Slug  string `json:"slug" bson:"slug"`
	Title string `json:"title" bson:"title"`
}

type TagInterface struct {
	Tag   *Tag `json:"tag" bson:"tag"`
	Total int  `json:"total" bson:"total"`
}