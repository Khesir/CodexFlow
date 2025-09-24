package models

import "time"

type User struct {
	ID                 int        `db:"id"`
	Email              string     `db:"email"`
	Password           string     `db:"password"`
	IsAdmin            bool       `db:"is_admin"`
	IsGoogleUser       bool       `db:"is_google_user"`
	LastLoginAt        time.Time  `db:"last_login_at"`
	HashedRefreshToken *string    `db:"refresh_token"`
	ModifiedAt         time.Time  `db:"modified_at"`
	CreatedAt          time.Time  `db:"created_at"`
	DeletedAt          *time.Time `db:"deleted_at"`
}
