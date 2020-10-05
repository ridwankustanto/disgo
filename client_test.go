package disgo

import (
	"log"
	"testing"
)

func openConnection(uri, username, apiKey string) (*Client, error) {

	client, err := NewClient(uri, username, apiKey)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func TestCreate(t *testing.T) {
	//want := 201
	username := "your_username"
	apiKey := "xxxxxx"
	uri := "http://localhost"

	client, err := openConnection(uri, username, apiKey)
	if err != nil {
		t.Errorf("NewClient() = got %q", err)
	}

	user := User{
		Name:     "New User",
		Email:    "newuser@example.com",
		Username: "new-user",
		Password: "123456",
		Active:   true,
		Approved: true,
	}

	res, err := client.CreateUser(user)
	if err != nil {
		t.Errorf("Error while executeRequest(). got %q", err)
	}

	log.Println(res)
	log.Println(res.(map[string]interface{})["message"])
}
