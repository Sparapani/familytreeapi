package services

import (
	"github.com/Sparapani/familytreeapi/graph"
	"github.com/Sparapani/familytreeapi/types"
)

var familyTree *graph.FamilyTree

func InitTree() {
	familyTree = graph.NewFamilyTree()
}

func GetAllMembers() (membersOut []types.MemberAPI) {
	if familyTree.Members == nil {
		return
	}
	for _, member := range familyTree.Members {
		memberTemp := familyTree.GetMemberByID(member.ID)
		memberOut := types.MemberAPI{Name: memberTemp.Name}
		for i, relationship := range memberTemp.Children {
			memberOut.Relationships = append(memberOut.Relationships, types.RelationShip{
				Name:         familyTree.GetMemberNameByID(i),
				Relationship: familyTree.GetRelationshipbyEnum(relationship),
			})
		}
		for i, relationship := range memberTemp.Parents {
			memberOut.Relationships = append(memberOut.Relationships, types.RelationShip{
				Name:         familyTree.GetMemberNameByID(i),
				Relationship: familyTree.GetRelationshipbyEnum(relationship),
			})
		}
		for i, relationship := range memberTemp.Spouse {
			memberOut.Relationships = append(memberOut.Relationships, types.RelationShip{
				Name:         familyTree.GetMemberNameByID(i),
				Relationship: familyTree.GetRelationshipbyEnum(relationship),
			})
		}
		for i, relationship := range memberTemp.Cousins {
			memberOut.Relationships = append(memberOut.Relationships, types.RelationShip{
				Name:         familyTree.GetMemberNameByID(i),
				Relationship: familyTree.GetRelationshipbyEnum(relationship),
			})
		}
		membersOut = append(membersOut, memberOut)
	}
	return
}

func GetMemberIDByName(name string) int {
	return familyTree.GetMemberIDByName(name)
}

func GetRelativesByID(member graph.Member) (relativesOut []types.MemberAPI) {
	relatives := familyTree.FindMemberRelatives(familyTree.GetMemberIDByName(member.Name))
	for _, relative := range relatives {
		relativeTemp := familyTree.GetMemberByID(relative.ID)
		relativeOut := types.MemberAPI{Name: relativeTemp.Name}
		for j := range relativeTemp.Parents {
			relativeOut.Relationships = append(relativeOut.Relationships, types.RelationShip{
				Name:         familyTree.GetMemberNameByID(j),
				Relationship: "Parents",
			})
		}
		relativesOut = append(relativesOut, relativeOut)
	}
	return
}

func AddMember(member string) int {
	if member == "" {
		return -1
	}
	return familyTree.AddMember(member)

}

func AddKinship(idMember, idRelative int, grade graph.Relative) {
	familyTree.AddKinship(idMember, idRelative, grade)
}

func UpdateMember(member graph.Member) (ok bool) {
	if member.Name == "" {
		return
	}
	familyTree.UpdateMember(member)
	ok = true
	return
}

func RemoveMember(member graph.Member) (ok bool) {
	if member.Name == "" {
		return
	}
	member.ID = familyTree.GetMemberIDByName(member.Name)
	if member.ID == -1 {
		return
	}
	familyTree.RemoveMember(member.ID)
	ok = true
	return
}
