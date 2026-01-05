# Task 8: Data Format Adapter Pattern

## Описание
Реализация паттерна Adapter для работы с различными форматами данных (JSON, XML, CSV).

## Паттерн
**Adapter (Адаптер)** - структурный паттерн проектирования, который позволяет объектам с несовместимыми интерфейсами работать вместе.

## Структура
- `DataStorage` - интерфейс для хранения данных
- `JSONStorage` - реализация для JSON
- `XMLStorage` - реализация для XML
- `CSVAdapter` - адаптер для CSV формата

## Запуск
```bash
cd task8-adapter
go run main.go
