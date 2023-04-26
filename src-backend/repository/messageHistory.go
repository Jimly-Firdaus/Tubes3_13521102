package repository

import (
  "database/sql"
  "TUBES3_13521102/src-backend/structs"
)

func GetAllHistoryMessage(db *sql.DB) (results structs.MessageHistory, err error) {
  rows, err := db.Query("SELECT historyID FROM History")

  if err != nil {
    panic(err)
  }

  defer rows.Close()

  for rows.Next() {
    // Set History structs
    var history = structs.History{}
    var historyId int64

    err = rows.Scan(&historyId)

    if err != nil {
      panic(err)
    }

    history, err = GetHistoryByHistoryID(db, historyId)

    if err != nil {
      panic(err)
    }

    results.MessagesHistory = append(results.MessagesHistory, history)

  }
  return
}

