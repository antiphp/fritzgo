package fritzgo

import (
	"context"
	"fmt"
	"os"

	"github.com/antiphp/fritzgo/internal/render"
	"github.com/antiphp/fritzgo/pkg/fritzclient"
	"github.com/antiphp/fritzgo/pkg/fritztypes"
	"github.com/hamba/logger/v2"
)

type Client interface {
	ListUsers(context.Context) ([]fritztypes.User, error)
}

type FritzGo struct {
	fritzBox Client
	log      *logger.Logger
}

func New(addr, user, pass string, log *logger.Logger) (*FritzGo, error) {
	client, err := fritzclient.New(addr, user, pass)
	if err != nil {
		return nil, err
	}

	return &FritzGo{
		fritzBox: client,
		log:      log,
	}, nil
}

func (f *FritzGo) ListUsers(ctx context.Context) error {
	users, err := f.fritzBox.ListUsers(ctx)
	if err != nil {
		return fmt.Errorf("getting users: %w", err)
	}

	if err = render.UsersList(os.Stdout, users); err != nil {
		return fmt.Errorf("rendering users: %w", err)
	}

	return nil
}