package models

type Roster interface{}

type RosterV2 struct {
	Name         string        `json:"name"`
	Faction      string        `json:"faction"`
	Subfaction   FactionRules  `json:"subfaction"`
	ForceLeader  ForceLeader   `json:"forceLeader"`
	TotalCreated int           `json:"totalCreated"`
	Cgs          []CombatGroup `json:"cgs"`
	Version      int           `json:"version"`
	RulesVersion string        `json:"rulesVersion"`
	IsEliteForce bool          `json:"isEliteForce"`
	Player       string        `json:"player"`
	WhenCreated  string        `json:"whenCreated"`
}

type RosterV3 struct {
	Name         string        `json:"name"`
	Faction      FactionRules  `json:"faction"`
	Subfaction   FactionRules  `json:"subfaction"`
	ForceLeader  ForceLeader   `json:"forceLeader"`
	TotalCreated int           `json:"totalCreated"`
	Cgs          []CombatGroup `json:"cgs"`
	Version      int           `json:"version"`
	RulesVersion string        `json:"rulesVersion"`
	IsEliteForce bool          `json:"isEliteForce"`
	Player       string        `json:"player"`
	WhenCreated  string        `json:"whenCreated"`
}

type CombatGroup struct {
	Primary        Group    `json:"primary"`
	Secondary      Group    `json:"secondary"`
	Name           string   `json:"name"`
	IsVet          bool     `json:"isVet"`
	EnabledOptions []string `json:"enabledOptions"`
}

type FactionRules struct {
	Name         string        `json:"name"`
	EnabledRules []FactionRule `json:"enabledRules"`
}

type FactionRule struct {
	ID      string   `json:"id"`
	Options []string `json:"options"`
}

type ForceLeader struct {
	Cg       *string `json:"cg"`
	Group    *string `json:"group"`
	Unit     *string `json:"unit"`
	Position *int    `json:"position"`
}
