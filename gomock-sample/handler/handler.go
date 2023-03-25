package handler

import "gomock-sample/store"

type Handler struct {
	store store.Store
}

func NewHandler(s store.Store) *Handler {
	return &Handler{store: s}
}

func (h *Handler) GetValue(key string) (string, error) {
	return h.store.Get(key)
}

func (h *Handler) SetValue(key, value string) error {
	return h.store.Set(key, value)
}
