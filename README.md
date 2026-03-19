# kanbn-mcp

MCP (Model Context Protocol) server for [Kan.bn](https://kan.bn) — exposes 51 tools for AI agents to manage boards, cards, lists, labels, workspaces, users, and attachments via the Kan.bn REST API.

## Requirements

- Go 1.22+
- A Kan.bn account with an API key

## Configuration

| Variable | Required | Description |
|----------|----------|-------------|
| `KAN_API_KEY` | ✓ | API key from Kan.bn account settings |
| `KAN_BASE_URL` | ✓ | Base URL e.g. `https://app.kan.bn/api/v1` |
| `MCP_HTTP_ADDR` | — | HTTP listen address (e.g. `:8080`). If set, serves Streamable HTTP for remote AI agents. If unset, uses stdio (Claude Desktop). Defaults to `:8080` in Docker image. |

## Build

```bash
git clone https://github.com/tuntran/kanbn-mcp
cd kanbn-mcp
go build -o kanbn-mcp ./cmd/server
```

## Docker / Coolify

```bash
docker build -t kanbn-mcp .

# Streamable HTTP (remote AI agents, Coolify) — default port 8080
docker run -p 8080:8080 \
  -e KAN_API_KEY=your_key \
  -e KAN_BASE_URL=https://app.kan.bn/api/v1 \
  kanbn-mcp

# stdio (local only — override MCP_HTTP_ADDR to empty)
docker run -i \
  -e KAN_API_KEY=your_key \
  -e KAN_BASE_URL=https://app.kan.bn/api/v1 \
  -e MCP_HTTP_ADDR="" \
  kanbn-mcp
```

**Coolify:** Set domain + port `8080`. Add env vars `KAN_API_KEY` và `KAN_BASE_URL`. MCP endpoint: `https://your-domain/mcp`.

## Usage with Claude Desktop

Add to `~/.config/claude/claude_desktop_config.json`:

```json
{
  "mcpServers": {
    "kanbn": {
      "command": "/path/to/kanbn-mcp",
      "env": {
        "KAN_API_KEY": "your_api_key",
        "KAN_BASE_URL": "https://app.kan.bn/api/v1"
      }
    }
  }
}
```

## Usage with Claude Code (CLI)

Add to `.claude/mcp.json` in your project:

```json
{
  "mcpServers": {
    "kanbn": {
      "command": "./kanbn-mcp",
      "env": {
        "KAN_API_KEY": "your_api_key",
        "KAN_BASE_URL": "https://app.kan.bn/api/v1"
      }
    }
  }
}
```

## Tools

### Boards (7)
`list_boards` · `create_board` · `get_board` · `update_board` · `delete_board` · `get_board_by_slug` · `check_board_slug_available`

### Cards (17)
`create_card` · `get_card` · `update_card` · `delete_card` · `add_comment` · `update_comment` · `delete_comment` · `add_label_to_card` · `remove_label_from_card` · `add_member_to_card` · `remove_member_from_card` · `get_card_activities` · `add_checklist` · `delete_checklist` · `add_checklist_item` · `update_checklist_item` · `delete_checklist_item`

### Lists (3)
`create_list` · `update_list` · `delete_list`

### Labels (4)
`create_label` · `get_label` · `update_label` · `delete_label`

### Workspaces (10)
`list_workspaces` · `create_workspace` · `get_workspace` · `update_workspace` · `delete_workspace` · `get_workspace_by_slug` · `check_workspace_slug_available` · `search_workspace` · `invite_member` · `remove_member`

### Invites (5)
`get_active_invite_link` · `create_invite_link` · `deactivate_invite_link` · `get_invite_by_code` · `accept_invite`

### Users (2)
`get_user` · `update_user`

### Attachments (3)
`generate_presigned_url` · `confirm_attachment_upload` · `delete_attachment`

> **Note:** File upload is a 2-step flow: call `generate_presigned_url` → upload to the presigned S3 URL → call `confirm_attachment_upload` with the returned `key`.

## API Reference

See [docs.kan.bn/api-reference](https://docs.kan.bn/api-reference/introduction).
