# Test task

### Локальный запуск:
___

#### Запуск docker-compose
```shell
make start-docker
```
#### Запуск http-сервера шлюза (Docker должен быть запущен `make start-docker`)
```shell
make run-gateway
```
#### Запуск сервиса (Docker должен быть запущен `make start-docker`)
```shell
make run-service
```
#### Остановка Docker
```shell
make stop-docker
```
### Описание работы:
___
Репозиторий содержит два микросервиса, общающихся между собой через gRPC.

Сервис **api-gateway** является шлюзом, содержит в себе http-сервер для общения с клиентом. Сервис получает данные из микросервиса **custshopsvc** посредством общения через gRPC.

Сервис **custshopsvc** хранит и обрабатывает данные о покупателях и магазинах. Для хранения используется БД Postgres, возможно хранение данных в runtime. Сервис содержит в себе gRPC-сервер, который содержит эндпоинты для операций с данными. В качестве входного значения сервис получает оба типа записи.

Стандартный адрес работы http-сервера `localhost:8282`

---
### Запросы к записям о покупателях:
```
// POST /api/v1/customer/
// PATCH /api/v1/customer/{id}
// DELETE /api/v1/customer/{id}
// GET /api/v1/customer/?surname={surname}
// GET /api/v1/customer?surname={surname}&field={field}
// GET /api/v1/customer/{id}?field={field}
```
Запросы к записям о магазинах:
```
// POST /api/v1/shop/
// PATCH /api/v1/shop/{id}
// DELETE /api/v1/shop/{id}
// GET /api/v1/shop/?name={name}
// GET /api/v1/shop?name={name}&field={field}
// GET /api/v1/shop/{id}?field={field}
```
---
# REST API для записей о покупателях
## Создание
### Создать покупателя
#### Request

`POST /api/v1/customer/`

Содержит raw body с json покупателя
```
{
    "surname":"Familiya",
    "firstname":"Imya",
    "patronym":"Otchesvovich",
    "age":"55"
}
```

### Response
Ответ содержит в себе json с созданным id. Возвращает статус 201:

    HTTP/1.1 201 Created

    {"id": "1cb1b52a-fdd0-4e22-8f88-80e16012ec95"}


## Получение
### Получение по id
#### Request
`GET /api/v1/customer/{id}?field={field}`

Получение по id покупателя.

Опционально можно получить только одно поле из записи, указав название поля (название поля из json) query параметром field. 

#### Response

Возвращает статус 200 при успешно найденной записи и json с полученными данными.

    HTTP/1.1 200 OK
    Date: Fri, 13 Jan 2023 20:45:15 GMT
    Status: 200 OK
    Content-Type: application/json
    Content-Length: 117

    {"surname":"Familiya","firstname":"Imya","patronym":"Otovich","age":"55","date_created":"2023-01-13T23:44:42.11825Z"}

Возвращает статус 404 если запись не найдена и json с текстом ошибки.

    {"error": "Not Found: Record not found"}

### Получение по параметру (фамилия)
#### Request
`GET /api/v1/customer?surname={surname}&field={field}`

Поиск покупателя по параметру surname.

Опционально можно получить только одно поле из записи, указав название поля (название поля из json) query параметром field.

### Response

Возвращает статус 200 при успешно найденной записи и json с полученными данными.

    HTTP/1.1 200 OK
    Date: Fri, 13 Jan 2023 20:45:15 GMT
    Status: 200 OK
    Content-Type: application/json
    Content-Length: 46

    {"date_created": "2023-01-13T23:21:09.131116Z"}

Возвращает статус 404 если запись не найдена и json с текстом ошибки.

    {"error": "Not Found: Record not found"}

## Изменение
### Изменить покупателя
#### Request

`PATCH /api/v1/customer/`

Содержит raw body с json полями, которые необходимо изменить
```
{
    "surname":"Familiyev",
    "firstname":"Test"
}
```

### Response
Ответ возвращает статус 200 при успешном изменении. :

    HTTP/1.1 200 OK
    Date: Fri, 13 Jan 2023 20:23:42 GMT

Возвращает статус 404 если запись не найдена и json с текстом ошибки.

    {"error": "Not Found: Record not found"}

## Удаление
### Удаление по id
#### Request
`DELETE /api/v1/customer/{id}`

Удаление по id записи с покупателем.

#### Response

Возвращает статус 204 при успешно удаленной записи.

    HTTP/1.1 204 No Content
    Date: Fri, 13 Jan 2023 20:45:15 GMT

Возвращает статус 404 если запись не найдена и json с текстом ошибки.

    {"error": "Not Found: Record not found"}

# REST API для записей о магазинах
## Создание
### Создать магазин
#### Request

`POST /api/v1/shop/`

Содержит raw body с json магазина
```
{
    "name":"Final test",
    "address":"Add ress",
    "work_status":true,
    "owner":"Ownerov Ovner"
}
```

### Response
Ответ содержит в себе json с созданным id. Возвращает статус 201:

    HTTP/1.1 201 Created

    {"id": "f5b6a4ad-85b3-4823-8d33-47249645577b"}


## Получение
### Получение по id
#### Request
`GET /api/v1/shop/{id}?field={field}`

Получение по id записи о магазине.

Опционально можно получить только одно поле из записи, указав название поля (название поля из json) query параметром field.

#### Response

Возвращает статус 200 при успешно найденной записи и json с полученными данными.

    HTTP/1.1 200 OK
    Date: Fri, 13 Jan 2023 20:45:15 GMT
    Status: 200 OK
    Content-Type: application/json
    Content-Length: 117

    {"name":"Final test","address":"Add ress","work_status":true,"owner":"Ownerov Ovner"}

Возвращает статус 404 если запись не найдена и json с текстом ошибки.

    {"error": "Not Found: Record not found"}

### Получение по параметру (название магазина)
#### Request
`GET /api/v1/shop?name={name}&field={field}`

Поиск магазина по параметру name.

Опционально можно получить только одно поле из записи, указав название поля (название поля из json) query параметром field.

### Response

Возвращает статус 200 при успешно найденной записи и json с полученными данными.

    HTTP/1.1 200 OK
    Date: Fri, 13 Jan 2023 20:45:15 GMT
    Status: 200 OK
    Content-Type: application/json
    Content-Length: 46
    
    {"address":"Add ress 33"}

Возвращает статус 404 если запись не найдена и json с текстом ошибки.

    {"error": "Not Found: Record not found"}

## Изменение
### Изменить запись о магазине
#### Request

`PATCH /api/v1/shop/`

Содержит raw body с json полями, которые необходимо изменить
```
{
    "name":"Final test Update",
    "address":"Add ress 33"
}
```

### Response
Ответ возвращает статус 200 при успешном изменении. :

    HTTP/1.1 200 OK
    Date: Fri, 13 Jan 2023 20:23:42 GMT

Возвращает статус 404 если запись не найдена и json с текстом ошибки.

    {"error": "Not Found: Record not found"}

## Удаление
### Удаление по id
#### Request
`DELETE /api/v1/shop/{id}`

Удаление по id записи с магазином.

#### Response

Возвращает статус 204 при успешно удаленной записи.

    HTTP/1.1 204 No Content
    Date: Fri, 13 Jan 2023 20:45:15 GMT

Возвращает статус 404 если запись не найдена и json с текстом ошибки.

    {"error": "Not Found: Record not found"}


---

## ТЗ:

Сделать HTTP  сервер (без применения фреймворков), умеющий создавать, изменять и удалять записи двух типов:

- Информация о покупателе

    ◦ Фамилия
    
    ◦ Имя
    
    ◦ Отчество
    
    ◦ Возраст (необязательное поле)

    ◦ Дата регистрации


- Информацию о магазине

    ◦ Название

    ◦ Адрес

    ◦ Работающий или нет

    ◦ Владелец (необязательное поле)

Непосредственную работу с записями необходимо осуществлять с использованием одной функции, умеющей принимать в качестве входного значения оба типа записи.

Хранить и накапливать информацию можно по выбору: в рантайм, СУБД, файлах.

Входными и выходными параметрами в  HTTP  запросах являются данные в формате  JSON.

На выходе сервис должен уметь возвращать всю запись
, либо одно поле из записи в зависимости от запроса пользователя
, осуществлять поиск по Фамилии и Названию  для соответствующих записей.



Можно немного усложнить ее, разбив на два микросервиса:

Сделать два микросервиса работающих через  GRPC:
1. HTTP  сервер для общения с пользователем
2. Сервис хранения и обработки данных, в который ходит HTTP  сервер за данными.
Остальные вводные те же.