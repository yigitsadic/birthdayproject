
![go version](https://img.shields.io/badge/Go-1.19-blue)
[![Go Report Card](https://goreportcard.com/badge/github.com/yigitsadic/birthday-api)](https://goreportcard.com/report/github.com/yigitsadic/birthday-api)

# 🎁 Birthday App Backend Project

Backend for birthdays app project.

## Issues & Roadmap

- [ ] Reading config for JWT secret, db connection
- [ ] Extend logs to handlers.
- [ ] Add database timeouts.
- [ ] Collect logs with a tool.
- [ ] Configure CI for building docker image.
- [ ] Configure CI for linting, code quality and coverage badges.
- [ ] Integrate automated OpenApi v3 spec generation.
- [ ] Add references for admin, front-end, db migrator, cron & background job execution repositories.
- [ ] Add request a demo handler, store and openapi spec.
- [ ] Extend & describe project stack on README file.
- [ ] sqlc integration?
- [ ] Add integration tests.
- [ ] Graceful shutdown.

## CLI Reference

Running the application:

```
make run
```

Running tests without db:

```
make test
```

Run all tests:

In order to run integration tests, docker should be installed on your system.

```
make test/db
```

## API Reference

Could be found on [openapi.yml](docs/openapi.yml)
