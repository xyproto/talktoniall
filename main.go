package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"

	"github.com/xyproto/niall"
)

func readln(prompt string) string {
	fmt.Print(prompt)
	in := bufio.NewReader(os.Stdin)
	var cmd []byte
	cmd, _ = in.ReadBytes('\n')
	buf := bytes.NewBuffer(cmd)
	return strings.TrimSpace(buf.String())
}

func main() {
	niall.Init()

	input := readln("> ")
	END: for {
		switch input {
		case "quit","exit":
			break END
		case "save":
			niall.SaveDictionary("niall.brain")
		case "load":
			niall.LoadDictionary("niall.brain")
		case "help":
			fmt.Println("Commands:")
			fmt.Println("  load - loads niall.brain")
			fmt.Println("  save - saves niall.brain")
			fmt.Println("  correct [from] [to] - corrects spelling")
			fmt.Println("  quit - exits")
			fmt.Println("  exit - exits")
		default:
			if strings.Index(input, "correct") == 0 {
				if strings.Count(input, " ") == 2 {
					fields := strings.Split(input, " ")
					from, to := fields[1], fields[2]
					niall.CorrectSpelling(from, to)
					fmt.Println("ok, corrected spelling")
				} else {
					fmt.Println("syntax for correcting spelling: correct [from] [to]")
				}
			} else {
				niall.Learn(input)
				fmt.Println(niall.Talk())
			}
		}
		input = readln("> ")
	}

	niall.Quit()
}
