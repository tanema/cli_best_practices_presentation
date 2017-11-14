package main

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var (
	configPath string
	title      string

	stdOut = log.New(os.Stdout, "", 0)
	stdErr = log.New(os.Stderr, "", 0)

	mainCmd = &cobra.Command{
		Use:   "meetup",
		Short: "A tool for managing your meetup ",
		Long: `Meetup is a tool for managing your meetup from the command line.
Meetup allows you to create, delete and edit events on your meetup`,
		RunE: func(cmd *cobra.Command, args []string) error {
			stdOut.Println("Your config directory is:", configPath)
			return nil
		},
	}

	newCmd = &cobra.Command{
		Use:   "new",
		Short: "creates a new event",
		RunE: func(cmd *cobra.Command, args []string) error {
			if title == "" {
				return fmt.Errorf("Name argument required")
			}
			stdOut.Println("Created your event", title)
			return nil
		},
	}
)

func init() {
	mainCmd.PersistentFlags().StringVarP(&configPath, "config", "c", "/var/app_config.yml", "path to config.yml")

	newCmd.Flags().StringVar(&title, "title", "", "a title for the event you are creating")

	mainCmd.AddCommand(newCmd)
}

func main() {
	if err := mainCmd.Execute(); err != nil {
		stdErr.Fatal(err)
	}
}
