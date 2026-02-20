package verify

type SendRequest struct {
	Email string `validate:"required,email"`
}
