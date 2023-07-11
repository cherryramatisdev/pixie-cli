package tmux

import (
	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/help"
)

var Cmd = &Z.Cmd{
	Name:     "tmux",
	Aliases:  []string{`tm`},
	Summary:  "A tmux branch for more complex but common actions",
	Commands: []*Z.Cmd{editCmd, help.Cmd},
}

var editCmd = &Z.Cmd{
	Name:     "edit",
	Aliases:  []string{`e`},
	Commands: []*Z.Cmd{verticalSplitCmd, horizontalSplitCmd, help.Cmd},
}

var verticalSplitCmd = &Z.Cmd{
	Name:    "vertical",
	Aliases: []string{`v`},
	MinArgs: 1,
	Call: func(x *Z.Cmd, args ...string) error {
		if err := Split("vertical"); err != nil {
			return err
		}

		if err := Edit(args[0]); err != nil {
			return err
		}

		return nil
	},
}

var horizontalSplitCmd = &Z.Cmd{
	Name:    "horizontal",
	Aliases: []string{`s`, `h`},
	MinArgs: 1,
	Call: func(x *Z.Cmd, args ...string) error {
		if err := Split("horizontal"); err != nil {
			return err
		}

		if err := Edit(args[0]); err != nil {
			return err
		}

		return nil
	},
}
