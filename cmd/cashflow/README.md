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
  -message: "Успешное изменение стоимости участия в городе {city}" <- typeof string
```

### Создание мероприятия

```sh
Url: http://localhost:8080/api-v1/ad/createAd
```js
{
  "body": {
    "title": string, <- Название
    "date_start": TIMESTAMP, <- Дата начала (yyyy-MM-dd hh:mm:ss)
    "city": string, <- Город
    "price": int, <- Цена
    "description": string, <- Описание
    "event_type": string, <- Тип объявления ('открытая', 'закрытая')
    "serial_number": int, <- Порядковый номер объявления
    "points_options": int, <- Параметры зачислений баллов, которые будут использоваться в мероприятии
    "is_visible": bool <- Если true, то объявление активное и оно становится видимым для всех пользователей. Если false, то оно является как черновик и не виден никому, кроме администратора
    "is_finished": bool <- Если true, то мероприятие считается законченным
  }
}
```sh
RETURN: 
  -status: 200 <- typeof int
  -message: "Успешное создание мероприятия" <- typeof string
```

### Активация мероприятия (сделать видимым мероприятие для всех)

```sh
Url: http://localhost:8080/api-v1/ad/activateAd
```js
{
  "body": {
    "id": int, <- Идентификатор мероприятия из таблицы ad поле id
  }
}
```sh
RETURN: 
  -status: 200 <- typeof int
  -message: "Успешная активация мероприятия" <- typeof string
```

### Завершение мероприятия (если мероприятие уже состоялось)

```sh
Url: http://localhost:8080/api-v1/ad/toCompleteAd
```js
{
  "body": {
    "id": int, <- Идентификатор мероприятия из таблицы ad поле id
  }
}
```sh
RETURN: 
  -status: 200 <- typeof int
  -message: "Успешное завершение мероприятия" <- typeof string

POSSIBLE MISTAKES:
  -message: "некорректно переданы данные в body";
  -message: "ошибка получения из базы данных";
  -message: "ошибка преобразования полученных данных, {err}";
  -message: "вы уже зарегистрированы на данное мероприятие";
  -message: "ошибка изменения из базы данных, {err}";
```

### Отмена мероприятия (если мероприятие по каким-то причинам было отменено)

```sh
Url: http://localhost:8080/api-v1/ad/cancelAd
```js
{
  "body": {
    "id": int, <- Идентификатор мероприятия из таблицы ad поле id
  }
}
```sh
RETURN: 
  -status: 200 <- typeof int
  -message: "Успешная отмена мероприятия" <- typeof string
```