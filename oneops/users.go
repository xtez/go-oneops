package oneops

import (
	"context"
	"net/http"
	"time"
)

// UsersService handles communication with the user related
// methods of the OneOps API.
type UsersService service

// User represents OneOps user
type User struct {
	ID              int          `json:"id,omitempty"`
	Email           string       `json:"email,omitempty"`
	SignInCount     int          `json:"sign_in_count,omitempty"`
	CurrentSignInAt time.Time    `json:"current_sign_in_at,omitempty"`
	CreatedAt       time.Time    `json:"created_at,omitempty"`
	UpdatedAt       time.Time    `json:"updated_at,omitempty"`
	Name            string       `json:"name,omitempty"`
	Username        string       `json:"username,omitempty"`
	OwnerEmail      string       `json:"owner_email,omitempty"`
	Organization    Organization `json:"organization,omitempty"`
}

func (u User) String() string {
	return Stringify(u)
}

// Get fetches the authenticated user.
func (s *UsersService) Get(ctx context.Context) (*User, *http.Response, error) {
	u := "account/profile"

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	user := new(User)
	resp, err := s.client.Do(ctx, req, user)
	if err != nil {
		return nil, resp, err
	}

	return user, resp, nil
}
