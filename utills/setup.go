package utills

import (
	"fmt"
	"os"

	"github.com/go-git/go-git/v5"
	"github.com/hoehwa/meok/constants"
)

func checkIfError(err error) {
	if err == nil {
		return
	}

	fmt.Printf("\x1b[31;1m%s\x1b[0m\n", fmt.Sprintf("error: %s", err))
	os.Exit(1)
}

// info should be used to describe the example commands that are about to run.
func info(format string, args ...interface{}) {
	fmt.Printf("\x1b[34;1m%s\x1b[0m\n", fmt.Sprintf(format, args...))
}

func Setup(owner string, repo string) {
	// Clone the given repository to the given directory
	targetURL := fmt.Sprintf("https://github.com/%s/%s", owner, repo)

	info(fmt.Sprintf("git clone %s", targetURL))

	_, err := git.PlainClone(constants.BaseDir, false, &git.CloneOptions{
		URL:      targetURL,
		Progress: os.Stdout,
	})

	checkIfError(err)
}
