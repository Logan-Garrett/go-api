package controllers

import (
	"errors"

	"net/http"

	"github.com/gin-gonic/gin"
)

type Job struct {
	ID         string `json:"id"`
	Restaurant string `json:"restaurant"`
	Good       bool   `json:"good"`
}

var Jobs = []Job{
	{ID: "1", Restaurant: "Taco Bell", Good: false},
	{ID: "2", Restaurant: "Wendy's", Good: false},
	{ID: "3", Restaurant: "Mcdonald's", Good: false},
	{ID: "4", Restaurant: "Arby's", Good: false},
}

func GetJobs(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, Jobs)
}

func AddJob(context *gin.Context) {
	var newJob Job

	if err := context.BindJSON(&newJob); err != nil {
		return
	}

	Jobs = append(Jobs, newJob)
	context.IndentedJSON(http.StatusCreated, newJob)
}

func GetJobById(id string) (*Job, error) {
	for i, d := range Jobs {
		if d.ID == id {
			return &Jobs[i], nil
		}
	}

	return nil, errors.New("Job Not Found")
}

func GetJob(context *gin.Context) {
	id := context.Param("id")
	job, err := GetJobById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Job Not Found."})
		return
	}

	context.IndentedJSON(http.StatusOK, job)
}

func ModifyJob(context *gin.Context) {
	id := context.Param("id")
	job, err := GetJobById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Job Not Found."})
		return
	}

	job.Good = !job.Good
	context.IndentedJSON(http.StatusOK, job)
}
