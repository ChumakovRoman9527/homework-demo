package orders

import (
	"6-order-api-cart/configs"
	"6-order-api-cart/internal/user"
	"6-order-api-cart/pkg/middleware"
	"6-order-api-cart/pkg/req"
	"6-order-api-cart/pkg/res"
	"net/http"
)

type orderHandler struct {
	OrderRepository *OrderRepository
	UserRepository  *user.UserRepository
	Config          *configs.Config
}

type OrderHandlerDeps struct {
	OrderRepository *OrderRepository
	UserRepository  *user.UserRepository
	Config          *configs.Config
}

func OrderHandler(router *http.ServeMux, deps OrderHandlerDeps) {

	handler := &orderHandler{
		OrderRepository: deps.OrderRepository,
		UserRepository:  deps.UserRepository,
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
		UserId, err := handler.UserRepository.GetIdByPhone(phone)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		body, err := req.HandleBody[CreateOrderRequest](&w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		createdOrder, err := handler.OrderRepository.CreateOrder(UserId, phone, body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		res.Json(w, createdOrder, http.StatusCreated)

	}
}
