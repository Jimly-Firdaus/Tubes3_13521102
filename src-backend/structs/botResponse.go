package structs

type BotResponse struct {
  ResponseID  int64   `json:"response_id"`
  Question    string  `json:"question"`
  Answer      string  `json:"answer"`
}
