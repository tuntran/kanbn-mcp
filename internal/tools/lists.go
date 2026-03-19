package tools

import (
	"context"
	"fmt"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/tuntran/kanbn-mcp/internal/kan"
)

// --- Input structs ---

type createListInput struct {
	BoardPublicId string `json:"boardPublicId" jsonschema:"Public ID of the board"`
	Name          string `json:"name" jsonschema:"List name"`
}

type updateListInput struct {
	ListPublicId string `json:"listPublicId" jsonschema:"List public ID"`
	Name         string `json:"name" jsonschema:"New list name"`
}

type deleteListInput struct {
	ListPublicId string `json:"listPublicId" jsonschema:"List public ID"`
}

// --- Handlers ---

func handleCreateList(c *kan.Client) func(context.Context, *mcp.CallToolRequest, createListInput) (*mcp.CallToolResult, any, error) {
	return func(ctx context.Context, _ *mcp.CallToolRequest, input createListInput) (*mcp.CallToolResult, any, error) {
		data, err := c.Post(ctx, "/lists", kan.CreateListInput{
			BoardPublicId: input.BoardPublicId,
			Name:          input.Name,
		})
		if err != nil {
			return nil, nil, err
		}
		return textResult(data), nil, nil
	}
}

func handleUpdateList(c *kan.Client) func(context.Context, *mcp.CallToolRequest, updateListInput) (*mcp.CallToolResult, any, error) {
	return func(ctx context.Context, _ *mcp.CallToolRequest, input updateListInput) (*mcp.CallToolResult, any, error) {
		data, err := c.Put(ctx, fmt.Sprintf("/lists/%s", input.ListPublicId), kan.UpdateListInput{
			Name: input.Name,
		})
		if err != nil {
			return nil, nil, err
		}
		return textResult(data), nil, nil
	}
}

func handleDeleteList(c *kan.Client) func(context.Context, *mcp.CallToolRequest, deleteListInput) (*mcp.CallToolResult, any, error) {
	return func(ctx context.Context, _ *mcp.CallToolRequest, input deleteListInput) (*mcp.CallToolResult, any, error) {
		data, err := c.Delete(ctx, fmt.Sprintf("/lists/%s", input.ListPublicId))
		if err != nil {
			return nil, nil, err
		}
		return textResult(data), nil, nil
	}
}

// --- Registration ---

func registerLists(s *mcp.Server, c *kan.Client) {
	mcp.AddTool(s, &mcp.Tool{Name: "create_list", Description: "Create a new list in a board"}, handleCreateList(c))
	mcp.AddTool(s, &mcp.Tool{Name: "update_list", Description: "Update a list name"}, handleUpdateList(c))
	mcp.AddTool(s, &mcp.Tool{Name: "delete_list", Description: "Delete a list by its public ID"}, handleDeleteList(c))
}
