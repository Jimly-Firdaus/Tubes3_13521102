package main

import (
    "fmt"
    "TUBES3_13521102/src-backend/FeatureDate"
    "TUBES3_13521102/src-backend/FeatureCalculator"
)

func main() {
    date := "25/02/2023"

    dayName := FeatureDate.DateDayName(date)
    fmt.Println(dayName)

    hasil, err := FeatureCalculator.CalculateExpression("((3+5)/2.13)")

    if (err != nil) {
      fmt.Println(err)
    } else {
      fmt.Println(hasil)
    }
}
