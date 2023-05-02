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
	func (c rune) bool {
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

	return newStr
}

func FilterMessage(req *structs.Request, stat *string, db *sql.DB) {
	// Define multiple regex patterns
	patterns := []string{
		`^(?i)Tambahkan pertanyaan (.*) dengan jawaban (.*)$`, // Tambah pertanyaan
		`^(?i)Hapus pertanyaan (.*)$`,                         // Hapus pertanyaan
		`(?i)^(Hari apa )?[0-9]{2}/[0-9]{2}/[0-9]{4}\??$`,     // Kalendar
		`^[\d()+\-*\/. ]+$`,                                   // Kalkulator
		`.*`,                                                  // Pertanyaan Teks
	}
	// Compile the patterns into regex objects
	regexes := make([]*regexp.Regexp, len(patterns))
	for i, pattern := range patterns {
		regex, err := regexp.Compile(pattern)
		if err != nil {
			panic(err)
		}
		regexes[i] = regex
	}
	// Filtering required feature for message.
	for i, regex := range regexes {
		if regex.MatchString(req.Text) {
			// fmt.Printf("String '%s' matches pattern %d\n", text, i+1)
			GetResponse(req, i+1, stat, db)
			return
		}
	}
}


func GetResponse(req *structs.Request, index int, stat *string, db *sql.DB) {
	// Fitur Tambah Pertanyaan
	if index == 1 {
		// Get only question and answer from string
		result := GetPertanyaanJawaban(req.Text)

		question := result[0]
		answer := result[1]

		// Checking if question already exist in the table or not
		if (repository.CheckQuestionExist(db, question)) {
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
		if req.Method != "KMP" {
			for _, question := range questions {
				dbq := strings.ToLower(question.Question)
				text := strings.ToLower(req.Text)
				if FeatureStringmatching.BmMatch(dbq, text) != -1 {
					req.Response = question.Answer
					return
				}
			}

		} else if req.Method == "BoyerMoore" { // Boyer-Moore
			for _, question := range questions {
				dbq := strings.ToLower(question.Question)
				text := strings.ToLower(req.Text)
				if FeatureStringmatching.KMP(dbq, text) != -1 {
					req.Response = question.Answer
					return
				}
			}
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
		req.Response = ""
		sort.Slice(qList, func(i, j int) bool {
			return qList[i].i > qList[j].i
		})
		if len(qList) > 3 {
			for i := 0; i < 3; i++ {
				req.Response = req.Response + qList[i].s + "<-|->"
			}
		} else {
			for i := 0; i < len(qList); i++ {
				req.Response = req.Response + qList[i].s + "<-|->"
			}
		}
		return
	} else {
		return
	}
}
