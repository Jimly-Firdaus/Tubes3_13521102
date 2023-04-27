package structs

type History struct {
  HistoryID    int64          `json:"historyId"`
  Topic        string         `json:"topic"`
  Conversation []Message      `json:"conversation"`
}
