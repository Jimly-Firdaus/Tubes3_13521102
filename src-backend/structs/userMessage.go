package structs

type UserMessage struct {
  UserQuestionID int64   `json:"user_question_id"`
  UserQuestion   string  `json:"user_question"`
  BotAnswer      string  `json:"bot_answer"`
}
