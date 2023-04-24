package e2e

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/yigitsadic/birthday-app-api/internal/sessions"
)

// Login makes given user login.
func Login(servUrl string) (string, error) {
	body := bytes.NewBufferString(`{
		"email":    "johndo@google.com",
		"password": "123456789"
	}`)

	endpoint := fmt.Sprintf("%s/sessions/create", servUrl)
	req, err := http.NewRequest(http.MethodPost, endpoint, body)
	if err != nil {
		return "", err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}

	if res.StatusCode != http.StatusCreated {
		return "", errors.New("status code is not 201")
	}

	var data sessions.AuthenticationModel

	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		return "", err
	}

	return data.AccessToken, nil
}
