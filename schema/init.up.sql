-- Вся информация о стоимости мероприятия в зависимости от выбранного города
CREATE TABLE IF NOT EXISTS ad_params(
    id SERIAL PRIMARY KEY NOT NULL,
    price INTEGER NOT NULL,
    city VARCHAR(255) NOT NULL
);

INSERT INTO ad_params(price, city) VALUES(5000, 'Москва');
INSERT INTO ad_params(price, city) VALUES(5000, 'Екатеринбург');

--  Вся информация о абонементах
CREATE TABLE IF NOT EXISTS subscriptions(
    id SERIAL PRIMARY KEY NOT NULL,
    name VARCHAR(255) NOT NULL,
    price INTEGER NOT NULL,
    discont INTEGER,
    by_days BOOLEAN NOT NULL DEFAULT 'f',
    by_visits BOOLEAN NOT NULL DEFAULT 'f'
);

INSERT INTO subscriptions(name, price, discont, by_days, by_visits) VALUES('buy_2_month', 50000, 62, true, false);
INSERT INTO subscriptions(name, price, discont, by_days, by_visits) VALUES('buy_3_month', 100000, 92, true, false);
INSERT INTO subscriptions(name, price, discont, by_days, by_visits) VALUES('3_games_discont', 0, 3, false, true);

-- Зачисление баллов после игры.
CREATE TABLE IF NOT EXISTS points_game(
    id SERIAL PRIMARY KEY NOT NULL,
    title VARCHAR(255) NOT NULL,
    city VARCHAR(255) NOT NULL,
    version INTEGER NOT NULL,
    scoring jsonb
);

INSERT INTO points_game(title, city, version, scoring) VALUES('Зачисление баллов после игры для Екатеринбурга', 'Екатеринбург', 1, '[{"id": "1", "place": 1, "numberPoints": 9}, {"id": "2", "place": 2, "numberPoints": 7}, {"id": "3", "place": 3, "numberPoints": 6}, {"id": "4", "place": 4, "numberPoints": 5}, {"id": "5", "place": 5, "numberPoints": 4}, {"id": "6", "place": 6, "numberPoints": 3}, {"id": "7", "place": 7, "numberPoints": 1}]');
INSERT INTO points_game(title, city, version, scoring) VALUES('Зачисление баллов после игры для Екатеринбурга', 'Екатеринбург', 2, '[{"id": "1", "place": 1, "numberPoints": 11}, {"id": "2", "place": 2, "numberPoints": 7}, {"id": "3", "place": 3, "numberPoints": 6}, {"id": "4", "place": 4, "numberPoints": 5}, {"id": "5", "place": 5, "numberPoints": 4}, {"id": "6", "place": 6, "numberPoints": 3}, {"id": "7", "place": 7, "numberPoints": 1}]');
INSERT INTO points_game(title, city, version, scoring) VALUES('Зачисление баллов после игры для Екатеринбурга', 'Москва', 1, '[{"id": "1", "place": 1, "numberPoints": 10}, {"id": "2", "place": 2, "numberPoints": 8}, {"id": "3", "place": 3, "numberPoints": 6}, {"id": "4", "place": 4, "numberPoints": 5}, {"id": "5", "place": 5, "numberPoints": 4}, {"id": "6", "place": 6, "numberPoints": 3}, {"id": "7", "place": 7, "numberPoints": 1}]');
INSERT INTO points_game(title, city, version, scoring) VALUES('Зачисление баллов после игры для Сочи', 'Сочи', 1, '[{"id": "1", "place": 1, "numberPoints": 9}, {"id": "2", "place": 2, "numberPoints": 8}, {"id": "3", "place": 3, "numberPoints": 6}, {"id": "4", "place": 4, "numberPoints": 5}, {"id": "5", "place": 5, "numberPoints": 4}, {"id": "6", "place": 6, "numberPoints": 3}, {"id": "7", "place": 7, "numberPoints": 1}]');

-- Зачисление баллов после игры.
CREATE TABLE IF NOT EXISTS ad(
    id SERIAL PRIMARY KEY NOT NULL,
    title VARCHAR(255) NOT NULL,
    date_start TIMESTAMP NOT NULL,
    created TIMESTAMP NOT NULL DEFAULT now(),
    city VARCHAR(255) NOT NULL,
    price INTEGER,
    description VARCHAR(255) NOT NULL,
    event_type VARCHAR(255) NOT NULL,
    participant INTEGER[],
    serial_number INTEGER NOT NULL,
    points_options INTEGER,
    is_visible BOOLEAN NOT NULL DEFAULT 'f',
    FOREIGN KEY (points_options) REFERENCES points_game (id),
    FOREIGN KEY (price) REFERENCES ad_params (id)
);

-- Таблица с распределёнными баллами участникам за пройденное мероприятие.
CREATE TABLE IF NOT EXISTS distribution_points(
    id SERIAL PRIMARY KEY NOT NULL,
    id_ad INTEGER,
    winners_part jsonb,
    isaccrued_points BOOLEAN NOT NULL DEFAULT 'f'
);

-- Вся информация о городе с его доступными абонементами и начислениями бонусов.
CREATE TABLE IF NOT EXISTS cities(
    id SERIAL PRIMARY KEY NOT NULL,
    name VARCHAR(255) NOT NULL,
    hide BOOLEAN NOT NULL DEFAULT 'f',
    allow_subscrip INTEGER[],
    allow_points INTEGER[]
);

INSERT INTO cities(name, hide, allow_subscrip, allow_points) VALUES('Москва', false, '{1, 3}', '{1, q2}');
INSERT INTO cities(name, hide, allow_subscrip, allow_points) VALUES('Екатеринбург', false, '{1, 2, 3}', '{1, 2}');
INSERT INTO cities(name, hide, allow_subscrip, allow_points) VALUES('Сочи', false, '{1, 2}', '{4}');

-- Пользователь Будет создание, изменение, отправка в бан пользователя
CREATE TABLE IF NOT EXISTS users(
    id SERIAL PRIMARY KEY NOT NULL,
    uid VARCHAR(255) NOT NULL UNIQUE,
    tele_id INTEGER NOT NULL UNIQUE,
    tele_name VARCHAR(255) NOT NULL,
    tele_family VARCHAR(255) NOT NULL,
    parent_id INTEGER NOT NULL,
    ok BOOLEAN NOT NULL DEFAULT 'f',
    utype VARCHAR(255) NOT NULL,
    date_delete TIMESTAMP,
    admin_comment VARCHAR(255) NOT NULL,
    created TIMESTAMP NOT NULL DEFAULT now(),
    lang VARCHAR NOT NULL DEFAULT 'ru',
    city INTEGER,
    register_form jsonb,
    is_delete_bot BOOLEAN NOT NULL DEFAULT 'f',
    is_ban BOOLEAN NOT NULL DEFAULT 'f',
    FOREIGN KEY (city) REFERENCES cities (id)
);

-- Состояние счёта пользователя
CREATE TABLE IF NOT EXISTS users_base(
    id INTEGER,
    balance INTEGER NOT NULL DEFAULT 0,
    downb INTEGER NOT NULL DEFAULT 0,
    passb INTEGER NOT NULL DEFAULT 0,
    waitb INTEGER NOT NULL DEFAULT 0,
    refund INTEGER NOT NULL DEFAULT 0,
    cash INTEGER NOT NULL DEFAULT 0,
    FOREIGN KEY (id) REFERENCES users (id)
);

-- Абонемент, подписка. Например: на 2-3 месяца. Абонемент для новичков на 3 игры и т.д
CREATE TABLE IF NOT EXISTS users_subscriptions(
    id INTEGER,
    user_id INTEGER,
    name VARCHAR(255) NOT NULL,
    price INTEGER NOT NULL,
    discont INTEGER NOT NULL,
    current INTEGER NOT NULL DEFAULT 0,
    city VARCHAR(255) NOT NULL,
    date_start TIMESTAMP NOT NULL DEFAULT now(),
    date_expiration TIMESTAMP,
    FOREIGN KEY (id) REFERENCES subscriptions (id),
    FOREIGN KEY (user_id) REFERENCES users (id)
);

-- Вся информация по статистике пользователя в каждом из городов
CREATE TABLE IF NOT EXISTS users_game_statistic(
    id INTEGER,
    name_city VARCHAR(255) NOT NULL,
    total_games INTEGER NOT NULL,
    count_victories INTEGER NOT NULL,
    ranking_year INTEGER NOT NULL,
    average_points INTEGER NOT NULL,
    rank INTEGER NOT NULL,
    count_points INTEGER NOT NULL,
    FOREIGN KEY (id) REFERENCES users (id)
);

-- Будут отображаться все события в приложении. Напрмер: покупка абонемента, вступление в мероприятие, бан пользователя и т.д
CREATE TABLE IF NOT EXISTS events(
    id SERIAL PRIMARY KEY NOT NULL,
    user_id INTEGER,
    ev_type VARCHAR(255) NOT NULL,
    user_data jsonb,
    created TIMESTAMP NOT NULL DEFAULT now(),
    FOREIGN KEY (user_id) REFERENCES users (id)
);

-- Где будут находиться пользователи которые были забанены
CREATE TABLE IF NOT EXISTS users_ban(
    id SERIAL PRIMARY KEY NOT NULL,
    user_id INTEGER,
    tele_id INTEGER,
    uid VARCHAR(255),
    tele_name VARCHAR(255) NOT NULL,
    tele_family VARCHAR(255) NOT NULL,
    admin_comment VARCHAR(255) NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users (id),
    FOREIGN KEY (tele_id) REFERENCES users (tele_id),
    FOREIGN KEY (uid) REFERENCES users (uid)
);

-- Таблица с платежами пользователя
CREATE TABLE IF NOT EXISTS users_pay(
    id SERIAL PRIMARY KEY NOT NULL,
    user_id INTEGER,
    tele_id INTEGER NOT NULL,
    type VARCHAR(255) NOT NULL,
    url VARCHAR(255) NOT NULL,
    amount INTEGER NOT NULL,
    success BOOLEAN NOT NULL DEFAULT 'f',
    date_success TIMESTAMP,
    balance BOOLEAN NOT NULL DEFAULT 'f',
    cancel BOOLEAN NOT NULL DEFAULT 'f',
    err VARCHAR(255),
    created TIMESTAMP NOT NULL DEFAULT now(),
    cmmnt VARCHAR(255),
    date_cancel TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users (id)
);