package user

type RegisterUserPayload struct {
	Name       string `json:"name" validate:"required"`
	GivenName  string `json:"givenName" validate:"required"`
	FamilyName string `json:"familyName" validate:"required"`
	Email      string `json:"email" validate:"required,email"`
	Role       string `json:"role" validate:"required"`
	Avatar     string `json:"avatar" validate:"required"`
}
