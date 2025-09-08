package handler

import (
	"encoding/json"
	"github.com/Microservices-Gopher-Hex/provider-service/api/dto"
	"github.com/Microservices-Gopher-Hex/provider-service/application/provider"
	"net/http"
	"strconv"
)

type ProviderHandler struct {
	Create *provider.CreateUseCase
	List   *provider.ListUseCase
}

func (h *ProviderHandler) Health(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("ok"))
}

func (h *ProviderHandler) Post(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateProviderRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	out, err := h.Create.Do(r.Context(), req.Name, req.Email)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	resp := dto.ProviderResponse{ID: out.ID, Name: out.Name, Email: out.Email}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(resp)
}

func (h *ProviderHandler) GetList(w http.ResponseWriter, r *http.Request) {
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	offset, _ := strconv.Atoi(r.URL.Query().Get("offset"))
	if limit <= 0 {
		limit = 50
	}
	if offset < 0 {
		offset = 0
	}

	items, err := h.List.Do(r.Context(), int32(limit), int32(offset))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	resp := make([]dto.ProviderResponse, 0, len(items))
	for _, it := range items {
		resp = append(resp, dto.ProviderResponse{ID: it.ID, Name: it.Name, Email: it.Email})
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(resp)
}
