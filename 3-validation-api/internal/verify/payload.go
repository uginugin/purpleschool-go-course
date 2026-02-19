package verify

type SendRequest struct {
	email string `validate:"email"`
}
