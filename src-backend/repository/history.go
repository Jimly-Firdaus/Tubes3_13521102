package repository

import (
	"TUBES3_13521102/src-backend/structs"
	"database/sql"
)

func InsertHistory(db *sql.DB, history structs.History) (error) {
  errs := db.QueryRow("INSERT INTO History VALUES ($1, $2)", history.HistoryID, history.HistoryTitle)

  return errs.Err()
}


func DeleteHistory(db *sql.DB, history structs.History) (error) {
  // Deleting from foreign key table first
  _, errs := db.Exec("DELETE FROM HistoryMessage WHERE historyID = $1", history.HistoryID)

  if errs != nil {
    panic(errs)
  }

  res, errs := db.Exec("DELETE FROM History WHERE historyID = $1", history.HistoryID)
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
  res, errs := db.Exec("UPDATE History SET historyTitle = $1 WHERE historyID = $2", history.HistoryTitle, history.HistoryID)
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
  rows, err := db.Query("SELECT * FROM History WHERE historyID = ?", historyID)

  if err != nil {
    panic(err)
  }

  defer rows.Close()

  for rows.Next() {

    err = rows.Scan(&results.HistoryID, &results.HistoryTitle)

    if err != nil {
      panic(err)
    }

    // Get all messages for current history
    var questionId int64

    messageRows, errs := db.Query("SELECT userQuestionID FROM HistoryMessage WHERE historyID = ?", results.HistoryID)

    if errs != nil {
      panic(errs)
    }

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

