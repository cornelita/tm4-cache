package main

import (
	"errors"

	"github.com/gin-gonic/gin"
)

var database []Student

func getStudentByNPM(npm string) (*Student, int, error) {
	for idx, student := range database {
		if student.NPM == npm {
			return &student, idx, nil
		}
	}
	return nil, -1, errors.New("Student not found!")
}

func updateHandler(c *gin.Context) {
	var studentData Student
	c.BindJSON(&studentData)
	_, idx, err := getStudentByNPM(studentData.NPM)
	if err != nil {
		database = append(database, studentData)
	} else {
		student := &database[idx]
		student.Name = studentData.Name
	}

	c.JSON(200, gin.H{
		"status": "OK",
	})
}

func readHandler(c *gin.Context) {
	npm := c.Param("npm")
	student, _, err := getStudentByNPM(npm)
	if err != nil {
		c.JSON(404, gin.H{
			"status": "error",
			"data":   err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"status": "OK",
		"npm":    student.NPM,
		"nama":   student.Name,
	})
}
