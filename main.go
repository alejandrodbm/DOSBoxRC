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

type config struct {
	path  string
	input string
	value string
	lines []string
}

func (c *config) setRes(value string) {
	c.lines = []string{}
	c.value = value
	if c.readFile() {
		c.writeFile()
	} else {
		checkOpt(4, "")
	}
}

func (c *config) readFile() bool {
	file, err := os.Open(c.path)
	if err != nil {
		checkOpt(2, "")
		return false
	}
	defer file.Close()

	var winResParamExist bool = false
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "windowresolution=") {
			c.lines = append(c.lines, "windowresolution="+c.value)
			winResParamExist = true
		} else {
			c.lines = append(c.lines, scanner.Text())
		}
	}
	return winResParamExist
}

func (c *config) writeFile() {
	file, err := os.Create(c.path)
	if err != nil {
		checkOpt(3, "")
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, v := range c.lines {
		fmt.Fprintln(w, v)
	}
	w.Flush()

	checkOpt(1, c.value)
}

var input config

func init() {
	input.path = "dosbox.conf"
}

func main() {
	for {
		clear()
		menu()
		fmt.Printf("Choose an option: ")
		_, err := fmt.Scan(&input.input)
		if err != nil {
			log.Fatal(err)
		}

		if govalidator.IsInt(input.input) {
			switch input.input {
			case "1":
				input.setRes("640x480")
			case "2":
				input.setRes("800x600")
			case "3":
				input.setRes("960x720")
			case "4":
				input.setRes("1024x768")
			case "5":
				input.setRes("1280x960")
			case "6":
				input.setRes("1400x1050")
			case "7":
				input.setRes("1440x1080")
			case "8":
				input.setRes("1600x1200")
			case "9":
				input.setRes("1856x1392")
			case "10":
				input.setRes("1920x1440")
			case "11":
				input.setRes("2048x1536")
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

func checkOpt(i int, value string) {
	switch i {
	case 0:
		clear()
		menu()
		fmt.Printf("Invalid Option...")
		time.Sleep(2 * time.Second)
	case 1:
		clear()
		menu()
		fmt.Printf("Resolution changed to: %s", value)
		time.Sleep(2 * time.Second)
	case 2:
		clear()
		menu()
		fmt.Printf("\"dosbox.conf\" file was not found...")
		time.Sleep(2 * time.Second)
	case 3:
		clear()
		menu()
		fmt.Printf("Error creating or modifying the \"dosbox.conf\" file...")
		time.Sleep(2 * time.Second)
	case 4:
		clear()
		menu()
		fmt.Printf("Missing window resolution parameter in \"dosbox.conf\" file...")
		time.Sleep(2 * time.Second)
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
