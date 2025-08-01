package entity

type Rules struct {
	MaxEntriesPerDay      int `json:"max_entries_per_day"`
	MinDurationFor1Point  int `json:"min_duration_one_point"`
	MinDurationFor2Points int `json:"min_duration_two_points"`
	MaxPointsPerDay       int `json:"max_points_per_day"`
}

type LeaderBoard struct {
	UserUuid     string `json:"user_uuid"`
	TotalPoints  string `json:"total_points"`
	LastActivity string `json:"last_activity"`
}

type Gamification struct {
	Uuid        string        `json:"uuid"`
	OwnerUuid   string        `json:"owner_uuid"`
	Name        string        `json:"name"`
	StartDate   string        `json:"start_date"`
	EndDate     string        `json:"end_date"`
	Photo       string        `json:"photo"`
	Rules       Rules         `json:"rules"`
	Users       []string      `json:"users"`
	LeaderBoard []LeaderBoard `json:"leaderboard"`
}

func NewGamification(
	uuid string,
	ownerUuid string,
	name string,
	startDate string,
	endDate string,
	photo string,
) *Gamification {
	entity := &Gamification{
		Uuid:      uuid,
		OwnerUuid: ownerUuid,
		Name:      name,
		StartDate: startDate,
		EndDate:   endDate,
		Rules: Rules{
			MaxEntriesPerDay:      2,
			MinDurationFor1Point:  1200,
			MinDurationFor2Points: 2400,
			MaxPointsPerDay:       4,
		},
		Users:       []string{ownerUuid},
		LeaderBoard: []LeaderBoard{},
		Photo:       photo,
	}

	return entity
}
