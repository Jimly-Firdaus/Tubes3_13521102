package structs

type HistoryPayload struct {
	HistoryCollection  []HistoryRequest `json:historyCollection`
}