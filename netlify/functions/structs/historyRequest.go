package structs

type HistoryRequest struct {
	HistoryID 	 int64	`json:historyId`
	HistoryTopic string `json:historyTopic`
}