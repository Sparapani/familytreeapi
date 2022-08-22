// Package models implements the models to support the tree
package models

type MemberAPI struct {
	Name          string         `json:"name,omitempty"`
	Relationships []RelationShip `json:"relationships,omitempty"`
}
type BaconsNumber struct {
	NameFrom string `json:"name_from,omitempty"`
	NameTo   string `json:"name_to,omitempty"`
}

type RelationShip struct {
	Name         string `json:"name,omitempty"`
	Relationship string `json:"relationship,omitempty"`
}
