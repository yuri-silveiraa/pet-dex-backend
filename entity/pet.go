package entity

type Pet struct {
	Id int
	Name string
	BreedId string
	Size string
	Weight float32
	AdoptionDate int
	Birthdate int
	Comorbidity *string
	Tags string
	Castrated bool
	AvaliableToAdoption bool
	UserId *string
}
type PetDetails struct {
	Breed string
	Age int
	Size string
}