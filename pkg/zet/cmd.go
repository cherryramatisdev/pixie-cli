package zet

import (
	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/help"
)

var Cmd = &Z.Cmd{
	Name:     "zet",
	Aliases:  []string{"z"},
	Summary:  "A zettelkaten branch",
	Commands: []*Z.Cmd{help.Cmd},
	Call: func(x *Z.Cmd, _ ...string) error {
		panic("not implemented")
	},
}
