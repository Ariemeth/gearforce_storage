package models

type Mod struct {
	ID       string    `json:"id"`
	Selected *Selected `json:"selected,omitempty"`
}

type Selected struct {
	Text     string    `json:"text,omitempty"`
	Selected *Selected `json:"selected"`
}
