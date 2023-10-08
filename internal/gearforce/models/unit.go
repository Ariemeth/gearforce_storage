package models

type Unit struct {
	Frame           string           `json:"frame"`
	Variant         string           `json:"variant"`
	Mods            map[string][]Mod `json:"mods"`
	Command         string           `json:"command"`
	FactionOverride *string          `json:"factionOverride"`
}
