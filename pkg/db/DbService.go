package db

import (
	"encoding/json"
	"errors"
	"github.com/go-redis/redis/v7"
	"github.com/google/uuid"
	"github.com/subramanya812/crud-operation/pkg/dto"
	"github.com/subramanya812/crud-operation/pkg/model"
	"time"
)

type redisCache struct {
	host string
	db   int
	exp  time.Duration
}

func (cache redisCache) GetMovie(id string) (*model.Movie, error) {
	//TODO implement me
	c := cache.getClient()
	res, err := c.HGet("movies", id).Result()
	if err != nil {
		return nil, err
	}
	movies := &model.Movie{}
	errs := json.Unmarshal([]byte(res), movies)
	if errs != nil {
		return nil, errs
	}

	return movies, nil
}

func (cache redisCache) GetMovies() ([]*model.Movie, error) {
	//TODO implement me
	c := cache.getClient()
	res, err := c.HGetAll("movies").Result()
	if err != nil {
		return nil, err
	}
	movies := []*model.Movie{}
	for _, data := range res {
		movie := &model.Movie{}
		err := json.Unmarshal([]byte(data), movie)
		if err != nil {
			return nil, err
		}
		movies = append(movies, movie)
	}
	return movies, nil

}

func (cache redisCache) CreateMovie(movie *model.Movie) (*model.Movie, error) {
	//TODO implement me
	c := cache.getClient()
	movie.ID = uuid.New().String()
	jsonData, err := json.Marshal(movie)
	c.HSet("movies", movie.ID, jsonData)
	if err != nil {
		return nil, err
	}
	return movie, err
}

func (cache redisCache) UpdateMovie(movie *model.Movie) (*model.Movie, error) {
	//TODO implement me
	c := cache.getClient()
	json, err := json.Marshal(movie)
	if err != nil {
		return nil, err
	}
	c.HSet("movies", movie.ID, json)
	if err != nil {
		return nil, err
	}

	return movie, err
}

func (cache redisCache) DeleteMovie(id string) error {
	//TODO implement me
	c := cache.getClient()
	isDeleted, err := c.HDel("movies", id).Result()
	if isDeleted == 0 {
		return errors.New("cannot deleted")
	}
	if err != nil {
		return err
	}
	return nil
}

func NewRedisCache(host string, db int, exp time.Duration) dto.MovieService {
	return &redisCache{
		host: host,
		db:   db,
		exp:  exp,
	}
}
func (cache redisCache) getClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     cache.host,
		Password: "",
		DB:       cache.db,
	})
}
