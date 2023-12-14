package gh

import (
	"fmt"
	"os"

	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/help"
)

var Cmd = &Z.Cmd{
	Name:     "github",
	Aliases:  []string{`gh`},
	Summary:  "A github branch to manipulate the gh cli with more complex actions",
	Commands: []*Z.Cmd{prCmd, profileCmd, help.Cmd},
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

		if len(args[0]) > 0 {
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

var profileCmd = &Z.Cmd{
	Name:     "profile",
	Aliases:  []string{`p`},
	Commands: []*Z.Cmd{profileSwitchCmd, help.Cmd},
}

var profileSwitchCmd = &Z.Cmd{
	Name:    "switch",
	Aliases: []string{`s`},
	Call: func(x *Z.Cmd, args ...string) error {
		Z.Exec("gh", "auth", "switch")

		home := os.Getenv("HOME")

		// Open the file (create or truncate)
		file, err := os.Create(fmt.Sprintf("%s/.local/share/gh_profile.txt", home))
		if err != nil {
			fmt.Println("Error creating/truncating file:", err)
			return nil
		}
		defer file.Close()

		content := Z.Out("gh", "auth", "status")

		// Write the content to the file
		_, err = file.WriteString(content)
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return nil
		}

		return nil
	},
}
