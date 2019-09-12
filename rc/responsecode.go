package rc

import "fmt"

const (
	//GENERAL
	Success              = "00"
	Failed               = "01"
	InternalServiceError = "50"
	RequestTimeOut       = "68"
	TemporaryUnavailable = "51"

	//AUTHORIZATION
	InvalidAuthHeader     = "11"
	ExpiredToken          = "13"
	InvalidUserCredential = "14"
	AuthRequired          = "15"
	InvalidToken          = "16"
	InvalidCredential     = "17"

	//Open Biller Payment Ecommerce
	InvalidMerchantId    = "20"
	InvalidDateFormat    = "21"
	ExpireCode           = "22"
	InvalidPaymentCode   = "23"
	DuplicatePaymentCode = "24"
	AlreadyPaid          = "88"
	//Other					= xx
	//Other					= xxx //bisa jadi rc dari http status code
)

var ResponseMessage = map[string]string{
	//GENERAL
	Success:              fmt.Sprintf("[%s] %s", "Transaksi Berhasil!"),
	InternalServiceError: fmt.Sprintf("[%s] %s", InternalServiceError, "Internal service error"),
	RequestTimeOut:       fmt.Sprintf("[%s] %s", RequestTimeOut, "Request timeout."),
	TemporaryUnavailable: fmt.Sprintf("[%s] %s", TemporaryUnavailable, "Service temporarily unavailable."),

	//AUTHORIZATION
	InvalidAuthHeader:     fmt.Sprintf("[%s] %s", InvalidAuthHeader, "Invalid authorization header"),
	InvalidCredential:     fmt.Sprintf("[%s] %s", InvalidCredential, "Invalid credential value"),
	ExpiredToken:          fmt.Sprintf("[%s] %s", ExpiredToken, "Token is expired"),
	InvalidUserCredential: fmt.Sprintf("[%s] %s", InvalidUserCredential, "Invalid user credential"),
	AuthRequired:          fmt.Sprintf("[%s] %s", AuthRequired, "An authorization header is required!"),
	InvalidToken:          fmt.Sprintf("[%s] %s", InvalidToken, "Invalid access token!"),

	//Open Biller Payment Ecommerce
	InvalidMerchantId:    fmt.Sprintf("[%s] %s", InvalidMerchantId, "Invalid merchant id"),
	InvalidPaymentCode:   fmt.Sprintf("[%s] %s", InvalidPaymentCode, "Kode pembayaran salah"),
	InvalidDateFormat:    fmt.Sprintf("[%s] %s", InvalidDateFormat, "Invalid date format"),
	ExpireCode:           fmt.Sprintf("[%s] %s", ExpireCode, "Kode pembayaran sudah expire"),
	AlreadyPaid:          fmt.Sprintf("[%s] %s", AlreadyPaid, "Tagihan sudah terbayar"),
	DuplicatePaymentCode: fmt.Sprintf("[%s] %s", DuplicatePaymentCode, "Kode pembayaran sudah pernah digunakan"),
}
