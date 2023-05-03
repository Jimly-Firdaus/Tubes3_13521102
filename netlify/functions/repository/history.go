package repository

import (
	"TUBES3_13521102/netlify/functions/structs"
	"database/sql"
	"time"
)

func InsertHistory(db *sql.DB, history structs.History, datetime string) (error) {

  layoutFormat := "2006-01-02 15:04:05"
  date, err := time.Parse(layoutFormat, datetime)

  if err != nil {
    panic(err)
  }

  errs := db.QueryRow("INSERT INTO History VALUES (?, ?, ?)", history.HistoryID, history.Topic, date)

  return errs.Err()
}


func DeleteHistory(db *sql.DB, historyID int) (error) {
  // Getting all usermessageID that is related to the history to delete
  messageRows, err := db.Query("SELECT userQuestionID FROM HistoryMessage WHERE historyID = ?", historyID)

  if err != nil {
    panic(err)
  }

  defer messageRows.Close()

  messageIDs := make([]int, 0, 5)
  for messageRows.Next() {
    var messageID int
    messageRows.Scan(&messageID)
    messageIDs = append(messageIDs, messageID)
  }

  // Deleting from foreign key table first
  _, errs := db.Exec("DELETE FROM HistoryMessage WHERE historyID = ?", historyID)

  if errs != nil {
    panic(errs)
  }

  // Deleting all question from userMessage table
  for i := 0; i < len(messageIDs); i++ {
    err = DeleteUserMessage(db, messageIDs[i])
    
    if err != nil {
      panic(err)
    }
  }

  res, errs := db.Exec("DELETE FROM History WHERE historyID = ?", historyID)
  n, _ := res.RowsAffected()

  if errs != nil {
    panic(errs)
  }

  if n == 0 {
    return sql.ErrNoRows
  }

  return nil
}

func UpdateHistory(db *sql.DB, history structs.History) (error) {
  res, errs := db.Exec("UPDATE History SET historyTitle = ? WHERE historyID = ?", history.Topic, history.HistoryID)
  n, _ := res.RowsAffected()

  if errs != nil {
    panic(errs)
  }

  if n == 0 {
    return sql.ErrNoRows
  }

  return nil
}

func GetHistoryByHistoryID(db *sql.DB, historyID int64) (results structs.History, err error) {
  rows, err := db.Query("SELECT historyID, historyTitle FROM History WHERE historyID = ?", historyID)

  if err != nil {
    panic(err)
  }

  defer rows.Close()

  for rows.Next() {
    // type DateType time.Time

    // var date DateType

    err = rows.Scan(&results.HistoryID, &results.Topic)

    if err != nil {
      panic(err)
    }

    // Get all messages for current history
    var questionId int64

    messageRows, errs := db.Query("SELECT userQuestionID FROM HistoryMessage WHERE historyID = ?", results.HistoryID)

    if errs != nil {
      panic(errs)
    }

    defer messageRows.Close()

    for messageRows.Next() {
      errs = messageRows.Scan(&questionId)

      if errs != nil {
        panic(errs)
      }

      message, err := GetUserMessageByID(db, questionId)

      if err != nil {
        panic(err)
      }

      results.Conversation = append(results.Conversation, message...)
    }
  }

  return
}

func GetOldestHistoryID(db *sql.DB) int {
  rows, err := db.Query("SELECT historyID FROM History ORDER BY createTime ASC LIMIT 1")

  if err != nil {
    panic(err)
  }

  defer rows.Close()

  var id int

  for rows.Next() {

    rows.Scan(&id)
  }

  return id
}


func CheckHistoryExist(db *sql.DB, historyID int) bool {
  rows, err := db.Query("SELECT * FROM History WHERE historyID = ?", historyID)

  if err != nil {
    panic(err)
  }

  defer rows.Close()

  var count = 0

  for rows.Next() {
    count++
  }

  return count != 0
}

func CountHistory(db *sql.DB) int {
  rows, err := db.Query("SELECT * FROM History")

  if err != nil {
    panic(err)
  }

  defer rows.Close()

  var count = 0

  for rows.Next() {
    count++
  }

  return count
}