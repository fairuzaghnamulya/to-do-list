# ✅ Prabogo To-Do List

> A modern, minimalist to-do list application built with Go, featuring a beautiful dark/light theme toggle.

![Go](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat-square&logo=go&logoColor=white)
![Fiber](https://img.shields.io/badge/Fiber-v2-00ACD7?style=flat-square)
![Tailwind](https://img.shields.io/badge/Tailwind-CSS-38B2AC?style=flat-square&logo=tailwind-css&logoColor=white)
![Alpine.js](https://img.shields.io/badge/Alpine.js-3.x-8BC0D0?style=flat-square)

---

## ✨ Features

- 🌓 **Dark/Light Mode** - Seamless theme switching
- ⚡ **Fast & Lightweight** - Built with Go Fiber framework
- 💾 **Persistent Storage** - JSON-based database
- 📱 **Responsive Design** - Works on all devices
- 🎨 **Clean UI** - Minimalist design with smooth animations

---

## 🛠️ Tech Stack

| Layer         | Technology                      |
| ------------- | ------------------------------- |
| **Backend**   | Go + Fiber                      |
| **Frontend**  | HTML + Tailwind CSS + Alpine.js |
| **Database**  | JSON File Storage               |
| **Framework** | Prabogo Style Architecture      |

---

## 🚀 Getting Started

### Prerequisites

- Go 1.21 or higher

### Installation

```bash
# Clone the repository
git clone https://github.com/yourusername/prabogo-todo.git
cd prabogo-todo

# Install dependencies
go mod tidy

# Run the application
go run cmd/main.go
```

### Access

Open your browser and navigate to:

```
http://localhost:3000
```

---

## 📁 Project Structure

```
prabogo-todo/
├── cmd/
│   └── main.go           # Application entry point
├── database/
│   ├── sawit.go          # Database operations
│   └── todos.json        # Data storage
├── views/
│   └── index.html        # Frontend template
├── go.mod
├── go.sum
└── README.md
```

---

## 🎯 API Endpoints

| Method   | Endpoint         | Description        |
| -------- | ---------------- | ------------------ |
| `GET`    | `/`              | Serve frontend     |
| `GET`    | `/api/todos`     | Get all todos      |
| `POST`   | `/api/todos`     | Create new todo    |
| `PUT`    | `/api/todos/:id` | Toggle todo status |
| `DELETE` | `/api/todos/:id` | Delete todo        |

---

## 🎨 Screenshots

### Dark Mode

Clean dark interface that's easy on the eyes.

### Light Mode

Bright and minimal for daytime use.

---

## 📝 Usage

1. **Add Task** - Type your task and press Enter or click Add
2. **Complete Task** - Click on a task to mark it as done
3. **Delete Task** - Click the × button to remove
4. **Toggle Theme** - Click the sun/moon icon in the header

---

## 🤝 Contributing

Contributions are welcome! Feel free to submit a Pull Request.

---

## 👤 Author

**arcynith**

---

## 📄 License

This project is open source and available under the [MIT License](LICENSE).

---

<p align="center">
  Made with ❤️ using <b>Prabogo</b>, <b>JokoUI</b>, and <b>SawitDB</b>
</p>
