package vkapiclient

type VkApiErrorRequestParam struct {
	Key   string `json:\"key\"`
	Value string `json:\"value\"`
}

type VkApiError struct {
	Code          int    `json:\"error_code\"`
	Message       string `json:\"error_msg\"`
	Text          string `json:\"error_text\"`
	RequestParams []VkApiErrorRequestParam
}

type AuthConfirmResponseBody struct {
	// 1 if success
	success int `json:\"success\"`
	// User ID
	userId int `json:\"user_id\"`
}

type AuthConfirmResponse struct {
	Error    VkApiError              `json:\"error\"`
	Response AuthConfirmResponseBody `json:\"response\"`
}
