package graph

import "fmt"

// FamilyTree represents a graph of a family tree
type FamilyTree struct {
	Members []*Member `json:"members,omitempty"`
}

// Member represents the struct of a family tree member
type Member struct {
	ID       int              `json:"id,omitempty"`       // member ID (unique)
	Name     string           `json:"name,omitempty"`     // name of the family member
	Parents  map[int]Relative `json:"parents,omitempty"`  // kinship of that member. key is id of relative and value is the kinship
	Children map[int]Relative `json:"children,omitempty"` // kinship of that member. key is id of relative and value is the kinship
}

func NewFamilyTree() *FamilyTree {
	return &FamilyTree{
		Members: []*Member{},
	}
}

func (f *FamilyTree) AddMember(name string) (id int) {
	id = len(f.Members)
	f.Members = append(f.Members, &Member{
		ID:       id,
		Name:     name,
		Parents:  make(map[int]Relative),
		Children: make(map[int]Relative),
	})
	return
}

func (f *FamilyTree) UpdateMember(member Member) {
	f.Members[member.ID].Name = member.Name
	f.Members[member.ID].Children = member.Children
	f.Members[member.ID].Parents = member.Parents
}

func (f *FamilyTree) RemoveMember(memberID int) {
}

func (f *FamilyTree) AddKinship(idMember, idRelative int, grade Relative) {

	if f.Members[idMember] == nil {
		return
	}

	if grade == Parents {
		f.Members[idMember].Parents[idRelative] = Parents
		return
	}
	f.Members[idMember].Children[idRelative] = Children
}

func (f *FamilyTree) GetMemberByID(idMember int) (memberOut Member) {
	if f.Members[idMember] == nil {
		return
	}
	memberOut = *f.Members[idMember]
	return
}

func (f *FamilyTree) GetMemberNameByID(idMember int) (memberName string) {
	if f.Members[idMember] == nil {
		return
	}
	memberName = f.Members[idMember].Name
	return
}

func (f *FamilyTree) GetRelationshipbyEnum(rel Relative) (meaning string) {
	if rel == Parents {
		meaning = "Parents"
		return
	}
	meaning = "Children"
	return
}

func (f *FamilyTree) GetMemberIDByName(name string) (id int) {
	id = -1
	if f.Members == nil {
		return
	}
	for _, member := range f.Members {
		if member.Name == name {
			id = member.ID
			break
		}
	}
	return
}

func (f *FamilyTree) GetMembers() (membersOut []Member) {
	for _, member := range f.Members {
		membersOut = append(membersOut, Member{
			ID:       member.ID,
			Name:     member.Name,
			Parents:  member.Parents,
			Children: member.Children,
		})
	}
	return
}

func (f *FamilyTree) FindMemberRelatives(idMember int) (parents []Member) {
	fmt.Println("idMember", idMember)
	if idMember == -1 || f.Members[idMember] == nil {
		return
	}
	parents = append(parents, *f.Members[idMember])
	parents = append(parents, f.findParents(f.Members[idMember].Parents)...)
	return
}

func (f *FamilyTree) findParents(parentsIn map[int]Relative) (parentsOut []Member) {
	if len(parentsIn) == 0 {
		return
	}
	for i := range parentsIn {
		member := f.GetMemberByID(i)
		parentsOut = append(parentsOut, member)
		parentsOut = append(parentsOut, f.findParents(f.Members[i].Parents)...)
	}
	return
}
