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

  // Deleting foreign key table value first
  _, err := db.Exec("Delete FROM Messages WHERE historyID = ?", historyID)

  if err != nil {
    panic(err)
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


func GetHistoryByHistoryID(db *sql.DB, historyID int64) (results structs.History, err error) {
  rows, err := db.Query("SELECT historyID, historyTitle FROM History WHERE historyID = ?", historyID)

  if err != nil {
    panic(err)
  }

  defer rows.Close()

  for rows.Next() {

    err = rows.Scan(&results.HistoryID, &results.Topic)

    if err != nil {
      panic(err)
    }

  }

  messages, err := GetHistoryMessages(db, int(historyID))

  if err != nil {
    panic(err)
  }

  results.Conversation = append(results.Conversation, messages...)

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