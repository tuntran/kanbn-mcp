package tools

import (
	"context"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/tuntran/kanbn-mcp/internal/kan"
)

// --- Input structs ---

type getUserInput struct{}

type updateUserInput struct {
	Name  string `json:"name,omitempty" jsonschema:"Display name"`
	Image string `json:"image,omitempty" jsonschema:"Profile image URL"`
}

// --- Handlers ---

func handleGetUser(c *kan.Client) func(context.Context, *mcp.CallToolRequest, getUserInput) (*mcp.CallToolResult, any, error) {
	return func(ctx context.Context, _ *mcp.CallToolRequest, _ getUserInput) (*mcp.CallToolResult, any, error) {
		data, err := c.Get(ctx, "/users/me")
		if err != nil {
			return nil, nil, err
		}
		return textResult(data), nil, nil
	}
}

func handleUpdateUser(c *kan.Client) func(context.Context, *mcp.CallToolRequest, updateUserInput) (*mcp.CallToolResult, any, error) {
	return func(ctx context.Context, _ *mcp.CallToolRequest, input updateUserInput) (*mcp.CallToolResult, any, error) {
		data, err := c.Put(ctx, "/users/me", kan.UpdateUserInput{
			Name:  input.Name,
			Image: input.Image,
		})
		if err != nil {
			return nil, nil, err
		}
		return textResult(data), nil, nil
	}
}

// --- Registration ---

func registerUsers(s *mcp.Server, c *kan.Client) {
	mcp.AddTool(s, &mcp.Tool{Name: "get_user", Description: "Get the current authenticated user profile"}, handleGetUser(c))
	mcp.AddTool(s, &mcp.Tool{Name: "update_user", Description: "Update the current user's display name or profile image"}, handleUpdateUser(c))
}
