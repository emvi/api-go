package emvi

import (
	"net/url"
	"strconv"
	"strings"
	"time"
)

const (
	DateFormat = "2006-01-02"
)

type Filter interface {
	addParams(values *url.Values)
}

type BaseSearch struct {
	CreatedStart time.Time `json:"created_start"`
	CreatedEnd   time.Time `json:"created_end"`
	UpdatedStart time.Time `json:"updated_start"`
	UpdatedEnd   time.Time `json:"updated_end"`
	SortCreated  string    `json:"sort_created"`
	SortUpdated  string    `json:"sort_updated"`
	Offset       int       `json:"offset"`
	Limit        int       `json:"limit"`
}

func (filter *BaseSearch) addParams(query *url.Values) {
	query.Add("created_start", filter.CreatedStart.Format(DateFormat))
	query.Add("created_end", filter.CreatedEnd.Format(DateFormat))
	query.Add("updated_start", filter.UpdatedStart.Format(DateFormat))
	query.Add("updated_end", filter.UpdatedEnd.Format(DateFormat))
	query.Add("sort_created", filter.SortCreated)
	query.Add("sort_updated", filter.SortUpdated)
	query.Add("offset", strconv.Itoa(filter.Offset))
	query.Add("limit", strconv.Itoa(filter.Limit))
}

type ArticleFilter struct {
	BaseSearch

	Archived         bool      `json:"archived"`
	WIP              bool      `json:"wip"`
	ClientAccess     bool      `json:"client_access"`
	Preview          bool      `json:"preview"`
	PreviewParagraph bool      `json:"preview_paragraph"`
	PreviewImage     bool      `json:"preview_image"`
	Title            string    `json:"title"`
	Content          string    `json:"content"`
	Tags             string    `json:"tags"`
	TagIds           []string  `json:"tag_ids"`
	AuthorUserIds    []string  `json:"authors"`
	Commits          string    `json:"commits"`
	PublishedStart   time.Time `json:"published_start"`
	PublishedEnd     time.Time `json:"published_end"`
	SortTitle        string    `json:"sort_title"`
	SortPublished    string    `json:"sort_published"`
}

func (filter *ArticleFilter) addParams(query *url.Values) {
	filter.BaseSearch.addParams(query)
	query.Add("archived", boolToString(filter.Archived))
	query.Add("wip", boolToString(filter.WIP))
	query.Add("client_access", boolToString(filter.ClientAccess))
	query.Add("preview", boolToString(filter.Preview))
	query.Add("preview_paragraph", boolToString(filter.PreviewParagraph))
	query.Add("preview_image", boolToString(filter.PreviewImage))
	query.Add("title", filter.Title)
	query.Add("content", filter.Content)
	query.Add("tags", filter.Tags)
	query.Add("tag_ids", sliceToString(filter.TagIds))
	query.Add("authors", sliceToString(filter.AuthorUserIds))
	query.Add("commits", filter.Commits)
	query.Add("published_start", filter.PublishedStart.Format(DateFormat))
	query.Add("published_end", filter.PublishedEnd.Format(DateFormat))
	query.Add("sort_title", filter.SortTitle)
	query.Add("sort_published", filter.SortPublished)
}

func boolToString(b bool) string {
	if b {
		return "true"
	}

	return "false"
}

func sliceToString(slice []string) string {
	return strings.Join(slice, ",")
}
