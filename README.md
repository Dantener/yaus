# Yet Another URL Shortener

Очередной сервис для сокращения URL'ов

## Установка и запуск
```
make build
make run
```

## Swagger
### `localhost:8080/swagger/index.html#/`

## Методы
### `POST localhost:8080/api/v1/short/url`:
Метод для сокращения полного URL в короткий
```JSON
{
  "originURL": "https://www.google.com/search?q=golang&oq=golang&aqs=chrome..69i57j69i59l3j69i60l4.1983j0j7&sourceid=chrome&ie=UTF-8"
}
```
#### Ответ:
```JSON
{
  "shortURL": "http://localhost:8080/4DuNiJ"
}
```

### `GET localhost:8080/api/v1/short/url?shortURL=http://localhost:8080/4DuNiJ`:
Метод для преобразования короткого URL в полный
#### Result:
```html
<a href="https://www.google.com/search?q=golang&amp;oq=golang&amp;aqs=chrome..69i57j69i59l3j69i60l4.1983j0j7&amp;sourceid=chrome&amp;ie=UTF-8">Moved
    Permanently</a>.
```