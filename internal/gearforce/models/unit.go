package models

type Unit interface {
	Version() int
}

type UnitV1 struct {
	Frame           string           `json:"frame"`
	Variant         string           `json:"variant"`
	Mods            map[string][]Mod `json:"mods"`
	Command         string           `json:"command"`
	FactionOverride *string          `json:"factionOverride"`
	UnitVersion     int              `json:"version"`
}

func (u UnitV1) Version() int {
	return 1
}

type UnitV2 struct {
	Frame           string      `json:"frame"`
	Variant         string      `json:"variant"`
	Mods            []StoredMod `json:"mods"`
	Command         string      `json:"command"`
	FactionOverride *string     `json:"factionOverride"`
	UnitVersion     int         `json:"version"`
}

func (u UnitV2) Version() int {
	return u.UnitVersion
}
