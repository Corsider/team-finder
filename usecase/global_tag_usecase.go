package usecase

import (
	"team-finder/domain"
	"time"
)

type globalTagUsecase struct {
	globalTagRepository domain.GlobalTagRepository
	contextTimeout      time.Duration
}

func NewGlobalTagUsecase(globalTagRepository domain.GlobalTagRepository, timeout time.Duration) domain.GlobalTagUsecase {
	return &globalTagUsecase{
		globalTagRepository: globalTagRepository,
		contextTimeout:      timeout,
	}
}

func (gt *globalTagUsecase) GetAll() ([]domain.GlobalTag, error) {
	return gt.globalTagRepository.GetAll()
}
