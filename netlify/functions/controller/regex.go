package controller

import (
	"TUBES3_13521102/netlify/functions/FeatureCalculator"
	"TUBES3_13521102/netlify/functions/FeatureDate"
	"TUBES3_13521102/netlify/functions/FeatureStringmatching"
	"TUBES3_13521102/netlify/functions/repository"
	"TUBES3_13521102/netlify/functions/structs"
	"database/sql"
	"regexp"
	"sort"
	"strings"
)

// Function to get only pertanyaan and jawaban from string
func GetPertanyaanJawaban(req string) []string {
	// Replacing unnecessary string value with null
	substr1 := "Tambahkan pertanyaan "
	substr2 := "dengan jawaban "
	newStr := strings.Replace(req, substr1, "", 1)
	newStr = strings.Replace(newStr, substr2, "", 1)

	result := strings.FieldsFunc(newStr,
		func(c rune) bool {
			return c == ' '
		})

	return result
}

// Function to get only pertanyaan from string
func GetPertanyaan(req string) string {
	// Replacing unnecessary string value with null
	substr1 := "Hapus pertanyaan "

	newStr := strings.Replace(req, substr1, "", 1)

	return newStr
}

// Function to get only date from string
func FilterDate(date string) string {
	// Replacing unnecessary string value with null
	substr := "Hari apa "
	newStr := strings.Replace(date, substr, "", 1)
	newStr = strings.Replace(newStr, "?", "", 1)

	return newStr
}

func FilterMessage(req *structs.Request, stat *string, db *sql.DB, regex []*regexp.Regexp) {
	// Filtering required feature for message.
	for i, reg := range regex {
		if reg.MatchString(req.Text) {
			// fmt.Printf("String '%s' matches pattern %d\n", text, i+1)
			GetResponse(req, i+1, stat, db)
			return
		}
	}
}
func StringMatching(req *structs.Request, db *sql.DB, questions []structs.BotResponse) bool {
	if req.Method == "KMP" {
		for _, question := range questions {
			dbq := strings.ToLower(question.Question)
			text := strings.ToLower(req.Text)
			if FeatureStringmatching.BmMatch(dbq, text) == -2 {
				req.Response = question.Answer
				return true
			}
		}
	} else if req.Method == "BoyerMoore" {
		for _, question := range questions {
			dbq := strings.ToLower(question.Question)
			text := strings.ToLower(req.Text)
			if FeatureStringmatching.KMP(dbq, text) == -2 {
				req.Response = question.Answer
				return true
			}
		}
	}
	return false
}
func GetResponse(req *structs.Request, index int, stat *string, db *sql.DB) {
	// Fitur Tambah Pertanyaan
	if index == 1 {
		// Get only question and answer from string
		result := GetPertanyaanJawaban(req.Text)

		question := result[0]
		answer := result[1]

		// Checking if question already exist in the table or not
		if repository.CheckQuestionExist(db, question) {
			repository.UpdateBotResponse(db, question, answer)

			// Set bot response
			req.Response = "Pertanyaan " + question + " sudah ada! jawaban di update ke " + answer
		} else {
			repository.InsertBotResponse(db, question, answer)

			// Set bot response
			req.Response = "Pertanyaan " + question + " telah ditambah"
		}

	} else if index == 2 { // Fitur Hapus Pertanyaan
		// Get only question from string
		question := GetPertanyaan(req.Text)

		// Delete question from table
		err := repository.DeleteBotResponse(db, question)

		// Set bot response
		if err != nil {
			req.Response = "Tidak ada pertanyaan " + question + " pada database!"
		} else {
			req.Response = "Pertanyaan " + question + " telah dihapus"
		}

	} else if index == 3 { // Fitur Kalendar
		// Split unnecessary string value
		date := FilterDate(req.Text)

		// Set bot response
		req.Response = FeatureDate.FindDayName(date)

	} else if index == 4 { //  Fitur Kalkulator

		// Get expression result
		result, err := FeatureCalculator.CalculateExpression(req.Text)

		if err != nil {
			panic(err)
		}

		// Set bot response
		req.Response = result

	} else if index == 5 { // Fitur Pertanyaan Teks
		var questions []structs.BotResponse
		questions, err := repository.GetAllBotResponse(db)
		if err != nil {
			panic(err)
		}
		//
		if StringMatching(req, db, questions) {
			return
		}

		// No match found
		*stat = "404"
		qList := []struct {
			i float64
			s string
		}{}
		for _, question := range questions {
			req := req.Text
			distance := FeatureStringmatching.LevenshteinDistance(question.Question, req)
			qList = append(qList, struct {
				i float64
				s string
			}{
				i: distance,
				s: question.Question,
			})
		}
		// Sort Descending by Percentage value
		sort.Slice(qList, func(i, j int) bool {
			return qList[i].i > qList[j].i
		})
		req.Response = "Pertanyaan tidak ditemukan di database.\nApakah maksud anda:\n"
		if len(qList) > 3 {
			for i := 0; i < 3; i++ {
				req.Response = req.Response + qList[i].s + "\n"
			}
		} else {
			for i := 0; i < len(qList); i++ {
				req.Response = req.Response + qList[i].s + "\n"
			}
		}
		return
	}
	// Adding user message to the database

	// First we check if the historyID is already in database or not

	if repository.CheckHistoryExist(db, int(req.HistoryId)) {
		// If history already exist then we only need to add to table UserMessage and HistoryMessage
		err := repository.InsertUserMessage(db, *req)

		if err != nil {
			panic(err)
		}
	} else {
		// IF history doesn't exist yet we check how many rows there are now in History table
		if repository.CountHistory(db) >= 10 {
			// We delete the oldest history from table
			oldestID := repository.GetOldestHistoryID(db)

			err := repository.DeleteHistory(db, oldestID)

			if err != nil {
				panic(err)
			}
		}

		// If there is space to insert row then insert row to history table
		var newHistory = structs.History{}
		newHistory.HistoryID = req.HistoryId
		newHistory.Topic = req.Text
		err := repository.InsertHistory(db, newHistory, req.HistoryTimeStamp)

		if err != nil {
			panic(err)
		}

		// Also don't forget to insert row to userMessage table

		err = repository.InsertUserMessage(db, *req)

		if err != nil {
			panic(err)
		}
	}
}
