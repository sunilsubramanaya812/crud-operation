package dto

import "github.com/subramanya812/crud-operation/pkg/model"

type MovieService interface {
	GetMovie(id string) (*model.Movie, error)
	GetMovies() ([]*model.Movie, error)
	CreateMovie(movie *model.Movie) (*model.Movie, error)
	UpdateMovie(movie *model.Movie) (*model.Movie, error)
	DeleteMovie(id string) error
}
