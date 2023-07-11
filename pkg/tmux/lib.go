package tmux

import (
	"errors"
	"fmt"
	"os"

	Z "github.com/rwxrob/bonzai/z"
)

// Split accept a direction parameter that's either "vertical" or
// "horizontal" and perform the correct tmux command.
func Split(direction string) error {
	switch direction {
	case "vertical":
		return Z.Exec("tmux", "splitw", "-h")
	case "horizontal":
		return Z.Exec("tmux", "splitw")
	default:
		return errors.New("direction should be either 'vertical' or 'horizontal'")
	}
}

func Edit(path string) error {
	editor := os.Getenv("EDITOR")
	return Z.Exec("tmux", "send-keys", fmt.Sprintf("%s %s", editor, path), "Enter")
}
