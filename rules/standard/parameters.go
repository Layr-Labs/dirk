// Copyright © 2020 Attestant Limited.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package standard

import (
	"errors"

	"github.com/rs/zerolog"
)

type parameters struct {
	logLevel        zerolog.Level
	storagePath     string
	adminIPs        []string
	periodicPruning bool
}

// Parameter is the interface for service parameters.
type Parameter interface {
	apply(p *parameters)
}

type parameterFunc func(*parameters)

func (f parameterFunc) apply(p *parameters) {
	f(p)
}

// WithLogLevel sets the log level for the module.
func WithLogLevel(logLevel zerolog.Level) Parameter {
	return parameterFunc(func(p *parameters) {
		p.logLevel = logLevel
	})
}

// WithStoragePath sets the storage path for the module.
func WithStoragePath(storagePath string) Parameter {
	return parameterFunc(func(p *parameters) {
		p.storagePath = storagePath
	})
}

// WithAdminIPs sets the administration IP addreses for the module.
func WithAdminIPs(adminIPs []string) Parameter {
	return parameterFunc(func(p *parameters) {
		p.adminIPs = adminIPs
	})
}

// WithPeriodicPruning enables periodic pruning of the rules database.
func WithPeriodicPruning(periodicPruning bool) Parameter {
	return parameterFunc(func(p *parameters) {
		p.periodicPruning = periodicPruning
	})
}

// parseAndCheckParameters parses and checks parameters to ensure that mandatory parameters are present and correct.
func parseAndCheckParameters(params ...Parameter) (*parameters, error) {
	parameters := parameters{
		logLevel: zerolog.GlobalLevel(),
	}
	for _, p := range params {
		if params != nil {
			p.apply(&parameters)
		}
	}

	if parameters.storagePath == "" {
		return nil, errors.New("no storage path specified")
	}

	return &parameters, nil
}
