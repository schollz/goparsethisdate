package main

import (
	"fmt"
	"time"
)

func ParseDateChannel(value string) (time.Time, error) {
	years := []string{"2006", "06", ""}
	months := []string{"Jan", "Jan.", "January", "01", "1", ""}
	days := []string{"Mon", "Mon.", "Monday", "2", "02", "2nd", "2st", "2rd", "2th", ""}
	delimiters := []string{"/", " ", "", ", "}
	timeNumsFirst := []string{"3:04pm", "15:04", "3:04:05pm", "15:04:05", "3pm", ""}
	c := make(chan time.Time, 2*len(delimiters)*len(delimiters)*len(delimiters)*len(days)*len(months)*len(years)*len(timeNumsFirst))
	for _, year := range years {
		for _, month := range months {
			for _, day := range days {
				for _, delimiter1 := range delimiters {
					for _, delimiter2 := range delimiters {
						for _, delimiter3 := range delimiters {
							for _, timeNumFirst := range timeNumsFirst {
								layout := timeNumFirst + delimiter3 + month + delimiter1 + day + delimiter2 + year

								go func(layout string, value string) {
									t, err := time.Parse(layout, value)
									//fmt.Println(layout)
									if err == nil {
										c <- t
									}
								}(layout, value)

								layout = month + delimiter1 + day + delimiter2 + year + delimiter3 + timeNumFirst

								go func(layout string, value string) {
									t, err := time.Parse(layout, value)
									//fmt.Println(layout)
									if err == nil {
										c <- t
									}
								}(layout, value)
							}
						}
					}
				}
			}
		}
	}
	x := <-c
	return x, nil
}

func ParseDate(value string) (time.Time, error) {
	years := []string{"2006", "06", ""}
	months := []string{"Jan", "Jan.", "January", "01", "1", ""}
	days := []string{"Mon", "Mon.", "Monday", "2", "02", "2nd", "2st", "2rd", "2th", ""}
	delimiters := []string{"/", " ", "", ", "}
	timeNumsFirst := []string{"3:04pm", "15:04", "3:04:05pm", "15:04:05", "3pm", ""}
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
									return t, err
								}
								layout = month + delimiter1 + day + delimiter2 + year + delimiter3 + timeNumFirst
								t, err = time.Parse(layout, value)
								//fmt.Println(layout)
								if err == nil {
									return t, err
								}
							}
						}
					}
				}
			}
		}
	}

	return time.Now(), fmt.Errorf("Could not find date")
}

func main() {
	// start := time.Now()
	// ParseDateChannel("Feb. 2nd 2016")
	// fmt.Println(time.Since(start))
	start := time.Now()
	ParseDate("Feb. 2nd 2016")
	fmt.Println(time.Since(start))
}
