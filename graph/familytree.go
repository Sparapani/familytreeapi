package graph

// FamilyTree represents a graph of a family tree
type FamilyTree struct {
	Members []*Member `json:"members,omitempty"`
}

// Member represents the struct of a family tree member
type Member struct {
	ID       int              `json:"id,omitempty"`       // member ID (unique)
	Name     string           `json:"name,omitempty"`     // name of the family member
	Parents  map[int]Relative `json:"parents,omitempty"`  // parents of that member. key is id of relative and value is the kinship
	Children map[int]Relative `json:"children,omitempty"` // children of that member. key is id of relative and value is the kinship
	Spouse   map[int]Relative `json:"spouse,omitempty"`   // spouse of that member. key is id of relative and value is the kinship
	Cousins  map[int]Relative `json:"cousins,omitempty"`  // cousins of that member. key is id of relative and value is the kinship
}

var relativeFound map[int]bool
var personSearch map[int]bool

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
		Spouse:   make(map[int]Relative),
		Cousins:  make(map[int]Relative),
	})
	return
}

func (f *FamilyTree) AddKinship(idMember, idRelative int, grade Relative) {

	if f.Members[idMember] == nil {
		return
	}

	switch grade {
	case Parents:
		f.Members[idMember].Parents[idRelative] = Parents
	case Children:
		f.Members[idMember].Children[idRelative] = Children
		f.checkSpouse(idMember, idRelative)
		f.checkCousins(idMember, idRelative)
	case Spouse:
		f.Members[idMember].Spouse[idRelative] = Spouse
		f.Members[idRelative].Spouse[idMember] = Spouse
	case Cousins:
		f.Members[idMember].Cousins[idRelative] = Cousins
		f.Members[idRelative].Cousins[idMember] = Cousins

	}
}

func (f *FamilyTree) checkSpouse(idMember, idRelative int) {
	for i, member := range f.Members {
		for j := range member.Children {
			if j == idRelative && i != idMember {
				f.AddKinship(i, idMember, Spouse)
			}
		}
	}
}

func (f *FamilyTree) checkCousins(idMember, idRelative int) {
	for _, member := range f.Members {
		for i := range member.Parents {
			for j := range f.Members[i].Children {
				for l := range f.Members[j].Children {
					if idRelative != l {
						f.AddKinship(idRelative, l, Cousins)
					}
				}
			}
		}
	}
}

func (f *FamilyTree) UpdateMember(member Member) {
	f.Members[member.ID].Name = member.Name
	f.Members[member.ID].Children = member.Children
	f.Members[member.ID].Parents = member.Parents
	f.Members[member.ID].Spouse = member.Spouse
	f.Members[member.ID].Cousins = member.Cousins
}

func (f *FamilyTree) RemoveMember(memberID int) {
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
	switch rel {
	case Parents:
		meaning = "Parents"
	case Children:
		meaning = "Children"
	case Spouse:
		meaning = "Spouse"
	case Cousins:
		meaning = "Cousins"
	}
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
			Spouse:   member.Spouse,
			Cousins:  member.Cousins,
		})
	}
	return
}

func (f *FamilyTree) GetBaconsNumber(idActually, idToFound, baconsNumber int, found bool, personActually map[int]bool) (int, bool) {
	if idActually == idToFound {
		baconsNumber++
		return baconsNumber, true
	}
	baconsNumber++
	personSearch = make(map[int]bool)
	personSearch = personActually
	personSearch[idActually] = true
	for i := range f.Members[idActually].Parents {
		if found {
			return baconsNumber, found
		}
		if !personSearch[i] {
			baconsNumber, found = f.GetBaconsNumber(i, idToFound, baconsNumber, found, personSearch)

		}
	}
	for i := range f.Members[idActually].Children {
		if found {
			return baconsNumber, found
		}
		if !personSearch[i] {
			baconsNumber, found = f.GetBaconsNumber(i, idToFound, baconsNumber, found, personSearch)
		}
	}
	baconsNumber--
	return baconsNumber, found
}

func (f *FamilyTree) FindMemberRelatives(idMember int) (relatives []Member) {
	if idMember == -1 || f.Members[idMember] == nil {
		return
	}
	relativeFound = make(map[int]bool)
	relatives = append(relatives, *f.Members[idMember])
	relativeFound[idMember] = true
	relatives = append(relatives, f.findChildren(f.Members[idMember].Children)...)
	relatives = append(relatives, f.findParents(f.Members[idMember].Parents)...)
	return
}

func (f *FamilyTree) findParents(parentsIn map[int]Relative) (parentsOut []Member) {
	if len(parentsIn) == 0 {
		return
	}
	for i := range parentsIn {
		member := f.GetMemberByID(i)
		if !relativeFound[member.ID] {
			parentsOut = append(parentsOut, member)
			relativeFound[member.ID] = true
		}
		parentsOut = append(parentsOut, f.findParents(f.Members[i].Parents)...)
		parentsOut = append(parentsOut, f.findChildren(f.Members[i].Children)...)
	}
	return
}

func (f *FamilyTree) findChildren(childrenIn map[int]Relative) (childrenOut []Member) {
	if len(childrenIn) == 0 {
		return
	}
	for i := range childrenIn {
		member := f.GetMemberByID(i)
		if !relativeFound[member.ID] {
			childrenOut = append(childrenOut, member)
			relativeFound[member.ID] = true
		}
		childrenOut = append(childrenOut, f.findChildren(f.Members[i].Children)...)
	}
	return
}
