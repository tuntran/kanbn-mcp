package tools

import (
	"context"
	"fmt"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/tuntran/kanbn-mcp/internal/kan"
)

// --- Input structs ---

type createCardInput struct {
	Title           string   `json:"title" jsonschema:"required,description=Card title"`
	ListPublicId    string   `json:"listPublicId" jsonschema:"required,description=List to place the card in"`
	Position        string   `json:"position" jsonschema:"required,description=Position in list: start or end"`
	Description     string   `json:"description,omitempty" jsonschema:"description=Card description (markdown supported)"`
	DueDate         string   `json:"dueDate,omitempty" jsonschema:"description=Due date in ISO 8601 format (e.g. 2026-12-31)"`
	LabelPublicIds  []string `json:"labelPublicIds,omitempty" jsonschema:"description=Label public IDs to assign"`
	MemberPublicIds []string `json:"memberPublicIds,omitempty" jsonschema:"description=Member public IDs to assign"`
}

type getCardInput struct {
	CardPublicId string `json:"cardPublicId" jsonschema:"required,description=Card public ID"`
}

type updateCardInput struct {
	CardPublicId string `json:"cardPublicId" jsonschema:"required,description=Card public ID"`
	Title        string `json:"title,omitempty" jsonschema:"description=New title"`
	Description  string `json:"description,omitempty" jsonschema:"description=New description"`
	DueDate      string `json:"dueDate,omitempty" jsonschema:"description=New due date (ISO 8601)"`
	ListPublicId string `json:"listPublicId,omitempty" jsonschema:"description=Move card to this list"`
	Position     *int   `json:"position,omitempty" jsonschema:"description=New position index in the list"`
}

type deleteCardInput struct {
	CardPublicId string `json:"cardPublicId" jsonschema:"required,description=Card public ID"`
}

type addCommentInput struct {
	CardPublicId string `json:"cardPublicId" jsonschema:"required,description=Card public ID"`
	Content      string `json:"content" jsonschema:"required,description=Comment text"`
}

type updateCommentInput struct {
	CardPublicId string `json:"cardPublicId" jsonschema:"required,description=Card public ID"`
	CommentId    string `json:"commentId" jsonschema:"required,description=Comment ID"`
	Content      string `json:"content" jsonschema:"required,description=New comment text"`
}

type deleteCommentInput struct {
	CardPublicId string `json:"cardPublicId" jsonschema:"required,description=Card public ID"`
	CommentId    string `json:"commentId" jsonschema:"required,description=Comment ID"`
}

type cardLabelInput struct {
	CardPublicId  string `json:"cardPublicId" jsonschema:"required,description=Card public ID"`
	LabelPublicId string `json:"labelPublicId" jsonschema:"required,description=Label public ID"`
}

type cardMemberInput struct {
	CardPublicId   string `json:"cardPublicId" jsonschema:"required,description=Card public ID"`
	MemberPublicId string `json:"memberPublicId" jsonschema:"required,description=Member public ID"`
}

type getCardActivitiesInput struct {
	CardPublicId string `json:"cardPublicId" jsonschema:"required,description=Card public ID"`
	Page         *int   `json:"page,omitempty" jsonschema:"description=Page number for pagination"`
}

type addChecklistInput struct {
	CardPublicId string `json:"cardPublicId" jsonschema:"required,description=Card public ID"`
	Name         string `json:"name" jsonschema:"required,description=Checklist name"`
}

type deleteChecklistInput struct {
	CardPublicId string `json:"cardPublicId" jsonschema:"required,description=Card public ID"`
	ChecklistId  string `json:"checklistId" jsonschema:"required,description=Checklist ID"`
}

type addChecklistItemInput struct {
	CardPublicId string `json:"cardPublicId" jsonschema:"required,description=Card public ID"`
	ChecklistId  string `json:"checklistId" jsonschema:"required,description=Checklist ID"`
	Name         string `json:"name" jsonschema:"required,description=Item name"`
}

type updateChecklistItemInput struct {
	CardPublicId string `json:"cardPublicId" jsonschema:"required,description=Card public ID"`
	ItemId       string `json:"itemId" jsonschema:"required,description=Checklist item ID"`
	Name         string `json:"name,omitempty" jsonschema:"description=New item name"`
	Completed    *bool  `json:"completed,omitempty" jsonschema:"description=Mark as completed (true) or not (false)"`
}

type deleteChecklistItemInput struct {
	CardPublicId string `json:"cardPublicId" jsonschema:"required,description=Card public ID"`
	ItemId       string `json:"itemId" jsonschema:"required,description=Checklist item ID"`
}

// --- Handlers ---

func handleCreateCard(c *kan.Client) func(context.Context, *mcp.CallToolRequest, createCardInput) (*mcp.CallToolResult, any, error) {
	return func(ctx context.Context, _ *mcp.CallToolRequest, input createCardInput) (*mcp.CallToolResult, any, error) {
		data, err := c.Post(ctx, "/cards", kan.CreateCardInput{
			Title:           input.Title,
			ListPublicId:    input.ListPublicId,
			Position:        input.Position,
			Description:     input.Description,
			DueDate:         input.DueDate,
			LabelPublicIds:  input.LabelPublicIds,
			MemberPublicIds: input.MemberPublicIds,
		})
		if err != nil {
			return nil, nil, err
		}
		return textResult(data), nil, nil
	}
}

func handleGetCard(c *kan.Client) func(context.Context, *mcp.CallToolRequest, getCardInput) (*mcp.CallToolResult, any, error) {
	return func(ctx context.Context, _ *mcp.CallToolRequest, input getCardInput) (*mcp.CallToolResult, any, error) {
		data, err := c.Get(ctx, fmt.Sprintf("/cards/%s", input.CardPublicId))
		if err != nil {
			return nil, nil, err
		}
		return textResult(data), nil, nil
	}
}

func handleUpdateCard(c *kan.Client) func(context.Context, *mcp.CallToolRequest, updateCardInput) (*mcp.CallToolResult, any, error) {
	return func(ctx context.Context, _ *mcp.CallToolRequest, input updateCardInput) (*mcp.CallToolResult, any, error) {
		data, err := c.Put(ctx, fmt.Sprintf("/cards/%s", input.CardPublicId), kan.UpdateCardInput{
			Title:        input.Title,
			Description:  input.Description,
			DueDate:      input.DueDate,
			ListPublicId: input.ListPublicId,
			Position:     input.Position,
		})
		if err != nil {
			return nil, nil, err
		}
		return textResult(data), nil, nil
	}
}

func handleDeleteCard(c *kan.Client) func(context.Context, *mcp.CallToolRequest, deleteCardInput) (*mcp.CallToolResult, any, error) {
	return func(ctx context.Context, _ *mcp.CallToolRequest, input deleteCardInput) (*mcp.CallToolResult, any, error) {
		data, err := c.Delete(ctx, fmt.Sprintf("/cards/%s", input.CardPublicId))
		if err != nil {
			return nil, nil, err
		}
		return textResult(data), nil, nil
	}
}

func handleAddComment(c *kan.Client) func(context.Context, *mcp.CallToolRequest, addCommentInput) (*mcp.CallToolResult, any, error) {
	return func(ctx context.Context, _ *mcp.CallToolRequest, input addCommentInput) (*mcp.CallToolResult, any, error) {
		data, err := c.Post(ctx, fmt.Sprintf("/cards/%s/comments", input.CardPublicId), kan.CommentInput{Content: input.Content})
		if err != nil {
			return nil, nil, err
		}
		return textResult(data), nil, nil
	}
}

func handleUpdateComment(c *kan.Client) func(context.Context, *mcp.CallToolRequest, updateCommentInput) (*mcp.CallToolResult, any, error) {
	return func(ctx context.Context, _ *mcp.CallToolRequest, input updateCommentInput) (*mcp.CallToolResult, any, error) {
		data, err := c.Put(ctx, fmt.Sprintf("/cards/%s/comments/%s", input.CardPublicId, input.CommentId), kan.CommentInput{Content: input.Content})
		if err != nil {
			return nil, nil, err
		}
		return textResult(data), nil, nil
	}
}

func handleDeleteComment(c *kan.Client) func(context.Context, *mcp.CallToolRequest, deleteCommentInput) (*mcp.CallToolResult, any, error) {
	return func(ctx context.Context, _ *mcp.CallToolRequest, input deleteCommentInput) (*mcp.CallToolResult, any, error) {
		data, err := c.Delete(ctx, fmt.Sprintf("/cards/%s/comments/%s", input.CardPublicId, input.CommentId))
		if err != nil {
			return nil, nil, err
		}
		return textResult(data), nil, nil
	}
}

func handleAddLabelToCard(c *kan.Client) func(context.Context, *mcp.CallToolRequest, cardLabelInput) (*mcp.CallToolResult, any, error) {
	return func(ctx context.Context, _ *mcp.CallToolRequest, input cardLabelInput) (*mcp.CallToolResult, any, error) {
		data, err := c.Post(ctx, fmt.Sprintf("/cards/%s/labels/%s", input.CardPublicId, input.LabelPublicId), nil)
		if err != nil {
			return nil, nil, err
		}
		return textResult(data), nil, nil
	}
}

func handleRemoveLabelFromCard(c *kan.Client) func(context.Context, *mcp.CallToolRequest, cardLabelInput) (*mcp.CallToolResult, any, error) {
	return func(ctx context.Context, _ *mcp.CallToolRequest, input cardLabelInput) (*mcp.CallToolResult, any, error) {
		data, err := c.Delete(ctx, fmt.Sprintf("/cards/%s/labels/%s", input.CardPublicId, input.LabelPublicId))
		if err != nil {
			return nil, nil, err
		}
		return textResult(data), nil, nil
	}
}

func handleAddMemberToCard(c *kan.Client) func(context.Context, *mcp.CallToolRequest, cardMemberInput) (*mcp.CallToolResult, any, error) {
	return func(ctx context.Context, _ *mcp.CallToolRequest, input cardMemberInput) (*mcp.CallToolResult, any, error) {
		data, err := c.Post(ctx, fmt.Sprintf("/cards/%s/members/%s", input.CardPublicId, input.MemberPublicId), nil)
		if err != nil {
			return nil, nil, err
		}
		return textResult(data), nil, nil
	}
}

func handleRemoveMemberFromCard(c *kan.Client) func(context.Context, *mcp.CallToolRequest, cardMemberInput) (*mcp.CallToolResult, any, error) {
	return func(ctx context.Context, _ *mcp.CallToolRequest, input cardMemberInput) (*mcp.CallToolResult, any, error) {
		data, err := c.Delete(ctx, fmt.Sprintf("/cards/%s/members/%s", input.CardPublicId, input.MemberPublicId))
		if err != nil {
			return nil, nil, err
		}
		return textResult(data), nil, nil
	}
}

func handleGetCardActivities(c *kan.Client) func(context.Context, *mcp.CallToolRequest, getCardActivitiesInput) (*mcp.CallToolResult, any, error) {
	return func(ctx context.Context, _ *mcp.CallToolRequest, input getCardActivitiesInput) (*mcp.CallToolResult, any, error) {
		path := fmt.Sprintf("/cards/%s/activities", input.CardPublicId)
		if input.Page != nil {
			path += fmt.Sprintf("?page=%d", *input.Page)
		}
		data, err := c.Get(ctx, path)
		if err != nil {
			return nil, nil, err
		}
		return textResult(data), nil, nil
	}
}

func handleAddChecklist(c *kan.Client) func(context.Context, *mcp.CallToolRequest, addChecklistInput) (*mcp.CallToolResult, any, error) {
	return func(ctx context.Context, _ *mcp.CallToolRequest, input addChecklistInput) (*mcp.CallToolResult, any, error) {
		data, err := c.Post(ctx, fmt.Sprintf("/cards/%s/checklists", input.CardPublicId), kan.ChecklistInput{Name: input.Name})
		if err != nil {
			return nil, nil, err
		}
		return textResult(data), nil, nil
	}
}

func handleDeleteChecklist(c *kan.Client) func(context.Context, *mcp.CallToolRequest, deleteChecklistInput) (*mcp.CallToolResult, any, error) {
	return func(ctx context.Context, _ *mcp.CallToolRequest, input deleteChecklistInput) (*mcp.CallToolResult, any, error) {
		data, err := c.Delete(ctx, fmt.Sprintf("/cards/%s/checklists/%s", input.CardPublicId, input.ChecklistId))
		if err != nil {
			return nil, nil, err
		}
		return textResult(data), nil, nil
	}
}

func handleAddChecklistItem(c *kan.Client) func(context.Context, *mcp.CallToolRequest, addChecklistItemInput) (*mcp.CallToolResult, any, error) {
	return func(ctx context.Context, _ *mcp.CallToolRequest, input addChecklistItemInput) (*mcp.CallToolResult, any, error) {
		// checklistId sent as body field per API spec
		body := struct {
			Name        string `json:"name"`
			ChecklistId string `json:"checklistId"`
		}{Name: input.Name, ChecklistId: input.ChecklistId}
		data, err := c.Post(ctx, fmt.Sprintf("/cards/%s/checklist-items", input.CardPublicId), body)
		if err != nil {
			return nil, nil, err
		}
		return textResult(data), nil, nil
	}
}

func handleUpdateChecklistItem(c *kan.Client) func(context.Context, *mcp.CallToolRequest, updateChecklistItemInput) (*mcp.CallToolResult, any, error) {
	return func(ctx context.Context, _ *mcp.CallToolRequest, input updateChecklistItemInput) (*mcp.CallToolResult, any, error) {
		data, err := c.Put(ctx, fmt.Sprintf("/cards/%s/checklist-items/%s", input.CardPublicId, input.ItemId), kan.ChecklistItemInput{
			Name:      input.Name,
			Completed: input.Completed,
		})
		if err != nil {
			return nil, nil, err
		}
		return textResult(data), nil, nil
	}
}

func handleDeleteChecklistItem(c *kan.Client) func(context.Context, *mcp.CallToolRequest, deleteChecklistItemInput) (*mcp.CallToolResult, any, error) {
	return func(ctx context.Context, _ *mcp.CallToolRequest, input deleteChecklistItemInput) (*mcp.CallToolResult, any, error) {
		data, err := c.Delete(ctx, fmt.Sprintf("/cards/%s/checklist-items/%s", input.CardPublicId, input.ItemId))
		if err != nil {
			return nil, nil, err
		}
		return textResult(data), nil, nil
	}
}

// --- Registration ---

func registerCards(s *mcp.Server, c *kan.Client) {
	mcp.AddTool(s, &mcp.Tool{Name: "create_card", Description: "Create a new card in a list"}, handleCreateCard(c))
	mcp.AddTool(s, &mcp.Tool{Name: "get_card", Description: "Get a card by its public ID"}, handleGetCard(c))
	mcp.AddTool(s, &mcp.Tool{Name: "update_card", Description: "Update card fields (title, description, due date, list, position)"}, handleUpdateCard(c))
	mcp.AddTool(s, &mcp.Tool{Name: "delete_card", Description: "Delete a card by its public ID"}, handleDeleteCard(c))
	mcp.AddTool(s, &mcp.Tool{Name: "add_comment", Description: "Add a comment to a card"}, handleAddComment(c))
	mcp.AddTool(s, &mcp.Tool{Name: "update_comment", Description: "Update a comment on a card"}, handleUpdateComment(c))
	mcp.AddTool(s, &mcp.Tool{Name: "delete_comment", Description: "Delete a comment from a card"}, handleDeleteComment(c))
	mcp.AddTool(s, &mcp.Tool{Name: "add_label_to_card", Description: "Add a label to a card"}, handleAddLabelToCard(c))
	mcp.AddTool(s, &mcp.Tool{Name: "remove_label_from_card", Description: "Remove a label from a card"}, handleRemoveLabelFromCard(c))
	mcp.AddTool(s, &mcp.Tool{Name: "add_member_to_card", Description: "Add a member to a card"}, handleAddMemberToCard(c))
	mcp.AddTool(s, &mcp.Tool{Name: "remove_member_from_card", Description: "Remove a member from a card"}, handleRemoveMemberFromCard(c))
	mcp.AddTool(s, &mcp.Tool{Name: "get_card_activities", Description: "Get paginated activity log for a card"}, handleGetCardActivities(c))
	mcp.AddTool(s, &mcp.Tool{Name: "add_checklist", Description: "Add a checklist to a card"}, handleAddChecklist(c))
	mcp.AddTool(s, &mcp.Tool{Name: "delete_checklist", Description: "Delete a checklist from a card"}, handleDeleteChecklist(c))
	mcp.AddTool(s, &mcp.Tool{Name: "add_checklist_item", Description: "Add an item to a checklist"}, handleAddChecklistItem(c))
	mcp.AddTool(s, &mcp.Tool{Name: "update_checklist_item", Description: "Update a checklist item (name or completed status)"}, handleUpdateChecklistItem(c))
	mcp.AddTool(s, &mcp.Tool{Name: "delete_checklist_item", Description: "Delete a checklist item"}, handleDeleteChecklistItem(c))
}
