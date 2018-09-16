package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/asaskevich/govalidator"
)

func main() {
	var input string
	for {
		clear()
		menu()
		fmt.Printf("Choose an option: ")
		_, err := fmt.Scan(&input)
		if err != nil {
			log.Fatal(err)
		}
		if govalidator.IsInt(input) {
			switch input {
			case "1":
				writeFile("dosbox.conf", readFile("dosbox.conf", "640x480"), "640x480")
			case "2":
				writeFile("dosbox.conf", readFile("dosbox.conf", "800x600"), "800x600")
			case "3":
				writeFile("dosbox.conf", readFile("dosbox.conf", "960x720"), "960x720")
			case "4":
				writeFile("dosbox.conf", readFile("dosbox.conf", "1024x768"), "1024x768")
			case "5":
				writeFile("dosbox.conf", readFile("dosbox.conf", "1280x960"), "1280x960")
			case "6":
				writeFile("dosbox.conf", readFile("dosbox.conf", "1400x1050"), "1400x1050")
			case "7":
				writeFile("dosbox.conf", readFile("dosbox.conf", "1440x1080"), "1440x1080")
			case "8":
				writeFile("dosbox.conf", readFile("dosbox.conf", "1600x1200"), "1600x1200")
			case "9":
				writeFile("dosbox.conf", readFile("dosbox.conf", "1856x1392"), "1856x1392")
			case "10":
				writeFile("dosbox.conf", readFile("dosbox.conf", "1920x1440"), "1920x1440")
			case "11":
				writeFile("dosbox.conf", readFile("dosbox.conf", "2048x1536"), "2048x1536")
			default:
				checkOpt(0, "")
				continue
			}
		} else {
			checkOpt(0, "")
			continue
		}
	}
}

func readFile(path string, value string) []string {
	var lines []string
	file, err := os.Open(path)
	if err != nil {
		checkOpt(2, "")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "windowresolution=") {
			lines = append(lines, "windowresolution="+value)
		} else {
			lines = append(lines, scanner.Text())
		}
	}
	return lines
}

func writeFile(path string, lines []string, value string) {
	file, err := os.Create(path)
	if err != nil {
		clear()
		log.Fatal(err)
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, v := range lines {
		fmt.Fprintln(w, v)
	}
	w.Flush()
	checkOpt(1, value)
}

func checkOpt(i int, value string) {
	if i == 0 {
		clear()
		menu()
		fmt.Printf("Invalid Option...")
		time.Sleep(2 * time.Second)
	} else if i == 1 {
		clear()
		menu()
		fmt.Printf("Resolution changed to: %s", value)
		time.Sleep(2 * time.Second)
	} else if i == 2 {
		clear()
		menu()
		fmt.Printf("dosbox.conf file was not found...\n\n")
		time.Sleep(1 * time.Second)
		fmt.Printf("Stopping DOSBoxRC...\n")
		time.Sleep(3 * time.Second)
		clear()
		os.Exit(0)
	}
}

func clear() {
	// cmd := exec.Command("clear") // For [Linux|MAC]
	cmd := exec.Command("cmd", "/c", "cls") // For Windows
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func menu() {
	info := `===================================
| DOSBox Resolution Changer V.1.0 |
===================================
Available resolutions:
                    1: 640x480
                    2: 800x600
                    3: 960x720
                    4: 1024x768
                    5: 1280x960
                    6: 1400x1050
                    7: 1440x1080
                    8: 1600x1200
                    9: 1856x1392
                   10: 1920x1440
                   11: 2048x1536
						  
[Info]: "Crtl + C" to quit...
===================================`
	os.Stdout.Write([]byte(info + "\n"))
}
