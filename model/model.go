package model

import (
	"errors"
	"strings"
	"time"
)

type User struct {
	User_id     int       `json:"user_id" bson:"user_id"`
	Username    string    `json:"username" bson:"username"`
	Email       string    `json:"email" bson:"email"`
	Create_time time.Time `json:"create_time" bson:"create_time"`
}
type Vinyl struct {
	//Release_info Release
	Genre string `json:"genre" bson:"genre"`
}
type Release struct {
	Release_id int       `json:"release_id" bson:"release_id"`
	Artist     string    `json:"artist" bson:"artist"`
	Name       string    `json:"name" bson:"name"`
	Date       time.Time `json:"date" bson:"date"`
	Label      string    `json:"label" bson:"label"`
}

func (u User) ValidateUser() error {
	if strings.TrimSpace(u.Username) == "" {
		return errors.New("username is empty")
	}
	if strings.TrimSpace(u.Email) == "" {
		return errors.New("email is empty")
	}

	return nil
}
