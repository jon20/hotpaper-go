package hotpepper

// Error
const (
	ErrorConnectionFailed = "Connection Failed"
)

type errConnectionFailed struct {
	host string
}

func (err errConnectionFailed) Error() string {
	return "Cannot connect"
}
