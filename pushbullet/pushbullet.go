/*
Package "pushbullet" provides interfaces for Pushbullet HTTP API.

Pushbullet is a web service,
which makes your devices work better together by allowing you to move things between them easily.

The official url: https://www.pushbullet.com/

Currently, this package supports only "pushes" except file.

See the API documentation for the details: https://docs.pushbullet.com/#http
*/
package pushbullet

type Pushbullet struct {
	token string
}

/*
Create an instance to call Pushbullet HTTP API.
This requires the access token.
Ihe token is found in account settings.

Account settings: https://www.pushbullet.com/account
*/
func New(token string) *Pushbullet {
	pb := &Pushbullet{
		token: token,
	}

	return pb
}
