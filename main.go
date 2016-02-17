package main

import (
	"fmt"
	"time"
)

func parseDate(value string) time.Time {
	years := []string{"2006", "06", ""}
	months := []string{"Jan", "Jan.", "January", "01", "1", ""}
	days := []string{"Mon", "Mon.", "Monday", "2", "02", "2nd", "2st", "2rd", "2th", ""}
	delimiters := []string{"/", " ", "", ", "}
	timeNumsFirst := []string{"3:04pm", "15:04", "3:04:05pm", "15:04:05", ""}
	for _, year := range years {
		for _, month := range months {
			for _, day := range days {
				for _, delimiter1 := range delimiters {
					for _, delimiter2 := range delimiters {
						for _, delimiter3 := range delimiters {
							for _, timeNumFirst := range timeNumsFirst {
								layout := timeNumFirst + delimiter3 + month + delimiter1 + day + delimiter2 + year
								t, err := time.Parse(layout, value)
								//fmt.Println(layout)
								if err == nil {
									return t
								}
								layout =  month + delimiter1 + day + delimiter2 + year + delimiter3 + timeNumFirst
								t, err = time.Parse(layout, value)
								//fmt.Println(layout)
								if err == nil {
									return t
								}
							}
						}
					}
				}
			}
		}
	}

	return time.Now()
}

func main() {
	t0 := time.Now()
	fmt.Println(parseDate("02/17/2016 3:05pm"))
	fmt.Println(parseDate("02/17/2016, 9:01am"))
	fmt.Println(parseDate("February 17th, 2016"))
	fmt.Println(parseDate("Feb. 17, 2016"))
	fmt.Printf("took %s", time.Since(t0))
}
