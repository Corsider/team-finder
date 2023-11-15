package usecase

import (
	"team-finder/domain"
	"time"
)

type teamUsecase struct {
	teamRepository domain.TeamRepository
	contextTimeout time.Duration
}

func NewTeamUsecase(teamRepository domain.TeamRepository, timeout time.Duration) domain.TeamUsecase {
	return &teamUsecase{
		teamRepository: teamRepository,
		contextTimeout: timeout,
	}
}

func (t *teamUsecase) GetAll() ([]domain.Team, error) {
	return t.teamRepository.GetAll()
}

func (t *teamUsecase) GetByTeamId(id int) (domain.Team, error) {
	return t.teamRepository.GetByTeamId(id)
}

func (t *teamUsecase) GetByUserId(id int) ([]domain.Team, error) {
	return t.teamRepository.GetByUserId(id)
}

func (t *teamUsecase) GetByEventId(id int) ([]domain.Team, error) {
	return t.teamRepository.GetByEventId(id)
}
