package types

type Command struct {
	Name        string
	Identifier  string
	Description string
	IsReqAction bool
	Actions		map[string]string
	Handler     func(args map[string]string, config Config)
}