package entities

import (
	errors2 "errors"
	"regexp"
	"strings"
	"time"

	"github.com/jizambrana5/quickfix-back/pkg/domain"
	"github.com/jizambrana5/quickfix-back/pkg/lib/errors"
	"github.com/jizambrana5/quickfix-back/pkg/utils"
)

const (
	Layout = "2006-01-02 15:04"
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
		Name     string `json:"name"`
		LastName string `json:"last_name"`
		Phone    string `json:"phone"`
		Address  string `json:"address"`
	}

	RegisterProfessionalRequest struct {
		Username    string   `json:"username"`
		Email       string   `json:"email"`
		Password    string   `json:"password"`
		Profession  string   `json:"profession"`
		Description string   `json:"description"`
		Location    Location `json:"location"`
		Name        string   `json:"name"`
		LastName    string   `json:"last_name"`
		Phone       string   `json:"phone"`
		Address     string   `json:"address"`
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
	parsedTime, err := time.Parse(Layout, co.ScheduleTo)
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
	if r.Name == "" {
		return errors.EmptyName
	}
	if r.LastName == "" {
		return errors.EmptyLastName
	}
	if r.Phone == "" {
		return errors.EmptyPhone
	}
	if r.Address == "" {
		return errors.EmptyAddress
	}
	if err := IsValidEmail(r.Email); err != nil {
		return err
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

	if rp.Name == "" {
		return errors.EmptyName
	}
	if rp.LastName == "" {
		return errors.EmptyLastName
	}
	if rp.Phone == "" {
		return errors.EmptyPhone
	}
	if rp.Address == "" {
		return errors.EmptyAddress
	}

	if err := IsValidEmail(rp.Email); err != nil {
		return err
	}

	if rp.Password == "" {
		return errors.EmptyPassword
	}

	// Cargar ubicaciones válidas
	locations, err := utils.LoadLocations()
	if err != nil {
		return errors2.New("failed to load locations")
	}

	// Validación de la ubicación
	err = utils.ValidateLocation(rp.Location.Department, rp.Location.District, locations)
	if err != nil {
		return errors.ErrInvalidLocation
	}

	switch strings.ToUpper(rp.Profession) {
	case string(domain.Plomero), string(domain.Gasista), string(domain.Electricista):
		return nil
	default:
		return errors.ErrInvalidProfession
	}
}

// IsValidEmail validates if a string is a valid email format.
func IsValidEmail(email string) error {
	if email == "" {
		return errors.EmptyEmail
	}
	// Regex to validate email format
	regex := `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`
	if matched, _ := regexp.MatchString(regex, email); !matched {
		return errors.ErrInvalidEmail
	}

	return nil
}
