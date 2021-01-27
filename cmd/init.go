package cmd

import (
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize goggl files",
	Long: `Init creates required folders in the user's home dir under .goggl/
Two folders will be created: running/ and history/`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := initialize(); err != nil {
			logrus.Fatalf("cannot init goggl files: %s", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}

func initialize() error {
	// Get user home directory
	homedir, err := homedir.Dir()
	if err != nil {
		return nil
	}
	gogglRoot := homedir + "/.goggl"
	logrus.Infof("creating goggl files at %s", gogglRoot)

	// Check if goggl has already been init
	if _, err := os.Stat(gogglRoot); !os.IsNotExist(err) {
		logrus.Warn("goggl seems to be already initialized. Skipping init.")
		return nil
	}

	// Create root dir
	if err := os.Mkdir(gogglRoot, os.ModePerm); err != nil {
		return err
	}

	// Create running dir
	runDir := gogglRoot + "/running"
	if err := os.Mkdir(runDir, os.ModePerm); err != nil {
		return err
	}

	// Create history dir
	histDir := gogglRoot + "/history"
	if err := os.Mkdir(histDir, os.ModePerm); err != nil {
		return err
	}

	logrus.Infof("initialization has been successful !")
	return nil
}
