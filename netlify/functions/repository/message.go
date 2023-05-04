package repository

import (
	"TUBES3_13521102/netlify/functions/structs"
	"database/sql"
)

func InsertMessages(db *sql.DB, newMessage structs.Request) (error) {
  err := db.QueryRow("INSERT INTO Messages(userQuestion, botAnswer, sentTime, historyID) VALUES(?, ?, ?, ?)", newMessage.Text, newMessage.Response, newMessage.SentTime, newMessage.HistoryId)

  if err != nil {
    return err.Err()
  }

  return nil
}

func DeleteMessages(db *sql.DB, userQuestionID int) (error) {
  err := db.QueryRow("DELETE FROM Messages WHERE userQuestionID = ?", userQuestionID)

  if err != nil {
    return err.Err()
  }

  return nil
}

func GetHistoryMessages(db *sql.DB, historyID int) (results []structs.Message, err error) {
  rows, err := db.Query("SELECT * FROM Messages WHERE historyID = ?", historyID)

  if err != nil {
    panic(err)
  }

  defer rows.Close()

  for rows.Next() {
    var message = structs.Message{}

    err = rows.Scan(&message.Id, &message.Text, &message.Response, &message.SentTime, &message.HistoryId)

    if err != nil {
      panic(err)
    }

    results = append(results, message)
  }

  return
}

func GetMessagesByID(db *sql.DB, userQuestionID int64) (structs.Message, error) {
  rows, err := db.Query("SELECT * FROM Messages WHERE userQuestionID = ?", userQuestionID)

  if err != nil {
    panic(err)
  }

  defer rows.Close()
  
  var Messages = structs.Message{}

  for rows.Next() {

    err := rows.Scan(&Messages.Id, &Messages.Text, &Messages.Response, &Messages.SentTime, &Messages.HistoryId)

    if err != nil {
      panic(err)
    }

  }

  return Messages, nil
}

