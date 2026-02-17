# 🔥 Hot Configuration Management

[![project](https://img.shields.io/badge/github-psaraiva%2Flab-hcm-blue)](https://img.shields.io/badge/github-psaraiva%2F-lab-hcm-blue)
[![License](https://img.shields.io/badge/license-MIT-%233DA639.svg)](https://opensource.org/licenses/MIT)

[![Go Report Card](https://goreportcard.com/badge/github.com/psaraiva/lab-hcm)](https://goreportcard.com/report/github.com/psaraiva/lab-hcm)

[![Language: Português](https://img.shields.io/badge/Language-Portugu%C3%AAs-blue?style=flat-square)](./README_pt_br.md)

Demonstration project (lab) of **Hot Configuration Management** in Go, using the Viper library to allow configuration changes at runtime without requiring application restart.

## 🎯 About the Project

This project demonstrates how to implement **real-time configuration reloading** in Go applications. The application monitors the `config.json` file and whenever it is changed, the new configurations are automatically applied without requiring a restart.

The example uses a dice roll simulation system with multiple concurrent goroutines, where the sleep time of each process can be dynamically changed through the configuration file.

## ✨ Features

- ✅ **Hot Reload**: Changes to the configuration file are detected and applied automatically
- ✅ **Default Configuration**: Fallback values if configurations are missing
- ✅ **Safe Concurrency**: Use of goroutines with `sync.WaitGroup`
- ✅ **File Watching**: Change monitoring with `fsnotify`

## 🛠 Technologies Used

- **Go 1.25.0**: Programming language
- **[Viper](https://github.com/spf13/viper)**: Configuration management
- **[fsnotify](https://github.com/fsnotify/fsnotify)**: File change monitoring

## 📦 Prerequisites

- Go 1.25.0 or higher
- Git

## 🚀 Installation

1. Clone the repository:
```bash
git clone https://github.com/psaraiva/lab-hcm.git
cd hcm
```

2. Install dependencies:
```bash
go mod download
```

## 💻 Usage

1. Run the application:
```bash
go run .
```

2. The application will start with the configuration defined in `config.json`

3. **Test Hot Reload**: With the application running, edit the `config.json` file and change the value of `sleep.duration`:
```json
{
    "name": "Example Configuration",
    "sleep": {
        "duration": 2
    }
}
```

4. Observe in the terminal that the change was detected and applied automatically:
```
Config file changed: config.json
Process #Ab3d2F - Sleep Config: 3
```

## 📂 Project Structure

```
lab-hcm/
├── config.json         # Configuration file
├── dice.go             # Dice implementation
├── process.go          # Processing logic
├── main.go             # Application entry point
├── go.mod              # Dependency management
└── README.md           # This file
```

## 📝 Output Example

```bash
Process #g61FDq - Sleep Config: 2
Process #jkQi8B - Sleep Config: 2
Process #g61FDq - Sleep Config: 2 - Value: 5
Process #tCmewX - Sleep Config: 2
Config file changed: /.../config.json
Config file changed: /.../config.json
Process #fVvMVS - Sleep Config: 5
Process #jkQi8B - Sleep Config: 2 - Value: 6
Process #tCmewX - Sleep Config: 2 - Value: 4
Process #sZJJM3 - Sleep Config: 5
Process #uZa3eK - Sleep Config: 5
Process #9IyD7j - Sleep Config: 5
Config file changed: /.../config.json
Config file changed: /.../config.json
Process #yIiCW2 - Sleep Config: 7
Process #fVvMVS - Sleep Config: 5 - Value: 6
Process #atdfWv - Sleep Config: 7
Process #sZJJM3 - Sleep Config: 5 - Value: 5
Process #wCauZa - Sleep Config: 7
Process #uZa3eK - Sleep Config: 5 - Value: 5
Process #9IyD7j - Sleep Config: 5 - Value: 3
Process #yIiCW2 - Sleep Config: 7 - Value: 6
Process #atdfWv - Sleep Config: 7 - Value: 5
Process #wCauZa - Sleep Config: 7 - Value: 1
```

## 📝 Output Example (sorted by #process)

```bash
Process #g61FDq - Sleep Config: 2
Process #g61FDq - Sleep Config: 2 - Value: 5
Process #jkQi8B - Sleep Config: 2
Process #jkQi8B - Sleep Config: 2 - Value: 6
Process #tCmewX - Sleep Config: 2
Process #tCmewX - Sleep Config: 2 - Value: 4
Config file changed: /.../config.json
Config file changed: /.../config.json
Process #fVvMVS - Sleep Config: 5
Process #fVvMVS - Sleep Config: 5 - Value: 6
Process #sZJJM3 - Sleep Config: 5
Process #sZJJM3 - Sleep Config: 5 - Value: 5
Process #uZa3eK - Sleep Config: 5
Process #uZa3eK - Sleep Config: 5 - Value: 5
Process #9IyD7j - Sleep Config: 5
Process #9IyD7j - Sleep Config: 5 - Value: 3
Config file changed: /.../config.json
Config file changed: /.../config.json
Process #yIiCW2 - Sleep Config: 7
Process #yIiCW2 - Sleep Config: 7 - Value: 6
Process #wCauZa - Sleep Config: 7
Process #wCauZa - Sleep Config: 7 - Value: 1
Process #atdfWv - Sleep Config: 7
Process #atdfWv - Sleep Config: 7 - Value: 5
```

Note: The application will use the default value **10** defined in `viper.SetDefault()`.

## 📄 License

This project is open source and is under the MIT license.

## 👤 Author

Developed by [@psaraiva](https://github.com/psaraiva)

---

⭐ If this project was useful to you, consider giving it a star!
