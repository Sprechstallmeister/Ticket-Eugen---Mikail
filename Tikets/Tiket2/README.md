# Tiket: Facade Pattern

## Описание
Фасад для упрощения работы с системой заказов (Корзина, Склад, Оплата, Доставка).

## Структура
- `ShoppingCart`
- `InventoryManager`
- `PaymentProcessor`
- `ShippingService`
- `OrderFacade` - единая точка входа

## Запуск
```bash
go run main.go
