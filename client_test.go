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
	apiKey := "xxxxxxxx"
	uri := "http://localhost"

	client, err := openConnection(uri, username, apiKey)
	if err != nil {
		t.Errorf("NewClient() = got %q", err)
	}

	user := User{
		Name:     "Hamid Sudargo",
		Email:    "hamid@asadad.com",
		Username: "hamid",
		Password: "123456",
		Active:   true,
		Approved: true,
	}

	res, err := client.CreateUser(user)
	if err != nil {
		t.Errorf("Error while executeRequest(). got %q", err)
	}

	log.Println(res)
}
