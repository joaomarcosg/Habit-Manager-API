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
HABIT_MANAGER_KEY=abcdefghijlmnopqrstuvwxyz
```

<h3>Starting</h3>

```bash
cd habit-manager-api
go run /cmd/api/main.go
```