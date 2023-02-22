package domain

type Championship struct {
	Clubs Clubs
}

type Clubs []Club

type Club struct {
	ID   string
	Name string
}

type ChampionshipRepo interface {
	GetClubs() (Clubs, error)
}
