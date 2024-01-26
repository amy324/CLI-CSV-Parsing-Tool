// cmd.go
package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

func init() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

var rootCmd = &cobra.Command{
	Use:   "data-analysis-tool",
	Short: "A CLI tool for data-driven business analysis",
	// Add other command configurations as needed
}

var parseCSVCommand = &cobra.Command{
	Use:   "parsecsv [filename]",
	Short: "Parse CSV file",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		filename := args[0]
		records, err := parseCSV(filename)
		if err != nil {
			log.Fatal(err)
		}

		// Display CSV summary
		fmt.Printf("CSV File: %s\n", filename)
		fmt.Printf("Number of Rows: %d\n", len(records))
		fmt.Printf("Number of Columns: %d\n", len(records[0]))

		// Connect to the database
		db, err := connectDB()
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		// Example: Save CSV data to PostgreSQL table (modify as needed)
		// Assuming you have a table named 'csv_data' with columns 'id', 'name', 'age', 'occupation'
		for i, row := range records {
			// Skip the first row (header)
			if i == 0 {
				continue
			}

			// Convert "Age" to integer
			age, err := strconv.Atoi(row[1])
			if err != nil {
				log.Fatal(err)
			}

			_, err = db.Exec("INSERT INTO csv_data (name, age, occupation) VALUES ($1, $2, $3)", row[0], age, row[2])
			if err != nil {
				log.Fatal(err)
			}
		}

		fmt.Println("CSV data processed and saved to PostgreSQL.")
	},
}

var connectDBCommand = &cobra.Command{
	Use:   "connectdb",
	Short: "Connect to PostgreSQL database",
	Run: func(cmd *cobra.Command, args []string) {
		db, err := connectDB()
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()
		fmt.Println("Connected to the database")
	},
}

func init() {
	rootCmd.AddCommand(parseCSVCommand)
	rootCmd.AddCommand(connectDBCommand)
	// Add other commands as needed
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}

func main() {
	// Execute your CLI commands
	Execute()
}
