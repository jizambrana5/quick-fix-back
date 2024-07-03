package rest

import (
	"time"

	"github.com/jizambrana5/quickfix-back/pkg/domain"
	"github.com/jizambrana5/quickfix-back/pkg/lib/errors"
)

const (
	layout = "2006-01-02 15:04"
)

type (
	CreateOrderRequest struct {
		UserID         uint64 `json:"user_id"`
		ProfessionalID uint64 `json:"professional_id"`
		ScheduleTo     string `json:"schedule_to"`
	}

	AdvanceOrderRequest struct {
		Status domain.Status `json:"status"`
	}

	RegisterUserRequest struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	RegisterProfessionalRequest struct {
		Username    string   `json:"username"`
		Email       string   `json:"email"`
		Password    string   `json:"password"`
		Profession  string   `json:"profession"`
		Description string   `json:"description"`
		Location    Location `json:"location"`
	}
	Location struct {
		Department string `json:"department"`
		District   string `json:"district"`
	}
)

func (r AdvanceOrderRequest) Validate() error {
	if r.Status != domain.OrderStatusPending && r.Status != domain.OrderStatusAccepted && r.Status != domain.OrderStatusCompleted &&
		r.Status != domain.OrderStatusCancelled {
		return errors.ErrInvalidInput
	}
	return nil
}

func (co CreateOrderRequest) Validate() error {
	if co.UserID == 0 {
		return errors.ErrInvalidUserID
	}
	if co.ProfessionalID == 0 {
		return errors.ErrInvalidProfessionalID
	}
	if co.ScheduleTo == "" {
		return errors.ErrInvalidScheduleTo
	}
	parsedTime, err := time.Parse(layout, co.ScheduleTo)
	if err != nil {
		return errors.ErrInvalidScheduleTo
	}

	loc, _ := time.LoadLocation("America/Sao_Paulo")
	timeInLoc := parsedTime.In(loc)
	if timeInLoc.Before(time.Now().In(loc)) {
		return errors.ErrInvalidScheduleTo
	}
	return nil
}

func (r RegisterUserRequest) Validate() error {
	if r.Username == "" {
		return errors.EmptyUserName
	}
	if r.Email == "" {
		return errors.EmptyEmail
	}
	if r.Password == "" {
		return errors.EmptyPassword
	}
	return nil
}

func (rp RegisterProfessionalRequest) Validate() error {
	if rp.Username == "" {
		return errors.EmptyUserName
	}
	if rp.Email == "" {
		return errors.EmptyEmail
	}
	if rp.Password == "" {
		return errors.EmptyPassword
	}

	// Cargar ubicaciones válidas
	/*
		locations, err := utils.LoadLocations("pkg/utils/mendoza_locations.json")
		if err != nil {
			return errors2.New("failed to load locations")
		}

		// Validación de la ubicación
		err = utils.ValidateLocation(rp.Location.Department, rp.Location.District, locations)
		if err != nil {
			return err
		}

		switch strings.ToUpper(rp.Profession) {
		case string(domain.Plomero), string(domain.Gasista), string(domain.Electricista):
			return nil
		default:
			return errors.ErrInvalidProfession
		}*/
	return nil
}
