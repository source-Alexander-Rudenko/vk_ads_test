# worker-pool

Простой воркер-пул на Go с динамическим добавлением и удалением воркеров.

## Структура проекта

- `cmd/main.go` — демонстрация использования пула.
- `pkg/pool/errors.go` — определения собственных ошибок.
- `pkg/pool/types.go` — базовые типы и сообщения управления.
- `pkg/pool/pool.go` — логика менеджера пула.
- `pkg/pool/worker.go` — реализация самих воркеров.

## Использование

```bash
go build -o workerpool ./cmd/workerpool
./workerpool