package api

import "github.com/wyw14/solo-036/internal/service"

type Handler struct{ runtime *service.Runtime }

func NewHandler(runtime *service.Runtime) *Handler { return &Handler{runtime: runtime} }
