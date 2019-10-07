package server

// APIResponse API.Response
type APIResponse struct {
	Page int	`json:"page"`
	Size int	`json:"size"`
	Count int	`json:"count"`
	Suffix string	`json:"suffix"`
	List []string	`json:"list"`
}