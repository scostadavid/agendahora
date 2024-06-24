package dto

type AddToWaitlistDTO struct {
	Email string `validate:"required,email"`
}
