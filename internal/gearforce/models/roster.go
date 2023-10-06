package models

type Roster struct {
	Player     string `json:"player"`
	Name       string `json:"name"`
	Faction    string `json:"faction"`
	Subfaction struct {
		Name         string `json:"name"`
		EnabledRules []struct {
			ID      string   `json:"id"`
			Options []string `json:"options"`
		} `json:"enabledRules"`
	} `json:"subfaction"`
	ForceLeader struct {
		Cg       string `json:"cg"`
		Group    string `json:"group"`
		Unit     string `json:"unit"`
		Position int    `json:"position"`
	} `json:"forceLeader"`
	TotalCreated int `json:"totalCreated"`
	Cgs          []struct {
		Primary struct {
			Role  string `json:"role"`
			Units []struct {
				Frame           string           `json:"frame"`
				Variant         string           `json:"variant"`
				Mods            map[string][]Mod `json:"mods"`
				Command         string           `json:"command"`
				FactionOverride *string          `json:"factionOverride"`
			} `json:"units"`
		} `json:"primary"`
		Secondary struct {
			Role  string `json:"role"`
			Units []struct {
				Frame           string           `json:"frame"`
				Variant         string           `json:"variant"`
				Mods            map[string][]Mod `json:"mods"`
				Command         string           `json:"command"`
				FactionOverride *string          `json:"factionOverride"`
			} `json:"units"`
		} `json:"secondary"`
		Name           string   `json:"name"`
		IsVet          bool     `json:"isVet"`
		EnabledOptions []string `json:"enabledOptions"`
	} `json:"cgs"`
	Version      int    `json:"version"`
	RulesVersion string `json:"rulesVersion"`
	IsEliteForce bool   `json:"isEliteForce"`
	WhenCreated  string `json:"whenCreated"`
}
