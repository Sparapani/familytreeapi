package services

import (
	"github.com/Sparapani/familytreeapi/graph"
)

func InitialCharge() {
	InitTree()

	AddMember("Sonny")
	AddMember("Martin")
	AddMember("Anastasia")
	AddMember("Ellen")
	AddMember("Oprah")
	AddMember("Mike")
	AddMember("Phoebe")
	AddMember("Ursula")
	AddMember("Eric")
	AddMember("Ariel")
	AddMember("Dunny")
	AddMember("Bruce")
	AddMember("Jacqueline")
	AddMember("Melody")

	AddKinship(GetMemberIDByName("Sonny"), GetMemberIDByName("Mike"), graph.Children)
	AddKinship(GetMemberIDByName("Martin"), GetMemberIDByName("Phoebe"), graph.Children)
	AddKinship(GetMemberIDByName("Martin"), GetMemberIDByName("Ursula"), graph.Children)
	AddKinship(GetMemberIDByName("Anastasia"), GetMemberIDByName("Phoebe"), graph.Children)
	AddKinship(GetMemberIDByName("Anastasia"), GetMemberIDByName("Ursula"), graph.Children)
	AddKinship(GetMemberIDByName("Ellen"), GetMemberIDByName("Eric"), graph.Children)
	AddKinship(GetMemberIDByName("Oprah"), GetMemberIDByName("Eric"), graph.Children)
	AddKinship(GetMemberIDByName("Mike"), GetMemberIDByName("Dunny"), graph.Children)
	AddKinship(GetMemberIDByName("Mike"), GetMemberIDByName("Bruce"), graph.Children)
	AddKinship(GetMemberIDByName("Phoebe"), GetMemberIDByName("Dunny"), graph.Children)
	AddKinship(GetMemberIDByName("Phoebe"), GetMemberIDByName("Bruce"), graph.Children)
	AddKinship(GetMemberIDByName("Ursula"), GetMemberIDByName("Jacqueline"), graph.Children)
	AddKinship(GetMemberIDByName("Eric"), GetMemberIDByName("Jacqueline"), graph.Children)
	AddKinship(GetMemberIDByName("Eric"), GetMemberIDByName("Melody"), graph.Children)
	AddKinship(GetMemberIDByName("Ariel"), GetMemberIDByName("Melody"), graph.Children)
}
