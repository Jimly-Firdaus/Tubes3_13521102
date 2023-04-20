package FeatureDate

import (
  "time"
)

func DateDayName(searchDate string) string {
  date, err := time.Parse("2006-01-02", searchDate)
  if err != nil {
      panic(err)
  }
  dayName := date.Weekday().String()
  return dayName
}

func ParseDate(originDate string) string {
  const layout = "02/01/2006"
  date, err := time.Parse(layout, originDate)
  if err != nil {
    panic(err)
  }
  formattedDate := date.Format("2006-01-02")
  return formattedDate
}
