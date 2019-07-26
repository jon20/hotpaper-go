package hotpepper

// Error
const (
	ErrorConnectionFailed = "Connection Failed"
	ErrorRequest          = "Cannnot Create Request"
)

type errConnectionFailed struct {
	host string
}

func (err errConnectionFailed) Error() string {
	return "Cannot connect"
}
