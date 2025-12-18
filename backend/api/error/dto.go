package error


type ErrorDTO struct {
	Type string `json:"type" binding:"required"`
	Title string `json:"title" binding:"required"`
	Status int	`json:"status" binding:"required"`
	Detail string `json:"detail" binding:"required"`
	Instance string `json:"instance"`

}

