package tools

import (
	"context"
	"fmt"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/tuntran/kanbn-mcp/internal/kan"
)

// --- Input structs ---

type inviteLinkInput struct {
	WorkspacePublicId string `json:"workspacePublicId" jsonschema:"Workspace public ID"`
}

type inviteCodeInput struct {
	InviteCode string `json:"inviteCode" jsonschema:"Invite link code"`
}

// --- Handlers ---

func handleGetActiveInviteLink(c *kan.Client) func(context.Context, *mcp.CallToolRequest, inviteLinkInput) (*mcp.CallToolResult, any, error) {
	return func(ctx context.Context, _ *mcp.CallToolRequest, input inviteLinkInput) (*mcp.CallToolResult, any, error) {
		data, err := c.Get(ctx, fmt.Sprintf("/workspaces/%s/invite", input.WorkspacePublicId))
		if err != nil {
			return nil, nil, err
		}
		return textResult(data), nil, nil
	}
}

func handleCreateInviteLink(c *kan.Client) func(context.Context, *mcp.CallToolRequest, inviteLinkInput) (*mcp.CallToolResult, any, error) {
	return func(ctx context.Context, _ *mcp.CallToolRequest, input inviteLinkInput) (*mcp.CallToolResult, any, error) {
		data, err := c.Post(ctx, fmt.Sprintf("/workspaces/%s/invite", input.WorkspacePublicId), nil)
		if err != nil {
			return nil, nil, err
		}
		return textResult(data), nil, nil
	}
}

func handleDeactivateInviteLink(c *kan.Client) func(context.Context, *mcp.CallToolRequest, inviteLinkInput) (*mcp.CallToolResult, any, error) {
	return func(ctx context.Context, _ *mcp.CallToolRequest, input inviteLinkInput) (*mcp.CallToolResult, any, error) {
		data, err := c.Put(ctx, fmt.Sprintf("/workspaces/%s/invite", input.WorkspacePublicId), nil)
		if err != nil {
			return nil, nil, err
		}
		return textResult(data), nil, nil
	}
}

func handleGetInviteByCode(c *kan.Client) func(context.Context, *mcp.CallToolRequest, inviteCodeInput) (*mcp.CallToolResult, any, error) {
	return func(ctx context.Context, _ *mcp.CallToolRequest, input inviteCodeInput) (*mcp.CallToolResult, any, error) {
		data, err := c.Get(ctx, fmt.Sprintf("/invites/%s", input.InviteCode))
		if err != nil {
			return nil, nil, err
		}
		return textResult(data), nil, nil
	}
}

func handleAcceptInvite(c *kan.Client) func(context.Context, *mcp.CallToolRequest, inviteCodeInput) (*mcp.CallToolResult, any, error) {
	return func(ctx context.Context, _ *mcp.CallToolRequest, input inviteCodeInput) (*mcp.CallToolResult, any, error) {
		data, err := c.Post(ctx, fmt.Sprintf("/invites/%s/accept", input.InviteCode), nil)
		if err != nil {
			return nil, nil, err
		}
		return textResult(data), nil, nil
	}
}

// --- Registration ---

func registerInvites(s *mcp.Server, c *kan.Client) {
	mcp.AddTool(s, &mcp.Tool{Name: "get_active_invite_link", Description: "Get the active invite link for a workspace"}, handleGetActiveInviteLink(c))
	mcp.AddTool(s, &mcp.Tool{Name: "create_invite_link", Description: "Create a new invite link for a workspace"}, handleCreateInviteLink(c))
	mcp.AddTool(s, &mcp.Tool{Name: "deactivate_invite_link", Description: "Deactivate the current invite link for a workspace"}, handleDeactivateInviteLink(c))
	mcp.AddTool(s, &mcp.Tool{Name: "get_invite_by_code", Description: "Get invite information by invite code"}, handleGetInviteByCode(c))
	mcp.AddTool(s, &mcp.Tool{Name: "accept_invite", Description: "Accept a workspace invite using an invite code"}, handleAcceptInvite(c))
}
