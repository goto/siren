package provider

import (
	"context"

	"github.com/goto/siren/pkg/errors"
)

// Service handles business logic
type Service struct {
	supportedTypes []string
	repository     Repository
}

// NewService returns repository struct
func NewService(repository Repository, supportedProviders []string) *Service {
	return &Service{
		repository:     repository,
		supportedTypes: supportedProviders,
	}
}

func (s *Service) List(ctx context.Context, flt Filter) ([]Provider, error) {
	return s.repository.List(ctx, flt)
}

func (s *Service) Create(ctx context.Context, prov *Provider) error {
	if prov == nil {
		return errors.ErrInvalid.WithMsgf("provider is nil")
	}

	if !s.isTypeSupported(prov.Type) {
		return errors.ErrInvalid.WithMsgf("type %s is not supported, supported types are: %+v", prov.Type, s.supportedTypes)
	}

	err := s.repository.Create(ctx, prov)
	if err != nil {
		if errors.Is(err, ErrDuplicate) {
			return errors.ErrConflict.WithMsgf(err.Error())
		}
		return err
	}

	return nil
}

func (s *Service) Get(ctx context.Context, id uint64) (*Provider, error) {
	prov, err := s.repository.Get(ctx, id)
	if err != nil {
		if errors.As(err, new(NotFoundError)) {
			return nil, errors.ErrNotFound.WithMsgf(err.Error())
		}
		return nil, err
	}
	return prov, nil
}

func (s *Service) Update(ctx context.Context, prov *Provider) error {
	if prov == nil {
		return errors.ErrInvalid.WithMsgf("provider is nil")
	}

	if !s.isTypeSupported(prov.Type) {
		return errors.ErrInvalid.WithMsgf("type %s is not supported, supported types are: %+v", prov.Type, s.supportedTypes)
	}

	err := s.repository.Update(ctx, prov)
	if err != nil {
		if errors.Is(err, ErrDuplicate) {
			return errors.ErrConflict.WithMsgf(err.Error())
		}
		if errors.As(err, new(NotFoundError)) {
			return errors.ErrNotFound.WithMsgf(err.Error())
		}
		return err
	}

	return nil
}

func (s *Service) Delete(ctx context.Context, id uint64) error {
	return s.repository.Delete(ctx, id)
}

func (s *Service) isTypeSupported(providerType string) bool {
	for _, st := range s.supportedTypes {
		if st == providerType {
			return true
		}
	}
	return false
}
