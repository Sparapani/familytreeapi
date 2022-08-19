package main

import (
	"fmt"

	"github.com/Sparapani/familytreeapi/graph"
)

func main() {
	familyTree := graph.NewFamilyTree()
	familyTree.AddMember("Sonny")
	familyTree.AddMember("Martin")
	familyTree.AddMember("Anastasia")
	familyTree.AddMember("Ellen")
	familyTree.AddMember("Oprah")

	familyTree.AddMember("Mike")
	familyTree.AddMember("Phoebe")
	familyTree.AddMember("Ursula")
	familyTree.AddMember("Eric")
	familyTree.AddMember("Ariel")

	familyTree.AddMember("Dunny")
	familyTree.AddMember("Bruce")
	familyTree.AddMember("Jacqueline")
	familyTree.AddMember("Melody")

	familyTree.AddKinship(familyTree.GetMemberIDByName("Sonny"), familyTree.GetMemberIDByName("Mike"), graph.Children)

	familyTree.AddKinship(familyTree.GetMemberIDByName("Martin"), familyTree.GetMemberIDByName("Phoebe"), graph.Children)
	familyTree.AddKinship(familyTree.GetMemberIDByName("Martin"), familyTree.GetMemberIDByName("Ursula"), graph.Children)

	familyTree.AddKinship(familyTree.GetMemberIDByName("Anastasia"), familyTree.GetMemberIDByName("Phoebe"), graph.Children)
	familyTree.AddKinship(familyTree.GetMemberIDByName("Anastasia"), familyTree.GetMemberIDByName("Ursula"), graph.Children)

	familyTree.AddKinship(familyTree.GetMemberIDByName("Ellen"), familyTree.GetMemberIDByName("Eric"), graph.Children)

	familyTree.AddKinship(familyTree.GetMemberIDByName("Oprah"), familyTree.GetMemberIDByName("Eric"), graph.Children)

	familyTree.AddKinship(familyTree.GetMemberIDByName("Mike"), familyTree.GetMemberIDByName("Sonny"), graph.Parents)
	familyTree.AddKinship(familyTree.GetMemberIDByName("Mike"), familyTree.GetMemberIDByName("Dunny"), graph.Children)
	familyTree.AddKinship(familyTree.GetMemberIDByName("Mike"), familyTree.GetMemberIDByName("Bruce"), graph.Children)

	familyTree.AddKinship(familyTree.GetMemberIDByName("Phoebe"), familyTree.GetMemberIDByName("Martin"), graph.Parents)
	familyTree.AddKinship(familyTree.GetMemberIDByName("Phoebe"), familyTree.GetMemberIDByName("Anastasia"), graph.Parents)
	familyTree.AddKinship(familyTree.GetMemberIDByName("Phoebe"), familyTree.GetMemberIDByName("Dunny"), graph.Children)
	familyTree.AddKinship(familyTree.GetMemberIDByName("Phoebe"), familyTree.GetMemberIDByName("Bruce"), graph.Children)

	familyTree.AddKinship(familyTree.GetMemberIDByName("Ursula"), familyTree.GetMemberIDByName("Martin"), graph.Parents)
	familyTree.AddKinship(familyTree.GetMemberIDByName("Ursula"), familyTree.GetMemberIDByName("Anastasia"), graph.Parents)
	familyTree.AddKinship(familyTree.GetMemberIDByName("Ursula"), familyTree.GetMemberIDByName("Jacqueline"), graph.Children)

	familyTree.AddKinship(familyTree.GetMemberIDByName("Eric"), familyTree.GetMemberIDByName("Ellen"), graph.Parents)
	familyTree.AddKinship(familyTree.GetMemberIDByName("Eric"), familyTree.GetMemberIDByName("Oprah"), graph.Parents)
	familyTree.AddKinship(familyTree.GetMemberIDByName("Eric"), familyTree.GetMemberIDByName("Jacqueline"), graph.Children)
	familyTree.AddKinship(familyTree.GetMemberIDByName("Eric"), familyTree.GetMemberIDByName("Melody"), graph.Children)

	familyTree.AddKinship(familyTree.GetMemberIDByName("Ariel"), familyTree.GetMemberIDByName("Melody"), graph.Children)

	familyTree.AddKinship(familyTree.GetMemberIDByName("Dunny"), familyTree.GetMemberIDByName("Mike"), graph.Parents)
	familyTree.AddKinship(familyTree.GetMemberIDByName("Dunny"), familyTree.GetMemberIDByName("Phoebe"), graph.Parents)

	familyTree.AddKinship(familyTree.GetMemberIDByName("Bruce"), familyTree.GetMemberIDByName("Mike"), graph.Parents)
	familyTree.AddKinship(familyTree.GetMemberIDByName("Bruce"), familyTree.GetMemberIDByName("Phoebe"), graph.Parents)

	familyTree.AddKinship(familyTree.GetMemberIDByName("Jacqueline"), familyTree.GetMemberIDByName("Ursula"), graph.Parents)
	familyTree.AddKinship(familyTree.GetMemberIDByName("Jacqueline"), familyTree.GetMemberIDByName("Eric"), graph.Parents)

	familyTree.AddKinship(familyTree.GetMemberIDByName("Melody"), familyTree.GetMemberIDByName("Eric"), graph.Parents)
	familyTree.AddKinship(familyTree.GetMemberIDByName("Melody"), familyTree.GetMemberIDByName("Ariel"), graph.Parents)

	familyMembers := familyTree.GetMembers()
	fmt.Println("Member ID - Member Name - Relatives")
	for i := range familyMembers {
		fmt.Println(familyTree.GetMemberByID(i))
	}

	familyTree.FindMemberRelatives(11)
}
