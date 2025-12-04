package orders

const (
	OrderStatusCreated          = "создан"              // Создан / Оформлен
	OrderStatusPendingPayment   = "ожидает_оплаты"      // В ожидании оплаты
	OrderStatusPaid             = "оплачен"             // Оплачен
	OrderStatusCancelledByBuyer = "отменён_покупателем" // Отменён покупателем

	OrderStatusAccepted           = "принят_в_работу" // Принят в работу
	OrderStatusProcessing         = "в_обработке"     // В обработке / На сборке
	OrderStatusReserved           = "в_резерве"       // Резерв
	OrderStatusSearching          = "на_поиске"       // На поиске
	OrderStatusPartiallyAssembled = "частично_собран" // Частично собран
	OrderStatusAssembled          = "собран"          // Собран

	OrderStatusReadyToShip     = "готов_к_отгрузке" // Готов к отгрузке
	OrderStatusHandedToCourier = "передан_курьеру"  // Передан курьеру / В службу доставки
	OrderStatusShipped         = "отправлен"        // Отправлен

	OrderStatusInTransit       = "в_пути"                 // В пути
	OrderStatusAtSortingCenter = "в_сортировочном_центре" // Прибыл в сортировочный центр
	OrderStatusAtCustoms       = "на_таможне"             // На таможне
	OrderStatusDeliveryDelayed = "задержка_доставки"      // Задержка доставки

	OrderStatusReadyForPickup      = "готов_к_выдаче"     // Готов к выдаче
	OrderStatusCourierOnTheWay     = "курьер_в_пути"      // Передан курьеру (курьер в пути)
	OrderStatusDelivered           = "доставлен"          // Доставлен / Вручен
	OrderStatusDeliveryFailed      = "неудачная_доставка" // Неудачная доставка
	OrderStatusReturnedToWarehouse = "возвращён_на_склад" // Возвращён на склад

	OrderStatusCancelledByStore  = "отменён_магазином"     // Отменён магазином
	OrderStatusReturnInProgress  = "на_возврате"           // На возврате
	OrderStatusReturnCompleted   = "возврат_завершён"      // Возврат завершён
	OrderStatusLost              = "утерян"                // Утерян
	OrderStatusAwaitConfirmation = "ожидает_подтверждения" // Ожидает подтверждения
)
