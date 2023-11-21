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
	PostTagToUser(userId int, tag int) error
	PostTagToTeam(teamId int, tag int) error
	PostTagToEvent(eventId int, tag int) error
	DeleteTagFromUser(userId, tagId int) error
	DeleteTagFromTeam(teamId, tagId int) error
	DeleteTagFromEvent(eventId, tagId int) error
}

type TagAllResponse struct {
	Tags []Tag `json:"tags"`
}

type PostTagsRequest struct {
	Tags []int `json:"tags"`
}

type TagUsecase interface {
	GetAll() ([]Tag, error)
	GetByUserId(id int) ([]Tag, error)
	GetByTeamId(id int) ([]Tag, error)
	GetByEventId(id int) ([]Tag, error)
	GetByGlobalTagId(id int) ([]Tag, error)
	PostTagsToUser(id int, tags []int) error
	PostTagsToTeam(id int, tags []int) error
	PostTagsToEvent(id int, tags []int) error
	DeleteTagsFromUser(request PostTagsRequest, userId int) error
	DeleteTagsFromTeam(request PostTagsRequest, teamId int) error
	DeleteTagsFromEvent(request PostTagsRequest, eventId int) error
}
