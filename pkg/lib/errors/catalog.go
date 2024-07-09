package errors

import (
	"fmt"
	"net/http"
)

type CustomError interface {
	error
	HTTPCode() int
	InternalCode() string
}

type AppError struct {
	Err          error
	httpCode     int
	internalCode string
}

func (e AppError) Error() string {
	return e.Err.Error()
}

func (e AppError) HTTPCode() int {
	return e.httpCode
}

func (e AppError) InternalCode() string {
	return e.internalCode
}

var (
	OrderGet                      = AppError{Err: fmt.Errorf("error getting order"), httpCode: http.StatusInternalServerError, internalCode: "error_getting_order"}
	OrdersGet                     = AppError{Err: fmt.Errorf("error getting orders"), httpCode: http.StatusInternalServerError, internalCode: "error_getting_orders"}
	OrderCompleted                = AppError{Err: fmt.Errorf("order completed"), httpCode: http.StatusInternalServerError, internalCode: "order completed"}
	OrderCanceled                 = AppError{Err: fmt.Errorf("order canceled"), httpCode: http.StatusInternalServerError, internalCode: "order canceled"}
	OrderNotFound                 = AppError{Err: fmt.Errorf("order not found"), httpCode: http.StatusNotFound, internalCode: "order_not_found"}
	OrderSave                     = AppError{Err: fmt.Errorf("error saving order"), httpCode: http.StatusInternalServerError, internalCode: "error_saving_order"}
	OrderUpdate                   = AppError{Err: fmt.Errorf("error updating order"), httpCode: http.StatusInternalServerError, internalCode: "error_updating_order"}
	OrderAlreadyExists            = AppError{Err: fmt.Errorf("error order already exist"), httpCode: http.StatusInternalServerError, internalCode: "error_order_already_exist"}
	ProfessionalGet               = AppError{Err: fmt.Errorf("error getting professional"), httpCode: http.StatusInternalServerError, internalCode: "error_getting_professional"}
	ProfessionalAlreadyExist      = AppError{Err: fmt.Errorf("error professional already exist"), httpCode: http.StatusInternalServerError, internalCode: "error_professional_already_exist"}
	ErrCreateProfessional         = AppError{Err: fmt.Errorf("error creating professional"), httpCode: http.StatusInternalServerError, internalCode: "error_creating_professional"}
	UserGet                       = AppError{Err: fmt.Errorf("error getting user"), httpCode: http.StatusInternalServerError, internalCode: "error_getting_user"}
	UserNotFound                  = AppError{Err: fmt.Errorf("user not found"), httpCode: http.StatusNotFound, internalCode: "user_not_found"}
	ProfessionalNotFound          = AppError{Err: fmt.Errorf("professional not found"), httpCode: http.StatusNotFound, internalCode: "professional_not_found"}
	UserAlreadyExists             = AppError{Err: fmt.Errorf("error user already exist"), httpCode: http.StatusInternalServerError, internalCode: "error_user_already_exist"}
	ErrInvalidRegisterUser        = AppError{Err: fmt.Errorf("invalid register user"), httpCode: http.StatusInternalServerError, internalCode: "invalid_register_user"}
	ErrCreateUser                 = AppError{Err: fmt.Errorf("error creating user"), httpCode: http.StatusInternalServerError, internalCode: "error_creating_user"}
	EmptyUserName                 = AppError{Err: fmt.Errorf("empty user name"), httpCode: http.StatusBadRequest, internalCode: "empty_user_name"}
	EmptyEmail                    = AppError{Err: fmt.Errorf("empty email"), httpCode: http.StatusBadRequest, internalCode: "empty_email"}
	EmptyPassword                 = AppError{Err: fmt.Errorf("empty password"), httpCode: http.StatusBadRequest, internalCode: "empty_password"}
	EmptyName                     = AppError{Err: fmt.Errorf("empty name"), httpCode: http.StatusBadRequest, internalCode: "empty_name"}
	EmptyLastName                 = AppError{Err: fmt.Errorf("empty last name"), httpCode: http.StatusBadRequest, internalCode: "empty_last_name"}
	EmptyPhone                    = AppError{Err: fmt.Errorf("empty phone"), httpCode: http.StatusBadRequest, internalCode: "empty_phone"}
	EmptyAddress                  = AppError{Err: fmt.Errorf("empty address"), httpCode: http.StatusBadRequest, internalCode: "empty_address"}
	ErrInvalidLocation            = AppError{Err: fmt.Errorf("invalid location"), httpCode: http.StatusBadRequest, internalCode: "invalid_location"}
	ErrInvalidProfession          = AppError{Err: fmt.Errorf("invalid profession"), httpCode: http.StatusBadRequest, internalCode: "invalid_profession"}
	ErrInvalidInput               = AppError{Err: fmt.Errorf("invalid input"), httpCode: http.StatusBadRequest, internalCode: "invalid_input"}
	ErrInvalidStatus              = AppError{Err: fmt.Errorf("invalid status"), httpCode: http.StatusBadRequest, internalCode: "invalid_status"}
	ErrInvalidUserID              = AppError{Err: fmt.Errorf("invalid user id"), httpCode: http.StatusBadRequest, internalCode: "invalid_user_id"}
	ErrInvalidOrderID             = AppError{Err: fmt.Errorf("invalid order id"), httpCode: http.StatusBadRequest, internalCode: "invalid_order_id"}
	ErrInvalidProfessionalID      = AppError{Err: fmt.Errorf("invalid professional id"), httpCode: http.StatusBadRequest, internalCode: "invalid_professional_id"}
	ErrInvalidScheduleTo          = AppError{Err: fmt.Errorf("invalid schedule to"), httpCode: http.StatusBadRequest, internalCode: "invalid_schedule_to"}
	ErrInvalidDay                 = AppError{Err: fmt.Errorf("invalid day"), httpCode: http.StatusBadRequest, internalCode: "invalid_day"}
	ErrInvalidCreateOrder         = AppError{Err: fmt.Errorf("invalid create order"), httpCode: http.StatusInternalServerError, internalCode: "invalid_create_order"}
	ErrInvalidAdvanceOrder        = AppError{Err: fmt.Errorf("invalid advance order"), httpCode: http.StatusInternalServerError, internalCode: "invalid_advance_order"}
	Validators                    = AppError{Err: fmt.Errorf("error building validators"), httpCode: http.StatusInternalServerError, internalCode: "error_building_validators"}
	ErrInvalidDepartment          = AppError{Err: fmt.Errorf("invalid department"), httpCode: http.StatusBadRequest, internalCode: "invalid_department"}
	ErrInvalidDistrict            = AppError{Err: fmt.Errorf("invalid district"), httpCode: http.StatusBadRequest, internalCode: "invalid_district"}
	OrderAlreadyInRequestedStatus = AppError{Err: fmt.Errorf("order already in requested status"), httpCode: http.StatusInternalServerError, internalCode: "order_already_in_requested_status"}
	ErrInvalidEmail               = AppError{Err: fmt.Errorf("invalid email"), httpCode: http.StatusBadRequest, internalCode: "invalid_email"}
)
