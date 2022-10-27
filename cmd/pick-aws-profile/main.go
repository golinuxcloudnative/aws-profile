package main

import (
	"fmt"
	"os"

	"github.com/golinuxcloudnative/aws-export-sso-profile/internal/profile"
	"github.com/golinuxcloudnative/aws-export-sso-profile/internal/prompt"
)

func main() {

	profile := profile.NewProfile("")

	cmd := prompt.NewPrompt(profile)

	if err := cmd.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}

}
