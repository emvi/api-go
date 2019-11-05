package emvi

import (
	"time"
)

type BaseEntity struct {
	Id      string    `json:"id"`
	DefTime time.Time `json:"def_time"`
	ModTime time.Time `json:"mod_time"`
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
