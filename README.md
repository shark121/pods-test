* Frontend → **HTML + JavaScript**
* Backend → **Go (Golang)**
* Purpose → Groups users into “pods” based on similar origins/destinations

Here’s a polished draft:

---

# Pods 🚗

*A rideshare app that groups users into shared rides ("pods") based on similar origins and destinations.*

## 📌 Overview

Pods is a smart rideshare platform that helps users share rides efficiently. Instead of random pairings, users are grouped into “pods” based on how close their **origins** and **destinations** are. This makes rides more:

* Affordable 💰
* Eco-friendly 🌍
* Convenient 🚀

The project uses a **simple HTML/JavaScript frontend** and a **Go backend** that powers the matching logic.

## ✨ Features

* Request a ride by entering origin and destination
* Grouping algorithm to match riders into pods
* REST API backend in Go
* Lightweight HTML/JS frontend
* Easy to extend with authentication, payments, or live tracking

## 🛠️ Tech Stack

* **Frontend:** HTML, CSS, JavaScript
* **Backend:** Go (Golang)
* **Database:** (to be added – PostgreSQL/MySQL/MongoDB)
* **Communication:** REST API (JSON)

## 🚀 Getting Started

### Prerequisites

* [Go 1.20+](https://go.dev/dl/)
* Git
* Any modern web browser

### Installation

1. **Clone the repo**

```bash
git clone https://github.com/shark121/pods-test.git
cd pods-test/pod-server-test
```

2. **Run the backend server**

```bash
go run main.go
```

3. **Open the frontend**

* Navigate to the `frontend` folder (or wherever `index.html` lives)
* Open it directly in your browser, or serve it with a simple server:

```bash
npx serve .
```

4. **Test it**
   Make a ride request and check if the backend groups you into a pod.

## 🔮 Roadmap

* [ ] Add database persistence for rides and users
* [ ] Implement WebSocket support for real-time ride tracking
* [ ] Build authentication (JWT) and user profiles
* [ ] Add mobile-friendly UI / app version
* [ ] Optimize matching algorithm with geolocation

## 🤝 Contributing

Contributions are welcome! Feel free to fork, submit pull requests, or open issues.

## 📜 License

MIT License – free to use and modify.

---

Do you want me to **make it more technical** (with API endpoint documentation and request/response examples), or keep it **lightweight for now**?
