package tools

import (
	"encoding/json"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/tuntran/kanbn-mcp/internal/kan"
)

// RegisterAll registers all Kan.bn tools on the MCP server.
func RegisterAll(s *mcp.Server, c *kan.Client) {
	registerBoards(s, c)
	registerLists(s, c)
	registerLabels(s, c)
	registerCards(s, c)
	registerWorkspaces(s, c)
	registerInvites(s, c)
	registerUsers(s, c)
	registerAttachments(s, c)
}

// textResult wraps raw JSON from the Kan API into an MCP tool result.
func textResult(data json.RawMessage) *mcp.CallToolResult {
	return &mcp.CallToolResult{
		Content: []mcp.Content{&mcp.TextContent{Text: string(data)}},
	}
}
