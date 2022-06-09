package handler

import (
	"context"
	"errors"
	"fmt"
	"html/template"
	"log"

	"github.com/DanielTitkov/lentils/internal/domain"

	"github.com/gorilla/mux"
	"github.com/jfyne/live"
)

const (
	// events
	// params
	paramTestCode = "testCode"
	// params values
)

type (
	TestInstance struct {
		*CommonInstance
		Test *domain.Test
	}
)

func (h *Handler) NewTestInstance(s live.Socket) *TestInstance {
	m, ok := s.Assigns().(*TestInstance)
	if !ok {
		return &TestInstance{
			CommonInstance: h.NewCommon(s, viewTest),
		}
	}

	return m
}

func (h *Handler) Test() live.Handler {
	t, err := template.ParseFiles(
		h.t + "base.layout.html",
		// h.t+"page.challenge_details.html",
		// h.t+"part.challenge_details_scale.html",
	)
	if err != nil {
		log.Fatal(err)
	}

	lvh := live.NewHandler(live.WithTemplateRenderer(t))
	// COMMON BLOCK START
	// this logic must be present in all handlers
	{
		constructor := h.NewTestInstance // NB: make sure constructor is correct
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
		// SAFE TO COPY END
	}
	// COMMON BLOCK END

	lvh.HandleMount(func(ctx context.Context, s live.Socket) (interface{}, error) {
		r := live.Request(ctx)
		testCode, ok := mux.Vars(r)[paramTestCode]
		if !ok {
			return nil, errors.New("test code is required")
		}
		fmt.Println("TEST CODE", testCode)

		instance := h.NewTestInstance(s)
		instance.fromContext(ctx)
		// challenge, err := h.app.GetChallengeByID(ctx, challengeID, instance.UserID)
		// if err != nil {
		// 	instance.Error = err
		// }
		// instance.Challenge = challenge

		return instance, nil
	})

	return lvh
}
