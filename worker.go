package httpapi

type Worker struct {
	ID         string  `json:"id"`
	Name       string  `json:"name"`
	Age        int16   `json:"age"`
	Salary     float32 `json:"salary"`
	Occupation string  `json:"occupation"`
}
