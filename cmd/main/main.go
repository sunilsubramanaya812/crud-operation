package main

import (
	"github.com/gin-gonic/gin"
	database "github.com/subramanya812/crud-operation/pkg/db"
	models "github.com/subramanya812/crud-operation/pkg/model"
	"net/http"
)

var (
	redisCache = database.NewRedisCache("localhost:6379", 0, 1)
)

func main() {
	r := gin.Default()

	r.POST("/movies", func(context *gin.Context) {
		var movie models.Movie
		if err := context.ShouldBindJSON(&movie); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		res, err := redisCache.CreateMovie(&movie)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		context.JSON(http.StatusOK, gin.H{
			"movie": res,
		})
	})

	r.PUT("/movies/:id", func(context *gin.Context) {
		id := context.Param("id")
		redisCache.GetMovie(id)

		res, err := redisCache.GetMovie(id)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		var movie models.Movie

		errs := context.ShouldBind(&movie)
		if errs != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": errs.Error(),
			})
			return
		}

		res.Title = movie.Title
		res.Description = movie.Description
		res, err = redisCache.UpdateMovie(res)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		context.JSON(http.StatusOK, gin.H{
			"movie": res,
		})
	})

	r.GET("/movies", func(context *gin.Context) {
		res, err := redisCache.GetMovies()
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		context.JSON(http.StatusOK, gin.H{
			"movie": res,
		})
	})

	r.GET("/movies/:id", func(context *gin.Context) {

		data := context.Param("id")
		movie, err := redisCache.GetMovie(data)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		context.JSON(http.StatusOK, gin.H{
			"movie": movie,
		})
	})

	r.DELETE("/movies/:id", func(context *gin.Context) {
		id := context.Param("id")
		err := redisCache.DeleteMovie(id)
		if err != nil {
			context.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}
		context.JSON(http.StatusOK, gin.H{
			"message": "movie deleted successfully",
		})
	})
	r.Run()
}
