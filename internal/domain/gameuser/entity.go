package gameuser

type GameUser struct {
	UUID     string
	Nickname string
	isActive bool
}

type GameUserGame struct {
	UserID string
	GameID string
}
