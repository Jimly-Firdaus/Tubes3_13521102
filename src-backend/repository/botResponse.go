package repository

import (
	"TUBES3_13521102/src-backend/structs"
	"database/sql"
)

func GetAllBotResponse(db *sql.DB) (results []structs.BotResponse, err error) {
  rows, err := db.Query("SELECT * FROM BotResponse")

  if err != nil {
    panic(err)
  }

  defer rows.Close()

  for rows.Next() {
    // Set BotResponse structs to be add to results
    var botResponse = structs.BotResponse{}
    err = rows.Scan(&botResponse.ResponseID, &botResponse.Question, &botResponse.Answer)

    if err != nil {
      panic(err)
    }

    results = append(results, botResponse)
  }

  return
}

func InsertBotResponse(db *sql.DB, newResponse structs.BotResponse) (error) {
  err := db.QueryRow("INSERT INTO BotResponse(question, answer) VALUES(?, ?)", newResponse.Question, newResponse.Answer)

  return err.Err()
}

// Deleting BotResponse based on its question
func DeleteBotResponse(db *sql.DB, botResponse structs.BotResponse) (error) {
  res, errs := db.Exec("DELETE FROM BotResponse WHERE question = ?", botResponse.Question)
  n, _ := res.RowsAffected()

  if errs != nil {
    panic(errs)
  }

  if n == 0 {
    return sql.ErrNoRows
  }
  return nil
}

// Updating BotResponse based on its question
func UpdateBotResponse(db *sql.DB, botResponse structs.BotResponse) (error) {
  res, errs := db.Exec("UPDATE BotResponse SET answer = ? WHERE question = ?", botResponse.Answer, botResponse.Question)
  n, _ := res.RowsAffected()

  if errs != nil {
    panic(errs)
  }

  if n == 0 {
    return sql.ErrNoRows
  }
  return nil
}

func GetBotResponseByResponseID(db *sql.DB, responseID int64) (results []structs.BotResponse, err error) {
  rows, err := db.Query("SELECT * FROM BotResponse WHERE responseID = ?", responseID)

  if err != nil {
    panic(err)
  }

  defer rows.Close()

  for rows.Next() {
    // Set BotResponse structs to be add to results
    var botResponse = structs.BotResponse{}
    err = rows.Scan(&botResponse.ResponseID, &botResponse.Question, &botResponse.Answer)

    if err != nil {
      panic(err)
    }

    results = append(results, botResponse)
  }

  return
}

// Check if a question is already in the BotResponse table or not
func CheckBotResponse(db *sql.DB, question string) bool {
  rows, err := db.Query("SELECT * FROM BotResponse WHERE question = ?", question)

  if err != nil {
    panic(err)
  }

  defer rows.Close()

  // Counting rows
  var count = 0

  for rows.Next() {
    count++
  }

  return count != 0
}


