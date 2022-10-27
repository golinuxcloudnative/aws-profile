package prompt

import (
	"fmt"
	"os"
	"os/user"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/golinuxcloudnative/aws-export-sso-profile/internal/domain"
	"github.com/golinuxcloudnative/aws-export-sso-profile/internal/profile"
)

type Prompt struct {
	profile *profile.Profile
}

func NewPrompt(profile *profile.Profile) *Prompt {
	return &Prompt{
		profile: profile,
	}
}

func (prompt *Prompt) Run() error {

	// get all AWS profiles
	profiles, err := prompt.profile.Profiles()
	if err != nil {
		return err
	}

	// generate the list
	items := prompt.generateList(profiles)

	l := list.New(items, list.NewDefaultDelegate(), 20, 20)
	l.Title = "List of AWS Profiles"

	m := model{list: l}

	p := tea.NewProgram(m)

	for {
		mod, err := p.StartReturningModel()
		if err != nil {
			return err
		}

		if mod.(model).quitting {
			return nil
		}

		errPrompt := prompt.NewShell(mod.(model).choice)
		if errPrompt != nil {
			return errPrompt
		}
	}
}

func (prompt *Prompt) NewShell(profile string) error {
	// current user
	me, err := user.Current()
	if err != nil {
		return err
	}

	// get the current working directory.
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	// set AWS Environment
	os.Setenv("AWS_PROFILE", profile)

	// stdin, stdout, and stderr to the new process
	pa := os.ProcAttr{
		Files: []*os.File{os.Stdin, os.Stdout, os.Stderr},
		Dir:   cwd,
	}

	// new shell.
	proc, err := os.StartProcess("/usr/bin/login", []string{"login", "-fpl", me.Username}, &pa)
	if err != nil {
		return err
	}

	// exit the shell
	state, err := proc.Wait()
	if err != nil {
		return err
	}

	fmt.Printf(">> Ending a shell with AWS profile %s, status: %s\n", profile, state.String())

	return nil
}

func (prompt *Prompt) generateList(profiles []*domain.Profile) []list.Item {
	items := []list.Item{}
	for _, profile := range profiles {
		items = append(items, item{
			name:      profile.Name(),
			region:    profile.Region(),
			url:       profile.Url(),
			role:      profile.Role(),
			accountId: profile.AccountId(),
		})
	}
	return items
}
