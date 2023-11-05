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
	WhenCreated  string `json:"whenCreated"`
}

func (r Roster) ToRosterStorage(key string) RosterStorage {
	return RosterStorage{Roster: r, Key: key}
}

type RosterStorage struct {
	Key string `json:"_key"`
	Roster
}
