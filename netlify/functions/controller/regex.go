package controller

import (
	"TUBES3_13521102/netlify/functions/FeatureStringmatching"
	"TUBES3_13521102/netlify/functions/repository"
	"TUBES3_13521102/netlify/functions/structs"
	"database/sql"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

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
		// Split
		return
	} else if index == 2 { // Fitur Hapus Pertanyaan
		// Split
		return
	} else if index == 3 { // Fitur Kalendar
		// Split ada
		return
	} else if index == 4 { //  Fitur Kalkulator
		return
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
				req.Response = strconv.Itoa(i+1) + ". " + req.Response + qList[i].s + "\n"
			}
		} else {
			for i := 0; i < len(qList); i++ {
				req.Response = strconv.Itoa(i+1) + ". " + req.Response + qList[i].s + "\n"
			}
		}
		return
	} else {
		return
	}
}
