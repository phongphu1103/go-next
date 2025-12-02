package component

type ErrorResponse struct {
	Error      string      `json:"error"`
    Message    string      `json:"message"`
	Code	   int		   `json:"code"`
	StatusCode int		   `json:"status_code,omitempty"`
    Details    interface{} `json:"details,omitempty"`
}