package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/MicBun/go-job-list-api/models"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strconv"
)

// GetJobListFromApi godoc
// @Summary Get job list from api
// @Description get job list from api
// @Tags Jobs
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Param description query string false "Description"
// @Param location query string false "Location"
// @Param full_time query string false "Full Time"
// @Param page query string false "Page"
// @Security BearerToken
// @Produce json
// @Success 200 {object} []models.Job
// @Router /jobs [get]
func GetJobListFromApi(ctx *gin.Context) {
	var err error
	url := "http://dev3.dansmultipro.co.id/api/recruitment/positions.json"
	description := ctx.Query("description")
	location := ctx.Query("location")
	fullTime := ctx.Query("full_time")
	pageStr := ctx.Query("page")
	page := 1
	if pageStr != "" {
		page, err = strconv.Atoi(pageStr)
		if err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}
	}
	if description != "" || location != "" || fullTime != "" {
		url += "?"
		if description != "" {
			url += "description=" + description + "&"
		}
		if location != "" {
			url += "location=" + location + "&"
		}
		if fullTime == "true" {
			url += "full_time=true"
		}
	}

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	var jobs models.Jobs
	err = json.Unmarshal(body, &jobs)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if len(jobs) > page*5 && jobs[page*5-1].ID != "" {
		jobs = jobs[(page-1)*5 : page*5]
	} else if len(jobs) > (page-1)*5 {
		jobs = jobs[(page-1)*5:]
	} else {
		jobs = []models.Job{}
	}

	if len(jobs) == 0 {
		ctx.JSON(404, gin.H{"error": "Not Found"})
		return
	}

	ctx.JSON(200, jobs)
}

// GetJobDetailFromApi godoc
// @Summary Get job detail from api
// @Description get job detail from api
// @Tags Jobs
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Param id path string true "Job ID"
// @Security BearerToken
// @Produce json
// @Success 200 {object} models.Job
// @Router /jobs/{id} [get]
func GetJobDetailFromApi(ctx *gin.Context) {
	var err error
	id := ctx.Param("id")
	url := "http://dev3.dansmultipro.co.id/api/recruitment/positions/" + id

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	var job models.Job
	err = json.Unmarshal(body, &job)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	ctx.JSON(200, job)
}
