package internals

import "github.com/spf13/cobra"
import "fmt"
import "ranma"
import "ranma/data"
import "os"
import "strings"

func init() {
	listCmd := &cobra.Command {
	Use: "candidates",
	Short: "Select a question from the bank of questions",
	Run: func(cmd *cobra.Command, args []string) {
		for key, _ := range ranma.QuestionBank {
			fmt.Println(key)
		}
	},
	}
	rootCmd.AddCommand(listCmd)
}

func init() {
	var questionId string
	pickCmd := &cobra.Command{
	Use: "pick",
	Short: "generate coding question",
	Run: func(cmd *cobra.Command, args[] string) {
		if questionId == "" {
			fmt.Println("missing question id, use candidates cmd to introspect")
			return
		}

		if _, ok := ranma.QuestionBank[questionId]; !ok {
			fmt.Println("pick a valid question id, use candidates cmd to introspect")
		}


		question := ranma.QuestionBank[questionId]
		dirPath := question.Package

		_, err := os.Stat(dirPath)
		fmt.Printf("Creating go package %s...\n", dirPath)
		if os.IsNotExist(err) {
			err := os.Mkdir(dirPath, 0755)
			if err != nil {
				fmt.Println("Error creating package:", err)
				return
			}
			fmt.Println("Package created successfully:", dirPath)
		} else if err == nil {
			// Directory already exists
			fmt.Println("Package already exists:", dirPath)
		} else {
			fmt.Println("Error checking package:", err)
			return
		}

		filePath := fmt.Sprintf("%s/%s_test.go", dirPath, strings.ToLower(question.FuncName))
		file, err := os.Create(filePath)

		if err != nil {
			fmt.Println("error creating file:", err)
			return
		}

		_, err = file.WriteString(fmt.Sprintf(data.GO_CODE, dirPath, question.FuncName, question.Question, question.FuncName))

		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}

		fmt.Println("happy hacking!")

	},
	}
	pickCmd.Flags().StringVarP(&questionId, "question", "q", "", "question to solve")
	rootCmd.AddCommand(pickCmd)
}