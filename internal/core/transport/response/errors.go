package core_http_respose

type ErorrResponse struct {
	Error   string `json:"error"   example:"full error text"`
	Message string `json:"message" example:"short human-readable message"`
}
