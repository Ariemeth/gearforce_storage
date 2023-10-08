package models

type Group struct {
	Role  string `json:"role"`
	Units []Unit `json:"units"`
}
