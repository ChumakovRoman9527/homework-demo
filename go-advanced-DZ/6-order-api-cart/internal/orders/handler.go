package orders

import (
	"6-order-api-cart/configs"
	"6-order-api-cart/pkg/middleware"
	"fmt"
	"net/http"
)

type orderHandler struct {
	OrderRepository *OrderRepository
	Config          *configs.Config
}

type OrderHandlerDeps struct {
	OrderRepository *OrderRepository
	Config          *configs.Config
}

func OrderHandler(router *http.ServeMux, deps OrderHandlerDeps) {

	handler := &orderHandler{
		OrderRepository: deps.OrderRepository,
		Config:          deps.Config,
	}
	authDeps := middleware.AuthDeps{Config: deps.Config}

	stack := middleware.Chain(
		authDeps.IsAuthed,
	)
	router.Handle("POST /order", stack(handler.CreateOrder()))

}

func (handler *orderHandler) CreateOrder() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//1. Надо по токену получить Юзера, а мы его записали в конекст !!! надо вытащить его из контекста

		// Получаем значение из контекста по ключу ConstPhoneKey
		phone, ok := r.Context().Value(middleware.ConstPhoneKey).(string)
		if !ok {
			http.Error(w, "Phone not found in context", http.StatusBadRequest)
			return
		}

		// Теперь можно использовать phone в логике обработчика
		fmt.Printf("Phone from context: %s\n", phone)

	}
}
