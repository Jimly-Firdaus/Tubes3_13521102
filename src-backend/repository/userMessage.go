package repository

import (
  "TUBES3_13521102/src-backend/structs"
	"database/sql"
)

func GetAllUserMessage(db *sql.DB) (results []structs.UserMessage, err error) {
  rows, err := db.Query("SELECT * FROM UserMessage")

  if err != nil {
    panic(err)
  }

  defer rows.Close()

  for rows.Next() {
    // Set UserMessage structs to be added to results
    var userMessage = structs.UserMessage{}

    err := rows.Scan(&userMessage.UserQuestionID, &userMessage.UserQuestion, &userMessage.BotAnswer)

    if err != nil {
      panic(err)
    }

    results = append(results, userMessage)
  }

  return
}

func InsertUserMessage(db *sql.DB, newMessage structs.UserMessage) (error) {
  err := db.QueryRow("INSERT INTO UserMessage(userQuestion, botAnswer) VALUES($1, $2)", newMessage.UserQuestion, newMessage.BotAnswer)

  return err.Err()
}

func DeleteUserMessage(db *sql.DB, userMessage structs.UserMessage) (error) {
  res, err := db.Exec("DELETE FROM UserMessage WHERE userQuestionID = $1", userMessage.UserQuestionID)
  n, _ := res.RowsAffected()

  if err != nil {
    panic(err)
  }

  if n == 0 {
    return sql.ErrNoRows
  }

  return nil
}

func UpdateUserMessage(db *sql.DB, userMessage structs.UserMessage) (error) {
  res, err := db.Exec("UPDATE UserMessage SET userQuestion = $1 WHERE botAnswer = $2", userMessage.UserQuestion, userMessage.BotAnswer)
  n, _ := res.RowsAffected()

  if err != nil {
    panic(err)
  }

  if n == 0 {
    return sql.ErrNoRows
  }

  return nil
}

func GetUserMessageByID(db *sql.DB, userQuestionID int64) (results []structs.UserMessage, err error) {
  rows, err := db.Query("SELECT * FROM UserMessage WHERE userQuestionID = $1", userQuestionID)

  if err != nil {
    panic(err)
  }

  defer rows.Close()

  for rows.Next() {
    // Set UserMessage structs to be added to results
    var userMessage = structs.UserMessage{}

    err := rows.Scan(&userMessage.UserQuestionID, &userMessage.UserQuestion, &userMessage.BotAnswer)

    if err != nil {
      panic(err)
    }

    results = append(results, userMessage)
  }

  return
}
