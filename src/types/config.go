package types

type Config struct {
	FirstRun bool               	`json:"first_run"`
	Profiles map[string]Profile 	`json:"profiles"`
	Preferences map[string]string 	`json:"preferences"`
}

type Profile struct {
	User  string 	`json:"user"`
	Email string 	`json:"email"`
}

type Preferences struct {
	SwapMethod string `json:"swap_method"` // @param: Used if no action is specified (menu, active, hotswap)
	// @todo: Add more preferences & implement SwapRules
	// SwapRules  string `json:"swap_rules"`
}