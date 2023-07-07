package todo

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/charmbracelet/glamour"
	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/help"
)

const TODO_ZET_ID = 33

var Cmd = &Z.Cmd{
	Name:     "todo",
	Aliases:  []string{`t`},
	Summary:  "A todo branch to provide a simple crud with markdown checkboxes",
	Commands: []*Z.Cmd{pendingCmd, editCmd, help.Cmd},
}

var pendingCmd = &Z.Cmd{
	Name:    "pending",
	Aliases: []string{`p`},
	Call: func(x *Z.Cmd, _ ...string) error {
		zetDir := Z.Out(`pixie`, `keg`, `directory`)
		todoFile, err := os.Open(fmt.Sprintf("%s/%d/README.md", zetDir, TODO_ZET_ID))

		if err != nil {
			return err
		}

		defer todoFile.Close()

		content, err := ioutil.ReadAll(todoFile)

		if err != nil {
			return err
		}

		strContent := strings.Replace(string(content), "# todos", "", 1)

		out, err := glamour.Render(strContent, "dark")

		if err != nil {
			return err
		}

		fmt.Print(out)

		return nil
	},
}

var editCmd = &Z.Cmd{
	Name:    "edit",
	Aliases: []string{`e`},
	Call: func(x *Z.Cmd, _ ...string) error {
		return Z.Exec("pixie", "keg", fmt.Sprint(TODO_ZET_ID))
	},
}
