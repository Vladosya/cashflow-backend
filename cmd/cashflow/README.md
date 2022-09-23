# http API

```sh
# dev
http://localhost:8080/
```

### Зарегистрировать пользователя на мероприятие бесплатно через службу поддержки (POST)

```sh
Url: http://localhost:8080/api-v1/support/entryToAdSupport
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
  -message: "успешное вступление пользователя в мероприятие" <- typeof string

POSSIBLE MISTAKES:
  -message: "некорректно переданы данные в body";
  -message: "ошибка получения из базы данных, {err}";
  -message: "ошибка преобразования полученных данных, {err}";
  -message: "данного мероприятия не существует";
  -message: "ошибка работы с json, {err}";
  -message: "данный пользователь уже зарегистрирован на данное мероприятие";
  -message: "данный пользователь не можете вступить в данное мероприятие, т.к оно полностью заполнено";
  -message: "ошибка обновления из базы данных, {err}";
```

### Создание мероприятия (POST)

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
  -message: "ошибка создания из базы данных и получение в Scan, {err}";
  -message: "ошибка создания из базы данных, {err}";
```

### Вступление пользователя в мероприятие (пока без оплаты) (POST)

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

### Изменение стоимости участия мероприятия (POST)

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

### Активация мероприятия (сделать видимым мероприятие для всех) (POST)

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

### Завершение мероприятия (если мероприятие уже состоялось) (POST)

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

### Отмена мероприятия (если мероприятие по каким-то причинам было отменено) (POST)

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

### Распределение баллов по участникам за пройденное мероприятие (POST)

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

### Пересаживание пользователей из одного стола в другой (POST)

```sh
Url: http://localhost:8080/api-v1/ad/replantAd
```js
{
  "body": {
    "id_ad": int, <- Идентификатор мероприятия из таблицы ad поле id или таблицы game поле id
    "seat_at_tables": [ <- Массив объектов с именами столов и игроками за этими столами
      {
        "id": int, <- идентификатор,
        "name": string, <- Название стола (Стол 1)
        "seat_at_table": [ <- Массив с игроками и их место за данным столом
          {"id": int, "name": string, "rank": int, "id_user": int},
        ]
      }
    ]
  }
}
```sh
RETURN:  
  -status: 200 <- typeof int
  -message: "успешное распределение пользователей по столам" <- typeof string

POSSIBLE MISTAKES:
  -message: "некорректно переданы данные в body";
  -message: "ошибка при кодировки данных в JSON, {err}";
  -message: "ошибка обновления из базы данных, {err}";
```

### Получить данные о том, кто за каким столом сидит и т.д (POST)

```sh
Url: http://localhost:8080/api-v1/ad/getInfoAbTables
```js
{
  "body": {
    "id_ad" int <- Идентификатор мероприятия из таблицы ad поле id или таблицы game поле id
  }
}
```sh
RETURN: 
  -status: 200 <- typeof int
  -message: "успешное получение информацию по столам (места игроков) в игре" <- typeof string
  -result: [] <- []appl.GameForm

POSSIBLE MISTAKES:
  -message: "некорректно переданы данные в body";
  -message: "ошибка получения из базы данных, {err}";
  -message: "ошибка преобразования полученных данных, {err}";
``` 

### Получение всех мероприятий с интервалом от сегодняшней даты + 30 дней (GET)

```sh
Url: http://localhost:8080/api-v1/ad/getAllAd
```js
{
  "body": {}
}
```sh
RETURN: 
  -status: 200 <- typeof int
  -message: "успешное получение всех мероприятий" <- typeof string
  -result: [] <- []appl.AdFull

POSSIBLE MISTAKES:
  -message: "некорректно переданы данные в body";
  -message: "ошибка получения из базы данных, {err}";
  -message: "ошибка преобразования полученных данных, {err}";
``` 