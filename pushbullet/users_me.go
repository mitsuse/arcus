package pushbullet

import (
	"encoding/json"
	"errors"
	"net/http"
)

// Get the current user.
func (pb *Pushbullet) GetUsersMe() (*UsersMeRes, error) {
	req, err := http.NewRequest("GET", ENDPOINT_USERS_ME, nil)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(pb.token, "")

	// TODO: Set the timeout.
	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	// TODO: Return an error value with human friendly message.
	if res.StatusCode != 200 {
		return nil, errors.New(res.Status)
	}

	var usersMeRes *UsersMeRes

	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&usersMeRes); err != nil {
		return nil, err
	}

	return usersMeRes, nil
}

type UsersMeRes struct {
	Iden            string        `json:"iden"`
	Email           string        `json:"email"`
	EmailNormalized string        `json:"email_normalized"`
	Name            string        `json:"name"`
	ImageUrl        string        `json:"image_url"`
	Created         float64       `json:"created"`
	Modified        float64       `json:"modified"`
	Preferences     *UsersMePrefs `json:"preferences"`
}

type UsersMePrefs struct {
	OnBoarding *UsersMeOnBording `json:"onboarding"`
	Social     bool              `json:"social"`
}

type UsersMeOnBording struct {
	App       bool `json:"app"`
	Friends   bool `json:"friends"`
	Extension bool `json:"extension"`
}
