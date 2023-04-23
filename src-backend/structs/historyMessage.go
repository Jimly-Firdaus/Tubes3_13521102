package structs

type HistoryMessage struct {
  UserQuestionID int64      `json:"user_question_id"`
  HistoryID      int64      `json:"history_id"`
  UserHistory   UserMessage `json:"user_history"`
}
