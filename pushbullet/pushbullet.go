package pushbullet

type Pushbullet struct {
	token string
}

func New(token string) *Pushbullet {
	pb := &Pushbullet{
		token: token,
	}

	return pb
}
