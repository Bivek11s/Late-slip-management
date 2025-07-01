package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"lateslip/initialializers"
	"lateslip/models"
)

func GetTodaySchedule(c *gin.Context) {
	userId, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	studentID, err := primitive.ObjectIDFromHex(userId.(string))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()

	// Get student to find their level
	var student models.Student
	studentCollection := initialializers.DB.Collection("students")
	err = studentCollection.FindOne(ctx, bson.M{"_id": studentID}).Decode(&student)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		return
	}

	// Get today's day name (e.g., "Monday")
	// today := time.Now().Weekday().String()
	today := "Monday"

	// Find schedules for this level and today
	scheduleCollection := initialializers.DB.Collection("schedules")
	filter := bson.M{
		"level": student.Level,
		"day":   today,
	}
	cursor, err := scheduleCollection.Find(ctx, filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch schedules"})
		return
	}
	defer cursor.Close(ctx)

	var schedules []models.Schedule
	if err := cursor.All(ctx, &schedules); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode schedules"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success":   true,
		"day":       today,
		"schedules": schedules,
	})
}
