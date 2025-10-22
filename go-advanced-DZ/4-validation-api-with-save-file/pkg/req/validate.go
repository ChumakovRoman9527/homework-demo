package req

import (
	"github.com/go-playground/validator/v10"
)

// IsValid проверяет структуру данных на соответствие правилам валидации.
//
// Параметры:
//   - payload: объект произвольного типа, который необходимо проверить.
//
// Возвращает:
//   - error: ошибку валидации, если проверка провалена; nil, если проверка успешна.
//
// Примечание:
// Для корректной работы структура payload должна содержать аннотации валидации (например, `validate:"required"`).
func IsValid[T any](payload T) error {
	validate := validator.New()
	err := validate.Struct(payload)
	if err != nil {

		return err
	}
	return nil
}
