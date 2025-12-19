package service

import (
	"time"

	"Form_1/database"
	"Form_1/models"

	"go.mongodb.org/mongo-driver/bson"
)

// lưu thông tin vào db
func SaveContestant(c models.Contestant) error {
	data := prepareContestantData(c)
	return database.InsertContestant(data)
}

func prepareContestantData(c models.Contestant) map[string]interface{} {
	now := time.Now()

	c.Management.SubmittedAt = now
	c.Management.UpdatedAt = now
	c.Management.IsActive = true

	data, _ := bson.Marshal(c)
	var result map[string]interface{}
	bson.Unmarshal(data, &result)

	return result
}
