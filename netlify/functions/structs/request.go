package structs

type Request struct {
  Id               int64   `json:"id"`
  Text             string  `json:"text"`
  Response         string  `json:"response"`
  SentTime         string  `json:"sentTime"`
  HistoryId        int64   `json:"historyId"`
  HistoryTimeStamp string  `json:"historyTimeStamp"`
  Method  string  `json:"method"`
}
