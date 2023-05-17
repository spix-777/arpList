package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

// Loggers
var (
	DoneLogger *log.Logger
	WarnLogger *log.Logger
)

// Initialize loggers to stdout
func init() {
	DoneLogger = log.New(os.Stdout, " [ + ] ", 0)
	WarnLogger = log.New(os.Stdout, " [ ! ] ", 0)
}

func banner() {
	banner := `
██   █▄▄▄▄ █ ▄▄      █    ▄█    ▄▄▄▄▄      ▄▄▄▄▀     
█ █  █  ▄▀ █   █     █    ██   █     ▀▄ ▀▀▀ █        
█▄▄█ █▀▀▌  █▀▀▀      █    ██ ▄  ▀▀▀▀▄       █        
█  █ █  █  █         ███▄ ▐█  ▀▄▄▄▄▀       █         
   █   █    █            ▀ ▐              ▀          
  █   ▀      ▀                                       
 ▀                                                   
                                      by SpiX-777`

	fmt.Println(banner)
	fmt.Println("")
}

func main() {
	banner()

	// Create or open a file
	file, err := os.Create("output.txt")
	if err != nil {
		errorFile := "Error creating file ..."
		WarnLogger.Fatalf("\x1b[31m%s\x1b[0m\n", errorFile, err)
		return
	}
	defer file.Close()

	// Create a writer to write to the file
	writer := bufio.NewWriter(file)

	// Get the lines from the memory
	lines := arpFile()

	// Fix the file by removing the first two lines and the last line
	FixLine := oneLast(lines)

	// Get the ip from the FixLine and put it in a string
	var ip string
	for _, line := range FixLine {
		count := strings.Index(line, ":") - 3
		ip += line[0:count] + "\n"
	}
	var ipList []string
	ipList = strings.Split(ip, "\n")

	// Create a map to store unique strings
	uniqueStrs := make(map[string]bool)

	// Iterate over the slice and add each string to the map
	// The map will automatically remove duplicates
	for _, str := range ipList {
		uniqueStrs[str] = true
	}

	// Create a new slice with the unique strings
	var uniqueSlice []string
	for str := range uniqueStrs {
		if str != "" {
			uniqueSlice = append(uniqueSlice, str)
		}
	}
	// Write the strings to the file
	for _, line := range uniqueSlice {
		// Write to the file
		_, err = writer.WriteString(line + "\n")
		if err != nil {
			errorFile := "Error writing to file ..."
			WarnLogger.Fatalf("\x1b[31m%s\x1b[0m\n", errorFile, err)
			return
		}
	}
	// Flush the writer to ensure all data is written to the file
	err = writer.Flush()
	if err != nil {
		errorFile := "Error flushing writer ..."
		WarnLogger.Fatalf("\x1b[31m%s\x1b[0m\n", errorFile, err)
	}

	// Done! :)
	done := "ARP Scan successfully wrote to file!"
	DoneLogger.Printf("\x1b[32m%s\x1b[0m\n", done)
	fmt.Println("")
}

func oneLast(lines []string) []string {
	var FixLines []string
	// Fix the file by removing the first two lines and the last line
	dis := 0
	for i, line := range lines {
		if i > 1 {
			if dis == 0 {
				conut := len(line)
				if conut == 0 {
					dis = 1
				} else {
					FixLines = append(FixLines, line)
				}
			}
		}
	}
	return FixLines
}

func arpFile() []string {
	// Define the command to run
	cmd := exec.Command("sudo", "arp-scan", "-l")

	// Run the command and capture its output
	output, err := cmd.Output()
	if err != nil {
		WarnLogger.Fatalln("Error running command:", err)
	}

	// Print the output
	var lines []string
	lines = strings.Split(string(output), "\n")
	return lines
}
