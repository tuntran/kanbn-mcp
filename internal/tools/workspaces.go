package tools

import (
	"context"
	"fmt"
	"net/url"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/tuntran/kanbn-mcp/internal/kan"
)

// --- Input structs ---

type listWorkspacesInput struct{}

type createWorkspaceInput struct {
	Name        string `json:"name" jsonschema:"required,description=Workspace name"`
	Slug        string `json:"slug" jsonschema:"required,description=Workspace URL slug"`
	Description string `json:"description,omitempty" jsonschema:"description=Workspace description"`
}

type getWorkspaceInput struct {
	WorkspacePublicId string `json:"workspacePublicId" jsonschema:"required,description=Workspace public ID"`
}

type updateWorkspaceInput struct {
	WorkspacePublicId string `json:"workspacePublicId" jsonschema:"required,description=Workspace public ID"`
	Name              string `json:"name,omitempty" jsonschema:"description=New workspace name"`
	Description       string `json:"description,omitempty" jsonschema:"description=New workspace description"`
}

type deleteWorkspaceInput struct {
	WorkspacePublicId string `json:"workspacePublicId" jsonschema:"required,description=Workspace public ID"`
}

type getWorkspaceBySlugInput struct {
	Slug string `json:"slug" jsonschema:"required,description=Workspace URL slug"`
}

type checkWorkspaceSlugInput struct {
	Slug string `json:"slug" jsonschema:"required,description=Slug to check for availability"`
}

type searchWorkspaceInput struct {
	WorkspacePublicId string `json:"workspacePublicId" jsonschema:"required,description=Workspace public ID"`
	Query             string `json:"query" jsonschema:"required,description=Search query string"`
}

type inviteMemberInput struct {
	WorkspacePublicId string `json:"workspacePublicId" jsonschema:"required,description=Workspace public ID"`
	Email             string `json:"email" jsonschema:"required,description=Email address to invite"`
}

type removeMemberInput struct {
	WorkspacePublicId string `json:"workspacePublicId" jsonschema:"required,description=Workspace public ID"`
	MemberPublicId    string `json:"memberPublicId" jsonschema:"required,description=Member public ID to remove"`
}

// --- Handlers ---

func handleListWorkspaces(c *kan.Client) func(context.Context, *mcp.CallToolRequest, listWorkspacesInput) (*mcp.CallToolResult, any, error) {
	return func(ctx context.Context, _ *mcp.CallToolRequest, _ listWorkspacesInput) (*mcp.CallToolResult, any, error) {
		data, err := c.Get(ctx, "/workspaces")
		if err != nil {
			return nil, nil, err
		}
		return textResult(data), nil, nil
	}
}

func handleCreateWorkspace(c *kan.Client) func(context.Context, *mcp.CallToolRequest, createWorkspaceInput) (*mcp.CallToolResult, any, error) {
	return func(ctx context.Context, _ *mcp.CallToolRequest, input createWorkspaceInput) (*mcp.CallToolResult, any, error) {
		data, err := c.Post(ctx, "/workspaces", kan.CreateWorkspaceInput{
			Name:        input.Name,
			Slug:        input.Slug,
			Description: input.Description,
		})
		if err != nil {
			return nil, nil, err
		}
		return textResult(data), nil, nil
	}
}

func handleGetWorkspace(c *kan.Client) func(context.Context, *mcp.CallToolRequest, getWorkspaceInput) (*mcp.CallToolResult, any, error) {
	return func(ctx context.Context, _ *mcp.CallToolRequest, input getWorkspaceInput) (*mcp.CallToolResult, any, error) {
		data, err := c.Get(ctx, fmt.Sprintf("/workspaces/%s", input.WorkspacePublicId))
		if err != nil {
			return nil, nil, err
		}
		return textResult(data), nil, nil
	}
}

func handleUpdateWorkspace(c *kan.Client) func(context.Context, *mcp.CallToolRequest, updateWorkspaceInput) (*mcp.CallToolResult, any, error) {
	return func(ctx context.Context, _ *mcp.CallToolRequest, input updateWorkspaceInput) (*mcp.CallToolResult, any, error) {
		data, err := c.Put(ctx, fmt.Sprintf("/workspaces/%s", input.WorkspacePublicId), kan.UpdateWorkspaceInput{
			Name:        input.Name,
			Description: input.Description,
		})
		if err != nil {
			return nil, nil, err
		}
		return textResult(data), nil, nil
	}
}

func handleDeleteWorkspace(c *kan.Client) func(context.Context, *mcp.CallToolRequest, deleteWorkspaceInput) (*mcp.CallToolResult, any, error) {
	return func(ctx context.Context, _ *mcp.CallToolRequest, input deleteWorkspaceInput) (*mcp.CallToolResult, any, error) {
		data, err := c.Delete(ctx, fmt.Sprintf("/workspaces/%s", input.WorkspacePublicId))
		if err != nil {
			return nil, nil, err
		}
		return textResult(data), nil, nil
	}
}

func handleGetWorkspaceBySlug(c *kan.Client) func(context.Context, *mcp.CallToolRequest, getWorkspaceBySlugInput) (*mcp.CallToolResult, any, error) {
	return func(ctx context.Context, _ *mcp.CallToolRequest, input getWorkspaceBySlugInput) (*mcp.CallToolResult, any, error) {
		data, err := c.Get(ctx, "/workspaces?slug="+url.QueryEscape(input.Slug))
		if err != nil {
			return nil, nil, err
		}
		return textResult(data), nil, nil
	}
}

func handleCheckWorkspaceSlug(c *kan.Client) func(context.Context, *mcp.CallToolRequest, checkWorkspaceSlugInput) (*mcp.CallToolResult, any, error) {
	return func(ctx context.Context, _ *mcp.CallToolRequest, input checkWorkspaceSlugInput) (*mcp.CallToolResult, any, error) {
		data, err := c.Get(ctx, "/workspaces/check-slug?slug="+url.QueryEscape(input.Slug))
		if err != nil {
			return nil, nil, err
		}
		return textResult(data), nil, nil
	}
}

func handleSearchWorkspace(c *kan.Client) func(context.Context, *mcp.CallToolRequest, searchWorkspaceInput) (*mcp.CallToolResult, any, error) {
	return func(ctx context.Context, _ *mcp.CallToolRequest, input searchWorkspaceInput) (*mcp.CallToolResult, any, error) {
		path := fmt.Sprintf("/workspaces/%s/search?query=%s", input.WorkspacePublicId, url.QueryEscape(input.Query))
		data, err := c.Get(ctx, path)
		if err != nil {
			return nil, nil, err
		}
		return textResult(data), nil, nil
	}
}

func handleInviteMember(c *kan.Client) func(context.Context, *mcp.CallToolRequest, inviteMemberInput) (*mcp.CallToolResult, any, error) {
	return func(ctx context.Context, _ *mcp.CallToolRequest, input inviteMemberInput) (*mcp.CallToolResult, any, error) {
		data, err := c.Post(ctx, fmt.Sprintf("/workspaces/%s/members/invite", input.WorkspacePublicId), kan.InviteMemberInput{Email: input.Email})
		if err != nil {
			return nil, nil, err
		}
		return textResult(data), nil, nil
	}
}

func handleRemoveMember(c *kan.Client) func(context.Context, *mcp.CallToolRequest, removeMemberInput) (*mcp.CallToolResult, any, error) {
	return func(ctx context.Context, _ *mcp.CallToolRequest, input removeMemberInput) (*mcp.CallToolResult, any, error) {
		data, err := c.Delete(ctx, fmt.Sprintf("/workspaces/%s/members/%s", input.WorkspacePublicId, input.MemberPublicId))
		if err != nil {
			return nil, nil, err
		}
		return textResult(data), nil, nil
	}
}

// --- Registration ---

func registerWorkspaces(s *mcp.Server, c *kan.Client) {
	mcp.AddTool(s, &mcp.Tool{Name: "list_workspaces", Description: "List all workspaces for the authenticated user"}, handleListWorkspaces(c))
	mcp.AddTool(s, &mcp.Tool{Name: "create_workspace", Description: "Create a new workspace"}, handleCreateWorkspace(c))
	mcp.AddTool(s, &mcp.Tool{Name: "get_workspace", Description: "Get a workspace by its public ID"}, handleGetWorkspace(c))
	mcp.AddTool(s, &mcp.Tool{Name: "update_workspace", Description: "Update workspace name or description"}, handleUpdateWorkspace(c))
	mcp.AddTool(s, &mcp.Tool{Name: "delete_workspace", Description: "Delete a workspace by its public ID"}, handleDeleteWorkspace(c))
	mcp.AddTool(s, &mcp.Tool{Name: "get_workspace_by_slug", Description: "Get a workspace by its URL slug"}, handleGetWorkspaceBySlug(c))
	mcp.AddTool(s, &mcp.Tool{Name: "check_workspace_slug_available", Description: "Check if a workspace URL slug is available"}, handleCheckWorkspaceSlug(c))
	mcp.AddTool(s, &mcp.Tool{Name: "search_workspace", Description: "Search boards and cards within a workspace"}, handleSearchWorkspace(c))
	mcp.AddTool(s, &mcp.Tool{Name: "invite_member", Description: "Invite a member to a workspace by email"}, handleInviteMember(c))
	mcp.AddTool(s, &mcp.Tool{Name: "remove_member", Description: "Remove a member from a workspace"}, handleRemoveMember(c))
}
