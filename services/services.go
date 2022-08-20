package services

import (
	"fmt"

	"github.com/Sparapani/familytreeapi/graph"
)

var familyTree *graph.FamilyTree

func InitTree() {
	familyTree = graph.NewFamilyTree()
}

func GetAllMembers() (members []graph.Member) {
	if familyTree.Members == nil {
		return
	}
	for i := range familyTree.Members {
		members = append(members, familyTree.GetMemberByID(i))
	}
	return
}

func GetMemberIDByName(name string) int {
	return familyTree.GetMemberIDByName(name)
}

func GetRelativeMember() {
	fmt.Println(familyTree.FindMemberRelatives(11))
}

func AddMember(member string) {
	familyTree.AddMember(member)
}

func AddKinship(idMember, idRelative int, grade graph.Relative) {
	familyTree.AddKinship(idMember, idRelative, grade)
}
