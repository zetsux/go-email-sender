package main

import (
	"fmt"
	"os"
)

func main() {
	rows, err := ReadCSV(fmt.Sprintf("%s/%s", DEFAULT_DATA_DIRECTORY, DEFAULT_DATA_NAME))
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, row := range rows {
		if len(row) >= 2 && IsValidEmail(row[0]) && row[1] != "" {
			email := row[0]
			name := row[1]

			mailBody, err := MakeEmailBody(name)
			if err != nil {
				fmt.Println("Failed to generate email body for : ", row, " | Error : ", err)
				continue
			}

			filePath := DEFAULT_FILE_DIRECTORY + DEFAULT_FILE_NAME
			if _, err := os.Stat(filePath); err != nil {
				fmt.Println("File unable to be checked for : ", row, " | Error : ", err)
				continue
			}

			subject := "Subject"
			err = SendEmail(email, subject, mailBody, filePath)
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Printf("Email sent successfully to %v\n", email)
		} else {
			fmt.Println("Invalid row : ", row)
		}
	}
}
