package domain

const (
	TableGlobalTag = "global_tag"
)

type GlobalTag struct {
	GlobalTagID int    `db:"global_tag_id" json:"global_tag_id"`
	Category    string `db:"category" json:"category"`
}

type GlobalTagRepository interface {
	GetAll() ([]GlobalTag, error)
}

type GlobalTagUsecase interface {
	GetAll() ([]GlobalTag, error)
}
