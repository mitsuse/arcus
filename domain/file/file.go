// Package `file` Provides the type to represent uploaded file and the service to upload a file.
package file

// `Type` represents the information of an uploaded file.
type Type struct {
	name string
	t    string
	url  string
}

func (f *Type) Name() string {
	return f.name
}

func (f *Type) Type() string {
	return f.t
}

func (f *Type) Url() string {
	return f.url
}
