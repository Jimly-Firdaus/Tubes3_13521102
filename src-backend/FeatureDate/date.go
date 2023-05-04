package FeatureDate

import (
	"strconv"
	"strings"
)

// Get all number in the date
func GetNumber(date string) []string {
  number := strings.FieldsFunc(date, func(c rune) bool {
    return c == '/'
  })
  return number
}

// Convert string to number
func StringtoInteger (number string) int {
  value, _ := strconv.Atoi(number)
  return value
}

// Get day's name
func GetDayName(number int) string {
  if (number == 0) {
    return "Hari Minggu"
  } else if (number == 1) {
    return "Hari Senin"
  } else if (number == 2) {
    return "Hari Selasa"
  } else if (number == 3) {
    return "Hari Rabu"
  } else if (number == 4) {
    return "Hari Kamis"
  } else if (number == 5) {
    return "Hari Jumat"
  } else {
    return "Hari Sabtu"
  }
}

// Check if a date is valid or not
func CheckDateValid(day int, month int, year int) bool {

  if (month == 1 || month == 3 || month == 5 || month == 7 || month == 8 || month == 10 || month == 12) {
    if (day < 0 || day > 31) {
      return false
    }
  } else if (month == 4 || month == 6 || month == 9 || month == 11) {

    if (day < 0 || day > 30) {
      return false
    }

  } else if (month == 2) {

    if (year % 4 == 0) {

      if (day < 0 || day > 29) {
        return false
      }

    } else if (month == 2) {

      if (day < 0 || day > 28) {
        return false
      }

    }

  } else {

    return false

  }

  return true
}

func FindDayName(date string) string {
  // Get all numbers value for the date
  numbers := GetNumber(date)
  var day int
  var month int
  var year int
  for index, value := range numbers  {

    if (index == 0) {

      day = StringtoInteger(value)

    } else if (index == 1) {

      month = StringtoInteger(value)

    } else {

      year = StringtoInteger(value)

    }

  }

  // First we check if a date is valid or not
  if (!CheckDateValid(day, month, year)) {
    return "Invalid Date"
  }

  // If the date is valid then we proceed to calculate the date's day name using Zeller's Rule
  // F=k+ [(13*m-1)/5] +D+ [D/4] +[C/4]-2*C where

  // k is  the day of the month.
  // m is the month number.
  // D is the last two digits of the year.
  // C is the first two digits of the year.

  /* According to Zeller’s rule the month is counted as follows:
    March is 1, April is 2….. January is 11 and February is 12.
    So the year starts from March and ends with February. So if the given date has month as January or February subtract 1 from the year. */

  // Translating month and year according to Zeller's rule

  if (month <= 2) {
    year--
  }

  if (month == 2) {
    month = 12
  } else {
    month = (month - 2) % 12
  }



  // We calculate the D and C first
  var D = year % 100
  var C = year / 100

  // Calculate the F value
  var F = day + ((13*month-1)/5) + D + (D/4) + (C/4)-2*C

  // Then we calculate the day by using F mod 7
  var dayNumber = F % 7
  if (dayNumber < 0) {
    dayNumber += 7
  }

  // Get the day's name
  return GetDayName(dayNumber)
}
