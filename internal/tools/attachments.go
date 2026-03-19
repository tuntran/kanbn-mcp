package tools

import (
	"context"
	"fmt"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/tuntran/kanbn-mcp/internal/kan"
)

// --- Input structs ---

type generatePresignedURLInput struct {
	CardPublicId string `json:"cardPublicId" jsonschema:"required,description=Card public ID"`
	Filename     string `json:"filename" jsonschema:"required,description=File name (1-255 characters)"`
	ContentType  string `json:"contentType" jsonschema:"required,description=MIME type (e.g. image/png)"`
	Size         int64  `json:"size" jsonschema:"required,description=File size in bytes (max 52428800)"`
}

type confirmAttachmentInput struct {
	CardPublicId string `json:"cardPublicId" jsonschema:"required,description=Card public ID"`
	Key          string `json:"key" jsonschema:"required,description=Upload key returned from generate_presigned_url"`
	Filename     string `json:"filename" jsonschema:"required,description=File name"`
}

type deleteAttachmentInput struct {
	CardPublicId string `json:"cardPublicId" jsonschema:"required,description=Card public ID"`
	AttachmentId string `json:"attachmentId" jsonschema:"required,description=Attachment ID"`
}

// --- Handlers ---

func handleGeneratePresignedURL(c *kan.Client) func(context.Context, *mcp.CallToolRequest, generatePresignedURLInput) (*mcp.CallToolResult, any, error) {
	return func(ctx context.Context, _ *mcp.CallToolRequest, input generatePresignedURLInput) (*mcp.CallToolResult, any, error) {
		data, err := c.Post(ctx, fmt.Sprintf("/cards/%s/attachments/upload-url", input.CardPublicId), kan.GeneratePresignedURLInput{
			Filename:    input.Filename,
			ContentType: input.ContentType,
			Size:        input.Size,
		})
		if err != nil {
			return nil, nil, err
		}
		return textResult(data), nil, nil
	}
}

func handleConfirmAttachmentUpload(c *kan.Client) func(context.Context, *mcp.CallToolRequest, confirmAttachmentInput) (*mcp.CallToolResult, any, error) {
	return func(ctx context.Context, _ *mcp.CallToolRequest, input confirmAttachmentInput) (*mcp.CallToolResult, any, error) {
		data, err := c.Post(ctx, fmt.Sprintf("/cards/%s/attachments", input.CardPublicId), kan.ConfirmAttachmentInput{
			Key:      input.Key,
			Filename: input.Filename,
		})
		if err != nil {
			return nil, nil, err
		}
		return textResult(data), nil, nil
	}
}

func handleDeleteAttachment(c *kan.Client) func(context.Context, *mcp.CallToolRequest, deleteAttachmentInput) (*mcp.CallToolResult, any, error) {
	return func(ctx context.Context, _ *mcp.CallToolRequest, input deleteAttachmentInput) (*mcp.CallToolResult, any, error) {
		data, err := c.Delete(ctx, fmt.Sprintf("/cards/%s/attachments/%s", input.CardPublicId, input.AttachmentId))
		if err != nil {
			return nil, nil, err
		}
		return textResult(data), nil, nil
	}
}

// --- Registration ---

func registerAttachments(s *mcp.Server, c *kan.Client) {
	mcp.AddTool(s, &mcp.Tool{
		Name:        "generate_presigned_url",
		Description: "Generate a presigned S3 URL for uploading a file attachment. Step 1 of 2 for attachments.",
	}, handleGeneratePresignedURL(c))
	mcp.AddTool(s, &mcp.Tool{
		Name:        "confirm_attachment_upload",
		Description: "Confirm a file attachment upload after uploading to the presigned URL. Step 2 of 2.",
	}, handleConfirmAttachmentUpload(c))
	mcp.AddTool(s, &mcp.Tool{
		Name:        "delete_attachment",
		Description: "Delete a file attachment from a card",
	}, handleDeleteAttachment(c))
}
