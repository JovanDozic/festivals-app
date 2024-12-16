package user

import dto "backend/internal/dto/user"

// ! Note: this is used as a contract for the order service to interact with the user service as our services are decoupled

type UserProvider interface {
	GetUserProfile(username string) (*dto.GetUserProfileResponse, error)
	GetUserProfileById(userId uint) (*dto.GetUserProfileResponse, error)
	GetUserID(username string) (uint, error)
}
