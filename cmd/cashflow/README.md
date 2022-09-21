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
  
POSSIBLE MISTAKES:
  -message: "некорректно переданы данные в body";
  -message: "ошибка изменения из базы данных, {err}";
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

POSSIBLE MISTAKES:
  -message: "некорректно переданы данные в body";
  -message: "ошибка создания из базы данных, {err}";
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
 
POSSIBLE MISTAKES:
  -message: "некорректно переданы данные в body";
  -message: "ошибка обновления из базы данных, {err}";
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
  -message: "ошибка обновления из базы данных, {err}";
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

POSSIBLE MISTAKES:
  -message: "некорректно переданы данные в body";
  -message: "ошибка обновления из базы данных, {err}";
```

### Вступление пользователя в мероприятие (пока без оплаты)

```sh
Url: http://localhost:8080/api-v1/user/entryToAd
```js
{
  "body": {
    "id_user": int, <- Идентификатор пользователя из таблицы users поле id
    "id_ad": int <- Идентификатор мероприятия из таблицы ad поле id
  }
}
```sh
RETURN: 
  -status: 200 <- typeof int
  -message: "успешное вступление в мероприятие" <- typeof string

POSSIBLE MISTAKES:
  -message: "некорректно переданы данные в body";
  -message: "ошибка получения из базы данных, {err}";
  -message: "ошибка преобразования полученных данных, {err}";
  -message: "данного мероприятия не существует";
  -message: "ошибка работы с json, {err}";
  -message: "вы уже зарегистрированы на данное мероприятие";
  -message: "вы не можете вступить в данное мероприятие, т.к оно полностью заполнено";
  -message: "ошибка обновления из базы данных, {err}";
```

### Распределение баллов по участникам за пройденное мероприятие

```sh
Url: http://localhost:8080/api-v1/ad/summarizingAd
```js
{
  "body": {
    "id_ad": int, <- Идентификатор мероприятия из таблицы ad поле id
    "winners_part": [ <- Массив объектов с именами столов и игроками за этими столами с их заработанными баллами
      {
        "name": string, <- Название стола (Стол 1)
        "winUser": [ <- Массив с игроками и их заработанное место в игре
          {"id": int, "place": int, "assigned": int},
        ]
      }
    ]
  }
}
```sh
RETURN: 
  -status: 200 <- typeof int
  -message: "успешное распределение баллов" <- typeof string

POSSIBLE MISTAKES:
  -message: "некорректно переданы данные в body";
  -message: "ошибка получения из базы данных, {err}";
  -message: "ошибка преобразования полученных данных, {err}";
  -message: "данного мероприятия не существует";
  -message: "ошибка работы с json, {err}";
```