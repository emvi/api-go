package emvi

import (
	"net/url"
	"strconv"
	"strings"
	"time"
)

const (
	DateFormat     = "2006-01-02"
	SortAscending  = SortDirection("asc")
	SortDescending = SortDirection("desc")
)

type SortDirection string

// Filter is used to filter results.
type Filter interface {
	addParams(values *url.Values)
}

type BaseSearch struct {
	CreatedStart time.Time     `json:"created_start"`
	CreatedEnd   time.Time     `json:"created_end"`
	UpdatedStart time.Time     `json:"updated_start"`
	UpdatedEnd   time.Time     `json:"updated_end"`
	SortCreated  SortDirection `json:"sort_created"`
	SortUpdated  SortDirection `json:"sort_updated"`
	Offset       int           `json:"offset"`
	Limit        int           `json:"limit"`
}

func (filter *BaseSearch) addParams(query *url.Values) {
	addParamIfNotEmpty(query, "created_start", dateToString(filter.CreatedStart))
	addParamIfNotEmpty(query, "created_end", dateToString(filter.CreatedEnd))
	addParamIfNotEmpty(query, "updated_start", dateToString(filter.UpdatedStart))
	addParamIfNotEmpty(query, "updated_end", dateToString(filter.UpdatedEnd))
	addParamIfNotEmpty(query, "sort_created", string(filter.SortCreated))
	addParamIfNotEmpty(query, "sort_updated", string(filter.SortUpdated))
	addParamIfNotEmpty(query, "offset", intToString(filter.Offset))
	addParamIfNotEmpty(query, "limit", intToString(filter.Limit))
}

type ArticleFilter struct {
	BaseSearch

	Archived         bool          `json:"archived"`
	WIP              bool          `json:"wip"`
	ClientAccess     bool          `json:"client_access"`
	Preview          bool          `json:"preview"`
	PreviewParagraph bool          `json:"preview_paragraph"`
	PreviewImage     bool          `json:"preview_image"`
	Title            string        `json:"title"`
	Content          string        `json:"content"`
	Tags             string        `json:"tags"`
	TagIds           []string      `json:"tag_ids"`
	AuthorUserIds    []string      `json:"authors"`
	Commits          string        `json:"commits"`
	PublishedStart   time.Time     `json:"published_start"`
	PublishedEnd     time.Time     `json:"published_end"`
	SortTitle        SortDirection `json:"sort_title"`
	SortPublished    SortDirection `json:"sort_published"`
}

func (filter *ArticleFilter) addParams(query *url.Values) {
	filter.BaseSearch.addParams(query)
	addParamIfNotEmpty(query, "archived", boolToString(filter.Archived))
	addParamIfNotEmpty(query, "wip", boolToString(filter.WIP))
	addParamIfNotEmpty(query, "client_access", boolToString(filter.ClientAccess))
	addParamIfNotEmpty(query, "preview", boolToString(filter.Preview))
	addParamIfNotEmpty(query, "preview_paragraph", boolToString(filter.PreviewParagraph))
	addParamIfNotEmpty(query, "preview_image", boolToString(filter.PreviewImage))
	addParamIfNotEmpty(query, "title", filter.Title)
	addParamIfNotEmpty(query, "content", filter.Content)
	addParamIfNotEmpty(query, "tags", filter.Tags)
	addParamIfNotEmpty(query, "tag_ids", sliceToString(filter.TagIds))
	addParamIfNotEmpty(query, "authors", sliceToString(filter.AuthorUserIds))
	addParamIfNotEmpty(query, "commits", filter.Commits)
	addParamIfNotEmpty(query, "published_start", dateToString(filter.PublishedStart))
	addParamIfNotEmpty(query, "published_end", dateToString(filter.PublishedEnd))
	addParamIfNotEmpty(query, "sort_title", string(filter.SortTitle))
	addParamIfNotEmpty(query, "sort_published", string(filter.SortPublished))
}

func addParamIfNotEmpty(query *url.Values, key, value string) {
	if value != "" {
		query.Add(key, value)
	}
}

func dateToString(date time.Time) string {
	if date.IsZero() {
		return "" // ignore in URL
	}

	return date.Format(DateFormat)
}

func boolToString(b bool) string {
	if b {
		return "true"
	}

	return "" // ignore in URL
}

func intToString(i int) string {
	if i == 0 {
		return ""
	}

	return strconv.Itoa(i)
}

func sliceToString(slice []string) string {
	return strings.Join(slice, ",")
}
