# Multi-functional Search Engine

Multi-functional Search Engine — легковесный и высокопроизводительный поисковой движок, написанный на Go с фронтендом на Vue.js. Он поддерживает индексирование и поиск по различным форматам текстовых документов благодаря реализации предварительной обработки текста (токенизация, нормализация, удаление стоп-слов, стемминг), построению инвертированного индекса, сохранённого в MongoDB, и ранжированию результатов с помощью TF-IDF и cosine similarity. Бэкенд использует goroutines языка Go для параллельной обработки, а приложение можно развернуть через Docker и Nginx.

## Основной функционал

* **Text Preprocessing**: токенизация, нормализация, удаление стоп-слов и стемминг для обоих языков.
* **Inverted Index**: построение и хранение инвертированного индекса, связывающего токены с идентификаторами документов для быстрого поиска.
* **Ranking**: вычисление TF-IDF весов и ранжирование документов по cosine similarity.
* **Concurrency**: использование goroutines для параллельного парсинга файлов, индексирования и обработки запросов.
* **Frontend UI**: интерфейс на Vue.js для загрузки документов, выполнения поисковых запросов и отображения результатов.
* **Testing**: unit tests для основных пакетов (Parser, Indexer, Search).
* **Deployment**: контейнеризация с помощью Docker и Docker Compose, деплой через Nginx.

## Технологический стек

* **Backend**: Go (Golang)
* **Frontend**: Vue.js
* **Database**: MongoDB
* **Containerization**: Docker, Docker Compose
* **Web Server**: Nginx

## Начало работы

### Требования

* Установленный Go 1.18+
* Node.js и npm
* Запущенный экземпляр MongoDB
* Docker и Docker Compose (опционально, для контейнеризированного развёртывания)

### Установка

1. **Клонирование репозитория**:

   ```bash
   git clone https://github.com/FoggGhostt/Multi-functional-Search-Engine.git
   cd Multi-functional-Search-Engine
   ```
2. **Настройка**:

   * Отредактируйте `config.yaml`, указав ваш MongoDB URI и имя базы данных.
3. **Сборка бэкенда**:

   ```bash
   go build -o search-engine main.go
   ```
4. **Установка зависимостей фронтенда**:

   ```bash
   cd frontend
   npm install
   ```

## Использование

1. **Запустите MongoDB** (если ещё не запущен).
2. **Запуск бэкенда**:

   ```bash
   ./search-engine
   ```
3. **Запуск фронтенда**:

   ```bash
   cd frontend
   npm run serve
   ```
4. **Откройте** [http://localhost:8080](http://localhost:8080) в браузере для доступа к UI.

## Запуск тестов

Для запуска unit tests основных пакетов выполните:

```bash
cd pkg/parser && go test
cd pkg/indexer && go test
cd pkg/search && go test
```

## Развёртывание в Docker

1. **Сборка и запуск контейнеров**:

   ```bash
   docker-compose up --build
   ```
2. Фронтенд будет обслуживаться через Nginx, который также проксирует API-запросы к Go-бэкенду.

## Структура проекта

```plaintext
.
├── Dockerfile
├── docker-compose.yml
├── config.yaml                  # MongoDB and other settings
├── main.go                      # Application entrypoint (Go)
├── pkg                          # Go packages
│   ├── API                      # HTTP handlers
│   ├── parser                   # File reading & tokenization
│   ├── indexer                  # Inverted index builder
│   ├── search                   # TF-IDF ranking & query logic
│   ├── models                   # Data models
│   ├── mongodb                  # MongoDB connection & operations
│   └── config                   # Configuration loader
└── frontend                     # Vue.js application
    ├── public
    └── src
        ├── components           # UI components (SearchBar, FileUpload, Results)
        └── App.vue, main.js     # App entrypoint
```
