package gameuser

type GameUser struct {
	UUID     string
	Nickname string
	IsActive bool
}

type GameUserGame struct {
	UserID string
	GameID string
}
