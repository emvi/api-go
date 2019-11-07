package emvi

import (
	"time"
)

type BaseEntity struct {
	Id      string    `json:"id"`
	DefTime time.Time `json:"def_time"`
	ModTime time.Time `json:"mod_time"`
}

type Article struct {
	BaseEntity
	OrganizationId       string          `json:"organization_id"`
	Views                uint            `json:"views"`
	WIP                  int             `json:"wip"`
	Archived             string          `json:"archived"`
	Published            time.Time       `json:"published"`
	Pinned               bool            `json:"pinned"`
	LatestArticleContent *ArticleContent `json:"latest_article_content"`
	Tags                 []Tag           `json:"tags"`
	PreviewImage         string          `json:"preview_image"`
}

type ArticleContent struct {
	BaseEntity
	Title      string `json:"title"`
	Content    string `json:"content"`
	Version    int    `json:"version"`
	Commit     string `json:"commit"`
	WIP        bool   `json:"wip"`
	ArticleId  string `json:"article_id"`
	LanguageId string `json:"language_id"`
	UserId     string `json:"user_id"` // user who created this commit
	Authors    []User `json:"authors"`
}

type Tag struct {
	BaseEntity
	OrganizationId string `json:"organization_id"`
	Name           string `json:"name"`
	Usages         int    `json:"usages"`
}

type User struct {
	BaseEntity
	Email              string              `json:"email"`
	Firstname          string              `json:"firstname"`
	Lastname           string              `json:"lastname"`
	Language           string              `json:"language"`
	Info               string              `json:"info"`
	Picture            string              `json:"picture"`
	OrganizationMember *OrganizationMember `json:"organization_member"`
}

type OrganizationMember struct {
	BaseEntity
	OrganizationId string `json:"organization_id"`
	UserId         string `json:"user_id"`
	LanguageId     string `json:"language_id"`
	Username       string `json:"username"`
	Phone          string `json:"phone"`
	Mobile         string `json:"mobile"`
	Info           string `json:"info"`
	User           *User  `json:"user"`
}

type Organization struct {
	BaseEntity
	Name             string `json:"name"`
	NameNormalized   string `json:"name_normalized"`
	Picture          string `json:"picture"`
	Expert           bool   `json:"expert"`
	CreateGroupAdmin bool   `json:"create_group_admin"`
	CreateGroupMod   bool   `json:"create_group_mod"`
	MemberCount      int    `json:"member_count"`
	ArticleCount     int    `json:"article_count"`
}

type Language struct {
	BaseEntity
	OrganizationId string `json:"organization_id"`
	Name           string `json:"name"`
	Code           string `json:"code"`
	Default        bool   `json:"default"`
}
