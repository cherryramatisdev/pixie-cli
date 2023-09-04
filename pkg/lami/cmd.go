package lami

import (
	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/help"
)

var Cmd = &Z.Cmd{
	Name:     "lami",
	Aliases:  []string{`l`},
	Summary:  "A work branch for common lami stuff",
	Commands: []*Z.Cmd{openCmd, help.Cmd},
}

var openCmd = &Z.Cmd{
	Name:     "open",
	Aliases:  []string{`o`},
	Commands: []*Z.Cmd{openDiscussCmd, openSlackTemplatesCmd, help.Cmd},
}

var openDiscussCmd = &Z.Cmd{
	Name: "discuss",
	Call: func(x *Z.Cmd, _ ...string) error {
		return Z.Exec("open", "https://github.com/lami-health/discussions/discussions")
	},
}

var openSlackTemplatesCmd = &Z.Cmd{
	Name:    "slack_templates",
	Aliases: []string{`st`},
	Call: func(x *Z.Cmd, _ ...string) error {
		return Z.Exec("open", "https://github.com/lami-health/discussions/discussions/25")
	},
}
