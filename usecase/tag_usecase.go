package usecase

import (
	"team-finder/domain"
	"time"
)

type tagUsecase struct {
	tagRepository  domain.TagRepository
	contextTimeout time.Duration
}

func NewTagUsecase(tagRepository domain.TagRepository, timeout time.Duration) domain.TagUsecase {
	return &tagUsecase{
		tagRepository:  tagRepository,
		contextTimeout: timeout,
	}
}

func (tu *tagUsecase) GetAll() ([]domain.Tag, error) {
	return tu.tagRepository.GetAll()
}

func (tu *tagUsecase) GetByUserId(id int) ([]domain.Tag, error) {
	return tu.tagRepository.GetByUserId(id)
}

func (tu *tagUsecase) GetByTeamId(id int) ([]domain.Tag, error) {
	return tu.tagRepository.GetByTeamId(id)
}

func (tu *tagUsecase) GetByEventId(id int) ([]domain.Tag, error) {
	return tu.tagRepository.GetByEventId(id)
}

func (tu *tagUsecase) GetByGlobalTagId(id int) ([]domain.Tag, error) {
	return tu.tagRepository.GetByGlobalTagId(id)
}

func (tu *tagUsecase) PostTagsToUser(id int, tags []int) error {
	for _, tag := range tags {
		count, err := tu.tagRepository.GetUserTagCount(id, tag)
		if err != nil {
			return err
		}
		if count > 0 {
			continue
		}
		err = tu.tagRepository.PostTagToUser(id, tag)
		if err != nil {
			return err
		}
	}
	return nil
}

func (tu *tagUsecase) PostTagsToTeam(teamId int, tags []int) error {
	for _, tag := range tags {
		count, err := tu.tagRepository.GetTeamTagCount(teamId, tag)
		if err != nil {
			return err
		}
		if count > 0 {
			continue
		}
		err = tu.tagRepository.PostTagToTeam(teamId, tag)
		if err != nil {
			return err
		}
	}
	return nil
}

func (tu *tagUsecase) PostTagsToEvent(eventId int, tags []int) error {
	for _, tag := range tags {
		count, err := tu.tagRepository.GetEventTagCount(eventId, tag)
		if err != nil {
			return err
		}
		if count > 0 {
			continue
		}
		err = tu.tagRepository.PostTagToEvent(eventId, tag)
		if err != nil {
			return err
		}
	}
	return nil
}

func (tu *tagUsecase) DeleteTagsFromUser(request domain.PostTagsRequest, userId int) error {
	for _, tag := range request.Tags {
		count, err := tu.tagRepository.GetUserTagCount(userId, tag)
		if err != nil {
			return err
		}
		if count == 0 {
			continue
		}
		err = tu.tagRepository.DeleteTagFromUser(userId, tag)
		if err != nil {
			return err
		}
	}
	return nil
}

func (tu *tagUsecase) DeleteTagsFromTeam(request domain.PostTagsRequest, teamId int) error {
	for _, tag := range request.Tags {
		count, err := tu.tagRepository.GetTeamTagCount(teamId, tag)
		if err != nil {
			return err
		}
		if count == 0 {
			continue
		}
		err = tu.tagRepository.DeleteTagFromTeam(teamId, tag)
		if err != nil {
			return err
		}
	}
	return nil
}

func (tu *tagUsecase) DeleteTagsFromEvent(request domain.PostTagsRequest, eventId int) error {
	for _, tag := range request.Tags {
		count, err := tu.tagRepository.GetEventTagCount(eventId, tag)
		if err != nil {
			return err
		}
		if count == 0 {
			continue
		}
		err = tu.tagRepository.DeleteTagFromEvent(eventId, tag)
		if err != nil {
			return err
		}
	}
	return nil
}
