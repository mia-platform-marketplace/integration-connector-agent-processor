// Copyright Mia srl
// SPDX-License-Identifier: Apache-2.0
//
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

package main

import (
	rpcprocessor "github.com/mia-platform/integration-connector-agent/adapters/rpc-processor"
	"github.com/mia-platform/integration-connector-agent/entities"
)

type CustomProcessor struct {
	logger rpcprocessor.Logger
	config []byte
}

func (g *CustomProcessor) Process(input entities.PipelineEvent) (entities.PipelineEvent, error) {
	output := input.Clone()
	output.WithData([]byte(`{"data":"processed by CustomProcessor"}`))

	g.logger.Trace("CustomProcessor sucessfully processed the input event")
	return output, nil
}

func (g *CustomProcessor) Init(config []byte) error {
	g.config = config

	g.logger.Info("CustomProcessor initialized")
	return nil
}
