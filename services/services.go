package services

import (
	"github.com/Sparapani/familytreeapi/graph"
	"github.com/Sparapani/familytreeapi/models"
)

var familyTree *graph.FamilyTree

func InitTree() {
	familyTree = graph.NewFamilyTree()
}

func GetAllMembers() (membersOut []models.MemberAPI) {
	if familyTree.Members == nil {
		return
	}
	for _, member := range familyTree.Members {
		memberTemp := familyTree.GetMemberByID(member.ID)
		memberOut := models.MemberAPI{Name: memberTemp.Name}
		for i, relationship := range memberTemp.Children {
			memberOut.Relationships = append(memberOut.Relationships, models.RelationShip{
				Name:         familyTree.GetMemberNameByID(i),
				Relationship: familyTree.GetRelationshipbyEnum(relationship),
			})
		}
		for i, relationship := range memberTemp.Parents {
			memberOut.Relationships = append(memberOut.Relationships, models.RelationShip{
				Name:         familyTree.GetMemberNameByID(i),
				Relationship: familyTree.GetRelationshipbyEnum(relationship),
			})
		}
		for i, relationship := range memberTemp.Spouse {
			memberOut.Relationships = append(memberOut.Relationships, models.RelationShip{
				Name:         familyTree.GetMemberNameByID(i),
				Relationship: familyTree.GetRelationshipbyEnum(relationship),
			})
		}
		for i, relationship := range memberTemp.Cousins {
			memberOut.Relationships = append(memberOut.Relationships, models.RelationShip{
				Name:         familyTree.GetMemberNameByID(i),
				Relationship: familyTree.GetRelationshipbyEnum(relationship),
			})
		}
		membersOut = append(membersOut, memberOut)
	}
	return
}

func GetRelativesByID(member graph.Member) (relativesOut []models.MemberAPI) {
	relatives := familyTree.FindMemberRelatives(familyTree.GetMemberIDByName(member.Name))
	for _, relative := range relatives {
		relativeTemp := familyTree.GetMemberByID(relative.ID)
		relativeOut := models.MemberAPI{Name: relativeTemp.Name}
		for j := range relativeTemp.Parents {
			relativeOut.Relationships = append(relativeOut.Relationships, models.RelationShip{
				Name:         familyTree.GetMemberNameByID(j),
				Relationship: "Parents",
			})
		}
		relativesOut = append(relativesOut, relativeOut)
	}
	return
}

func GetBaconsNumber(members models.BaconsNumber) (baconNumbes int) {
	if members.NameFrom == "" || members.NameTo == "" {
		return -1
	}
	baconNumbes, _ = familyTree.GetBaconsNumber(familyTree.GetMemberIDByName(members.NameFrom), familyTree.GetMemberIDByName(members.NameTo), -1, false, make(map[int]bool))
	return
}

func GetMemberIDByName(name string) int {
	return familyTree.GetMemberIDByName(name)
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
