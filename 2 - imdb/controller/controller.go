package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"imdb/config"
	"imdb/model"

	"github.com/gin-gonic/gin"
)

func GetMovie(c *gin.Context) {
	searchWord := c.Query("search_word")
	page := c.DefaultQuery("pagination", "1")
	client := &http.Client{}
	db := config.Connect()
	defer config.Close(db)

	newLog := model.Log{
		SearchWord: searchWord,
		Page:       page,
	}

	db.Table("logs").Create(&newLog)
	url := "http://www.omdbapi.com/?apikey=faf7e5bb&s=" + searchWord + "&page=" + page
	request, _ := http.NewRequest("GET", url, nil)

	q := request.URL.Query()

	request.URL.RawQuery = q.Encode()
	response, _ := client.Do(request)

	var searchResponses model.SearchResponse
	_ = json.NewDecoder(response.Body).Decode(&searchResponses)

	result := model.Result{
		Results: searchResponses.Search,
		Page:    page,
	}
	fmt.Println(page)
	c.JSON(http.StatusOK, result)

}

func ShowLogs(c *gin.Context) {

	db := config.Connect()
	defer config.Close(db)

	var logs []model.Log

	db.Table("logs").Find(&logs)

	c.JSON(http.StatusOK, logs)

}
