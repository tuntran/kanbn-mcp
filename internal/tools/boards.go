package tools

import (
	"context"
	"fmt"
	"net/url"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/tuntran/kanbn-mcp/internal/kan"
)

// --- Input structs ---

type listBoardsInput struct {
	WorkspacePublicId string `json:"workspacePublicId" jsonschema:"required,description=Public ID of the workspace"`
}

type createBoardInput struct {
	WorkspacePublicId string `json:"workspacePublicId" jsonschema:"required,description=Public ID of the workspace"`
	Name              string `json:"name" jsonschema:"required,description=Board name"`
	Slug              string `json:"slug" jsonschema:"required,description=Board URL slug"`
	Description       string `json:"description,omitempty" jsonschema:"description=Board description"`
}

type getBoardInput struct {
	BoardPublicId string `json:"boardPublicId" jsonschema:"required,description=Board public ID"`
}

type updateBoardInput struct {
	BoardPublicId string `json:"boardPublicId" jsonschema:"required,description=Board public ID"`
	Name          string `json:"name,omitempty" jsonschema:"description=New board name"`
	Description   string `json:"description,omitempty" jsonschema:"description=New board description"`
}

type deleteBoardInput struct {
	BoardPublicId string `json:"boardPublicId" jsonschema:"required,description=Board public ID"`
}

type getBoardBySlugInput struct {
	Slug string `json:"slug" jsonschema:"required,description=Board URL slug"`
}

type checkBoardSlugInput struct {
	Slug string `json:"slug" jsonschema:"required,description=Slug to check for availability"`
}

// --- Handlers ---

func handleListBoards(c *kan.Client) func(context.Context, *mcp.CallToolRequest, listBoardsInput) (*mcp.CallToolResult, any, error) {
	return func(ctx context.Context, _ *mcp.CallToolRequest, input listBoardsInput) (*mcp.CallToolResult, any, error) {
		data, err := c.Get(ctx, fmt.Sprintf("/workspaces/%s/boards", input.WorkspacePublicId))
		if err != nil {
			return nil, nil, err
		}
		return textResult(data), nil, nil
	}
}

func handleCreateBoard(c *kan.Client) func(context.Context, *mcp.CallToolRequest, createBoardInput) (*mcp.CallToolResult, any, error) {
	return func(ctx context.Context, _ *mcp.CallToolRequest, input createBoardInput) (*mcp.CallToolResult, any, error) {
		data, err := c.Post(ctx, "/boards", kan.CreateBoardInput{
			WorkspacePublicId: input.WorkspacePublicId,
			Name:              input.Name,
			Slug:              input.Slug,
			Description:       input.Description,
		})
		if err != nil {
			return nil, nil, err
		}
		return textResult(data), nil, nil
	}
}

func handleGetBoard(c *kan.Client) func(context.Context, *mcp.CallToolRequest, getBoardInput) (*mcp.CallToolResult, any, error) {
	return func(ctx context.Context, _ *mcp.CallToolRequest, input getBoardInput) (*mcp.CallToolResult, any, error) {
		data, err := c.Get(ctx, fmt.Sprintf("/boards/%s", input.BoardPublicId))
		if err != nil {
			return nil, nil, err
		}
		return textResult(data), nil, nil
	}
}

func handleUpdateBoard(c *kan.Client) func(context.Context, *mcp.CallToolRequest, updateBoardInput) (*mcp.CallToolResult, any, error) {
	return func(ctx context.Context, _ *mcp.CallToolRequest, input updateBoardInput) (*mcp.CallToolResult, any, error) {
		data, err := c.Put(ctx, fmt.Sprintf("/boards/%s", input.BoardPublicId), kan.UpdateBoardInput{
			Name:        input.Name,
			Description: input.Description,
		})
		if err != nil {
			return nil, nil, err
		}
		return textResult(data), nil, nil
	}
}

func handleDeleteBoard(c *kan.Client) func(context.Context, *mcp.CallToolRequest, deleteBoardInput) (*mcp.CallToolResult, any, error) {
	return func(ctx context.Context, _ *mcp.CallToolRequest, input deleteBoardInput) (*mcp.CallToolResult, any, error) {
		data, err := c.Delete(ctx, fmt.Sprintf("/boards/%s", input.BoardPublicId))
		if err != nil {
			return nil, nil, err
		}
		return textResult(data), nil, nil
	}
}

func handleGetBoardBySlug(c *kan.Client) func(context.Context, *mcp.CallToolRequest, getBoardBySlugInput) (*mcp.CallToolResult, any, error) {
	return func(ctx context.Context, _ *mcp.CallToolRequest, input getBoardBySlugInput) (*mcp.CallToolResult, any, error) {
		data, err := c.Get(ctx, "/boards?slug="+url.QueryEscape(input.Slug))
		if err != nil {
			return nil, nil, err
		}
		return textResult(data), nil, nil
	}
}

func handleCheckBoardSlug(c *kan.Client) func(context.Context, *mcp.CallToolRequest, checkBoardSlugInput) (*mcp.CallToolResult, any, error) {
	return func(ctx context.Context, _ *mcp.CallToolRequest, input checkBoardSlugInput) (*mcp.CallToolResult, any, error) {
		data, err := c.Get(ctx, "/boards/check-slug?slug="+url.QueryEscape(input.Slug))
		if err != nil {
			return nil, nil, err
		}
		return textResult(data), nil, nil
	}
}

// --- Registration ---

func registerBoards(s *mcp.Server, c *kan.Client) {
	mcp.AddTool(s, &mcp.Tool{Name: "list_boards", Description: "List all boards in a workspace"}, handleListBoards(c))
	mcp.AddTool(s, &mcp.Tool{Name: "create_board", Description: "Create a new board in a workspace"}, handleCreateBoard(c))
	mcp.AddTool(s, &mcp.Tool{Name: "get_board", Description: "Get a board by its public ID"}, handleGetBoard(c))
	mcp.AddTool(s, &mcp.Tool{Name: "update_board", Description: "Update board name or description"}, handleUpdateBoard(c))
	mcp.AddTool(s, &mcp.Tool{Name: "delete_board", Description: "Delete a board by its public ID"}, handleDeleteBoard(c))
	mcp.AddTool(s, &mcp.Tool{Name: "get_board_by_slug", Description: "Get a board by its URL slug"}, handleGetBoardBySlug(c))
	mcp.AddTool(s, &mcp.Tool{Name: "check_board_slug_available", Description: "Check if a board URL slug is available"}, handleCheckBoardSlug(c))
}
