package utils

import (
	"app/models/projects"
	"app/models/tags"
	"sort"
	"time"
)

type TagInterface struct {
	ID        uint64    `json:"id"`
	Slug      string    `json:"slug"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Count     int       `json:"count"`
}

type TagData struct {
	ID        uint64
	Index     int
	Title     string
	Slug      string
	CreatedAt string
	UpdatedAt string
}

type ProjectTagData struct {
	Title   string
	Slug    string
	Checked bool
}

// Given a tag, check it exists amongst an array of tags
func CheckTagExistsInTags(needle *tags.Tag, haystack []*tags.Tag) bool {
	for _, needleInHay := range haystack {
		if needleInHay.ID == needle.ID {
			return true
		}
	}
	return false
}

func TagOccurrencesInProjects(projects []*projects.Project, tag *tags.Tag) int {
	count := 0
	for _, project := range projects {
		for _, projectTag := range project.Tags {
			if projectTag.ID == tag.ID {
				count = count + 1
			}
		}
	}

	return count
}

func TagsOccurrencesInProjects(projects []*projects.Project) []*TagInterface {
	var tags []*TagInterface
	tagsMap := make(map[string]*TagInterface)
	for _, project := range projects {
		for _, tag := range project.Tags {
			if tagInterface, ok := tagsMap[tag.Slug]; ok {
				tagInterface.Count = tagInterface.Count + 1
			} else {
				tagsMap[tag.Slug] = &TagInterface{
					ID:        tag.ID,
					Slug:      tag.Slug,
					Title:     tag.Title,
					CreatedAt: tag.CreatedAt,
					UpdatedAt: tag.UpdatedAt,
					Count:     1,
				}
			}
		}
	}

	keys := make([]string, 0, len(tagsMap))
	for key := range tagsMap {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	for _, key := range keys {
		tags = append(tags, tagsMap[key])
	}

	return tags
}

func GetTagsData() ([]*TagData, error) {
	tags, t_err := tags.All()
	if t_err != nil {
		return nil, t_err
	}

	var tag_data []*TagData
	zone, _ := time.LoadLocation("Australia/Perth")

	for i, tag := range tags {
		tag_data = append(tag_data, &TagData{
			ID:        tag.ID,
			Index:     i + 1,
			Title:     tag.Title,
			Slug:      tag.Slug,
			CreatedAt: tag.CreatedAt.In(zone).Format(time.RFC822),
			UpdatedAt: tag.UpdatedAt.In(zone).Format(time.RFC822),
		})
	}

	return tag_data, nil
}

func GetTagData(id string) (*TagData, error) {
	tag, not_found := tags.Find(id)

	if not_found != nil {
		return nil, not_found
	}

	tag_data := &TagData{
		ID:    tag.ID,
		Title: tag.Title,
		Slug:  tag.Slug,
	}

	return tag_data, nil
}
