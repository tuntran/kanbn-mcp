package tools

import (
	"context"
	"fmt"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/tuntran/kanbn-mcp/internal/kan"
)

// --- Input structs ---

type createLabelInput struct {
	WorkspacePublicId string `json:"workspacePublicId" jsonschema:"Public ID of the workspace"`
	Name              string `json:"name" jsonschema:"Label name"`
	ColourCode        string `json:"colourCode,omitempty" jsonschema:"Hex colour code (e.g. #FF5733)"`
}

type getLabelInput struct {
	LabelPublicId string `json:"labelPublicId" jsonschema:"Label public ID"`
}

type updateLabelInput struct {
	LabelPublicId string `json:"labelPublicId" jsonschema:"Label public ID"`
	Name          string `json:"name,omitempty" jsonschema:"New label name"`
	ColourCode    string `json:"colourCode,omitempty" jsonschema:"New hex colour code"`
}

type deleteLabelInput struct {
	LabelPublicId string `json:"labelPublicId" jsonschema:"Label public ID"`
}

// --- Handlers ---

func handleCreateLabel(c *kan.Client) func(context.Context, *mcp.CallToolRequest, createLabelInput) (*mcp.CallToolResult, any, error) {
	return func(ctx context.Context, _ *mcp.CallToolRequest, input createLabelInput) (*mcp.CallToolResult, any, error) {
		data, err := c.Post(ctx, "/labels", kan.CreateLabelInput{
			WorkspacePublicId: input.WorkspacePublicId,
			Name:              input.Name,
			ColourCode:        input.ColourCode,
		})
		if err != nil {
			return nil, nil, err
		}
		return textResult(data), nil, nil
	}
}

func handleGetLabel(c *kan.Client) func(context.Context, *mcp.CallToolRequest, getLabelInput) (*mcp.CallToolResult, any, error) {
	return func(ctx context.Context, _ *mcp.CallToolRequest, input getLabelInput) (*mcp.CallToolResult, any, error) {
		data, err := c.Get(ctx, fmt.Sprintf("/labels/%s", input.LabelPublicId))
		if err != nil {
			return nil, nil, err
		}
		return textResult(data), nil, nil
	}
}

func handleUpdateLabel(c *kan.Client) func(context.Context, *mcp.CallToolRequest, updateLabelInput) (*mcp.CallToolResult, any, error) {
	return func(ctx context.Context, _ *mcp.CallToolRequest, input updateLabelInput) (*mcp.CallToolResult, any, error) {
		data, err := c.Put(ctx, fmt.Sprintf("/labels/%s", input.LabelPublicId), kan.UpdateLabelInput{
			Name:       input.Name,
			ColourCode: input.ColourCode,
		})
		if err != nil {
			return nil, nil, err
		}
		return textResult(data), nil, nil
	}
}

func handleDeleteLabel(c *kan.Client) func(context.Context, *mcp.CallToolRequest, deleteLabelInput) (*mcp.CallToolResult, any, error) {
	return func(ctx context.Context, _ *mcp.CallToolRequest, input deleteLabelInput) (*mcp.CallToolResult, any, error) {
		data, err := c.Delete(ctx, fmt.Sprintf("/labels/%s", input.LabelPublicId))
		if err != nil {
			return nil, nil, err
		}
		return textResult(data), nil, nil
	}
}

// --- Registration ---

func registerLabels(s *mcp.Server, c *kan.Client) {
	mcp.AddTool(s, &mcp.Tool{Name: "create_label", Description: "Create a new label in a workspace"}, handleCreateLabel(c))
	mcp.AddTool(s, &mcp.Tool{Name: "get_label", Description: "Get a label by its public ID"}, handleGetLabel(c))
	mcp.AddTool(s, &mcp.Tool{Name: "update_label", Description: "Update a label name or colour"}, handleUpdateLabel(c))
	mcp.AddTool(s, &mcp.Tool{Name: "delete_label", Description: "Delete a label by its public ID"}, handleDeleteLabel(c))
}
