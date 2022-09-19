# http API

```sh
# dev
http://localhost:8080/
```

### Изменение стоимости участия мероприятия

```sh
Url: http://localhost:8080/api-v1/ad/changeParams
```js
{
  "body": {
    "city": string, <- Название города
    "price": int, <- Стоимость
  }
}
```sh
RETURN: 
  -status: 200 <- typeof int
  -message: "Успешное изменение стоимости участия в городе Екатеринбург" <- typeof string
```