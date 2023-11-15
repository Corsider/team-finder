package repository

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"team-finder/domain"
	"team-finder/postgres"
)

type tagRepository struct {
	database postgres.Database //*sql.DB
	table    string
}

func NewTagRepository(db *sql.DB, table string) domain.TagRepository {
	return &tagRepository{
		database: &postgres.PostgresDB{DB: db},
		table:    table,
	}
}

func (t *tagRepository) GetByUserId(id int) ([]domain.Tag, error) {
	//TODO implement me
	panic("implement me")
}

func (t *tagRepository) GetByTeamId(id int) ([]domain.Tag, error) {
	//TODO implement me
	panic("implement me")
}

func (t *tagRepository) GetAll() ([]domain.Tag, error) {
	//TODO implement me
	panic("implement me")
}

func (t *tagRepository) GetByGlobalTagId(id int) ([]domain.Tag, error) {
	//TODO implement me
	panic("implement me")
}

func (t *tagRepository) GetByEventId(id int) ([]domain.Tag, error) {
	//TODO implement me
	panic("implement me")
}

func (t *tagRepository) AddTagToUser(c *gin.Context) error {
	//TODO implement me
	panic("implement me")
}

func (t *tagRepository) AddTagToTeam(c *gin.Context) error {
	//TODO implement me
	panic("implement me")
}

func (t *tagRepository) AddTagToEvent(c *gin.Context) error {
	//TODO implement me
	panic("implement me")
}

func (t *tagRepository) CreateTag(c *gin.Context) error {
	//TODO implement me
	panic("implement me")
}

func (t *tagRepository) GetById(id int) (domain.Tag, error) {
	//TODO implement me
	panic("implement me")
}

func (t *tagRepository) GetUserTagCount(userId, tagId int) (int, error) {
	// todo
	panic("")
}

func (t *tagRepository) GetTeamTagCount(teamId, tagId int) (int, error) {
	// todo
	panic("")
}

func (t *tagRepository) GetEventTagCount(eventId, tagId int) (int, error) {
	// todo
	panic("")
}

func (t *tagRepository) PostTagToUser(userId int, tag domain.Tag) error {
	// todo
	panic("")
}

func (t *tagRepository) PostTagToTeam(teamId int, tag domain.Tag) error {
	// todo
	panic("")
}

func (t *tagRepository) PostTagToEvent(eventId int, tag domain.Tag) error {
	// todo
	panic("")
}
