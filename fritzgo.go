// Package fritzgo contains CLI tools to access fritz data.
package fritzgo

import (
	"context"
	"fmt"

	"github.com/antiphp/fritzgo/pkg/fritztypes"
	"github.com/hamba/logger/v2"
)

// Client represents an HTTP client to access fritz data.
type Client interface {
	ListUsers(context.Context) ([]fritztypes.User, error)
	Info(context.Context) (fritztypes.Info, error)
}

// Renderer render fritz data.
type Renderer interface {
	ListUsers([]fritztypes.User) error
	Info(fritztypes.Info) error
}

// FritzGo is the core application to retrieve, manage and render fritz data.
type FritzGo struct {
	fritzBox Client
	render   Renderer
	log      *logger.Logger
}

// New returns a new application.
func New(client Client, render Renderer) *FritzGo {
	return &FritzGo{
		fritzBox: client,
		render:   render,
	}
}

// ListUsers retrieves and renders fritz users.
func (f *FritzGo) ListUsers(ctx context.Context) error {
	users, err := f.fritzBox.ListUsers(ctx)
	if err != nil {
		return fmt.Errorf("getting users: %w", err)
	}

	if err = f.render.ListUsers(users); err != nil {
		return fmt.Errorf("rendering users: %w", err)
	}

	return nil
}

// Info retrieves and renders basic information.
func (f *FritzGo) Info(ctx context.Context) error {
	info, err := f.fritzBox.Info(ctx)
	if err != nil {
		return fmt.Errorf("retrieving info: %w", err)
	}

	if err = f.render.Info(info); err != nil {
		return fmt.Errorf("retrieving info: %w", err)
	}

	return nil
}
