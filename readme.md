# 🦊 Go RedFox 🦊
A simpler and more elegant way to handle exception messages in Go! 🚀

RedFox helps you create clean, readable, and maintainable error messages with minimal effort. Perfect for both small projects and large-scale applications. ✨

Features:
- 🎯 Easy to use API
- 📦 Zero dependencies
- 🔒 Concurency safe
- 💾 Zero memory allocation
- 🔍 Clear error tracing
- 🛠 Customizable message templates

## 📚 Table of Contents
- [🦊 Go RedFox 🦊](#-go-redfox-)
	- [📚 Table of Contents](#-table-of-contents)
	- [📥 Installation](#-installation)
	- [🚀 Usage](#-usage)
		- [Basic Usage](#basic-usage)
		- [🔧 Manager Usage](#-manager-usage)
		- [🌐 API Server Usage](#-api-server-usage)
	- [🤝 Contributing](#-contributing)
	- [📃 License](#-license)

## 📥 Installation
```bash
go get github.com/evilmagics/go-redfox
```

## 🚀 Usage
### Basic Usage
The simplest way to use RedFox is to create a new error with a predefined message template.
```go
import "github.com/evilmagics/go-redfox"

func main() {
    err := redfox.New("SERVER_ERROR", "internal server error")
}
```

### 🔧 Manager Usage
If you need more control over your error messages, you can use the Manager.
```go
import (
	"encoding/json"
	"fmt"

	"github.com/evilmagics/go-redfox"
)

func main() {
	manager := redfox.NewManager[string]()

	// Add many exceptions
	// Using Set function carefully, all registered exceptions will be overwritten
	manager.Set(map[string]redfox.Exception[string]{
		"USERNAME_REQUIRED": redfox.New("USERNAME_REQUIRED", "username must filled"),
		"PASSWORD_REQUIRED": redfox.New("PASSWORD_REQUIRED", "password must filled"),
	})

	// Add multiple exceptions
	manager.AddAll(
		redfox.New("AUTHORIZATION_INVALID", "authorization invalid"),
		redfox.New("SIGNATURE_INVALID", "signature invalid"),
	)

	// Add a new error template
	manager.Add(redfox.New("SERVER_ERROR", "internal server error"))

	// Add safe exceptions make sure that new exception will not overwrite existing exceptions
	err := manager.SafeAdd(redfox.New("DATABASE_ERROR", "database not connected"))
	if err != nil {
		panic(err)
	}

	for _, v := range manager.GetAll() {
		j, err := json.Marshal(v.View())
		if err != nil {
			panic(err)
		}

		fmt.Println(string(j))
	}
}
```

### 🌐 API Server Usage
If you're building an API server, you can use the API Server to handle errors.
```go
import (
	"net/http"

	"github.com/evilmagics/go-redfox"
)

type InternalException redfox.Exception[string]

// Assume this is on your internal exception module
// This is just an simple example without using Manager to handle uncached exceptions
var (
	USERNAME_REQUIRED = redfox.NewForAPI("USERNAME_REQUIRED", "username must filled", http.StatusBadRequest)
	PASSWORD_REQUIRED = redfox.NewForAPI("PASSWORD_REQUIRED", "password must filled", http.StatusBadRequest)
	SERVER_ERROR      = redfox.NewForAPI("SERVER_ERROR", "internal server error", http.StatusInternalServerError)
)

func validation() InternalException {
	// Assume username is empty
	return USERNAME_REQUIRED
}

func http_handler(w http.ResponseWriter, r *http.Request) {
	// Validation simulation
	err := validation()

	w.WriteHeader(err.StatusCode())
	w.Write([]byte(err.Message()))
}

func main() {
	http.HandleFunc("/", http_handler)
	http.ListenAndServe(":8080", nil)
}

```

## 🤝 Contributing
Contributions are welcome! Please feel free to open an issue or submit a pull request.

## 📃 License
This project is licensed under the MIT License.
