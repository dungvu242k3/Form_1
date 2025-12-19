package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type PublicContestantData struct {
	ID          string `bson:"_id"`
	FullName    string `bson:"fullName"`
	DateOfBirth string `bson:"dateOfBirth"`
	City        string `bson:"city"`
	District    string `bson:"district"`
	Street      string `bson:"street"`
}

// lấy danh sách công khai (id, name, dateOfBirth, address)
func GetPublicContestantsList() ([]PublicContestantData, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{
		"management.isVisibleToUsers": true,
		"management.isActive":         true,
	}

	projection := bson.M{
		"_id":                           1,
		"personalInfo.fullName":         1,
		"personalInfo.dateOfBirth":      1,
		"personalInfo.address.city":     1,
		"personalInfo.address.district": 1,
		"personalInfo.address.street":   1,
	}

	opts := options.Find().SetProjection(projection)

	cursor, err := ContestantsCollection.Find(ctx, filter, opts)
	if err != nil {
		log.Printf("lỗi khi truy vấn: %v", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	var results []PublicContestantData

	for cursor.Next(ctx) {
		var temp struct {
			ID           string `bson:"_id"`
			PersonalInfo struct {
				FullName    string `bson:"fullName"`
				DateOfBirth string `bson:"dateOfBirth"`
				Address     struct {
					City     string `bson:"city"`
					District string `bson:"district"`
					Street   string `bson:"street"`
				} `bson:"address"`
			} `bson:"personalInfo"`
		}

		if err := cursor.Decode(&temp); err != nil {
			log.Printf("lỗi khi decode: %v", err)
			return nil, err
		}

		results = append(results, PublicContestantData{
			ID:          temp.ID,
			FullName:    temp.PersonalInfo.FullName,
			DateOfBirth: temp.PersonalInfo.DateOfBirth,
			City:        temp.PersonalInfo.Address.City,
			District:    temp.PersonalInfo.Address.District,
			Street:      temp.PersonalInfo.Address.Street,
		})
	}

	return results, nil
}
