package file

// `Remote` represents the information of an uploaded file.
type Remote struct {
	name string
	t    string
	url  string
}

func (r *Remote) Name() string {
	return r.name
}

func (r *Remote) Type() string {
	return r.t
}

func (r *Remote) Url() string {
	return r.url
}
