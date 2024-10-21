package httpapi

type Worker struct {
	ID         string  `json:"id"`
	Name       string  `json:"name" binding:"required"`
	Age        int16   `json:"age" binding:"required"`
	Salary     float32 `json:"salary" binding:"required"`
	Occupation string  `json:"occupation" binding:"required"`
}
