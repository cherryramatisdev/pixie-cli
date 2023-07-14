package pixie

import (
	"github.com/cherryramatisdev/pixie-cli/pkg/gh"
	"github.com/cherryramatisdev/pixie-cli/pkg/tmux"
	"github.com/cherryramatisdev/pixie-cli/pkg/todo"
	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/help"
	"github.com/rwxrob/keg"
	"github.com/rwxrob/pomo"
)

var Cmd = &Z.Cmd{
	Name: "pixie",
	Shortcuts: map[string][]string{
		`1v1`: {`keg`, `1v1`},
	},
	Summary:     "A general purpose command line tool",
	Version:     "0.0.1",
	Description: "A general purpose command line tool",
	Site:        "https://github.com/cherryramatisdev/pixie-cli",
	Source:      "https://github.com/cherryramatisdev/pixie-cli",
	Issues:      "https://github.com/cherryramatisdev/pixie-cli/issues",
	Commands: []*Z.Cmd{
		help.Cmd,
		pomo.Cmd,
		keg.Cmd,
		todo.Cmd,
		tmux.Cmd,
		gh.Cmd,
	},
}
