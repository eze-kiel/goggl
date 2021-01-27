package cmd

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/eze-kiel/goggl/session"
	"github.com/mitchellh/go-homedir"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// stopCmd represents the stop command
var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop a running working session",
	Long: `Stop and archive a running session. The JSON file will be moved to 
history/$YEAR/$MONTH/`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := stop(cmd, args); err != nil {
			logrus.Fatalf("cannot stop running session: %s", err)
		}
	},
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		if len(args) != 0 {
			return nil, cobra.ShellCompDirectiveNoFileComp
		}
		return getRunningSessions(toComplete), cobra.ShellCompDirectiveNoFileComp
	},
}

func init() {
	rootCmd.AddCommand(stopCmd)
}

func stop(cmd *cobra.Command, args []string) error {
	// Get user home directory
	homedir, err := homedir.Dir()
	if err != nil {
		return err
	}

	f := homedir + "/.goggl/running/" + args[0]
	data, err := ioutil.ReadFile(f)
	if err != nil {
		return err
	}

	// Unmarshall file
	s := session.New()
	json.Unmarshal(data, &s)

	// Set end
	s.EndTime = time.Now().Format(TimeFormat)

	// Set duration
	begin, err := time.Parse(TimeFormat, s.StartTime)
	if err != nil {
		return err
	}
	end, err := time.Parse(TimeFormat, s.EndTime)
	if err != nil {
		return err
	}
	s.Duration = end.Sub(begin).String()

	logrus.Infof("work session duration: %s", s.Duration)

	y, m, _ := begin.Date()

	// If the year folder doesn't exists, create it
	if _, err := os.Stat(homedir + "/.goggl/history/" + strconv.Itoa(y)); os.IsNotExist(err) {
		os.Mkdir(homedir+"/.goggl/history/"+strconv.Itoa(y), os.ModePerm)
	}

	// If the month folder doesn't exists, create it
	if _, err := os.Stat(homedir + "/.goggl/history/" + strconv.Itoa(y) + "/" + m.String()); os.IsNotExist(err) {
		os.Mkdir(homedir+"/.goggl/history/"+strconv.Itoa(y)+"/"+m.String(), os.ModePerm)
	}

	fullPath := homedir + "/.goggl/history/" + strconv.Itoa(y) + "/" + m.String()

	// Write session in JSON file
	jsonData, err := json.Marshal(s)
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(fullPath+"/"+s.Name+".json", jsonData, os.ModePerm); err != nil {
		return err
	}
	logrus.Infof("archived %s into history", s.Name)

	// Delete file from running/
	os.Remove(f)

	return nil
}

func getRunningSessions(pfx string) []string {
	// Get user home directory
	homedir, err := homedir.Dir()
	if err != nil {
		logrus.Fatalf("cannot get user's homedir: %s", err)
	}

	// Read running/ content
	var files []string
	fileInfo, err := ioutil.ReadDir(homedir + "/.goggl/running/")
	if err != nil {
		logrus.Fatalf("error reading running directory: %s", err)
	}
	for _, file := range fileInfo {
		files = append(files, file.Name())
	}

	res := []string{}
	for _, k := range files {
		if strings.HasPrefix(k, pfx) {
			res = append(res, k)
		}
	}

	return res
}
