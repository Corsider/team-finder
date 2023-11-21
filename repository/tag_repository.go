package repository

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"strconv"
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
	rows, err := t.database.SelectAllFromXWhereYeqZ("users_tags", "user_id", strconv.Itoa(id))
	if err != nil {
		return nil, err
	}
	tags := []domain.Tag{}
	var usr_tag domain.UsersTags
	for rows.Next() {
		var tag domain.Tag
		rows.Scan(&usr_tag.TagId, &usr_tag.UserId)
		row := t.database.Select1FromXWhereYeqZ("tag", "tag_id", strconv.Itoa(usr_tag.TagId))
		row.Scan(&tag.TagID, &tag.Activity, &tag.GlobalTagID)
		tags = append(tags, tag)
	}
	return tags, nil
}

func (t *tagRepository) GetByTeamId(id int) ([]domain.Tag, error) {
	rows, err := t.database.SelectAllFromXWhereYeqZ("team_tags", "team_id", strconv.Itoa(id))
	if err != nil {
		return nil, err
	}
	tags := []domain.Tag{}
	var teamTag domain.TeamsTags
	for rows.Next() {
		var tag domain.Tag
		rows.Scan(&teamTag.TagId, &teamTag.TeamId)
		row := t.database.Select1FromXWhereYeqZ("tag", "tag_id", strconv.Itoa(teamTag.TagId))
		row.Scan(&tag.TagID, &tag.Activity, &tag.GlobalTagID)
		tags = append(tags, tag)
	}
	return tags, nil
}

func (t *tagRepository) GetAll() ([]domain.Tag, error) {
	rows, err := t.database.SelectAllFromX("tag")
	if err != nil {
		return nil, err
	}
	tags := []domain.Tag{}
	for rows.Next() {
		var tag domain.Tag
		rows.Scan(&tag.TagID, &tag.Activity, &tag.GlobalTagID)
		tags = append(tags, tag)
	}
	return tags, nil
}

func (t *tagRepository) GetByGlobalTagId(id int) ([]domain.Tag, error) {
	rows, err := t.database.SelectAllFromXWhereYeqZ(t.table, "global_tag_id", strconv.Itoa(id))
	if err != nil {
		return nil, err
	}
	tags := []domain.Tag{}
	for rows.Next() {
		var tag domain.Tag
		rows.Scan(&tag.TagID, &tag.Activity, &tag.GlobalTagID)
		tags = append(tags, tag)
	}
	return tags, nil
}

func (t *tagRepository) GetByEventId(id int) ([]domain.Tag, error) {
	rows, err := t.database.SelectAllFromXWhereYeqZ("events_tags", "event_id", strconv.Itoa(id))
	if err != nil {
		return nil, err
	}
	tags := []domain.Tag{}
	var eventTag domain.EventTags
	for rows.Next() {
		var tag domain.Tag
		rows.Scan(&eventTag.EventId, &eventTag.TagId)
		row := t.database.Select1FromXWhereYeqZ("tag", "tag_id", strconv.Itoa(eventTag.TagId))
		row.Scan(&tag.TagID, &tag.Activity, &tag.GlobalTagID)
		tags = append(tags, tag)
	}
	return tags, nil
}

//func (t *tagRepository) AddTagToUser(c *gin.Context) error {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (t *tagRepository) AddTagToTeam(c *gin.Context) error {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (t *tagRepository) AddTagToEvent(c *gin.Context) error {
//	//TODO implement me
//	panic("implement me")
//}

func (t *tagRepository) CreateTag(c *gin.Context) error {
	//TODO implement me
	panic("implement me")
}

func (t *tagRepository) GetById(id int) (domain.Tag, error) {
	row := t.database.Select1FromXWhereYeqZ(t.table, "tag_id", strconv.Itoa(id))
	var tag domain.Tag
	err := row.Scan(&tag.TagID, &tag.Activity, &tag.GlobalTagID)
	if err != nil {
		return domain.Tag{}, err
	}
	return tag, nil
}

func (t *tagRepository) GetUserTagCount(userId, tagId int) (int, error) {
	count, err := t.database.SelectCountFromXWhereYeqZ("users_tags", "user_id=$1 and tag_id=$2", strconv.Itoa(userId), strconv.Itoa(tagId))
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (t *tagRepository) GetTeamTagCount(teamId, tagId int) (int, error) {
	count, err := t.database.SelectCountFromXWhereYeqZ("team_tags", "team_id=$1 and tag_id=$2", strconv.Itoa(teamId), strconv.Itoa(tagId))
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (t *tagRepository) GetEventTagCount(eventId, tagId int) (int, error) {
	count, err := t.database.SelectCountFromXWhereYeqZ("events_tags", "event_id=$1 and tag_id=$2", strconv.Itoa(eventId), strconv.Itoa(tagId))
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (t *tagRepository) PostTagToUser(userId int, tag int) error {
	z := strconv.Itoa(tag) + ", " + strconv.Itoa(userId)
	err := t.database.InsertIntoXYValuesZ("users_tags", "tag_id, user_id", z)
	if err != nil {
		return err
	}
	return nil
}

func (t *tagRepository) PostTagToTeam(teamId int, tag int) error {
	z := strconv.Itoa(tag) + ", " + strconv.Itoa(teamId)
	err := t.database.InsertIntoXYValuesZ("team_tags", "tag_id, team_id", z)
	if err != nil {
		return err
	}
	return nil
}

func (t *tagRepository) PostTagToEvent(eventId int, tag int) error {
	z := strconv.Itoa(eventId) + ", " + strconv.Itoa(tag)
	err := t.database.InsertIntoXYValuesZ("events_tags", "event_id, tag_id", z)
	if err != nil {
		return err
	}
	return nil
}

func (t *tagRepository) DeleteTagFromUser(userId, tagId int) error {
	return t.database.DeleteFromXWhereCond(domain.TableUsersTags, "user_id=$1 and tag_id=$2", userId, tagId)
}

func (t *tagRepository) DeleteTagFromTeam(teamId, tagId int) error {
	return t.database.DeleteFromXWhereCond(domain.TableTeamTags, "team_id=$1 and tag_id=$2", teamId, tagId)
}

func (t *tagRepository) DeleteTagFromEvent(eventId, tagId int) error {
	return t.database.DeleteFromXWhereCond(domain.TableEventsTags, "event_id=$1 and tag_id=$2", eventId, tagId)
}
