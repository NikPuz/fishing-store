В сервисе дополнительно реализованы Graceful shutdown, Migrations, Middlewares, Configs, Logging

При разработке использовалось:
- Среда разработки GoLand IDEA
- Библиотеки pgx, chi, zap, viper, goose
- Система управления базами данных PostgreSQL
- Система контроля версий GitHub
- Система контейнеризации Docker

Сервис поддерживает запросы:
- CRUD /products
- CRUD /categories
- CRUD /manufacturers
- POST /sales - Совершить продажу
- GET /sales
- POST /supplies - Совершить поставку
- GET /supplies

Схема базы данных:
![image](https://github.com/NikPuz/fishing-store/assets/84337089/a9d38493-3a1d-439f-94de-caebc8b3a3d3)

