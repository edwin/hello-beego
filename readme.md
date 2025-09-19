# hello-beego

**Hello World REST App using Beego**

A simple RESTful “Hello World” application built in Go using the Beego framework.  
Perfect as a starter project or template to explore Beego, REST APIs, or Go web development.

---

## Table of Contents

- [Features](#features)  
- [Prerequisites](#prerequisites)  
- [Installation](#installation)  
- [Configuration](#configuration)  
- [Running the App](#running-the-app)  
- [API Endpoints](#api-endpoints)  
- [Project Structure](#project-structure)  
- [License](#license)  
- [Contributing](#contributing)  
- [Contact](#contact)

---

## Features

- A minimal REST API built with Beego  
- Basic “Hello World” endpoint  
- Demonstrates project structure with configuration file and main application bootstrap  

---

## Prerequisites

Make sure you have:

- Go (≥ 1.16) installed  
- Git  

---

## Installation

1. Clone this repository

   ```bash
   git clone https://github.com/edwin/hello-beego.git
   cd hello-beego
   ```

2. Fetch dependencies

   ```bash
   go mod tidy
   ```

---

## Configuration

There is a `conf` directory (or configuration file) — you can adjust settings there (ports, app settings, etc.).  

Make sure any relevant environment variables or file paths are set up as needed.

---

## Running the App

To start the server:

```bash
go run main.go
```

The application by default listens on port (e.g. `:8080`) — check your `conf` file or code settings.  

---

## API Endpoints

Here are the sample endpoints provided:

| Method | Path | Description |
|--------|------|-------------|
| `GET /` | `/` | Returns a greeting message (Hello World) |

You can test using `curl`, Postman, or your browser. For example:

```bash
curl http://localhost:8080/
```

Expected response:

```json
{
  "message": "Hello World Beego",
  "status": "success"
}
```

---

## Project Structure

```
hello-beego/
├── conf/           # Configuration files
├── main.go         # Entry point
└── go.mod          # Go module file
```

---

## License

This project is open source. You can do anything you want with it.

---

## Contributing

Contributions, issues, and feature requests are welcome!  
Please open an issue first to discuss what you’d like to change.

---

## Contact

Maintainer: **edwin**  
GitHub: [edwin](https://github.com/edwin)  