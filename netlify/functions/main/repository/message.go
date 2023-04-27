package repository

import (
	"TUBES3_13521102/netlify/functions/main/structs"
	"database/sql"
)

func InsertUserMessage(db *sql.DB, newMessage structs.Message) (error) {
  err := db.QueryRow("INSERT INTO UserMessage(userQuestion, botAnswer, sentTime) VALUES(?, ?, ?)", newMessage.Text, newMessage.Response, newMessage.SentTime)

  if err != nil {
    return err.Err()
  }

  // Checking if History table already made for current message
  res, errs := db.Query("SELECT * FROM History WHERE historyID = ?", newMessage.HistoryId)

  if errs != nil {
    panic(errs)
  }

  defer res.Close()

  count := 0

  for res.Next() {
    count++
  }

  if count == 0 {
    var history = structs.History{}

    history.HistoryID = newMessage.HistoryId
    history.Topic = newMessage.Text

    InsertHistory(db, history, newMessage.HistoryTimeStamp)
  }

  // Add to HistoryMessage table
  err = db.QueryRow("INSERT INTO HistoryMessage VALUES(?, ?)", newMessage.HistoryId, newMessage.Id)

  return err.Err()
}


func GetUserMessageByID(db *sql.DB, userQuestionID int64) (results []structs.Message, err error) {
  rows, err := db.Query("SELECT * FROM UserMessage WHERE userQuestionID = ?", userQuestionID)

  if err != nil {
    panic(err)
  }

  defer rows.Close()

  for rows.Next() {
    // Set UserMessage structs to be added to results
    var userMessage = structs.Message{}

    err := rows.Scan(&userMessage.Id, &userMessage.Text, &userMessage.Response, &userMessage.SentTime)

    if err != nil {
      panic(err)
    }

    // Query to search for historyID
    row, err := db.Query("SELECT historyID FROM HistoryMessage WHERE userQuestionID = ?", userMessage.Id)

    if err != nil {
      panic(err)
    }

    for row.Next() {
      row.Scan(&userMessage.HistoryId);
    }

    results = append(results, userMessage)
  }

  return
}
