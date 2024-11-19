# Задача: написать сервис, который предоставляет http api, у api есть 3 handlera.
-  В первый приходит json запрос с полями name, age, salary,         
     occupation, запись заносится в postgres, где ей присваивается уникальный id, 
     в response лежит id
- Во второй приходит id, возвращается 3                              
     string поля - json, xml, toml, 
     который содержат сериализованную запись с данным id
- В третий приходит пустой request возвращается                      
     список сериализованных записей(как во втором handlere)

### Требования: 
- Dockerfile, все должно подниматься с помощью docker compose
- Конфигурация порта, на котором будет запускаться сервис, 
  конфигурация доступа к postgres(address, user, password, database)
-  Сериализация в разные форматы должна происходить одновременно
-  Должно осуществляться логирование
-  Наличие swagger с описанием api.
##### Будет плюсом (в твоём случае отсутствия ограничения времени это must have):
- Наличие тестов 
- Graceful shutdown  




## Установка и запуск: 
```
make build && make run
```
### Миграции: 
```
make migrate
```

# Поднятие сервиса с помощью docker-compose, без make
```
docker-compose up --build httpapi
```
#### Пример POST
```
{
    "name": "Georgiy Zavozin", 
    "age": 23, 
    "salary": 14523.56, 
    "occupation": "Грузчик"
}
```
Через Postman - норм


#### Пример GET id   /workers/id
```
     http://localhost:8080/workers/2

curl http://localhost:8080/workers/2
```
#### Пример GET  /workers
```
     http://localhost:8080/workers

curl http://localhost:8080/workers
```

# Дополнительные команды
Установка и запуск образа postgres в docker
```
docker pull postgres
docker run --name=worker-db -e POSTGRES_PASSWORD=qwerty -p 5436:5432 -d --rm postgres
```
Просмотр существующих образов
```
docker ps
```
Скачать и установить Make и Migrate можно с помощью scoop https://scoop.sh/ Команда в ps, что то из этих двух должно работать
```
Set-ExecutionPolicy -ExecutionPolicy RemoteSigned -Scope CurrentUser
Invoke-RestMethod -Uri https://get.scoop.sh | Invoke-Expression


irm get.scoop.sh -outfile 'install.ps1'
.\install.ps1 -RunAsAdmin [-OtherParameters ...]
# I don't care about other parameters and want a one-line command
iex "& {$(irm get.scoop.sh)} -RunAsAdmin"
```
Пример использования scoop
```
scoop install migrate
scoop install make
```
Создание файлов миграций
```
migrate create -ext sql -dir ./schema -seq init
```
Сама миграция, для этого должны быть сами файлы миграции, сгенерированные командой выше
```
migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5436/postgres?sslmode=disable' up
```
Подключение к контейнеру по id, который можно получить в docker ps
```
docker exec -it 9a063294aa50 /bin/bash
```
Посмотреть что внутри БД и вызов списка relation
```
psql -U postgres
\d
```


Для тестирования используется следующая библиотека 
https://github.com/uber-go/mock

Для тестирования репозитория использовалась 
https://github.com/zhashkevych/go-sqlxmock

Также Тестифай, для ассертов, итого 16 тестов
