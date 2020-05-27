// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ExpandedProject expanded project
// swagger:model ExpandedProject
type ExpandedProject struct {

	// Creation date of the project
	CreationDate string `json:"creationDate,omitempty"`

	// Git remote URI
	GitRemoteURI string `json:"gitRemoteURI,omitempty"`

	// Git User
	GitUser string `json:"gitUser,omitempty"`

	// last event context
	LastEventContext *EventContext `json:"lastEventContext,omitempty"`

	// Project name
	ProjectName string `json:"projectName,omitempty"`

	// Shipyard file content
	Shipyard string `json:"shipyard,omitempty"`

	// stages
	Stages []*ExpandedStage `json:"stages"`
}

// Validate validates this expanded project
func (m *ExpandedProject) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLastEventContext(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStages(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ExpandedProject) validateLastEventContext(formats strfmt.Registry) error {

	if swag.IsZero(m.LastEventContext) { // not required
		return nil
	}

	if m.LastEventContext != nil {
		if err := m.LastEventContext.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("lastEventContext")
			}
			return err
		}
	}

	return nil
}

func (m *ExpandedProject) validateStages(formats strfmt.Registry) error {

	if swag.IsZero(m.Stages) { // not required
		return nil
	}

	for i := 0; i < len(m.Stages); i++ {
		if swag.IsZero(m.Stages[i]) { // not required
			continue
		}

		if m.Stages[i] != nil {
			if err := m.Stages[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("stages" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ExpandedProject) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ExpandedProject) UnmarshalBinary(b []byte) error {
	var res ExpandedProject
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}