package charger

// Code generated by github.com/evcc-io/evcc/cmd/tools/decorate.go. DO NOT EDIT.

import (
	"github.com/evcc-io/evcc/api"
)

func decorateCustom(base *Charger, identifier func() (string, error), phaseSwitcher func(int) error) api.Charger {
	switch {
	case identifier == nil && phaseSwitcher == nil:
		return base

	case identifier != nil && phaseSwitcher == nil:
		return &struct {
			*Charger
			api.Identifier
		}{
			Charger: base,
			Identifier: &decorateCustomIdentifierImpl{
				identifier: identifier,
			},
		}

	case identifier == nil && phaseSwitcher != nil:
		return &struct {
			*Charger
			api.PhaseSwitcher
		}{
			Charger: base,
			PhaseSwitcher: &decorateCustomPhaseSwitcherImpl{
				phaseSwitcher: phaseSwitcher,
			},
		}

	case identifier != nil && phaseSwitcher != nil:
		return &struct {
			*Charger
			api.Identifier
			api.PhaseSwitcher
		}{
			Charger: base,
			Identifier: &decorateCustomIdentifierImpl{
				identifier: identifier,
			},
			PhaseSwitcher: &decorateCustomPhaseSwitcherImpl{
				phaseSwitcher: phaseSwitcher,
			},
		}
	}

	return nil
}

type decorateCustomIdentifierImpl struct {
	identifier func() (string, error)
}

func (impl *decorateCustomIdentifierImpl) Identify() (string, error) {
	return impl.identifier()
}

type decorateCustomPhaseSwitcherImpl struct {
	phaseSwitcher func(int) error
}

func (impl *decorateCustomPhaseSwitcherImpl) Phases1p3p(phases int) error {
	return impl.phaseSwitcher(phases)
}