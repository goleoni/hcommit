package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	firstArg := os.Args[1]
	fullMessageArray := os.Args[2:]

	if firstArg == "help" {
		fmt.Println(
			`
			Usage

			Permitted intentions
				-b --Bug fix (bugfix)
				-s --Start project
				-f --Finish project
				-d --Documentation or anyone comment on code only
				-w --Work in progress
				-r --Code review suggestion changes
				-p --Performance related changes
				-m --Maintenance changes: linter, config updates, etc.
				-rem --Removing code
			`)

		return
	}

	var currentCtx string

	if os.Getenv("CTX") != "" {
		currentCtx = os.Getenv("CTX")
	} else if os.Getenv("GLOBAL_COMMIT_CTX") != "" {
		currentCtx = os.Getenv("GLOBAL_COMMIT_CTX")
	} else {
		currentCtx = "default"
	}

	if len(currentCtx) > 10 {
		fmt.Println("The lenth of context should not be greather than 10.")
	}

	var templatemsg string

	switch firstArg {
	case "-b":
		templatemsg = ":bug:	bugfix(" + currentCtx + ")"
	case "-s":
		templatemsg = ":tada: start(" + currentCtx + ")"
	case "-f":
		templatemsg = ":package: finish(" + currentCtx + ")"
	case "-w":
		templatemsg = ":construction: wip(" + currentCtx + ")"
	case "-d":
		templatemsg = ":books: doc(" + currentCtx + ")"
	case "-p":
		templatemsg = ":racehorse: performance(" + currentCtx + ")"
	case "-m":
		templatemsg = ":wrench: maintenance(" + currentCtx + ")"
	case "-rem":
		templatemsg = ":fire: codingremoval(" + currentCtx + ")"
	case "-r":
		templatemsg = ":pencil: codereview(" + currentCtx + ")"
	default:
		os.Exit(1)
	}

	fullMessageStringified := strings.Join(fullMessageArray, " ")
	templatemsg = templatemsg + " " + fullMessageStringified

	out, err := exec.Command("git", "commit", "-m", "\""+templatemsg+"\"").Output()
	if err != nil {
		fmt.Printf("%s", err)
	}

	output := string(out[:])
	fmt.Println(output)

	fmt.Println("\n Commited successfully!")
}
