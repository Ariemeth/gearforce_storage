package models

type StoredMod struct {
	Type  string `json:"type"`
	Order int    `json:"order"`
	Mod   Mod    `json:"mod"`
}

type Mod struct {
	ID       string    `json:"id"`
	Selected *Selected `json:"selected,omitempty"`
}

type Selected struct {
	Text     string    `json:"text,omitempty"`
	Selected *Selected `json:"selected"`
}
