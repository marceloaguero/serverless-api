package users

import "context"

// User describes a user in the system
type User struct {
	ID    string `json:"id"`
	Email string `json:"email" validate:"email,required"`
	Name  string `json:"name" validate:"required,gte=1,lte=50"`
	Age   uint32 `json:"age" validate:"required,gte=0,lte=130"`
}

// UpdateUser is used when updating
type UpdateUser struct {
	Email string `json:"email" validate:"email,required"`
	Name  string `json:"name" validate:"gte=1,lte=50"`
	Age   uint32 `json:"age" validate:"gte=0,lte=130"`
}

// Repository represent the user's repository contract
type Repository interface {
	Create(ctx context.Context, user *User) error
	Get(ctx context.Context, id string) (*User, error)
	GetAll(ctx context.Context) ([]*User, error)
	Update(ctx context.Context, id string, user *UpdateUser) error
	Delete(ctx context.Context, id string) error
}
