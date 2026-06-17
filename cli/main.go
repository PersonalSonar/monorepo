package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/cobra"
)

type ServiceStatus struct {
	Name      string    `json:"name"`
	Status    string    `json:"status"`
	Port      int       `json:"port"`
	CheckedAt time.Time `json:"checked_at"`
}

var rootCmd = &cobra.Command{
	Use:   "monorepo",
	Short: "Monorepo CLI tool",
	Long:  `CLI tool for managing monorepo services`,
}

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Check status of all services",
	Run: func(cmd *cobra.Command, args []string) {
		services := []ServiceStatus{
			{Name: "API", Status: "running", Port: 3000, CheckedAt: time.Now()},
			{Name: "Worker", Status: "running", Port: 0, CheckedAt: time.Now()},
		}

		fmt.Println("Service Status:")
		fmt.Println("==============")
		for _, svc := range services {
			fmt.Printf("%-15s %s", svc.Name, svc.Status)
			if svc.Port > 0 {
				fmt.Printf(" (port %d)", svc.Port)
			}
			fmt.Println()
		}
	},
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all apps in the monorepo",
	Run: func(cmd *cobra.Command, args []string) {
		apps := []map[string]string{
			{"name": "api", "language": "TypeScript/Node.js", "description": "REST API"},
			{"name": "worker", "language": "Python", "description": "Job processor"},
			{"name": "cli", "language": "Go", "description": "CLI tool"},
		}

		fmt.Println("Monorepo Apps:")
		fmt.Println("==============")
		for _, app := range apps {
			fmt.Printf("%-12s %s\n", app["name"], app["description"])
			fmt.Printf("%-12s Language: %s\n", "", app["language"])
			fmt.Println()
		}
	},
}

var logCmd = &cobra.Command{
	Use:   "logs [service]",
	Short: "View service logs",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Usage: monorepo logs [api|worker]")
			return
		}

		service := args[0]
		fmt.Printf("Logs for %s service:\n", service)
		fmt.Println("======================")
		fmt.Printf("[%s] Service started\n", time.Now().Format("15:04:05"))
		fmt.Printf("[%s] Listening on port\n", time.Now().Format("15:04:05"))
	},
}

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Show monorepo info",
	Run: func(cmd *cobra.Command, args []string) {
		info := map[string]interface{}{
			"name":       "monorepo",
			"version":    "1.0.0",
			"apps":       3,
			"languages":  []string{"TypeScript", "Python", "Go"},
			"created_at": "2024-01-01",
		}

		data, _ := json.MarshalIndent(info, "", "  ")
		fmt.Println(string(data))
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(logCmd)
	rootCmd.AddCommand(infoCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
