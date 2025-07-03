package haos

type State string

const (
	Disarmed State = "disarmed"
	Away     State = "away"
	Home     State = "home"
)

func (s State) IsValid() bool {
	switch s {
	case Disarmed:
		return true
	case Away:
		return true
	case Home:
		return true
	default:
		return false
	}
}

type Entity struct {
	EntityId string `json:"entity_id"`
	State    State  `json:"state"`
}
