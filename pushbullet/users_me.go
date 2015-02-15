package pushbullet

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/mitsuse/bullet/pushbullet/responses"
)

// Get the current user.
func (pb *Pushbullet) GetUsersMe() (*responses.UsersMe, error) {
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

	var me *responses.UsersMe

	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&me); err != nil {
		return nil, err
	}

	return me, nil
}
