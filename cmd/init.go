/*
Copyright © 2020 Tomasz Martin <tomasz.martin@gmail.com>

*/
package cmd

import (
	"fmt"
	"log"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	gockify "github.com/tomaszmartin/gockify/src"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Set up the app variables.",
	Run: func(cmd *cobra.Command, args []string) {
		err := viper.ReadInConfig()
		if err != nil {
			log.Fatal(err)
		}
		var apiKey, chosenIndex string
		fmt.Print("Insert API key: ")
		fmt.Scanln(&apiKey)
		client := gockify.NewClient(apiKey)
		workspaces, err := client.GetWorkspaces()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Initialized. Available workspaces:")
		for index, space := range workspaces {
			fmt.Printf("\t(%d) %s\n", index, space.Name)
		}
		fmt.Print("Choose default workspace: ")
		fmt.Scanln(&chosenIndex)
		workspaceIndex, err := strconv.ParseInt(chosenIndex, 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		chosen := workspaces[workspaceIndex]
		viper.Set("WorkspaceID", chosen.ID)
		viper.Set("ApiKey", apiKey)
		err = viper.WriteConfig()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Workspace [%s] set as defualt.\n", chosen.Name)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
