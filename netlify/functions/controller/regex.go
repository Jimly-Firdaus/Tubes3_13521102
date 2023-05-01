package controller

import (
	"TUBES3_13521102/netlify/functions/structs"
	"database/sql"
	"regexp"
)

func FilterMessage(req structs.Request, db *sql.DB) {
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
			GetResponse(req, i+1, db)
			return
		}
	}
}

func GetResponse(req structs.Request, index int, db *sql.DB) {
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
		return
	} else {
		return
	}
}
