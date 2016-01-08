package resource

// `Location` represents the location of resource such as file, web page etc.
type Location string

// `IsLocal` returns `true` when the raw value represents the location of a file local file system.
func (l Location) IsLocal() bool {
	// TODO:
	return false
}

// `IsRemote` returns `true` when the raw value represents the location of resource on HTTP server..
func (l Location) IsRemote() bool {
	// TODO:
	return false
}
