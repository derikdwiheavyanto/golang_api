package exception

type Unauthorized struct {
	Error string
}

func NewUnauthorized() Unauthorized {
	return Unauthorized{Error: "Unauthorized"}
}
