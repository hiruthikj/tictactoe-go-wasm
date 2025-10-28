package game

import "encoding/json"

type boxSymbol int

const (
	noSymbol boxSymbol = iota
	xSymbol
	oSymbol
)

func (b boxSymbol) String() string {
	mapping := map[boxSymbol]string{
		noSymbol: "",
		xSymbol: "X",
		oSymbol: "O",
	}
	return mapping[b]
}

type Action struct {
	Player   int
	Position int
}

type Game struct {
	// State
	Board               [9]boxSymbol         `json:"board"`
	PlayerSymbolMapping map[int]boxSymbol `json:"playerSymbolMapping"`
	CurrentPlayerId     int                 `json:"currentPlayer"`
	// Events
	PossibleActions []Action `json:"possibleActions"`
	// Store
	IsGameOver   bool `json:"isGameOver"`
	WinnerPlayer int  `json:"winnerPlayer"`
}

func (g *Game) getPossibleActions() []Action {
	if g.IsGameOver {
		return []Action{}
	}

	possibleActions := []Action{}
	for i := 0; i < 9; i++ {
		if g.Board[i] == noSymbol {
			action := Action{Player: g.CurrentPlayerId, Position: i}
			possibleActions = append(possibleActions, action)
		}
	}
	return possibleActions
}

func (g *Game) Serialize() ([]byte, error) {
	bytes, err := json.Marshal(g)
	return bytes, err
}

func NewGame() *Game {
	g := Game{
		Board: [9]boxSymbol{},
		PlayerSymbolMapping: map[int]boxSymbol{
			0: xSymbol,
			1: oSymbol,
		},
		CurrentPlayerId: 0,
		IsGameOver:      false,
	}
	g.PossibleActions = g.getPossibleActions()

	return &g
}
