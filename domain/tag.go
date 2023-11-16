package domain

import "github.com/gin-gonic/gin"

const (
	TableTag = "tag"
)

type Tag struct {
	TagID       int    `db:"tag_id" json:"tag_id"`
	Activity    string `db:"activity" json:"activity"`
	GlobalTagID int    `db:"global_tag_id" json:"global_tag_id"`
}

type TagRepository interface {
	CreateTag(c *gin.Context) error
	GetByUserId(id int) ([]Tag, error)
	GetById(id int) (Tag, error)
	GetByTeamId(id int) ([]Tag, error)
	GetAll() ([]Tag, error)
	GetByGlobalTagId(id int) ([]Tag, error)
	GetByEventId(id int) ([]Tag, error)
	//AddTagToUser(c *gin.Context) error
	//AddTagToTeam(c *gin.Context) error
	//AddTagToEvent(c *gin.Context) error
	GetUserTagCount(userId, tagId int) (int, error)
	GetTeamTagCount(teamId, tagId int) (int, error)
	GetEventTagCount(eventId, tagId int) (int, error)
	PostTagToUser(userId int, tag Tag) error
	PostTagToTeam(teamId int, tag Tag) error
	PostTagToEvent(eventId int, tag Tag) error
}

type TagAllResponse struct {
	Tags []Tag `json:"tags"`
}

type PostTagsRequest struct {
	Tags []Tag `json:"tags"`
}

type TagUsecase interface {
	GetAll() ([]Tag, error)
	GetByUserId(id int) ([]Tag, error)
	GetByTeamId(id int) ([]Tag, error)
	GetByEventId(id int) ([]Tag, error)
	GetByGlobalTagId(id int) ([]Tag, error)
	PostTagsToUser(id int, tags []Tag) error
	PostTagsToTeam(id int, tags []Tag) error
	PostTagsToEvent(id int, tags []Tag) error
}
