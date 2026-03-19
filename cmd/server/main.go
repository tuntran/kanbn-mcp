package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/tuntran/kanbn-mcp/internal/kan"
	"github.com/tuntran/kanbn-mcp/internal/tools"
)

func main() {
	apiKey := os.Getenv("KAN_API_KEY")
	if apiKey == "" {
		log.Fatal("KAN_API_KEY environment variable is required")
	}

	baseURL := os.Getenv("KAN_BASE_URL")
	if baseURL == "" {
		log.Fatal("KAN_BASE_URL environment variable is required")
	}

	client := kan.NewClient(baseURL, apiKey)

	server := mcp.NewServer(&mcp.Implementation{
		Name:    "kanbn-mcp",
		Version: "1.0.0",
	}, nil)

	tools.RegisterAll(server, client)

	// MCP_HTTP_ADDR: if set, serve over Streamable HTTP (for remote AI agents / Coolify).
	// If not set, fall back to stdio (for local MCP clients like Claude Desktop).
	httpAddr := os.Getenv("MCP_HTTP_ADDR")
	if httpAddr != "" {
		log.Printf("kanbn-mcp HTTP server listening on %s", httpAddr)
		handler := mcp.NewStreamableHTTPHandler(func(*http.Request) *mcp.Server {
			return server
		}, nil)
		if err := http.ListenAndServe(httpAddr, handler); err != nil {
			log.Fatalf("http server error: %v", err)
		}
		return
	}

	log.Printf("kanbn-mcp stdio server starting (base URL: %s)", baseURL)
	if err := server.Run(context.Background(), &mcp.StdioTransport{}); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
