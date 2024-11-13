package response

// ! Conflict from response.Paging with storage.Paging
type Paging struct {
	Page        int `json:"page,omitempty"`
	TotalPage   int `json:"total_page,omitempty"`
	TotalRecord int `json:"total_record,omitempty"`
}

type pagingResponse[T any] struct {
	Timestamp  int    `json:"timestamp,omitempty"`
	Code       int    `json:"code,omitempty"`
	StatusCode int    `json:"status_code,omitempty"`
	Message    string `json:"message,omitempty"`
	Payload    T      `json:"payload,omitempty"`
	Paging     Paging `json:"paging,omitempty"`
}

func NewPaging(page int, totalPage int, totalRecord int) Paging {
	return Paging{Page: page, TotalPage: totalPage, TotalRecord: totalRecord}
}

func NewPagingResponse[T any](timestamp int, code int, statusCode int, message string, payload T, paging Paging) pagingResponse[T] {
	return pagingResponse[T]{Timestamp: timestamp, Code: code, StatusCode: statusCode, Message: message, Payload: payload, Paging: paging}
}
