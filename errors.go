package hotpepper

type errConnectionFailed struct {
	host string
}

func (err errConnectionFailed) Error() string {
	return "Cannot connect"
}
