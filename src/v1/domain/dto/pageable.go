package dto

type Page struct {
	Page int `form:"page" binding:"required,gte=1"`
	Size int `form:"size" binding:"required,gte=5,lte=50"`
}

type Paging struct {
	Page        int `json:"page,omitempty"`
	TotalPage   int `json:"total_page,omitempty"`
	TotalRecord int `json:"total_record,omitempty"`
}

func (p Page) GetOffSet() int {
	return (p.Page - 1) * p.Size
}
