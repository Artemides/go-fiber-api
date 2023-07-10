package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        *uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Name      string     `gorm:"type:varchar(100);not null"`
	Email     string     `gorm:"type:varchar(100);uniqueIndex;not null"`
	Password  string     `gorm:"type:varchar(100);not null"`
	Role      *string    `gorm:"type:varchar(50);default:'user';not null"`
	Provider  *string    `gorm:"type:varchar(50);default:'local';not null"`
	Photo     *string    `gorm:"default:'photo.png';not null"`
	Verified  *bool      `gorm:"default:false;not null"`
	CreatedAt *time.Time `gorm:"default:now();not null"`
	UpdatedAt *time.Time `grom:"default:now()"`
}

type SignUpInput struct {
	Name            string `json:"name" validate:"required"`
	Email           string `json:"email" validate:"required,email"`
	Password        string `json:"password" validate:"required,min=8"`
	ConfirmPassword string `json:"conafirmPassword" validate:"required,min=8"`
	Photo           string `json:"photo"`
}

type SignInInput struct {
	Email    string `json:"email" validate:"email,required"`
	Password string `json:"password" validate:"required"`
}

type UserResponse struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	Provider  string    `json:"provider"`
	Photo     string    `json:"photo,omitempty"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}

func FilterUserRecord(user *User) UserResponse {
	return UserResponse{
		ID:        *user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Role:      *user.Role,
		Photo:     *user.Photo,
		Provider:  *user.Provider,
		CreatedAt: *user.CreatedAt,
		UpdatedAt: *user.UpdatedAt,
	}
}
