package user

import (
	"database/sql"
	"fmt"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetUserById(id string) (*User, error) {
	rows, err := s.db.Query(`SELECT * FROM "users" WHERE id = $1`, id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	u := new(User)
	for rows.Next() {
		u, err = scanRowIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}

	if u.UID == "" {
		return nil, fmt.Errorf("user not found")
	}

	return u, nil
}

func (s *Store) GetUserByEmail(email string) (*User, error) {
	rows, err := s.db.Query(`SELECT * FROM "users" WHERE email = $1`, email)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	u := new(User)
	for rows.Next() {
		u, err = scanRowIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}

	if u.UID == "" {
		return nil, fmt.Errorf("user not found")
	}

	return u, nil
}

func (s *Store) CreateUser(user RegisterUserPayload) error {
	_, err := s.db.Exec(`
		INSERT INTO "users" 
		(uid, name, family_name, given_name, email, avatar, role, created_at) 
		VALUES (uuid_generate_v4(), $1, $2, $3, $4, $5, $6, NOW())`,
		user.Name, user.FamilyName, user.GivenName, user.Email, user.Avatar, user.Role,
	)
	if err != nil {
		return err
	}
	return nil
}

func scanRowIntoUser(rows *sql.Rows) (*User, error) {
	user := new(User)

	err := rows.Scan(
		&user.UID,
		&user.ID,
		&user.Name,
		&user.GivenName,
		&user.FamilyName,
		&user.Email,
		&user.DateOfBirth,
		&user.Religion,
		&user.EthnicGroup,
		&user.Address,
		&user.Resident,
		&user.Faculty,
		&user.Class,
		&user.AcademicYear,
		&user.Department,
		&user.Position,
		&user.Role,
		&user.PartyMember,
		&user.Avatar,
		&user.Signature,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.VerifiedAt,
		&user.ApprovedAt,
		&user.DeletedAt,
		&user.BlockedAt,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}
