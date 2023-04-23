package repository

import (
  "database/sql"
  "TUBES3_13521102/src-backend/structs"
)

func GetAllHistoryMessage(db *sql.DB) (results []structs.HistoryMessage, err error) {
  rows, err := db.Query("SELECT * FROM HistoryMessage")

  if err != nil {
    panic(err)
  }

  defer rows.Close()

  for rows.Next() {
    // Set UserMessage structs to be added to results
    var historyUser = structs.HistoryMessage{}

    err := rows.Scan(&historyUser.HistoryID, &historyUser.UserQuestionID)

    if err != nil {
      panic(err)
    }

    var userMessage = structs.UserMessage{}

    errs := db.QueryRow("SELECT * FROM UserMessage WHERE userQuestionID = $1", historyUser.UserQuestionID).Scan(&userMessage.UserQuestionID, &userMessage.UserQuestion, &userMessage.BotAnswer)

    if errs != nil {
      panic(errs)
    }

    historyUser.UserHistory = userMessage

    results = append(results, historyUser)
  }

  return
}

func InsertHistoryMessage(db *sql.DB, newHistory structs.HistoryMessage) (error) {
  err := db.QueryRow("INSERT INTO HistoryMessage VALUES($1, $2)", newHistory.HistoryID, newHistory.UserQuestionID)

  return err.Err()
}

// Deleting all message for specific history
func DeleteHistoryMessage(db *sql.DB, userHistory structs.HistoryMessage) (error) {
  res, err := db.Exec("DELETE FROM HistoryMessage WHERE historyID = $1", userHistory.HistoryID)
  n, _ := res.RowsAffected()

  if err != nil {
    panic(err)
  }

  if n == 0 {
    return sql.ErrNoRows
  }

  return nil
}

func GetHistoryMessageByID(db *sql.DB, historyID int64) (results []structs.HistoryMessage, err error) {
  rows, err := db.Query("SELECT * FROM HistoryMessage WHERE historyID = $1", historyID)

  if err != nil {
    panic(err)
  }

  defer rows.Close()

  for rows.Next() {
    // Set UserMessage structs to be added to results
    var historyUser = structs.HistoryMessage{}

    err := rows.Scan(&historyUser.HistoryID, &historyUser.UserQuestionID)

    if err != nil {
      panic(err)
    }

    var userMessage = structs.UserMessage{}

    errs := db.QueryRow("SELECT * FROM UserMessage WHERE userQuestionID = $1", historyUser.UserQuestionID).Scan(&userMessage.UserQuestionID, &userMessage.UserQuestion, &userMessage.BotAnswer)

    if errs != nil {
      panic(errs)
    }

    historyUser.UserHistory = userMessage

    results = append(results, historyUser)
  }

  return
}


