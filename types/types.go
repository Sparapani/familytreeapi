package types

type MemberAPI struct {
	Name          string         `json:"name,omitempty"`
	Relationships []RelationShip `json:"relationships,omitempty"`
}

type RelationShip struct {
	Name         string `json:"name,omitempty"`
	Relationship string `json:"relationship,omitempty"`
}
