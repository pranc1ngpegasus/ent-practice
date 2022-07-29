// Code generated by ent, DO NOT EDIT.

package ogent

import (
	"context"
	"net/http"

	"github.com/Pranc1ngPegasus/ent-practice/ent"
	"github.com/Pranc1ngPegasus/ent-practice/ent/account"
	"github.com/Pranc1ngPegasus/ent-practice/ent/organization"
	"github.com/go-faster/jx"
)

// OgentHandler implements the ogen generated Handler interface and uses Ent as data layer.
type OgentHandler struct {
	client *ent.Client
}

// NewOgentHandler returns a new OgentHandler.
func NewOgentHandler(c *ent.Client) *OgentHandler { return &OgentHandler{c} }

// rawError renders err as json string.
func rawError(err error) jx.Raw {
	var e jx.Encoder
	e.Str(err.Error())
	return e.Bytes()
}

// CreateAccount handles POST /accounts requests.
func (h *OgentHandler) CreateAccount(ctx context.Context, req CreateAccountReq) (CreateAccountRes, error) {
	b := h.client.Account.Create()
	// Add all fields.
	b.SetName(req.Name)
	b.SetCreatedAt(req.CreatedAt)
	b.SetUpdatedAt(req.UpdatedAt)
	if v, ok := req.DeletedAt.Get(); ok {
		b.SetDeletedAt(v)
	}
	// Add all edges.
	// Persist to storage.
	e, err := b.Save(ctx)
	if err != nil {
		switch {
		case ent.IsNotSingular(err):
			return &R409{
				Code:   http.StatusConflict,
				Status: http.StatusText(http.StatusConflict),
				Errors: rawError(err),
			}, nil
		case ent.IsConstraintError(err):
			return &R409{
				Code:   http.StatusConflict,
				Status: http.StatusText(http.StatusConflict),
				Errors: rawError(err),
			}, nil
		default:
			// Let the server handle the error.
			return nil, err
		}
	}
	// Reload the entity to attach all eager-loaded edges.
	q := h.client.Account.Query().Where(account.ID(e.ID))
	e, err = q.Only(ctx)
	if err != nil {
		// This should never happen.
		return nil, err
	}
	return NewAccountCreate(e), nil
}

// ReadAccount handles GET /accounts/{id} requests.
func (h *OgentHandler) ReadAccount(ctx context.Context, params ReadAccountParams) (ReadAccountRes, error) {
	q := h.client.Account.Query().Where(account.IDEQ(params.ID))
	e, err := q.Only(ctx)
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			return &R404{
				Code:   http.StatusNotFound,
				Status: http.StatusText(http.StatusNotFound),
				Errors: rawError(err),
			}, nil
		case ent.IsNotSingular(err):
			return &R409{
				Code:   http.StatusConflict,
				Status: http.StatusText(http.StatusConflict),
				Errors: rawError(err),
			}, nil
		default:
			// Let the server handle the error.
			return nil, err
		}
	}
	return NewAccountRead(e), nil
}

// UpdateAccount handles PATCH /accounts/{id} requests.
func (h *OgentHandler) UpdateAccount(ctx context.Context, req UpdateAccountReq, params UpdateAccountParams) (UpdateAccountRes, error) {
	b := h.client.Account.UpdateOneID(params.ID)
	// Add all fields.
	if v, ok := req.Name.Get(); ok {
		b.SetName(v)
	}
	if v, ok := req.UpdatedAt.Get(); ok {
		b.SetUpdatedAt(v)
	}
	if v, ok := req.DeletedAt.Get(); ok {
		b.SetDeletedAt(v)
	}
	// Add all edges.
	// Persist to storage.
	e, err := b.Save(ctx)
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			return &R404{
				Code:   http.StatusNotFound,
				Status: http.StatusText(http.StatusNotFound),
				Errors: rawError(err),
			}, nil
		case ent.IsConstraintError(err):
			return &R409{
				Code:   http.StatusConflict,
				Status: http.StatusText(http.StatusConflict),
				Errors: rawError(err),
			}, nil
		default:
			// Let the server handle the error.
			return nil, err
		}
	}
	// Reload the entity to attach all eager-loaded edges.
	q := h.client.Account.Query().Where(account.ID(e.ID))
	e, err = q.Only(ctx)
	if err != nil {
		// This should never happen.
		return nil, err
	}
	return NewAccountUpdate(e), nil
}

// DeleteAccount handles DELETE /accounts/{id} requests.
func (h *OgentHandler) DeleteAccount(ctx context.Context, params DeleteAccountParams) (DeleteAccountRes, error) {
	err := h.client.Account.DeleteOneID(params.ID).Exec(ctx)
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			return &R404{
				Code:   http.StatusNotFound,
				Status: http.StatusText(http.StatusNotFound),
				Errors: rawError(err),
			}, nil
		case ent.IsConstraintError(err):
			return &R409{
				Code:   http.StatusConflict,
				Status: http.StatusText(http.StatusConflict),
				Errors: rawError(err),
			}, nil
		default:
			// Let the server handle the error.
			return nil, err
		}
	}
	return new(DeleteAccountNoContent), nil

}

// ListAccount handles GET /accounts requests.
func (h *OgentHandler) ListAccount(ctx context.Context, params ListAccountParams) (ListAccountRes, error) {
	q := h.client.Account.Query()
	page := 1
	if v, ok := params.Page.Get(); ok {
		page = v
	}
	itemsPerPage := 30
	if v, ok := params.ItemsPerPage.Get(); ok {
		itemsPerPage = v
	}
	q.Limit(itemsPerPage).Offset((page - 1) * itemsPerPage)

	es, err := q.All(ctx)
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			return &R404{
				Code:   http.StatusNotFound,
				Status: http.StatusText(http.StatusNotFound),
				Errors: rawError(err),
			}, nil
		case ent.IsNotSingular(err):
			return &R409{
				Code:   http.StatusConflict,
				Status: http.StatusText(http.StatusConflict),
				Errors: rawError(err),
			}, nil
		default:
			// Let the server handle the error.
			return nil, err
		}
	}
	r := NewAccountLists(es)
	return (*ListAccountOKApplicationJSON)(&r), nil
}

// CreateOrganization handles POST /organizations requests.
func (h *OgentHandler) CreateOrganization(ctx context.Context, req CreateOrganizationReq) (CreateOrganizationRes, error) {
	b := h.client.Organization.Create()
	// Add all fields.
	b.SetName(req.Name)
	b.SetCreatedAt(req.CreatedAt)
	b.SetUpdatedAt(req.UpdatedAt)
	if v, ok := req.DeletedAt.Get(); ok {
		b.SetDeletedAt(v)
	}
	// Add all edges.
	// Persist to storage.
	e, err := b.Save(ctx)
	if err != nil {
		switch {
		case ent.IsNotSingular(err):
			return &R409{
				Code:   http.StatusConflict,
				Status: http.StatusText(http.StatusConflict),
				Errors: rawError(err),
			}, nil
		case ent.IsConstraintError(err):
			return &R409{
				Code:   http.StatusConflict,
				Status: http.StatusText(http.StatusConflict),
				Errors: rawError(err),
			}, nil
		default:
			// Let the server handle the error.
			return nil, err
		}
	}
	// Reload the entity to attach all eager-loaded edges.
	q := h.client.Organization.Query().Where(organization.ID(e.ID))
	e, err = q.Only(ctx)
	if err != nil {
		// This should never happen.
		return nil, err
	}
	return NewOrganizationCreate(e), nil
}

// ReadOrganization handles GET /organizations/{id} requests.
func (h *OgentHandler) ReadOrganization(ctx context.Context, params ReadOrganizationParams) (ReadOrganizationRes, error) {
	q := h.client.Organization.Query().Where(organization.IDEQ(params.ID))
	e, err := q.Only(ctx)
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			return &R404{
				Code:   http.StatusNotFound,
				Status: http.StatusText(http.StatusNotFound),
				Errors: rawError(err),
			}, nil
		case ent.IsNotSingular(err):
			return &R409{
				Code:   http.StatusConflict,
				Status: http.StatusText(http.StatusConflict),
				Errors: rawError(err),
			}, nil
		default:
			// Let the server handle the error.
			return nil, err
		}
	}
	return NewOrganizationRead(e), nil
}

// UpdateOrganization handles PATCH /organizations/{id} requests.
func (h *OgentHandler) UpdateOrganization(ctx context.Context, req UpdateOrganizationReq, params UpdateOrganizationParams) (UpdateOrganizationRes, error) {
	b := h.client.Organization.UpdateOneID(params.ID)
	// Add all fields.
	if v, ok := req.Name.Get(); ok {
		b.SetName(v)
	}
	if v, ok := req.UpdatedAt.Get(); ok {
		b.SetUpdatedAt(v)
	}
	if v, ok := req.DeletedAt.Get(); ok {
		b.SetDeletedAt(v)
	}
	// Add all edges.
	// Persist to storage.
	e, err := b.Save(ctx)
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			return &R404{
				Code:   http.StatusNotFound,
				Status: http.StatusText(http.StatusNotFound),
				Errors: rawError(err),
			}, nil
		case ent.IsConstraintError(err):
			return &R409{
				Code:   http.StatusConflict,
				Status: http.StatusText(http.StatusConflict),
				Errors: rawError(err),
			}, nil
		default:
			// Let the server handle the error.
			return nil, err
		}
	}
	// Reload the entity to attach all eager-loaded edges.
	q := h.client.Organization.Query().Where(organization.ID(e.ID))
	e, err = q.Only(ctx)
	if err != nil {
		// This should never happen.
		return nil, err
	}
	return NewOrganizationUpdate(e), nil
}

// DeleteOrganization handles DELETE /organizations/{id} requests.
func (h *OgentHandler) DeleteOrganization(ctx context.Context, params DeleteOrganizationParams) (DeleteOrganizationRes, error) {
	err := h.client.Organization.DeleteOneID(params.ID).Exec(ctx)
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			return &R404{
				Code:   http.StatusNotFound,
				Status: http.StatusText(http.StatusNotFound),
				Errors: rawError(err),
			}, nil
		case ent.IsConstraintError(err):
			return &R409{
				Code:   http.StatusConflict,
				Status: http.StatusText(http.StatusConflict),
				Errors: rawError(err),
			}, nil
		default:
			// Let the server handle the error.
			return nil, err
		}
	}
	return new(DeleteOrganizationNoContent), nil

}

// ListOrganization handles GET /organizations requests.
func (h *OgentHandler) ListOrganization(ctx context.Context, params ListOrganizationParams) (ListOrganizationRes, error) {
	q := h.client.Organization.Query()
	page := 1
	if v, ok := params.Page.Get(); ok {
		page = v
	}
	itemsPerPage := 30
	if v, ok := params.ItemsPerPage.Get(); ok {
		itemsPerPage = v
	}
	q.Limit(itemsPerPage).Offset((page - 1) * itemsPerPage)

	es, err := q.All(ctx)
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			return &R404{
				Code:   http.StatusNotFound,
				Status: http.StatusText(http.StatusNotFound),
				Errors: rawError(err),
			}, nil
		case ent.IsNotSingular(err):
			return &R409{
				Code:   http.StatusConflict,
				Status: http.StatusText(http.StatusConflict),
				Errors: rawError(err),
			}, nil
		default:
			// Let the server handle the error.
			return nil, err
		}
	}
	r := NewOrganizationLists(es)
	return (*ListOrganizationOKApplicationJSON)(&r), nil
}
