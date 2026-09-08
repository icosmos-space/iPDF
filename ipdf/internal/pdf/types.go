package pdf

// ToolResult is the unified response for tool operations.
type ToolResult struct {
	OutputPaths []string `json:"outputPaths"`
	Message     string   `json:"message"`
}
