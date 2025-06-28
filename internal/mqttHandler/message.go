package mqtthandler

import "fmt"

type Message struct {
	Mac      string  `json:"mac"`
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	IDType   int     `json:"idType"`
	RSSIat1m int     `json:"rssi@1m"`
	RSSI     int     `json:"rssi"`
	Raw      float32 `json:"raw"`
	Distance float32 `json:"distance"`
	Var      float32 `json:"var"`
	Int      int     `json:"int"`
}

func (m Message) String() string {
	return fmt.Sprintf("(id: %s, name: %s, rssi: %d)", m.ID, m.Name, m.RSSI)
}
