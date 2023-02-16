package domain

type Championship struct {
	Clubs Clubs
}

type Clubs []Club

type Club struct {
	ID   uint
	Name string
}
