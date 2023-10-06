package ranma


type QuestionMetadata struct {
	Question string
	Level string
	Package string
	FuncName string
}


var QuestionBank = map[string]*QuestionMetadata {
	"reverse_linked_list": &QuestionMetadata {
		Question: "Write a function that reverses a linked list",
		Level: "medium",
		Package: "linked_list",
		FuncName: "ReverseList",
	},
	"dfs_search": &QuestionMetadata {
		Question: "Find a key K in a graph based data structure",
		Level: "medium",
		Package: "graph",
		FuncName: "FindNodeDFS",
	},
}