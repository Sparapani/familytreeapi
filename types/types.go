package types

type MembersResponse struct {
	Name          string
	Relationships []RelationShip
}

type RelationShip struct {
	Name         string
	Relationship string
}
