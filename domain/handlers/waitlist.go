package handlers

import (
	"agendahora/domain/dto"
	"agendahora/domain/services"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

// o handler pode receber todos os services que quiser
type WaitlistHandler struct {
	waitlistService services.WaitlistService
}

func NewWaitlistHandler(waitlistService services.WaitlistService) *WaitlistHandler {
	return &WaitlistHandler{
		waitlistService: waitlistService,
	}
}

func (h *WaitlistHandler) AddToWaitList(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}

	email := r.FormValue("email")

	validate := validator.New()

	err = validate.Struct(&dto.AddToWaitlistDTO{Email: email})

	if err != nil {
		fmt.Fprintf(w, `<div class="toast-container bg-warning"><p class="text-warning-content">E-mail inválido</p></div>`)
		return
	}

	found := h.waitlistService.FindByEmail(r.Context(), email)

	if found.Email == email {
		fmt.Fprintf(w, `<div class="toast-container bg-warning"><p class="text-warning-content">O e-mail já foi adicionado</p></div>`)
		return
	}

	err = h.waitlistService.Create(r.Context(), email)

	// [] extrair isso em componentes
	// [] resposnvo mobile
	// [] impedir a pessoa de fazer varias requisições travando o botão ou algo assim
	// [] subida do v1

	if err != nil {
		fmt.Fprintf(w, `<div class="toast-container bg-error"><p class="text-error-content">Falha ao adicionar o e-mail %s</p></div>`, email)
		return
	}

	fmt.Fprintf(w, `<div class="toast-container bg-success"><p class="text-success-content">%s adicionado a lista de espera</p></div>`, email)
}
