package entity

import "time"

type Professor struct {
	ID             string    `db:"id"`
	Name           string    `db:"name"`
	Faculty        string    `db:"faculty"`
	Major          string    `db:"major"`
	ProfileImgLink string    `db:"profile_img_link"`
	CreatedAt      time.Time `db:"created_at"`
	UpdatedAt      time.Time `db:"updated_at"`
}
