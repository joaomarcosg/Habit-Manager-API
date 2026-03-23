<h1 align="center" style="font-weight:bold;">Habit Manager API</h1>

<p align="center">
 <a href="#tech">Technologies</a> • 
 <a href="#started">Getting Started</a> • 
  <a href="#routes">API Endpoints</a>
</p>

<p align="center">
        The Habit Manager API is an application for controlling and managing personal habits.
        The entire API was developed in Go and with the help of the Chi framework to accelerate project development.
</p>

<h2 id="technologies">💻 Technologies</h2>

| Technology | Description |
| ---------- | ----------- |
| Go | Statically typed programming language |
| Chi | Go framework for creating HTTP servers |
| Postgres | Relational database |
| Docker | Software platform for deploying containerized applications |
| Gorilla Websocket | Library for implementing real-time communication |
| SCS - Session Manager | Session-based authentication |

<h2 id="started">🚀 Getting started</h2>

<h3>Prerequisites</h3>

- Go 1.20+

<h3>clone the repository</h3>

```bash
git clone https://github.com/joaomarcosg/Habit-Manager-API.git
```

<h3>Install the dependencies.</h3>

```bash
go mod tidy
```

<h3>Config .env variables</h2>

Use the `.env.example` as reference to create your configuration file `.env`

```yaml
HABIT_MANAGER_DATABASE_PORT=5432
HABIT_MANAGER_DATABASE_NAME=habit
HABIT_MANAGER_DATABASE_USER=postgres
HABIT_MANAGER_DATABASE_PASSWORD=123
HABIT_MANAGER_DATABASE_HOST=localhost
HABIT_MANAGER_KEY=abcdefghijlmnopqrstuvwxyz1234567
```

📌 **The environment variable ```HABIT_MANAGER_KEY``` is a 32 bits key. Use [random.org](https://www.random.org/) to generate a string.**

<h3>Starting</h3>

```bash
cd habit-manager-api
go run /cmd/api/main.go
```

<h2 id="routes">📍 API Endpoints</h2>

| Route | Description |
| ----- | ----------- |
| <kbd>GET /api/v1/csrftoken</kbd> | Get authentication token [response details](#get-auth-detail) |
| <kbd>POST /api/v1/users/signupuser | User registration [request details](#post-signup-user) |
| <kbd>POST /api/v1/users/loginuser | User login [request details](#post-login-user)  |
| <kbd>POST /api/v1/users/logout | User logout [response details](#post-logout-user)  |
| <kbd>POST /api/v1/categories/ | Create a category for the habit [response details](#post-logout-user)  |
| <kbd>POST /api/v1/categories/getCategory | Get a category by the name [response details](#post-create-category)  |
| <kbd>GET /api/v1/categories/getCategoryEnties | Get category entries by the name [response details](#post-category-entries)  |
| <kbd>POST /api/v1/categories/deleteCategory | delete a category by the name [response details](#post-delete-category)  |
| <kbd>POST /api/v1/habits/ | Create a habit [response details](#post-create-habit)  |