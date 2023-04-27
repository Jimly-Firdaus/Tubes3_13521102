package controller

import (
	"TUBES3_13521102/netlify/functions/main/repository"
	"database/sql"
	"log"
	"net/http"
    "encoding/json"
)
func GetAllHistoryMessage(w http.ResponseWriter, r *http.Request) {
    db, err := sql.Open("mysql", "root:PNGO6atNekbjjq4g2yPy@tcp(containers-us-west-13.railway.app:6330)/railway")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    pingErr := db.Ping()
    if pingErr != nil {
        log.Fatal(pingErr)
    }

    defer db.Close()

    historyMessageList, err := repository.GetAllHistoryMessage(db)

    // Write response as JSON
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]interface{}{
        "historyMessage": historyMessageList,
    })
}
