package model

import "time"

// MSCommunication represents the ms_communications tbl
type MSCommunication struct {
	ID          int       `db:"id"`
	Category    string    `db:"category"`
	Subcategory string    `db:"subcategory"`
	Action      string    `db:"action"`
	TmplID      int       `db:"tmpl_id"`
	TmplName    string    `db:"tmpl_name"`
	TmplMessage string    `db:"tmpl_message"`
	IsActive    bool      `db:"is_active"`
	ComType     string    `db:"com_type"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

type JRequest struct {
	TemplateID int                    `json:"templateId"`
	Variables  map[string]interface{} `json:"variables"`
	Data       map[string]string      `json:"data"`
}
