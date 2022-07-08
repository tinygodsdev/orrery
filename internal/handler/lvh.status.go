package handler

import (
	"context"
	"html/template"
	"log"

	"github.com/jfyne/live"
)

type (
	StatusInstance struct {
		*CommonInstance
		Errors []error
	}
)

func (ins *StatusInstance) withError(err error) *StatusInstance {
	ins.Error = err
	return ins
}

func (h *Handler) NewStatusInstance(s live.Socket) *StatusInstance {
	m, ok := s.Assigns().(*StatusInstance)
	if !ok {
		return &StatusInstance{
			CommonInstance: h.NewCommon(s, viewStatus),
			Errors:         nil,
		}
	}

	return m
}

func (h *Handler) Status() live.Handler {
	t, err := template.ParseFiles(
		h.t+"base.layout.html",
		h.t+"page.status.html",
	)
	if err != nil {
		log.Fatal(err)
	}

	lvh := live.NewHandler(live.WithTemplateRenderer(t))
	// COMMON BLOCK START
	// this logic must be present in all handlers
	{
		constructor := h.NewStatusInstance // NB: make sure constructor is correct
		// SAFE TO COPY
		lvh.HandleEvent(eventCloseAuthModals, func(ctx context.Context, s live.Socket, p live.Params) (interface{}, error) {
			instance := constructor(s)
			instance.CloseAuthModals()
			return instance, nil
		})

		lvh.HandleEvent(eventOpenLogoutModal, func(ctx context.Context, s live.Socket, p live.Params) (interface{}, error) {
			instance := constructor(s)
			instance.OpenLogoutModal()
			return instance, nil
		})

		lvh.HandleEvent(eventOpenLoginModal, func(ctx context.Context, s live.Socket, p live.Params) (interface{}, error) {
			instance := constructor(s)
			instance.OpenLoginModal()
			return instance, nil
		})

		lvh.HandleEvent(eventCloseError, func(ctx context.Context, s live.Socket, p live.Params) (interface{}, error) {
			instance := constructor(s)
			instance.CloseError()
			return instance, nil
		})

		lvh.HandleEvent(eventCloseMessage, func(ctx context.Context, s live.Socket, p live.Params) (interface{}, error) {
			instance := constructor(s)
			instance.CloseMessage()
			return instance, nil
		})

		lvh.HandleEvent(eventToggleDark, func(ctx context.Context, s live.Socket, p live.Params) (i interface{}, err error) {
			instance := constructor(s)
			if instance.User != nil {
				instance.User.UseDarkTheme = !instance.User.UseDarkTheme
			}
			instance.User, err = h.app.UpdateUser(ctx, instance.User)
			if err != nil {
				return instance.withError(err), nil
			}
			return instance, nil
		})
		// SAFE TO COPY END
	}
	// COMMON BLOCK END

	lvh.HandleMount(func(ctx context.Context, s live.Socket) (interface{}, error) {
		instance := h.NewStatusInstance(s)
		instance.fromContext(ctx)

		instance.Errors = h.app.Errors

		return instance, nil
	})

	return lvh
}