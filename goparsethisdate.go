package goparsethisdate

import (
	"fmt"
	"time"
)

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
