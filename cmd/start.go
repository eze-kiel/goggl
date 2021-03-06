package cmd

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"time"

	"github.com/eze-kiel/goggl/session"
	"github.com/mitchellh/go-homedir"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start a new working session",
	Long: `Start a new working session. A JSON file will be created under running/
with the start time, the tag and the name of the session.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := start(cmd, args); err != nil {
			logrus.Fatalf("cannot start a new work session: %s", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
	startCmd.Flags().StringP("tag", "t", "untagged", "Work session tag")
}

func start(cmd *cobra.Command, args []string) error {
	// Get user home directory
	homedir, err := homedir.Dir()
	if err != nil {
		return err
	}

	// Set running path
	runDir := homedir + "/.goggl/running/"

	// Parse tag
	tag, err := cmd.Flags().GetString("tag")
	if err != nil {
		return err
	}

	// Get date
	date := time.Now().Format(TimeFormat)

	// Create session
	s := session.New()
	s.Name = tag + "_" + date
	s.Tag = tag
	s.StartTime = date

	// Write session in JSON file
	jsonData, err := json.Marshal(s)
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(runDir+s.Name+".json", jsonData, os.ModePerm); err != nil {
		return err
	}

	logrus.Infof("work session %s created successfully ! GLHF", s.Name)
	return nil
}
