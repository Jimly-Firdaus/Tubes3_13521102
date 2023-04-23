package repository

import (
	"TUBES3_13521102/src-backend/structs"
	"database/sql"
)

func GetAllTopik(db *sql.DB) (results []structs.Topik, err error) {
  // Getting data from MySQL
  query := "SELECT * FROM Topik"

  rows, err := db.Query(query)

  if err != nil {
    panic(err)
  }

  defer rows.Close()

  for rows.Next() {
    // Set Topik structs to be add to results
    var topik = structs.Topik{}
    err = rows.Scan(&topik.IdTopik, &topik.TopikUtama)
    if err != nil {
      panic(err)
    }

    results = append(results, topik)
  }

  return
}

func InsertTopik(db *sql.DB, topik structs.Topik) (error) {
  errs := db.QueryRow("INSERT INTO TopiK (topikUtama) VALUES (?)", topik.TopikUtama)

  return errs.Err()
}


func DeleteTopik(db *sql.DB, topik structs.Topik) (error) {
  res, errs := db.Exec("DELETE FROM Topik WHERE idTopik = ?", topik.IdTopik)
  n, _ := res.RowsAffected()

  if errs != nil {
    panic(errs)
  }

  if n == 0 {
    return sql.ErrNoRows
  }

  return nil
}

func UpdateTopik(db *sql.DB, topik structs.Topik) (error) {
  res, errs := db.Exec("UPDATE Topik SET topikUtama = ? WHERE idTopik = ?", topik.TopikUtama, topik.IdTopik)
  n, _ := res.RowsAffected()

  if errs != nil {
    panic(errs)
  }

  if n == 0 {
    return sql.ErrNoRows
  }

  return nil
}

func GetTopikByIdTopik(db *sql.DB, idTopik int64) (results []structs.Topik, err error) {
  rows, err := db.Query("SELECT * FROM Topik WHERE idTopik = ?", idTopik)

  if err != nil {
    panic(err)
  }

  defer rows.Close()

  for rows.Next() {
    var topik = structs.Topik{}

    err = rows.Scan(&topik.IdTopik, &topik.TopikUtama)

    if err != nil {
      panic(err)
    }

    results = append(results, topik)
  }

  return
}

