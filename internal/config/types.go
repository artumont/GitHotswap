package config

type Config struct {
	FirstRun    bool               `json:"first_run"`
	Profiles    map[string]Profile `json:"profiles"`
	Preferences Preferences        `json:"preferences"`
}

type Profile struct {
	User  string `json:"user"`
	Email string `json:"email"`
}

type Preferences struct {
	SwapMethod string `json:"swap_method"` // @param: Used if no action is specified (menu or hotswap)
	// @todo: Add more preferences & implement SwapRules (will add more swap methods when SwapRules is done)
	// SwapRules  string `json:"swap_rules"`
}
