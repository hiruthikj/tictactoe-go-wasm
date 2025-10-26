package game

import "encoding/json"

type boxState int

const (
	emptyBox boxState = iota
	x
	o
)

type Action struct {
	Player   int
	Position int
}

type Game struct {
	// State
	Board           [9]boxState         `json:"board"`
	SymbolMapping   map[boxState]string `json:"symbolMapping"`
	CurrentPlayerId int                 `json:"currentPlayer"`
	// Events
	PossibleActions []Action `json:"possibleActions"`
	// Store
	IsGameOver bool `json:"isGameOver"`
}

func (g *Game) getPossibleActions() []Action {
	if g.IsGameOver {
		return []Action{}
	}

	possibleActions := []Action{}
	for i := 0; i < 9; i++ {
		if g.Board[i] == emptyBox {
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
		Board: [9]boxState{},
		SymbolMapping: map[boxState]string{
			emptyBox: "",
			x:        "X",
			o:        "O",
		},
		CurrentPlayerId: 0,
		IsGameOver:      false,
	}
	g.PossibleActions = g.getPossibleActions()

	return &g
}
