package handlers

import (
	"agendahora/domain/services"
	"fmt"
	"net/http"
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

	// buscar o email anta
	found := h.waitlistService.FindByEmail(r.Context(), email)

	if found.Email == email {
		fmt.Fprintf(w, `<div class="toast-container bg-warning toast-error">email %s. já foi adicionado Tente novamente mais tarde.</div>`, email)
		return
	}

	err = h.waitlistService.Create(r.Context(), email)

	// [] extrair isso em componentes
	// [] resposnvo mobile
	// [] impedir a pessoa de fazer varias requisições travando o botão ou algo assim
	// validação de email
	// [] subida do v1

	if err != nil {
		fmt.Fprintf(w, `<div class="toast-container bg-error toast-error">Falha ao adicionar o email %s. Tente novamente mais tarde.</div>`, email)
		return
	}

	fmt.Fprintf(w, `<div class="toast-container bg-success">Email %s adicionado com sucesso!</div>`, email)
}
