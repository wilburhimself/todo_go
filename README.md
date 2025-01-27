# Go HTMX Todo Application

A modern, server-side rendered Todo application built with Go and HTMX. This application demonstrates how to create a dynamic, interactive web application without writing any JavaScript, using HTMX for seamless server interactions.

## Features

- Create new todos
- Toggle todo completion status
- Edit existing todos
- Delete todos
- Server-side rendering with Go templates
- Real-time updates using HTMX
- SQLite database using GORM

## Tech Stack

- **Backend**: Go
- **Frontend**: HTMX + HTML Templates
- **Database**: SQLite
- **ORM**: GORM
- **Template Engine**: Go's built-in html/template

## Project Structure

```
todo_go/
├── handlers/     # HTTP request handlers
├── lib/         # Library code (database initialization)
├── models/      # Data models
├── templates/   # HTML templates
└── main.go      # Application entry point
```

## Prerequisites

- Go 1.16 or higher
- SQLite

## Installation

1. Clone the repository:
```bash
git clone https://github.com/yourusername/todo_go.git
cd todo_go
```

2. Install dependencies:
```bash
go mod download
```

3. Run the application:
```bash
go run main.go
```

4. Open your browser and visit `http://localhost:8080`

## API Endpoints

- `GET /` - Display the main todo list
- `POST /add` - Add a new todo
- `POST /todos/{id}/toggle` - Toggle todo completion status
- `GET /todos/{id}/edit` - Get todo edit form
- `POST /todos/{id}/update` - Update todo
- `DELETE /todos/{id}/delete` - Delete todo

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.
