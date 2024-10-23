package httpapi

type Worker struct {
	ID         int     `json:"id" xml:"id" toml:"id" db:"id"`                  //чтобы искало в FindWorkerById
	Name       string  `json:"name" xml:"name" toml:"name" binding:"required"` //Валидируют наличие данных полей в теле запроса
	Age        int16   `json:"age" xml:"age" toml:"age" binding:"required"`    //являются реализацией фреймворка gin
	Salary     float32 `json:"salary" xml:"salary" toml:"salary" binding:"required"`
	Occupation string  `json:"occupation" xml:"occupation" toml:"occupation" binding:"required"`
}
