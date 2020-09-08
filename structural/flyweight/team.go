package flyweight

type Team struct {
	ID             uint64
	Name           uint64
	Shield         []byte
	Players        []Player
	HistoricalData []HistoricalData
}

const (
	TeamA = iota
	TeamB
)
