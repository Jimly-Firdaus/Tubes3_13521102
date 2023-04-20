package FeatureDate

import (
  "time"
)

func DateDayName(searchDate string) string {
  date, err := time.Parse("02/01/2006", searchDate)
  if err != nil {
      panic(err)
  }
  dayName := date.Weekday().String()
  return dayName
}
