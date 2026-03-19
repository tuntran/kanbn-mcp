package kan

// --- Boards ---

// CreateBoardInput is the request body for creating a board.
type CreateBoardInput struct {
	WorkspacePublicId string `json:"workspacePublicId"`
	Name              string `json:"name"`
	Slug              string `json:"slug"`
	Description       string `json:"description,omitempty"`
}

// UpdateBoardInput is the request body for updating a board.
type UpdateBoardInput struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

// --- Cards ---

// CreateCardInput is the request body for creating a card.
type CreateCardInput struct {
	Title           string   `json:"title"`
	ListPublicId    string   `json:"listPublicId"`
	Position        string   `json:"position"` // "start" or "end"
	Description     string   `json:"description,omitempty"`
	DueDate         string   `json:"dueDate,omitempty"`
	LabelPublicIds  []string `json:"labelPublicIds,omitempty"`
	MemberPublicIds []string `json:"memberPublicIds,omitempty"`
}

// UpdateCardInput is the request body for updating a card.
type UpdateCardInput struct {
	Title        string `json:"title,omitempty"`
	Description  string `json:"description,omitempty"`
	DueDate      string `json:"dueDate,omitempty"`
	ListPublicId string `json:"listPublicId,omitempty"`
	Position     *int   `json:"position,omitempty"`
}

// CommentInput is the request body for creating or updating a comment.
type CommentInput struct {
	Content string `json:"content"`
}

// ChecklistInput is the request body for creating a checklist.
type ChecklistInput struct {
	Name string `json:"name"`
}

// ChecklistItemInput is the request body for creating or updating a checklist item.
type ChecklistItemInput struct {
	Name      string `json:"name,omitempty"`
	Completed *bool  `json:"completed,omitempty"`
}

// --- Lists ---

// CreateListInput is the request body for creating a list.
type CreateListInput struct {
	BoardPublicId string `json:"boardPublicId"`
	Name          string `json:"name"`
}

// UpdateListInput is the request body for updating a list.
type UpdateListInput struct {
	Name string `json:"name"`
}

// --- Labels ---

// CreateLabelInput is the request body for creating a label.
type CreateLabelInput struct {
	WorkspacePublicId string `json:"workspacePublicId"`
	Name              string `json:"name"`
	ColourCode        string `json:"colourCode,omitempty"`
}

// UpdateLabelInput is the request body for updating a label.
type UpdateLabelInput struct {
	Name       string `json:"name,omitempty"`
	ColourCode string `json:"colourCode,omitempty"`
}

// --- Workspaces ---

// CreateWorkspaceInput is the request body for creating a workspace.
type CreateWorkspaceInput struct {
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Description string `json:"description,omitempty"`
}

// UpdateWorkspaceInput is the request body for updating a workspace.
type UpdateWorkspaceInput struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

// InviteMemberInput is the request body for inviting a member to a workspace.
type InviteMemberInput struct {
	Email string `json:"email"`
}

// --- Users ---

// UpdateUserInput is the request body for updating the current user.
type UpdateUserInput struct {
	Name  string `json:"name,omitempty"`
	Image string `json:"image,omitempty"`
}

// --- Attachments ---

// GeneratePresignedURLInput is the request body for generating a presigned S3 upload URL.
type GeneratePresignedURLInput struct {
	Filename    string `json:"filename"`
	ContentType string `json:"contentType"`
	Size        int64  `json:"size"`
}

// ConfirmAttachmentInput is the request body for confirming an attachment upload.
type ConfirmAttachmentInput struct {
	Key      string `json:"key"`
	Filename string `json:"filename"`
}
