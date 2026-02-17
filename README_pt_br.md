# 🔥 Hot Configuration Management

[![project](https://img.shields.io/badge/github-psaraiva%2Flab-hcm)](https://img.shields.io/badge/github-psaraiva%2F-lab-hcm)
[![License](https://img.shields.io/badge/license-MIT-%233DA639.svg)](https://opensource.org/licenses/MIT)

[![Go Report Card](https://goreportcard.com/badge/github.com/psaraiva/lab-hcm)](https://goreportcard.com/report/github.com/psaraiva/lab-hcm)

[![Language: English](https://img.shields.io/badge/Idioma-English-blue?style=flat-square)](./README.md)

Projeto demonstrativo (lab) de **Gestão de Configuração Quente** (Hot Configuration Management) em Go, utilizando a biblioteca Viper para permitir alterações de configuração em tempo de execução sem necessidade de reiniciar a aplicação.

## 🎯 Sobre o Projeto

Este projeto demonstra como implementar **recarregamento de configuração em tempo real** em aplicações Go. A aplicação monitora o arquivo `config.json` e, sempre que este é alterado, as novas configurações são aplicadas automaticamente sem necessidade de restart.

O exemplo utiliza um sistema de simulação de lançamento de dados com múltiplas goroutines concorrentes, onde o tempo de sleep de cada processo pode ser alterado dinamicamente através do arquivo de configuração.

## ✨ Características

- ✅ **Hot Reload**: Alterações no arquivo de configuração são detectadas e aplicadas automaticamente
- ✅ **Configuração Padrão**: Valores fallback caso configurações estejam ausentes
- ✅ **Concorrência Segura**: Uso de goroutines com `sync.WaitGroup`
- ✅ **File Watching**: Monitoramento de mudanças com `fsnotify`

## 🛠 Tecnologias Utilizadas

- **Go 1.25.0**: Linguagem de programação
- **[Viper](https://github.com/spf13/viper)**: Gestão de configuração
- **[fsnotify](https://github.com/fsnotify/fsnotify)**: Monitoramento de mudanças em arquivos

## 📦 Pré-requisitos

- Go 1.25.0 ou superior
- Git

## 🚀 Instalação

1. Clone o repositório:
```bash
git clone https://github.com/psaraiva/lab-hcm.git
cd hcm
```

2. Instale as dependências:
```bash
go mod download
```

## 💻 Uso

1. Execute a aplicação:
```bash
go run .
```

2. A aplicação iniciará com a configuração definida em `config.json`

3. **Teste o Hot Reload**: Com a aplicação em execução, edite o arquivo `config.json` e altere o valor de `sleep.duration`:
```json
{
    "name": "Example Configuration",
    "sleep": {
        "duration": 2
    }
}
```

4. Observe no terminal que a alteração foi detectada e aplicada automaticamente:
```
Config file changed: config.json
Process #Ab3d2F - Sleep Config: 3
```

## 📂 Estrutura do Projeto

```
lab-hcm/
├── config.json         # Arquivo de configuração
├── dice.go             # Implementação do dado
├── process.go          # Lógica de processamento
├── main.go             # Ponto de entrada da aplicação
├── go.mod              # Gerenciamento de dependências
└── README_pt_br.md     # Este arquivo
```

## 📝 Exemplo de saída

```bash
Processo #g61FDq - Sleep Config: 2
Processo #jkQi8B - Sleep Config: 2
Processo #g61FDq - Sleep Config: 2 - Value: 5
Processo #tCmewX - Sleep Config: 2
Arquivo de configuração alterado: /.../config.json
Arquivo de configuração alterado: /.../config.json
Processo #fVvMVS - Sleep Config: 5
Processo #jkQi8B - Sleep Config: 2 - Value: 6
Processo #tCmewX - Sleep Config: 2 - Value: 4
Processo #sZJJM3 - Sleep Config: 5
Processo #uZa3eK - Sleep Config: 5
Processo #9IyD7j - Sleep Config: 5
Arquivo de configuração alterado: /.../config.json
Arquivo de configuração alterado: /.../config.json
Processo #yIiCW2 - Sleep Config: 7
Processo #fVvMVS - Sleep Config: 5 - Value: 6
Processo #atdfWv - Sleep Config: 7
Processo #sZJJM3 - Sleep Config: 5 - Value: 5
Processo #wCauZa - Sleep Config: 7
Processo #uZa3eK - Sleep Config: 5 - Value: 5
Processo #9IyD7j - Sleep Config: 5 - Value: 3
Processo #yIiCW2 - Sleep Config: 7 - Value: 6
Processo #atdfWv - Sleep Config: 7 - Value: 5
Processo #wCauZa - Sleep Config: 7 - Value: 1
```

## 📝 Exemplo de saída (ordenado por #processo)

```bash
Processo #g61FDq - Sleep Config: 2
Processo #g61FDq - Sleep Config: 2 - Value: 5
Processo #jkQi8B - Sleep Config: 2
Processo #jkQi8B - Sleep Config: 2 - Value: 6
Processo #tCmewX - Sleep Config: 2
Processo #tCmewX - Sleep Config: 2 - Value: 4
Arquivo de configuração alterado: /.../config.json
Arquivo de configuração alterado: /.../config.json
Processo #fVvMVS - Sleep Config: 5
Processo #fVvMVS - Sleep Config: 5 - Value: 6
Processo #sZJJM3 - Sleep Config: 5
Processo #sZJJM3 - Sleep Config: 5 - Value: 5
Processo #uZa3eK - Sleep Config: 5
Processo #uZa3eK - Sleep Config: 5 - Value: 5
Processo #9IyD7j - Sleep Config: 5
Processo #9IyD7j - Sleep Config: 5 - Value: 3
Arquivo de configuração alterado: /.../config.json
Arquivo de configuração alterado: /.../config.json
Processo #yIiCW2 - Sleep Config: 7
Processo #yIiCW2 - Sleep Config: 7 - Value: 6
Processo #wCauZa - Sleep Config: 7
Processo #wCauZa - Sleep Config: 7 - Value: 1
Processo #atdfWv - Sleep Config: 7
Processo #atdfWv - Sleep Config: 7 - Value: 5
```

Obs: A aplicação usará o valor padrão **10** definido em `viper.SetDefault()`.

## 📄 Licença

Este projeto é open source e está sob a licença MIT.

## 👤 Autor

Developed by [@psaraiva](https://github.com/psaraiva)

---

⭐ Se este projeto foi útil para você, considere dar uma estrela!
