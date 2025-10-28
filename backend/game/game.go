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

func (g *Game) hasGameCompleted() {
	completedBoardStates := [][3]int{
		// Rows
		{0, 1, 2}, {3, 4, 5}, {6, 7, 8},
		// Columns
		{0, 3, 6}, {1, 4, 7}, {2, 5, 8},
		// Diagnols
		{0, 4, 8}, {2, 4, 6},
	}

	for _, completedBoard := range completedBoardStates {
		if g.Board[completedBoard[0]] == g.Board[completedBoard[1]] && g.Board[completedBoard[1]] == g.Board[completedBoard[2]] {
			g.IsGameOver = true
			winningSymbol := g.Board[completedBoard[0]]
			for playerId, symbol := range g.PlayerSymbolMapping {
				if symbol == winningSymbol {
					g.WinnerPlayer = playerId
				}
			} 
		}
	}

	// check for draw
	for _, box := range g.Board {
		if box == noSymbol {
			return
		}
	}

	// board is full
	g.IsGameOver = true
	g.WinnerPlayer = -1
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
