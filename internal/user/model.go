package user

import "time"

type UserStore interface {
	GetUserById(id string) (*User, error)
	GetUserByEmail(email string) (*User, error)
	CreateUser(user RegisterUserPayload) error
}

type User struct {
	UID          string     `json:"uid"`
	ID           *string    `json:"id,omitempty"`
	Name         string     `json:"name"`
	FamilyName   string     `json:"familyName"`
	GivenName    string     `json:"givenName"`
	Email        string     `json:"email"`
	DateOfBirth  *time.Time `json:"dateOfBirth,omitempty"`
	Religion     *string    `json:"religion,omitempty"`
	EthnicGroup  *string    `json:"ethnicGroup,omitempty"`
	Address      *string    `json:"address,omitempty"`
	Resident     *string    `json:"resident,omitempty"`
	Faculty      *string    `json:"faculty,omitempty"`
	Class        *string    `json:"class,omitempty"`
	AcademicYear *int       `json:"academicYear,omitempty"`
	Department   *string    `json:"department,omitempty"`
	Position     *string    `json:"position,omitempty"`
	Role         string     `json:"role"`
	PartyMember  *string    `json:"partyMember,omitempty"`
	Avatar       *string    `json:"avatar,omitempty"`
	Signature    *string    `json:"signature,omitempty"`
	CreatedAt    time.Time  `json:"createdAt"`
	UpdatedAt    *time.Time `json:"updatedAt,omitempty"`
	VerifiedAt   *time.Time `json:"verifiedAt,omitempty"`
	ApprovedAt   *time.Time `json:"approvedAt,omitempty"`
	DeletedAt    *time.Time `json:"deletedAt,omitempty"`
	BlockedAt    *time.Time `json:"blockedAt,omitempty"`
}
