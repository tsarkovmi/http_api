package httpapi

type Worker struct {
	ID         int     `json:"id" db:"id"`              //чтобы искало в FindWorkerById
	Name       string  `json:"name" binding:"required"` //Валидируют наличие данных полей в теле запроса
	Age        int16   `json:"age" binding:"required"`  //являются реализацией фреймворка gin
	Salary     float32 `json:"salary" binding:"required"`
	Occupation string  `json:"occupation" binding:"required"`
}
