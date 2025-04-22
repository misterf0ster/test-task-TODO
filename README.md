📌 Задание
Описание проекта
Необходимо разработать REST API для управления задачами (TODO-лист).
API должно позволять:
📂 Требования к реализации
Создавать задачу.
Читать список задач.
Обновлять задачу.
Удалять задачу.

1. Используемый стек:
   Go + Fiber
   PostgreSQL (через pgx )
   Среда выполнения – локальная (Docker не обязателен)
2. Структура БД
   Создайте таблицу tasks с полями:
   id SERIAL PRIMARY KEY
   title TEXT NOT NULL
   description TEXT
   status TEXT CHECK (status IN ('new', 'in_progress', 'done')) DEFAULT 'new'
   created_at TIMESTAMP DEFAULT now()
   updated_at TIMESTAMP DEFAULT now()

3. Реализовать API-эндпоинты:

   ✅ POST /tasks – создание задачи.

   ✅ GET /tasks – получение списка всех задач.

   ✅ PUT /tasks/:id – обновление задачи.

   ✅ DELETE /tasks/:id – удаление задачи.

4. Дополнительные требования:
   Корректная обработка ошибок.
   Код должен быть читаемым и структурированным.
5. Будет плюсом:
   Использование Docker
   Наличие документации
   Наличие миграций
