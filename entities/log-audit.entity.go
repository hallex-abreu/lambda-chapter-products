package entities

type LogAuditEntity struct {
	Id        string `json:"id"`
	UserId    int    `json:"user_id"`
	Name      string `json:"name"`
	Target    string `json:"target"`
	Action    string `json:"action"`
	OldValue  string `json:"old_value"`
	NewValue  string `json:"new_value"`
	CreatedAt string `json:"created_at"`
}
