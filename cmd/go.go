/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/mattismoel/registry/internal/service"
	"github.com/mattismoel/registry/internal/storage/sqlite"
	"github.com/spf13/cobra"
)

// goCmd represents the go command
var goCmd = &cobra.Command{
	Use:   "go",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		store, err := sqlite.New("database.db")
		if err != nil {
			log.Fatal(err)
		}

		err = store.Initialise()
		if err != nil {
			log.Fatal(err)
		}

		service := service.NewProjectService(store)

		search := args[0]
		project, err := service.ByTitle(search)
		if err != nil {
			id, _ := strconv.Atoi(search)
			project, err = service.ByID(int64(id))
			if err != nil {
				log.Fatal(err)
			}
		}

		io.WriteString(os.Stdout, fmt.Sprintf("%s\n", project.Path))
		// err = os.Chdir(project.Path)
		// if err != nil {
		// 	log.Fatal(err)
		// }

		// err = exec.Command("cd", project.Path).Run()
		// if err != nil {
		// 	log.Fatal(err)
		// }
	},
}

func init() {
	rootCmd.AddCommand(goCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// goCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// goCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
