# ToGo List

This is a todo list backend HTTP API written in Go. It is developed mainly for
self-learning different practices in writing in Go programming languages.

## Features

In addition to basic CRUD of todo list and items, the following are implemented:

- Health Checking
- Input Validation
- Logging
- API documentation

## Configuration

### Config Location

The application will read configuration from the below locations:

- `/etc/togolist/config.yaml`
- `$XDG_CONFIG_HOME/togolist/config.yaml`
- `./config/config.yaml`

## Libraries Used

- Dependency Injection: [Uber Fx](https://github.com/uber-go/fx)
- Logger: [Zap](https://github.com/uber-go/zap)
- Config: [Viper](https://github.com/spf13/viper)
- ORM: [GORM](https://github.com/go-gorm/gorm)
- HTTP: [Gin](https://github.com/gin-gonic/gin)
- Documentation: [Swagger](https://github.com/swaggo/swag)

## References

- [Simplify dependency injection using Uber FX](https://medium.com/@erez.levi/using-uber-fx-to-simplify-dependency-injection-875363245c4c)
