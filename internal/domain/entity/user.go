package entity

type User struct {
	Uuid         string   `json:"uuid"`
	Email        string   `json:"email"`
	BirthDate    string   `json:"birth_date"`
	Name         string   `json:"name"`
	Password     *[]byte  `json:"password,omitempty"`
	Sports       []string `json:"sports"`
	AuthProvider string   `json:"auth_provider"`
	GoogleSub    *string  `json:"google_sub,omitempty"`
	Photo        *string  `json:"photo,omitempty"`
}

func NewUser(
	uuid string,
	email string,
	birthDate string,
	name string,
	password *[]byte,
	sports []string,
	authProvider string,
	googleSub *string,
	photo *string,
) *User {
	entity := &User{
		Uuid:         uuid,
		Email:        email,
		Name:         name,
		Sports:       sports,
		Password:     password,
		BirthDate:    birthDate,
		AuthProvider: authProvider,
		GoogleSub:    googleSub,
		Photo:        photo,
	}
	return entity
}
