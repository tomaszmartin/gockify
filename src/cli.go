package gockify

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func Init() *Client {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Insert API key: ")
	scanner.Scan()
	apiKey := scanner.Text()
	client := NewClient(apiKey)
	workspaces, err := client.GetWorkspaces()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Initialized. Available workspaces:")
	for index, workspace := range workspaces {
		fmt.Printf("\t(%d) %s\n", index, workspace.Name)
	}
	fmt.Print("Choose default workspace: ")
	scanner.Scan()
	chosenIndex := scanner.Text()
	workspaceIndex, err := strconv.ParseInt(chosenIndex, 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	chosenWorkspace := workspaces[workspaceIndex]
	fmt.Printf("Workspace [%s] set as defualt.\n", chosenWorkspace.Name)
	return client
}
