package entity

import (
	"time"
)

type Location struct {
	Lat  string `json:"lat"`
	Long string `json:"long"`
}

type EmojiReaction struct {
	UserUuid  string `json:"user_uuid" bson:"user_uuid"`
	Emoji     string `json:"emoji" bson:"emoji"`
	Timestamp string `json:"timestamp" bson:"timestamp"`
}

type Workout struct {
	Uuid            string          `json:"uuid" bson:"uuid"`
	UserUuid        string          `json:"user_uuid" bson:"user_uuid"`
	Title           string          `json:"title" bson:"title"`
	Location        Location        `json:"location" bson:"location"`
	Duration        int             `json:"duration" bson:"duration"`
	Pace            *string         `json:"pace" bson:"pace"`
	Distance        *float64        `json:"distance" bson:"distance"`
	Comment         *string         `json:"comment" bson:"comment"`
	Type            string          `json:"type" bson:"type"`
	Photo           string          `json:"photo" bson:"photo"`
	CreatedAt       string          `json:"created_at" bson:"created_at"`
	UpdatedAt       string          `json:"updated_at" bson:"updated_at"`
	Reactions       []EmojiReaction `json:"reactions,omitempty" bson:"reactions,omitempty"`
	ReactionSummary map[string]int  `json:"reaction_summary,omitempty" bson:"reaction_summary,omitempty"`
}

func NewWorkout(
	uuid string,
	userUuid string,
	title string,
	location Location,
	duration int,
	distance *float64,
	comment *string,
	workoutType string,
	photo string,
	pace *string,
) *Workout {
	entity := &Workout{
		Uuid:      uuid,
		UserUuid:  userUuid,
		Title:     title,
		Location:  location,
		Duration:  duration,
		Distance:  distance,
		Comment:   comment,
		Type:      workoutType,
		Photo:     photo,
		Pace:      pace,
		CreatedAt: time.Now().Format(time.RFC3339),
		UpdatedAt: "",
		Reactions: []EmojiReaction{},
	}
	return entity
}
