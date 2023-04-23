package repository

import (
  "database/sql"
)

func GetAllTopik(db *sql.DB) (err error, results []structs.Pertanyaan)
