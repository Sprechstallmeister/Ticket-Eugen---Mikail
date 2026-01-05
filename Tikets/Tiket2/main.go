package main

import (
	"fmt"
)

// --- Subsystems ---

type ShoppingCart struct {
	items map[string]float64
}

func (c *ShoppingCart) AddItem(name string, price float64) {
	if c.items == nil {
		c.items = make(map[string]float64)
	}
	c.items[name] = price
	fmt.Printf("[CART] Added %s: $%.2f\n", name, price)
}

func (c *ShoppingCart) GetTotal() float64 {
	total := 0.0
	for _, price := range c.items {
		total += price
	}
	return total
}

type InventoryManager struct{}

func (i *InventoryManager) CheckStock(item string) bool {
	fmt.Printf("[INVENTORY] Checking stock for %s\n", item)
	return true // Всегда есть в наличии для демо
}

func (i *InventoryManager) ReserveItem(item string) {
	fmt.Printf("[INVENTORY] Reserved item: %s\n", item)
}

func (i *InventoryManager) ReleaseItem(item string) {
	fmt.Printf("[INVENTORY] Released item: %s\n", item)
}

type PaymentProcessor struct{}

func (p *PaymentProcessor) ProcessPayment(amount float64, card string) bool {
	if amount > 1000 {
		fmt.Printf("[PAYMENT] Payment failed (limit exceeded) for card %s\n", card)
		return false
	}
	fmt.Printf("[PAYMENT] Processed $%.2f from card %s\n", amount, card)
	return true
}

type ShippingService struct{}

func (s *ShippingService) CreateShipment(address string) string {
	tracking := "TRACK-" + address[:3] + "-123"
	fmt.Printf("[SHIPPING] Shipment created: %s\n", tracking)
	return tracking
}

// --- Facade ---

type OrderFacade struct {
	cart      *ShoppingCart
	inventory *InventoryManager
	payment   *PaymentProcessor
	shipping  *ShippingService
}

func NewOrderFacade() *OrderFacade {
	return &OrderFacade{
		cart:      &ShoppingCart{},
		inventory: &InventoryManager{},
		payment:   &PaymentProcessor{},
		shipping:  &ShippingService{},
	}
}

func (f *OrderFacade) AddToCart(item string, price float64) {
	f.cart.AddItem(item, price)
}

// PlaceOrder - метод фасада с ошибками логики
func (f *OrderFacade) PlaceOrder(card string, address string) {
	fmt.Println("\n--- Starting Order Process ---")
	
	// Ошибка 1: Нет проверки, пустая ли корзина
	
	total := f.cart.GetTotal()
	
	// Резервируем товары
	for item := range f.cart.items {
		if f.inventory.CheckStock(item) {
			f.inventory.ReserveItem(item)
		}
	}
	
	// Оплата
	success := f.payment.ProcessPayment(total, card)
	if !success {
		// Ошибка 2: Если оплата не прошла, мы просто выходим.
		// Товар остается "Зарезервированным" навечно! (Memory leak в бизнес-логике)
		fmt.Println("Error: Payment failed")
		return
	}
	
	// Доставка
	tracking := f.shipping.CreateShipment(address)
	fmt.Printf("Order completed successfully! Tracking: %s\n", tracking)
}

func main() {
	facade := NewOrderFacade()
	
	// Сценарий 1: Успешный заказ
	facade.AddToCart("Laptop", 800)
	facade.PlaceOrder("1234-5678", "NY-Street")
	
	// Сценарий 2: Провал оплаты (сумма > 1000)
	// Здесь будет видна ошибка логики: товар зарезервируется, но не освободится
	facade = NewOrderFacade()
	facade.AddToCart("Gaming PC", 2000)
	facade.PlaceOrder("1234-5678", "LA-Avenue")
}
