package gh

import (
	"os"

	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/help"
)

var Cmd = &Z.Cmd{
	Name:     "github",
	Aliases:  []string{`gh`},
	Summary:  "A github branch to manipulate the gh cli with more complex actions",
	Commands: []*Z.Cmd{prCmd, help.Cmd},
}

var prCmd = &Z.Cmd{
	Name:     "pull-request",
	Aliases:  []string{`pr`},
	Commands: []*Z.Cmd{prCreateCmd, help.Cmd},
}

var prCreateCmd = &Z.Cmd{
	Name:    "create",
	Aliases: []string{`c`},
	Call: func(x *Z.Cmd, args ...string) error {
		var branch string
		reviewers := os.Getenv("REVIEWERS")

		if args[0] == "main" || args[0] == "develop" {
			branch = args[0]
		} else {
			err := Z.Exec("git", "show-ref", "--verify", "--quiet", "refs/heads/develop")

			if err != nil {
				branch = "main"
			} else {
				branch = "develop"
			}
		}

		cmd := []string{"gh", "pr", "-a", "@me", "create", "-B", branch, "-r", reviewers, "--draft"}

		branchNumber, err := extractBranchNumber()
		if err == nil {
			title := Z.Out("jira", "pr-title", branchNumber)
			cmd = append(cmd, "-t")
			cmd = append(cmd, title)
		}

		if err := Z.Exec(cmd...); err != nil {
			return err
		}

		return Z.Exec("gh", "pr", "view", "-w")
	},
}
