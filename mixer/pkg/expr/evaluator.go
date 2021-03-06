// Copyright 2016 Istio Authors
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

package expr

import (
	pb "istio.io/api/policy/v1beta1"
)

type (
	// TypeChecker validates a given expression for type safety.
	TypeChecker interface {
		// EvalType produces the type of an expression or an error if the type cannot be evaluated.
		// TODO: we probably want to use a golang type rather than pb.ValueType (a proto).
		EvalType(expr string, finder AttributeDescriptorFinder) (pb.ValueType, error)

		// AssertType evaluates the type of expr using the attribute set; if the evaluated type is equal to
		// the expected type we return nil, and return an error otherwise.
		AssertType(expr string, finder AttributeDescriptorFinder, expectedType pb.ValueType) error
	}
)
