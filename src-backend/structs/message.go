package structs

type Message struct {
  Id        int64   `json:"id"`
  Text      string  `json:"text"`
  Response  string  `json:"response"`
  SentTime  string  `json:"sentTime"`
  HistoryId int64   `json:"historyId"`
}
