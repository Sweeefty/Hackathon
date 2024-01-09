package tools

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func CreateLog() {
	// Create log file
	f, err := os.Create("./data/log.txt")
	if err != nil {
		WriteErr("Error creating log file tools/Log.go l14")
		return
	}
	defer f.Close()
}

func WriteLog(ToWrite string) {
	// We get the current time
	t := time.Now()
	// We format the time to a string
	time := t.Format("2006-01-02 15:04:05")
	// We create the string to write
	ToWrite = time + " " + ToWrite
	// Write to log file
	f, err := os.OpenFile("./data/log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		WriteErr("Error opening log file tools/Log.go l30")
	}
	defer f.Close()
	if _, err := f.WriteString(ToWrite + "\n"); err != nil {
		WriteErr("Error writing log file tools/Log.go l34")
	}
}
func WriteErr(ToWrite string) {
	// We get the current time
	t := time.Now()
	// We format the time to a string
	time := t.Format("2006-01-02 15:04:05")
	// We create the string to write
	ToWrite = "e" + time + " " + ToWrite
	// Write to log file
	f, err := os.OpenFile("./data/log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	if _, err := f.WriteString(ToWrite + "\n"); err != nil {
		fmt.Println(err)
	}
}

func PrintLog() {
	colorReset := "\033[0m"
	colorRed := "\033[31m"
	file, err := os.Open("./data/log.txt")
	if err != nil {
		WriteErr("Error opening log file tools/Log.go l58")
	}
	scanner := bufio.NewScanner(file) //We take make a array with all the lines of the DB
	for scanner.Scan() {              //We explore the scanner
		text := scanner.Text() //We explore the scanner
		if text[0] == 'e' {
			fmt.Println(colorRed + text[1:] + colorReset)
		} else {
			fmt.Println(text)
		}
	}
}

func PrintErr() {
	colorReset := "\033[0m"
	colorRed := "\033[31m"
	file, err := os.Open("./data/log.txt")
	if err != nil {
		WriteErr("Error opening log file tools/Log.go l58")
	}
	scanner := bufio.NewScanner(file) //We take make a array with all the lines of the DB
	for scanner.Scan() {
		text := scanner.Text() //We explore the scanner
		if text[0] == 'e' {
			fmt.Println(colorRed + text[1:] + colorReset)
		}
	}
}

func DeleteLog() {
	// Delete log file
	err := os.Remove("./data/log.txt")
	if err != nil {
		WriteErr("Error deleting log file tools/Log.go l70")
	}
	fmt.Println("Log file deleted")
}

func Help() {
	fmt.Println("Tags:")
	fmt.Println("r: Recover the database")
	fmt.Println("d: Delete the log file")
	fmt.Println("p: Print the log file")
	fmt.Println("h: Print this help")
}
