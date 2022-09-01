/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

//go:generate oapi-codegen --config=openapi.cfg.yaml ../../../../api/openapi.yaml
//go:generate mockgen -destination controller_mocks_test.go -self_package mocks -package verifier_test -source=controller.go -mock_names profileService=MockProfileService

package verifier

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"

	"github.com/trustbloc/vcs/pkg/verifier"
)

var _ ServerInterface = (*Controller)(nil) // make sure Controller implements ServerInterface

type profileService interface {
	Create(profile *verifier.Profile) (*verifier.Profile, error)
	Update(profile *verifier.ProfileUpdate) (*verifier.Profile, error)
	Delete(profileID verifier.ProfileID) error
	GetProfile(profileID verifier.ProfileID) (*verifier.Profile, error)
	ActivateProfile(profileID verifier.ProfileID) error
	DeactivateProfile(profileID verifier.ProfileID) error
	GetAllProfiles(orgID string) ([]*verifier.Profile, error)
}

// Controller for Verifier Profile Management API.
type Controller struct {
	profileSvc profileService
}

// NewController creates a new controller for Verifier Profile Management API.
func NewController(profileSvc profileService) *Controller {
	return &Controller{
		profileSvc: profileSvc,
	}
}

// GetVerifierProfiles gets all verifier profiles for organization.
// GET /verifier/profiles.
func (c *Controller) GetVerifierProfiles(ctx echo.Context) error {
	// TODO: resolve orgID from auth token
	authHeader := ctx.Request().Header.Get("Authorization")
	if authHeader == "" || !strings.Contains(authHeader, "Bearer") {
		return echo.NewHTTPError(http.StatusUnauthorized, "missing authorization")
	}

	orgID := authHeader[len("Bearer "):] // for now assume that token is just plain orgID

	profiles, err := c.profileSvc.GetAllProfiles(orgID)
	if err != nil {
		return fmt.Errorf("failed to get verifier profiles: %w", err)
	}

	var verifierProfiles []*VerifierProfile

	for _, profile := range profiles {
		var vp VerifierProfile
		copier.Copy(&vp, profile) //nolint:errcheck

		verifierProfiles = append(verifierProfiles, &vp)
	}

	return ctx.JSON(http.StatusOK, verifierProfiles)
}

// PostVerifierProfiles creates a new verifier profile.
// POST /verifier/profiles.
func (c *Controller) PostVerifierProfiles(ctx echo.Context) error {
	var body CreateVerifierProfileData

	if err := ctx.Bind(&body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	var p verifier.Profile
	copier.Copy(&p, &body) //nolint:errcheck

	createdProfile, err := c.profileSvc.Create(&p)
	if err != nil {
		return fmt.Errorf("failed to create verifier profile: %w", err)
	}

	var vp VerifierProfile
	copier.Copy(&vp, createdProfile) //nolint:errcheck

	return ctx.JSON(http.StatusOK, &vp)
}

// DeleteVerifierProfilesProfileID deletes profile from VCS storage.
// DELETE /verifier/profiles/{profileID}.
func (c *Controller) DeleteVerifierProfilesProfileID(_ echo.Context, profileID string) error {
	if err := c.profileSvc.Delete(profileID); err != nil {
		return fmt.Errorf("failed to delete verifier profile: %w", err)
	}

	return nil
}

// GetVerifierProfilesProfileID gets profile by ID.
// GET /verifier/profiles/{profileID}.
func (c *Controller) GetVerifierProfilesProfileID(ctx echo.Context, profileID string) error {
	profile, err := c.profileSvc.GetProfile(profileID)
	if err != nil {
		return fmt.Errorf("failed to get verifier profile: %w", err)
	}

	var vp VerifierProfile
	copier.Copy(&vp, profile) //nolint:errcheck

	return ctx.JSON(http.StatusOK, &vp)
}

// PutVerifierProfilesProfileID updates profile.
// PUT /verifier/profiles/{profileID}.
func (c *Controller) PutVerifierProfilesProfileID(ctx echo.Context, profileID string) error {
	var body UpdateVerifierProfileData

	if err := ctx.Bind(&body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	var profileUpdate verifier.ProfileUpdate
	copier.Copy(&profileUpdate, &body) //nolint:errcheck

	profileUpdate.ID = profileID

	updatedProfile, err := c.profileSvc.Update(&profileUpdate)
	if err != nil {
		return fmt.Errorf("failed to update verifier profile: %w", err)
	}

	var vp VerifierProfile
	copier.Copy(&vp, updatedProfile) //nolint:errcheck

	return ctx.JSON(http.StatusOK, &vp)
}

// PostVerifierProfilesProfileIDActivate activates profile.
// POST /verifier/profiles/{profileID}/activate.
func (c *Controller) PostVerifierProfilesProfileIDActivate(_ echo.Context, profileID string) error {
	if err := c.profileSvc.ActivateProfile(profileID); err != nil {
		return fmt.Errorf("failed to activate verifier profile: %w", err)
	}

	return nil
}

// PostVerifierProfilesProfileIDDeactivate deactivates profile.
// POST /verifier/profiles/{profileID}/deactivate.
func (c *Controller) PostVerifierProfilesProfileIDDeactivate(_ echo.Context, profileID string) error {
	if err := c.profileSvc.DeactivateProfile(profileID); err != nil {
		return fmt.Errorf("failed to deactivate verifier profile: %w", err)
	}

	return nil
}