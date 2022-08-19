package graph

import "fmt"

// FamilyTree represents a graph of a family tree
type FamilyTree struct {
	Members []*Member
}

// Member represents the struct of a family tree member
type Member struct {
	ID       int            // member ID (unique)
	Name     string         // name of the family member
	Parents  map[int]string // kinship of that member. key is id of relative and value is the kinship
	Children map[int]string // kinship of that member. key is id of relative and value is the kinship
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
		Parents:  make(map[int]string),
		Children: make(map[int]string),
	})
	return
}

func (f *FamilyTree) AddKinship(idMember, idRelative int, grade string) {
	if grade == Parents {
		f.Members[idMember].Parents[idRelative] = Parents
		return
	}
	f.Members[idMember].Children[idRelative] = Children
}

func (f *FamilyTree) GetMemberByID(idMember int) (memberOut Member) {
	for _, member := range f.Members {
		if member.ID == idMember {
			memberOut = Member{
				ID:       member.ID,
				Name:     member.Name,
				Parents:  member.Parents,
				Children: member.Children,
			}
			break
		}
	}
	return
}

func (f *FamilyTree) GetTotalMemberByID(idMember int) (memberOut string) {
	for _, member := range f.Members {
		if member.ID == idMember {
			memberOut = member.Name
			break
		}
	}
	return
}

func (f *FamilyTree) GetMemberIDByName(name string) (id int) {
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

func (f *FamilyTree) FindMemberRelatives(idMember int) {
	var parents string
	parents += f.Members[idMember].Name
	parents += f.FindParents(f.Members[idMember].Parents)
	fmt.Println("Parents:", parents)
}

func (f *FamilyTree) FindParents(parentsIn map[int]string) (parentsOut string) {
	if len(parentsIn) == 0 {
		return
	}
	for i := range parentsIn {
		member := f.GetMemberByID(i)
		parentsOut += " " + member.Name
		parentsOut += f.FindParents(f.Members[i].Parents)
	}
	return
}
