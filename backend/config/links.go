package config

// Links are the addresses of other micro services
// with which the interaction takes place.
type Links struct {
}

func (links Links) Validate() error {
	return nil
}
