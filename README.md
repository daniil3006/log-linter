# log-linter

Линтер поддерживает следующие логгеры:

- `log/slog`
- `go.uber.org/zap`

# Правила

### 1. Лог-сообщения должны начинаться со строчной буквы

**Неправильно**  
`slog.Info("Service started")`

**Правильно**  
`slog.Info("service started")`

### 2. Лог-сообщения должны быть только на английском языке

**Неправильно**  
`slog.Info("ошибка подключения к базе данных")`

**Правильно**  
`slog.Info("failed to connect to database")`

### 3. Лог-сообщения не должны содержать спецсимволы или эмодзи

**Неправильно**  
`slog.Info("service started!")`

**Правильно**  
`slog.Info("service started")`

### 4. Лог-сообщения не должны содержать потенциально чувствительные данные

**Неправильно**  
`slog.Info(token + "token")`

**Правильно**  
`slog.Info("token is valid)`

# Установка

```
git clone <repo-url>
cd log-linter
```

Линтер можно запустить напрямую:
```
go run ./cmd/app ./...
```

**Линтер можно использовать как плагин для golangci-lint**  

Скачать golangci-lint
```
go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@latest
```

В корне проекта
```
golangci-lint custom
```

Запуск линтера
```
./custom-gcl run ./...
```

Линтер поддерживает авто-исправление для случая с заглавной буквой
```
./custom-gcl run ./... --fix
```

# Конфигурация
Линтер поддерживает конфигурационный файл. В конфигурационном файле можно настраивать правила и добавлять свои паттерны чувствительных данных. 
```
lowercase: true
latin: true
special_chars: false
sensitive_data: true

sensitive_patterns:
  - secret
```

По умолчанию используется конфигурационный файл `loglinter.yaml`, либо можно указать свой:
```
go run ./cmd/app ./... --config=config.yaml
```

# CI

Настроен CI. Запускается при push и pull request. Выполняет:
- загрузку зависимостей
- запуск тестов
- сборку проекта
