package repository

import (
	"TUBES3_13521102/netlify/functions/structs"
	"database/sql"
)

func GetAllHistory(db *sql.DB) (results structs.HistoryPayload, err error) {
	rows, err := db.Query("SELECT historyID, historyTitle FROM History ORDER BY createTime DESC")

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var historyReq = structs.HistoryRequest{}

		err = rows.Scan(&historyReq.HistoryID, &historyReq.HistoryTopic)

		if err != nil {
			panic(err)
		}

		results.HistoryCollection = append(results.HistoryCollection, historyReq)
	}

	return
}