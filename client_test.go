package disgo

import (
	"encoding/json"
	"io/ioutil"
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
	apiKey := "xxxxx"
	uri := "http://localhost"

	client, err := openConnection(uri, username, apiKey)
	if err != nil {
		t.Errorf("NewClient() = got %q", err)
	}

	user := User{
		Name:     "New Member",
		Email:    "newmember@example.com",
		Username: "new-member",
		Password: "member1234567",
		Active:   true,
		Approved: true,
	}

	body, err := json.Marshal(user)
	if err != nil {
		t.Errorf("Error while json marchal data user. got %q", err)
	}

	req, err := newRequestWithBody(client, "POST", "users", body)
	if err != nil {
		t.Errorf("Error while newRequestWithBody(). got %q", err)
	}

	res, err := executeRequest(client, req)
	if err != nil {
		t.Errorf("Error while executeRequest(). got %q", err)
	}

	bodyResp, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(bodyResp))
}
