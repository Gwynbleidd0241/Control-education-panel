-- Заполняем таблицу students из JSON-массива

WITH src AS (
    SELECT jsonb_array_elements($$
                                    [
                                    {
                                    "fullName": "Дмитрий Орлов",
                                "email": "dmitriy.orlov@gmail.com",
                                "age": 21,
                                "performance": "учится средне",
                                "course": {
      "id": "3",
                                "title": "Frontend-разработка с React",
                                "price": 99,
                                "level": "Начальный"
               },
    "photoUrl": "https://randomuser.me/api/portraits/men/10.jpg",
    "city": "Москва",
    "phone": "8 988-555-35-35",
    "format": "Очно",
    "id": "1",
    "progress": 53
  },
  {
    "id": "2",
    "fullName": "Мария Смирнова",
    "email": "maria.smirnova@yandex.ru",
    "age": 20,
    "photoUrl": "https://randomuser.me/api/portraits/women/10.jpg",
    "city": "Казань",
    "phone": "8 988-555-34-34",
    "format": "Онлайн",
    "course": {
      "id": "6",
      "title": "Машинное обучение с нуля",
      "price": 149,
      "level": "Средний"
    },
    "performance": "учится легко",
    "progress": 93
  },
  {
    "id": "5",
    "fullName": "Михаил Зайцев",
    "email": "mikhail.zaytsev@ya.ru",
    "age": 22,
    "photoUrl": "https://randomuser.me/api/portraits/men/30.jpg",
    "city": "Москва",
    "phone": "8 988-555-36-36",
    "format": "Очно",
    "course": {
      "id": "3",
      "title": "Frontend-разработка с React",
      "price": 99,
      "level": "Средний"
    },
    "performance": "учится средне",
    "progress": 62
  },
  {
    "id": "3",
    "fullName": "Наталья Кузнецова",
    "email": "natalya.kuznetsova@ya.ru",
    "age": 20,
    "photoUrl": "https://randomuser.me/api/portraits/women/18.jpg",
    "city": "Тула",
    "phone": "8 988-555-33-33",
    "format": "Онлайн",
    "course": {
      "id": "8",
      "title": "Базы данных и SQL",
      "price": 89,
      "level": "Средний"
    },
    "performance": "учится легко",
    "progress": 87
  },
  {
    "id": "4",
    "fullName": "Сергей Волков",
    "email": "sergey.volkov@gmail.com",
    "age": 24,
    "photoUrl": "https://randomuser.me/api/portraits/men/33.jpg",
    "city": "Астрахань",
    "phone": "8 988-555-34-34",
    "format": "Онлайн",
    "course": {
      "id": "2",
      "title": "Разработка веб-приложений на Django",
      "price": 119,
      "level": "Средний"
    },
    "performance": "учится средне",
    "progress": 43
  },
  {
    "id": "7",
    "fullName": "Алиса Панина",
    "email": "alisa.panina@ya.ru",
    "age": 21,
    "photoUrl": "https://randomuser.me/api/portraits/women/30.jpg",
    "city": "Дмитров",
    "phone": "8 988-555-30-30",
    "format": "Онлайн",
    "course": {
      "id": "2",
      "title": "Разработка веб-приложений на Django",
      "price": 119,
      "level": "Средний"
    },
    "performance": "учится тяжело",
    "progress": 11
  },
  {
    "id": "8",
    "fullName": "Владимир Сидоров",
    "email": "vladimir.sidorov@gmail.com",
    "age": 21,
    "photoUrl": "https://randomuser.me/api/portraits/men/34.jpg",
    "city": "Москва",
    "phone": "8 988-555-20-29",
    "format": "Очно",
    "course": {
      "id": "7",
      "title": "JavaScript для начинающих",
      "price": 79,
      "level": "Начальный"
    },
    "performance": "учится легко",
    "progress": 90
  },
  {
    "id": "21",
    "fullName": "Олег Сидоров",
    "email": "oleg.sidorov@gmail.com",
    "age": 22,
    "photoUrl": "https://randomuser.me/api/portraits/men/39.jpg",
    "city": "Москва",
    "phone": "8 988-555-20-19",
    "format": "Очно",
    "course": {
      "id": "14",
      "title": "Data Science: анализ данных на Python",
      "price": 120,
      "level": "Средний"
    },
    "performance": "учится средне",
    "progress": 48
  },
  {
    "id": "9",
    "fullName": "Татьяна Савельева",
    "email": "tatyana.saveleva@yandex.ru",
    "age": 20,
    "photoUrl": "https://randomuser.me/api/portraits/women/34.jpg",
    "city": "Москва",
    "phone": "8 988-555-20-28",
    "format": "Очно",
    "course": {
      "id": "10",
      "title": "Мобильная разработка на Flutter",
      "price": 109,
      "level": "Средний"
    },
    "performance": "учится средне",
    "progress": 50
  },
  {
    "id": "10",
    "fullName": "Людмила Сафонова",
    "email": "lyudmila.safonova@mail.ru",
    "age": 30,
    "photoUrl": "https://randomuser.me/api/portraits/women/24.jpg",
    "city": "Москва",
    "phone": "8 988-555-20-27",
    "format": "Очно",
    "course": {
      "id": "11",
      "title": "Кибербезопасность и защита данных",
      "price": 105,
      "level": "Средний"
    },
    "performance": "учится тяжело",
    "progress": 15
  },
  {
    "id": "11",
    "fullName": "Арина Захарова",
    "email": "arina.zakharova@gmail.com",
    "age": 20,
    "photoUrl": "https://randomuser.me/api/portraits/women/54.jpg",
    "city": "Москва",
    "phone": "8 988-555-20-20",
    "format": "Очно",
    "course": {
      "id": "13",
      "title": "Карьера в IT: как найти первую работу",
      "price": 60,
      "level": "Начальный"
    },
    "performance": "учится легко",
    "progress": 80
  },
  {
    "id": "12",
    "fullName": "Олег Олегов",
    "email": "vladimir.sidoov@gmail.com",
    "age": 21,
    "photoUrl": "https://randomuser.me/api/portraits/men/64.jpg",
    "city": "Томск",
    "phone": "8 988-555-22-22",
    "format": "Онлайн",
    "course": {
      "id": "4",
      "title": "Алгоритмы и структуры данных",
      "price": 89,
      "level": "Средний"
    },
    "performance": "учится легко",
    "progress": 75
  },
  {
    "id": "13",
    "fullName": "Владимир Топалов",
    "email": "vladimir.topalov@gmail.com",
    "age": 20,
    "photoUrl": "https://randomuser.me/api/portraits/men/63.jpg",
    "city": "Москва",
    "phone": "8 988-555-98-29",
    "format": "Очно",
    "course": {
      "id": "9",
      "title": "DevOps и CI/CD на практике",
      "price": 129,
      "level": "Средний"
    },
    "performance": "учится тяжело",
    "progress": 10
  },
  {
    "id": "14",
    "fullName": "Владимир Сидиков",
    "email": "vladimir.sidicov@gmail.com",
    "age": 28,
    "photoUrl": "https://randomuser.me/api/portraits/men/62.jpg",
    "city": "Москва",
    "phone": "8 988-555-99-29",
    "format": "Очно",
    "course": {
      "id": "10",
      "title": "Мобильная разработка на Flutter",
      "price": 109,
      "level": "Средний"
    },
    "performance": "учится легко",
    "progress": 95
  },
  {
    "id": "15",
    "fullName": "Никита Громов",
    "email": "nikita.gromov@ya.ru",
    "age": 21,
    "photoUrl": "https://randomuser.me/api/portraits/men/18.jpg",
    "city": "Москва",
    "phone": "8 988-555-90-29",
    "format": "Очно",
    "course": {
      "id": "14",
      "title": "Data Science: анализ данных на Python",
      "price": 120,
      "level": "Средний"
    },
    "performance": "учится средне",
    "progress": 50
  },
  {
    "id": "16",
    "fullName": "Виктория Белова",
    "email": "victoria.belova@ya.ru",
    "age": 21,
    "photoUrl": "https://randomuser.me/api/portraits/women/60.jpg",
    "city": "Москва",
    "phone": "8 988-555-70-29",
    "format": "Очно",
    "course": {
      "id": "15",
      "title": "Unity и разработка игр",
      "price": 100,
      "level": "Средний"
    },
    "performance": "учится легко",
    "progress": 86
  },
  {
    "id": "17",
    "fullName": "Маргарита Лебедева",
    "email": "margarita.lebedeva@mail.ru",
    "age": 31,
    "photoUrl": "https://randomuser.me/api/portraits/women/52.jpg",
    "city": "Краснодар",
    "phone": "8 988-555-60-29",
    "format": "Онлайн",
    "course": {
      "id": "12",
      "title": "Продвинутый курс по C++",
      "price": 115,
      "level": "Средний"
    },
    "performance": "учится средне",
    "progress": 72
  },
  {
    "id": "18",
    "fullName": "Максим Куликов",
    "email": "maksim.kulikov@yandex.ru",
    "age": 25,
    "photoUrl": "https://randomuser.me/api/portraits/men/53.jpg",
    "city": "Москва",
    "phone": "8 988-555-50-29",
    "format": "Очно",
    "course": {
      "id": "14",
      "title": "Data Science: анализ данных на Python",
      "price": 120,
      "level": "Средний"
    },
    "performance": "учится тяжело",
    "progress": 15
  },
  {
    "id": "19",
    "fullName": "Степан Ермаков",
    "email": "stepan.ermakov@ya.ru",
    "age": 23,
    "photoUrl": "https://randomuser.me/api/portraits/men/52.jpg",
    "city": "Москва",
    "phone": "8 988-555-19-29",
    "format": "Очно",
    "course": {
      "id": "15",
      "title": "Unity и разработка игр",
      "price": 100,
      "level": "Средний"
    },
    "performance": "учится средне",
    "progress": 60
  },
  {
    "id": "20",
    "fullName": "Злата Боброва",
    "email": "zlata.bobrova@gmail.com",
    "age": 18,
    "photoUrl": "https://randomuser.me/api/portraits/women/40.jpg",
    "city": "Москва",
    "phone": "8 988-555-10-29",
    "format": "Очно",
    "course": {
      "id": "9",
      "title": "DevOps и CI/CD на практике",
      "price": 118,
      "level": "Средний"
    },
    "performance": "учится тяжело",
    "progress": 20
  }
]
    $$::jsonb) AS x
)
INSERT INTO students (
    full_name,
    email,
    age,
    performance,
    photo_url,
    city,
    phone,
    study_format,
    progress
)
SELECT
    x->>'fullName',
    x->>'email',
    (x->>'age')::int,
    x->>'performance',
    x->>'photoUrl',
    x->>'city',
    x->>'phone',
    x->>'format',
    (x->>'progress')::int
FROM src;
