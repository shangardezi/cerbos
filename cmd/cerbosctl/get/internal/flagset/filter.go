// Copyright 2021-2023 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package flagset

import (
	"fmt"

	"github.com/cerbos/cerbos/internal/policy"
)

//nolint:govet
type Filters struct {
	Name            []string `help:"Filter policies by name"`
	NameRegexp      string   `help:"Filter policies by name, using regular expression"`
	Version         []string `help:"Filter policies by version"`
	VersionRegexp   string   `help:"Filter policies by version, using regular expression"`
	ScopeRegexp     string   `help:"Filter policies by scope, using regular expression"`
	IncludeDisabled bool     `help:"Include disabled policies"`
}

func (f Filters) Validate(kind policy.Kind, listing bool) error {
	if !listing && f.IncludeDisabled {
		return fmt.Errorf("--include-disabled is only available when listing")
	}

	if !listing && (len(f.Name) > 0 || len(f.Version) > 0) {
		return fmt.Errorf("--name and --version flags are only available when listing")
	}

	if !listing && (f.NameRegexp != "" || f.VersionRegexp != "" || f.ScopeRegexp != "") {
		return fmt.Errorf("--{name|version|scope}-regexp flags are only available when listing")
	}

	if kind == policy.DerivedRolesKind && (len(f.Version) > 0 || f.VersionRegexp != "") {
		return fmt.Errorf("--version and --version-regexp flags are not available when listing derived roles")
	}

	if len(f.Name) > 0 && f.NameRegexp != "" {
		return fmt.Errorf("--name and --name-regexp flags cannot be used together")
	}

	if len(f.Version) > 0 && f.VersionRegexp != "" {
		return fmt.Errorf("--version and --version-regexp flags cannot be used together")
	}

	return nil
}
