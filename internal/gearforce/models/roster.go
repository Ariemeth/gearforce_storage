package models

type RosterBase struct {
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
		Cg       *string `json:"cg"`
		Group    *string `json:"group"`
		Unit     *string `json:"unit"`
		Position *int    `json:"position"`
	} `json:"forceLeader"`
	TotalCreated int `json:"totalCreated"`
	Cgs          []struct {
		Primary        Group    `json:"primary"`
		Secondary      Group    `json:"secondary"`
		Name           string   `json:"name"`
		IsVet          bool     `json:"isVet"`
		EnabledOptions []string `json:"enabledOptions"`
	} `json:"cgs"`
	Version      int    `json:"version"`
	RulesVersion string `json:"rulesVersion"`
	IsEliteForce bool   `json:"isEliteForce"`
}

type RosterMetadata struct {
	Player      string `json:"player"`
	WhenCreated string `json:"whenCreated"`
}

type Roster struct {
	RosterBase
	RosterMetadata
}

func (r Roster) ToRosterStorage(key string) RosterStorage {
	return RosterStorage{RosterBase: r.RosterBase, Key: key}
}

type RosterStorage struct {
	Key string `json:"_key"`
	RosterBase
}
