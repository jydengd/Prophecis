// Code generated by go-swagger; DO NOT EDIT.

package groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetNamespaceByGroupIDAndNamespaceHandlerFunc turns a function with the right signature into a get namespace by group Id and namespace handler
type GetNamespaceByGroupIDAndNamespaceHandlerFunc func(GetNamespaceByGroupIDAndNamespaceParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetNamespaceByGroupIDAndNamespaceHandlerFunc) Handle(params GetNamespaceByGroupIDAndNamespaceParams) middleware.Responder {
	return fn(params)
}

// GetNamespaceByGroupIDAndNamespaceHandler interface for that can handle valid get namespace by group Id and namespace params
type GetNamespaceByGroupIDAndNamespaceHandler interface {
	Handle(GetNamespaceByGroupIDAndNamespaceParams) middleware.Responder
}

// NewGetNamespaceByGroupIDAndNamespace creates a new http.Handler for the get namespace by group Id and namespace operation
func NewGetNamespaceByGroupIDAndNamespace(ctx *middleware.Context, handler GetNamespaceByGroupIDAndNamespaceHandler) *GetNamespaceByGroupIDAndNamespace {
	return &GetNamespaceByGroupIDAndNamespace{Context: ctx, Handler: handler}
}

/*GetNamespaceByGroupIDAndNamespace swagger:route GET /cc/v1/groups/{groupId}/namespace/{namespaceId} Groups getNamespaceByGroupIdAndNamespace

get namespace by groupId and namespaceId.

Optional extended description in Markdown.

*/
type GetNamespaceByGroupIDAndNamespace struct {
	Context *middleware.Context
	Handler GetNamespaceByGroupIDAndNamespaceHandler
}

func (o *GetNamespaceByGroupIDAndNamespace) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetNamespaceByGroupIDAndNamespaceParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
