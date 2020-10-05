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
	username := "hamid"
	apiKey := "2f0e3398000a0a862fcc74eeb2aa2e97922bf32d2a497854f0eb15f98c036e56"
	uri := "https://community.sparks.id"

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
	log.Println(res.(map[string]interface{})["message"])
}
