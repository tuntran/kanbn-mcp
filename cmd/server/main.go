package main

import (
	"context"
	"log"
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

	log.Printf("kanbn-mcp server starting (base URL: %s)", baseURL)

	if err := server.Run(context.Background(), &mcp.StdioTransport{}); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
