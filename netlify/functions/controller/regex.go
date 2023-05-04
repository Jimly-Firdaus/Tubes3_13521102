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
	"strconv"
	"strings"
)

// Function to get only pertanyaan and jawaban from string
func GetPertanyaanJawaban(req string) []string {
	// Replacing unnecessary string value with null
	substrings := strings.Split(req, " dengan ")

	substrings[0] = strings.Replace(substrings[0], "Tambahkan pertanyaan ", "", 1)
	substrings[1] = strings.Replace(substrings[1], "jawaban ", "", 1)

	return substrings
}

// Function to get only pertanyaan from string
func GetPertanyaan(req string) string {
	// Replacing unnecessary string value with null

	newStr := strings.Replace(req, "Hapus pertanyaan ", "", 1)
	return newStr
}

// Function to get only date from string
func FilterDate(date string) string {
	// Replacing unnecessary string value with null
	newStr := strings.Replace(date, "?", "", 1)
	newStr = strings.Replace(newStr, "hari apa ", "", 1)
	newStr = strings.Replace(newStr, "hari apakah ", "", 1)

	return newStr
}

// Function get only equation from string
func FilterEquation(eq string) string {
	newStr := strings.Replace(eq, "berapakah ", "", 1)
	newStr = strings.Replace(newStr, "berapa ", "", 1)
	newStr = strings.Replace(newStr, "hasil dari ", "", 1)
	newStr = strings.Replace(newStr, "?", "", 1)
	return newStr
}

func FilterMessage(req *structs.Request, stat *string, db *sql.DB, regex []*regexp.Regexp) {
	// Split \n if exists in text
	var answers []string
	questions := strings.Split(req.Text, "\n")
	// Filtering required feature for message.
	for _, question := range questions {
		for i, reg := range regex {
			if reg.MatchString(question) {
				// fmt.Printf("String '%s' matches pattern %d\n", text, i+1)
				answer := GetResponse(req, question, i+1, stat, db)
				answers = append(answers, answer)
				break
			}
		}
	}

	// Building response based on text from user
	if len(answers) > 1 { // For multiple feature text
		for _, result := range answers {
			req.Response = req.Response + result + "\n\n"
		}
	} else if len(answers) == 1 { // Single feature text
		req.Response = answers[0]
	}

	// Adding user message to the database
	AddMessage(req, db)
}

func StringMatching(req *structs.Request, text string, db *sql.DB, questions []structs.BotResponse) (string, bool) {
	var exact bool
	var first string
	reqtext := strings.ToLower(text)
	if req.Method == "KMP" {
		for _, question := range questions {
			dbq := strings.ToLower(question.Question)
			if FeatureStringmatching.KMP(dbq, reqtext) == -2 { // Found Same string
				return question.Answer, true
			} else if FeatureStringmatching.KMP(dbq, reqtext) > -1 {
				if exact {
					return "", false // Found more than one exact match
				} else {
					first = question.Answer
					exact = true
				}
			}
		}
	} else if req.Method == "BoyerMoore" {
		for _, question := range questions {
			dbq := strings.ToLower(question.Question)
			if FeatureStringmatching.BmMatch(dbq, reqtext) == -2 { // Found Same string
				return question.Answer, true
			} else if FeatureStringmatching.BmMatch(dbq, reqtext) > -1 {
				if exact {
					return "", false // Found more than one exact match
				} else {
					first = question.Answer
					exact = true
				}
			}
		}
	}
	if exact {
		return first, true
	} else {
		return "", false
	}
}
func ValidateQuestion(req *structs.Request, text string, db *sql.DB, questions []structs.BotResponse) (string, bool) {
	var exact bool
	var first string
	reqtext := strings.ToLower(text)
	if req.Method == "KMP" {
		for _, question := range questions {
			dbq := strings.ToLower(question.Question)
			if FeatureStringmatching.KMP(dbq, reqtext) == -2 { // Found Same string
				return question.Question, true
			} else if FeatureStringmatching.KMP(dbq, reqtext) > -1 {
				if exact {
					return "", false // Found more than one exact match
				} else {
					first = question.Question
					exact = true
				}
			}
		}
	} else if req.Method == "BoyerMoore" {
		for _, question := range questions {
			dbq := strings.ToLower(question.Question)
			if FeatureStringmatching.BmMatch(dbq, reqtext) == -2 { // Found Same string
				return question.Question, true
			} else if FeatureStringmatching.BmMatch(dbq, reqtext) > -1 {
				if exact {
					return "", false // Found more than one exact match
				} else {
					first = question.Question
					exact = true
				}
			}
		}
	}
	if exact {
		return first, true
	} else {
		return "", false
	}
}
func LevenshteinValidation(text string, db *sql.DB, questions []structs.BotResponse) (string, bool) {
	qList := []struct {
		i float64
		s string
		a string
	}{}
	req := strings.ToLower(text)
	for _, question := range questions {
		q := strings.ToLower(question.Question)
		distance := FeatureStringmatching.LevenshteinDistance(q, req)
		qList = append(qList, struct {
			i float64
			s string
			a string
		}{
			i: distance,
			s: question.Question,
			a: question.Answer,
		})
	}
	// Sort Descending by Percentage value
	sort.Slice(qList, func(i, j int) bool {
		return qList[i].i > qList[j].i
	})
	if len(qList) != 0 {
		if qList[0].i > 0.9 {
			return qList[0].s, true
		} else {
			return "", false
		}
	} else {
		return "", false
	}
}

func LevenshteinController(text string, stat *string, db *sql.DB, questions []structs.BotResponse) string {
	// No match found
	var result string
	qList := []struct {
		i float64
		s string
		a string
	}{}
	req := strings.ToLower(text)
	for _, question := range questions {
		q := strings.ToLower(question.Question)
		distance := FeatureStringmatching.LevenshteinDistance(q, req)
		qList = append(qList, struct {
			i float64
			s string
			a string
		}{
			i: distance,
			s: question.Question,
			a: question.Answer,
		})
	}
	// Sort Descending by Percentage value
	sort.Slice(qList, func(i, j int) bool {
		return qList[i].i > qList[j].i
	})
	if len(qList) != 0 {
		if qList[0].i > 0.9 {
			return qList[0].a
		} else {
			var x int
			for i := 0; i < len(qList); {
				if qList[i].i > 0.5 {
					x = x + 1
				}
				i++
			}
			if x != 0 {
				*stat = "404"
				result = "Pertanyaan tidak ditemukan di database.\nApakah maksud anda:\n"
			} else {
				return "Tidak ada pertanyaan tersebut dalam database."
			}

			if x > 3 {
				for i := 0; i < 3; i++ {
					result = result + strconv.Itoa(i+1) + ". " + qList[i].s + "\n"
				}
			} else {
				for i := 0; i < x; i++ {
					result = result + strconv.Itoa(i+1) + ". " + qList[i].s + "\n"
				}
			}
			return result
		}
	} else {
		return "Tidak ada pertanyaan tersebut dalam database."
	}
}

func GetResponse(req *structs.Request, text string, index int, stat *string, db *sql.DB) string {
	var response string
	var questions []structs.BotResponse
	questions, err := repository.GetAllBotResponse(db)
	if err != nil {
		panic(err)
	}
	// Fitur Tambah Pertanyaan
	if index == 1 {
		// Get only question and answer from string
		result := GetPertanyaanJawaban(text)

		question := result[0]
		answer := result[1]

		match, exist := ValidateQuestion(req, question, db, questions)
		// Checking if question already exist in the table or not
		if !exist {
			// Validate with levenshtein
			match, exist := LevenshteinValidation(question, db, questions)
			if exist {
				repository.UpdateBotResponse(db, match, answer)

				// Set bot response
				response = "Pertanyaan " + question + " sudah ada! jawaban di update ke " + answer
			} else {
				repository.InsertBotResponse(db, question, answer)

				// Set bot response
				response = "Pertanyaan " + question + " telah ditambah"
			}
		} else {
			repository.UpdateBotResponse(db, match, answer)

			// Set bot response
			response = "Pertanyaan " + question + " sudah ada! jawaban di update ke " + answer
		}

	} else if index == 2 { // Fitur Hapus Pertanyaan
		// Get only question from string
		question := GetPertanyaan(text)

		// Delete question from table
		del, exist := StringMatching(req, question, db, questions)
		if exist {
			err := repository.DeleteBotResponse(db, del)
			if err == nil {
				response = "Pertanyaan " + question + " telah dihapus"
			}
		} else {
			response = "Tidak ada pertanyaan " + question + " pada database!"
		}

		// Set bot response

	} else if index == 3 { // Fitur Kalendar
		// Split unnecessary string value

		date := FilterDate(strings.ToLower(text))

		// Set bot response
		response = FeatureDate.FindDayName(date)

	} else if index == 4 { //  Fitur Kalkulator

		equation := FilterEquation(strings.ToLower(text))
		// Get expression result
		result, err := FeatureCalculator.CalculateExpression(equation)

		if err != nil {
			panic(err)
		}

		// Set bot response
		response = result

	} else if index == 5 { // Fitur Pertanyaan Teks
		// String matching with KMP/BoyerMoore based on method request
		result, match := StringMatching(req, text, db, questions)

		if !match {
			result = LevenshteinController(text, stat, db, questions) // no match found, then use levenshtein
		}
		response = result
	}
	return response
}

func AddMessage(req *structs.Request, db *sql.DB) {

	// First we check if the historyID is already in database or not

	if repository.CheckHistoryExist(db, int(req.HistoryId)) || req.HistoryTimeStamp == "" {
		// If history already exist then we only need to add to table Messages
		err := repository.InsertMessages(db, *req)

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

		// Also don't forget to insert row to Messages table

		err = repository.InsertMessages(db, *req)

		if err != nil {
			panic(err)
		}
	}
}
