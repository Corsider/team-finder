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

func (t *teamUsecase) RegTeam(request domain.TeamsRegRequest) (int, error) {
	return t.teamRepository.RegTeam(request)
}

func (t *teamUsecase) AddUserToTeam(userId int, teamId int) error {
	return t.teamRepository.AddUserToTeam(userId, teamId)
}

func (t *teamUsecase) DeleteTeamById(teamId int) error {
	return t.teamRepository.DeleteTeamById(teamId)
}

func (t *teamUsecase) Filter(onlyUser bool, tags []int, myTeam int, sortBy string, asc bool, from, to int) ([]domain.Team, error) {
	var order string
	switch sortBy {
	case "count":
		order = "(select count(*) from user_team where team_id = team.team_id)"
	case "activity":
		order = "(select count(*) from team_event where team_id = team.team_id)"
	case "name":
		order = "name"
	}
	if onlyUser {
		teams, err := t.teamRepository.FilterTeamUser(order, tags, myTeam, asc, from, to)
		if err != nil {
			return nil, err
		}
		return teams, nil
	} else {
		teams, err := t.teamRepository.FilterTeamNoUser(order, tags, myTeam, asc, from, to)
		if err != nil {
			return nil, err
		}
		return teams, nil
	}
}
