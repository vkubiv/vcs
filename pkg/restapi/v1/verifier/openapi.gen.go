// Package verifier provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.11.0 DO NOT EDIT.
package verifier

import (
	"fmt"
	"net/http"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/labstack/echo/v4"
)

// Defines values for VerifierChecksCredentialFormat.
const (
	JwtVc VerifierChecksCredentialFormat = "jwt_vc"
	LdpVc VerifierChecksCredentialFormat = "ldp_vc"
)

// Defines values for VerifierChecksPresentationFormat.
const (
	JwtVp VerifierChecksPresentationFormat = "jwt_vp"
	LdpVp VerifierChecksPresentationFormat = "ldp_vp"
)

// Model for creating verifier profile.
type CreateVerifierProfileData struct {
	// Type of checks to be performed and formats supported.
	Checks map[string]interface{} `json:"checks"`

	// Verifier’s display name.
	Name string `json:"name"`

	// Configuration for participating in OIDC4VC credential interaction operations.
	OidcConfig *map[string]interface{} `json:"oidcConfig,omitempty"`

	// Unique identifier of the organization.
	OrganizationID string `json:"organizationID"`

	// URI of the verifier.
	Url *string `json:"url,omitempty"`
}

// Model for updating verifier profile data.
type UpdateVerifierProfileData struct {
	// Type of checks to be performed and formats supported.
	Checks *map[string]interface{} `json:"checks,omitempty"`

	// Verifier’s display name.
	Name *string `json:"name,omitempty"`

	// Configuration for participating in OIDC4VC credential interaction operations.
	OidcConfig *map[string]interface{} `json:"oidcConfig,omitempty"`

	// URI of the verifier.
	Url *string `json:"url,omitempty"`
}

// Checks to be performed by a verifier profile for verifying credentials and presentations.
type VerifierChecks struct {
	// Checks to be performed during credential verification.
	Credential struct {
		// Supported credential formats.
		Format []VerifierChecksCredentialFormat `json:"format"`

		// Proof check for credential.
		Proof bool `json:"proof"`

		// Status check for credential.
		Status *bool `json:"status,omitempty"`
	} `json:"credential"`

	// Checks to be performed during presentation verification.
	Presentation struct {
		// Supported presentation formats.
		Format []VerifierChecksPresentationFormat `json:"format"`

		// Proof check for presentation.
		Proof bool `json:"proof"`
	} `json:"presentation"`
}

// VerifierChecksCredentialFormat defines model for VerifierChecks.Credential.Format.
type VerifierChecksCredentialFormat string

// VerifierChecksPresentationFormat defines model for VerifierChecks.Presentation.Format.
type VerifierChecksPresentationFormat string

// Model for verifier profile.
type VerifierProfile struct {
	// Defines if profile is enabled.
	Active bool `json:"active"`

	// Checks to be performed by a verifier profile for verifying credentials and presentations.
	Checks VerifierChecks `json:"checks"`

	// Short unique string across the VCS platform, to be used as a reference to this profile.
	Id string `json:"id"`

	// Verifier’s display name.
	Name string `json:"name"`

	// Configuration for OIDC4VC credential interaction operations.
	OidcConfig *map[string]interface{} `json:"oidcConfig,omitempty"`

	// Unique identifier of the organization.
	OrganizationID string `json:"organizationID"`

	// URI of the verifier.
	Url *string `json:"url,omitempty"`
}

// PostVerifierProfilesJSONBody defines parameters for PostVerifierProfiles.
type PostVerifierProfilesJSONBody = CreateVerifierProfileData

// PutVerifierProfilesProfileIDJSONBody defines parameters for PutVerifierProfilesProfileID.
type PutVerifierProfilesProfileIDJSONBody = UpdateVerifierProfileData

// PostVerifierProfilesJSONRequestBody defines body for PostVerifierProfiles for application/json ContentType.
type PostVerifierProfilesJSONRequestBody = PostVerifierProfilesJSONBody

// PutVerifierProfilesProfileIDJSONRequestBody defines body for PutVerifierProfilesProfileID for application/json ContentType.
type PutVerifierProfilesProfileIDJSONRequestBody = PutVerifierProfilesProfileIDJSONBody

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Get All Profiles
	// (GET /verifier/profiles)
	GetVerifierProfiles(ctx echo.Context) error
	// Create Profile
	// (POST /verifier/profiles)
	PostVerifierProfiles(ctx echo.Context) error
	// Delete Profile
	// (DELETE /verifier/profiles/{profileID})
	DeleteVerifierProfilesProfileID(ctx echo.Context, profileID string) error
	// Get Profile
	// (GET /verifier/profiles/{profileID})
	GetVerifierProfilesProfileID(ctx echo.Context, profileID string) error
	// Update Profile
	// (PUT /verifier/profiles/{profileID})
	PutVerifierProfilesProfileID(ctx echo.Context, profileID string) error
	// Activate Profile
	// (POST /verifier/profiles/{profileID}/activate)
	PostVerifierProfilesProfileIDActivate(ctx echo.Context, profileID string) error
	// Deactivate Profile
	// (POST /verifier/profiles/{profileID}/deactivate)
	PostVerifierProfilesProfileIDDeactivate(ctx echo.Context, profileID string) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GetVerifierProfiles converts echo context to params.
func (w *ServerInterfaceWrapper) GetVerifierProfiles(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetVerifierProfiles(ctx)
	return err
}

// PostVerifierProfiles converts echo context to params.
func (w *ServerInterfaceWrapper) PostVerifierProfiles(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PostVerifierProfiles(ctx)
	return err
}

// DeleteVerifierProfilesProfileID converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteVerifierProfilesProfileID(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "profileID" -------------
	var profileID string

	err = runtime.BindStyledParameterWithLocation("simple", false, "profileID", runtime.ParamLocationPath, ctx.Param("profileID"), &profileID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter profileID: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.DeleteVerifierProfilesProfileID(ctx, profileID)
	return err
}

// GetVerifierProfilesProfileID converts echo context to params.
func (w *ServerInterfaceWrapper) GetVerifierProfilesProfileID(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "profileID" -------------
	var profileID string

	err = runtime.BindStyledParameterWithLocation("simple", false, "profileID", runtime.ParamLocationPath, ctx.Param("profileID"), &profileID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter profileID: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetVerifierProfilesProfileID(ctx, profileID)
	return err
}

// PutVerifierProfilesProfileID converts echo context to params.
func (w *ServerInterfaceWrapper) PutVerifierProfilesProfileID(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "profileID" -------------
	var profileID string

	err = runtime.BindStyledParameterWithLocation("simple", false, "profileID", runtime.ParamLocationPath, ctx.Param("profileID"), &profileID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter profileID: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PutVerifierProfilesProfileID(ctx, profileID)
	return err
}

// PostVerifierProfilesProfileIDActivate converts echo context to params.
func (w *ServerInterfaceWrapper) PostVerifierProfilesProfileIDActivate(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "profileID" -------------
	var profileID string

	err = runtime.BindStyledParameterWithLocation("simple", false, "profileID", runtime.ParamLocationPath, ctx.Param("profileID"), &profileID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter profileID: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PostVerifierProfilesProfileIDActivate(ctx, profileID)
	return err
}

// PostVerifierProfilesProfileIDDeactivate converts echo context to params.
func (w *ServerInterfaceWrapper) PostVerifierProfilesProfileIDDeactivate(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "profileID" -------------
	var profileID string

	err = runtime.BindStyledParameterWithLocation("simple", false, "profileID", runtime.ParamLocationPath, ctx.Param("profileID"), &profileID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter profileID: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PostVerifierProfilesProfileIDDeactivate(ctx, profileID)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET(baseURL+"/verifier/profiles", wrapper.GetVerifierProfiles)
	router.POST(baseURL+"/verifier/profiles", wrapper.PostVerifierProfiles)
	router.DELETE(baseURL+"/verifier/profiles/:profileID", wrapper.DeleteVerifierProfilesProfileID)
	router.GET(baseURL+"/verifier/profiles/:profileID", wrapper.GetVerifierProfilesProfileID)
	router.PUT(baseURL+"/verifier/profiles/:profileID", wrapper.PutVerifierProfilesProfileID)
	router.POST(baseURL+"/verifier/profiles/:profileID/activate", wrapper.PostVerifierProfilesProfileIDActivate)
	router.POST(baseURL+"/verifier/profiles/:profileID/deactivate", wrapper.PostVerifierProfilesProfileIDDeactivate)

}
