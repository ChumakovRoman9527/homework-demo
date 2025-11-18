package auth

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	TOKEN string `json:"token"`
}

type RegisterRequest struct {
	LoginRequest
	Name string `json:"name" validate:"required"`
}

type RegisterResponse struct {
	TOKEN string `json:"token"`
}

type LoginPhoneRequest struct {
	Phone string `json:"phone" validate:"required,numeric,startswith=8,"`
}

type LoginPhoneResponse struct {
	SessionID string `json:"sessionId" validate:"required"`
}

type LoginSMSRequest struct {
	LoginPhoneResponse
	Code string `json:"code" validate:"required,numeric"`
}

type LoginSMSResonse struct {
	TOKEN string `json:"token"`
}
