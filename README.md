# Эксперименты с Go
## Ковыряю и тыкаю:
- ```net/http```
- [```air```](https://github.com/cosmtrek/air) - live reload
- [```Chi```](https://go-chi.io/) - Router
- [```SqlX```](https://github.com/jmoiron/sqlx) - DB adapter

## To-Do
- [x] Реализация небольшого приложения
- [ ] Накидать пару моделей со связью One-to-Many
- [ ] Накидать пару моделей со связью Many-to-Many
- [ ] Подключить и поковырятся с Redis
- [ ] Подключить и поковырятся с RabbitMQ
- [ ] Логирование
- [ ] Middleware
- [ ] Авторизация

## Run and Build
- Install ```AIR``` for run in live reload mode and run it
```bash
# Install it into ./bin
curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s

# Run with live reload
./bin/air
```

- Go run:
```bash
make run
```

- Build
```bash
make build
```
