package repository

import (
	"TUBES3_13521102/netlify/functions/structs"
	"database/sql"
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


