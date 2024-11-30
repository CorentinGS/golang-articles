package generator

import (
	"fmt"
	"math/rand"
)

type EcommerceExample struct{}

type Product struct {
	ID    int
	Name  string
	Price float64
}

type Order struct {
	ID       int
	UserID   int
	Products []Product
	Total    float64
}

func productGenerator(count int) <-chan Product {
	out := make(chan Product)
	go func() {
		defer close(out)
		for i := 0; i < count; i++ {
			out <- Product{
				ID:    i + 1,
				Name:  fmt.Sprintf("Product-%d", i+1),
				Price: 10.0 + float64(i),
			}
		}
	}()
	return out
}

func orderGenerator(userCount, orderPerUser int, products <-chan Product) <-chan Order {
	out := make(chan Order)
	go func() {
		defer close(out)
		var orderID int
		for userID := 1; userID <= userCount; userID++ {
			for i := 0; i < orderPerUser; i++ {
				orderID++
				var orderProducts []Product
				var total float64
				for j := 0; j < rand.Intn(5)+1; j++ {
					product := <-products
					orderProducts = append(orderProducts, product)
					total += product.Price
				}
				out <- Order{
					ID:       orderID,
					UserID:   userID,
					Products: orderProducts,
					Total:    total,
				}
			}
		}
	}()
	return out
}

func (g EcommerceExample) Execute() {
	productChan := productGenerator(1000)
	orderChan := orderGenerator(100, 5, productChan)

	// Simulate sending orders to an API
	for order := range orderChan {
		// In a real scenario, you'd send this to your API
		fmt.Printf("Sending order %d for user %d with total $%.2f\n", order.ID, order.UserID, order.Total)
	}
}
