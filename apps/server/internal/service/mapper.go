package service

import (
	"github.com/1001001010/messanger/gen/common"
	"github.com/1001001010/messanger/internal/database"
)

func mapUser(u database.User) *common.User {
	user := &common.User{
		Id:            u.ID.String(),
		Username:      u.Username.String,
		Email:         u.Email,
		FirstName:     u.FirstName,
		LastName:      u.LastName.String,
		Bio:           u.Bio.String,
		FileId:        u.FileID.String,
		EmailVerified: u.EmailVerified,
		IsOnline:      u.IsOnline,
	}

	if u.LastSeen.Valid {
		user.LastSeen = u.LastSeen.Time.Unix()
	}

	return user
}
