package models

type Mod struct {
	ID       string `json:"id"`
	Selected *struct {
		Text     string `json:"text,omitempty"`
		Selected any    `json:"selected,omitempty"`
	} `json:"selected,omitempty"`
}
