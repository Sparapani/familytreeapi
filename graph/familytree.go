package graph

// FamilyTree represents a graph of a family tree
type FamilyTree struct {
	members []*Member
}

// Member represents the struct of a family tree member
type Member struct {
	id      int            // member ID (unique)
	name    string         // name of the family member
	kinship map[int]string // kinship of that member. key is id of relative and value is the kinship
}

func NewFamilyTree() *FamilyTree {
	return &FamilyTree{
		members: []*Member{},
	}
}

func (f *FamilyTree) AddMember(name string) (id int) {
	id = len(f.members)
	f.members = append(f.members, &Member{
		id:      id,
		name:    name,
		kinship: make(map[int]string),
	})
	return
}

func (f *FamilyTree) AddKinship(idMember, idRelative int, kinship string) {
	f.members[idMember].kinship[idRelative] = kinship
}

func (f *FamilyTree) GetKinship(idMember int) (kinshipMembers []int) {
	for _, member := range f.members {
		for kinship := range member.kinship {
			if member.id == idMember {
				kinshipMembers = append(kinshipMembers, kinship)
			}
			if kinship == idMember {
				kinshipMembers = append(kinshipMembers, member.id)
			}
		}
	}
	return kinshipMembers
}
