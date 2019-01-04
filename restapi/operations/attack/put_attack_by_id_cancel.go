// Code generated by go-swagger; DO NOT EDIT.

package attack

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// PutAttackByIDCancelHandlerFunc turns a function with the right signature into a put attack by ID cancel handler
type PutAttackByIDCancelHandlerFunc func(PutAttackByIDCancelParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PutAttackByIDCancelHandlerFunc) Handle(params PutAttackByIDCancelParams) middleware.Responder {
	return fn(params)
}

// PutAttackByIDCancelHandler interface for that can handle valid put attack by ID cancel params
type PutAttackByIDCancelHandler interface {
	Handle(PutAttackByIDCancelParams) middleware.Responder
}

// NewPutAttackByIDCancel creates a new http.Handler for the put attack by ID cancel operation
func NewPutAttackByIDCancel(ctx *middleware.Context, handler PutAttackByIDCancelHandler) *PutAttackByIDCancel {
	return &PutAttackByIDCancel{Context: ctx, Handler: handler}
}

/*PutAttackByIDCancel swagger:route PUT /attack/{attackID}/cancel attack putAttackByIdCancel

Cancel an attack by ID

*/
type PutAttackByIDCancel struct {
	Context *middleware.Context
	Handler PutAttackByIDCancelHandler
}

func (o *PutAttackByIDCancel) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPutAttackByIDCancelParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
