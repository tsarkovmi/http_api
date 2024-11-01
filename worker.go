package httpapi

/*
чтобы искало в FindWorkerById
валидируют наличие данных полей в теле запроса
являются реализацией фреймворка gin
*/

type Worker struct {
	ID         int     `json:"id" xml:"id" toml:"id" db:"id"`
	Name       string  `json:"name" xml:"name" toml:"name" binding:"required,max=50"`
	Age        int16   `json:"age" xml:"age" toml:"age" binding:"required,gte=18,lte=70"`
	Salary     float32 `json:"salary" xml:"salary" toml:"salary" binding:"required,gt=0"`
	Occupation string  `json:"occupation" xml:"occupation" toml:"occupation" binding:"required,max=100"`
}
