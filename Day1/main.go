package main

import (
  "fmt"
  "time"
)

func main()  {
  year := 2020
  months := make([]int, 12)
  wDays := []string{"S","M","T","W","T","F","S"}
  // Looping Month
  for key, _ :=range months{
    // Declare Month
    Month := Date(year, key + 1, 1)
    // Declare total day
    Day := Date(year, key + 2, 0)
    // Show Month From January
    fmt.Println("========",Month.Month(),"========")
    // Show Week Day
    for _, val :=range wDays{
      fmt.Print(val + "   ")
    }
    fmt.Println()
    // Show Day
    s := 0
    for i := 1;  i<= Day.Day(); i++ {
      // Get Week Day
      weekDay := Date(year, key + 1, i)
      if i == 1 {
        for x := 0; x < int(weekDay.Weekday()); x++ {
          fmt.Print(">>  ")
          s++
        }
      }
      if i > 9 {
        fmt.Print(i, "  ")
      }else {
        fmt.Print(i, "   ")
      }
      s++
      if s % 7 == 0 {
        fmt.Println()
      }
    }
    fmt.Println()
  }
}
func Date(year, month, day int) time.Time {
  return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}
