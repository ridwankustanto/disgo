package disgo

import (
	"encoding/json"
)

type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
	Active   bool   `json:"active"`
	Approved bool   `json:"approved"`
}

// CreateUser create new user.
func (c *Client) CreateUser(user User) (res interface{}, err error) {
	body, err := json.Marshal(user)
	if err != nil {
		return nil, err
	}

	req, err := newRequestWithBody(c, "PUT", "users", body)
	if err != nil {
		return nil, err
	}

	if res, err = executeRequest(c, req); err != nil {
		return nil, err
	}

	return res, nil
}