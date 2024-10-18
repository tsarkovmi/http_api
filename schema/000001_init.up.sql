CREATE TABLE workers (
    id SERIAL PRIMARY KEY,        -- Уникальный идентификатор, автоматически увеличивающийся
    name VARCHAR(100) NOT NULL,   -- Имя и фамилия, строка до 100 символов
    age INT NOT NULL,             -- Возраст, целое число
    salary REAL NOT NULL,         -- Зарплата, тип данных с плавающей точкой (float32)
    occupation VARCHAR(100) NOT NULL    -- Должность, строка до 100 символов, допускается пробел
);
