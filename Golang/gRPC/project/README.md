# My gRPC Project

## Описание

Этот проект представляет собой простой пример клиент-серверного приложения, использующего gRPC для обмена данными. Сервер отдает поток данных (просто возрастающее число) на клиента. Через какое-то время, клиент посылает запрос на остановку передачи.

## Установка

### Tребования
- Go: Вам нужно установить Go, чтобы запустить этот проект. Проверьте версию Go, используя команду go version в командной строке. Этот проект требует Go 1.16 или выше.    
- gRPC: Вам также нужно установить gRPC. Вы можете установить его, следуя инструкциям на официальном сайте gRPC.    
- Protocol Buffers: Этот проект использует Protocol Buffers для определения сервисов gRPC, поэтому вам нужно установить компилятор protoc и плагин Go для него. Вы можете найти инструкции по установке на странице Protocol Buffers для Go.   
- Библиотеки Go: Вам нужно установить все зависимости Go проекта. Вы можете сделать это, перейдя в корневую директорию проекта и запустив go mod download.    

## Запуск

### Сервер

Запустите сервер, используя следующую команду:

```bash
$ go run cmd/server/main.go --port=50051
```

### Клиент

Запустите клиента, используя следующую команду:

```bash
$ go run cmd/client/main.go --address=localhost:50051 --username=test --password=test --interval=1000 --duration=8
```

- address=localhost:50051: Этот флаг указывает адрес сервера, к которому клиент будет подключаться.     
- username=test: Этот флаг указывает имя пользователя для входа в систему.
- password=test: Этот флаг указывает пароль для входа в систему.   
- interval=1000: Этот флаг указывает интервал (в миллисекундах), с которым сервер будет отправлять данные клиенту.   
- duration=8: Этот флаг указывает продолжительность (в секундах), в течение которой клиент будет получать данные от сервера.