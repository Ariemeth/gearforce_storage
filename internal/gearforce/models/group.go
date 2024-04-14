package models

import "encoding/json"

type Group struct {
	Role  string `json:"role"`
	Units []Unit `json:"units"`
}

func (g *Group) UnmarshalJSON(data []byte) error {
	type Alias Group
	aux := &struct {
		Units []json.RawMessage `json:"units"`
		*Alias
	}{
		Alias: (*Alias)(g),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	for _, unitData := range aux.Units {

		var version ModelVersion
		if err := json.Unmarshal(unitData, &version); err != nil {
			return err
		}
		switch version.Version {
		case 0, 1:
			var unitV1 UnitV1
			if err := json.Unmarshal(unitData, &unitV1); err != nil {
				return err
			}
			g.Units = append(g.Units, unitV1)
		default:
			var unitV2 UnitV2
			if err := json.Unmarshal(unitData, &unitV2); err != nil {
				return err
			}
			g.Units = append(g.Units, unitV2)
		}
	}
	return nil
}
func (g *Group) MarshalJSON() ([]byte, error) {
	type Alias Group
	aux := &struct {
		Units []json.RawMessage `json:"units"`
		*Alias
	}{
		Alias: (*Alias)(g),
	}
	for _, unit := range g.Units {
		unitData, err := json.Marshal(unit)
		if err != nil {
			return nil, err
		}
		aux.Units = append(aux.Units, unitData)
	}
	return json.Marshal(aux)
}
