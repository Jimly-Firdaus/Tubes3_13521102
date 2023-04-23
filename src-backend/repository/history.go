package repository

import (
	"TUBES3_13521102/src-backend/structs"
	"database/sql"
)

func GetAllHistory(db *sql.DB) (results []structs.History, err error) {
  // Getting data from MySQL
  query := "SELECT * FROM History"

  rows, err := db.Query(query)

  if err != nil {
    panic(err)
  }

  defer rows.Close()

  for rows.Next() {
    // Set History structs to be add to results
    var history = structs.History{}
    err = rows.Scan(&history.HistoryID, &history.HistoryTitle)
    if err != nil {
      panic(err)
    }

    results = append(results, history)
  }

  return
}

func InsertHistory(db *sql.DB, history structs.History) (error) {
  errs := db.QueryRow("INSERT INTO History (historyTitle) VALUES ($1)", history.HistoryTitle)

  return errs.Err()
}


func DeleteHistory(db *sql.DB, history structs.History) (error) {
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

func GetHistoryByHistoryID(db *sql.DB, historyID int64) (results []structs.History, err error) {
  rows, err := db.Query("SELECT * FROM History WHERE historyID = $1", historyID)

  if err != nil {
    panic(err)
  }

  defer rows.Close()

  for rows.Next() {
    var history = structs.History{}

    err = rows.Scan(&history.HistoryID, &history.HistoryTitle)

    if err != nil {
      panic(err)
    }

    results = append(results, history)
  }

  return
}

