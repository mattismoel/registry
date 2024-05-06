/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/mattismoel/registry/internal/model"
	"github.com/mattismoel/registry/internal/service"
	"github.com/mattismoel/registry/internal/storage/sqlite"
	"github.com/spf13/cobra"
)

var (
	title string
	path  string
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add entry to project registry.",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("add called")
		store, err := sqlite.New("database.db")
		if err != nil {
			log.Fatal(err)
		}

		err = store.Initialise()
		if err != nil {
			log.Fatal(err)
		}

		projectSrv := service.NewProjectService(store)

		if strings.HasPrefix(path, "~") {
			homeDir, err := os.UserHomeDir()
			if err != nil {
				log.Fatal(err)
			}
			path = strings.Replace(path, "~", homeDir, 1)
		}

		absPath, err := filepath.Abs(path)
		if err != nil {
			log.Fatalf("could not set abs path: %v", err)
		}

		p := model.Project{
			Title: title,
			Path:  absPath,
		}

		err = projectSrv.Add(p)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().StringVarP(&title, "title", "t", "", "Project title")
	addCmd.Flags().StringVarP(&path, "path", "p", ".", "Project path (absolute)")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
