// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package dialogflow_test

import (
	"context"

	dialogflow "cloud.google.com/go/dialogflow/cx/apiv3beta1"
	"google.golang.org/api/iterator"
	cxpb "google.golang.org/genproto/googleapis/cloud/dialogflow/cx/v3beta1"
)

func ExampleNewFlowsClient() {
	ctx := context.Background()
	c, err := dialogflow.NewFlowsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use client.
	_ = c
}

func ExampleFlowsClient_CreateFlow() {
	// import cxpb "google.golang.org/genproto/googleapis/cloud/dialogflow/cx/v3beta1"

	ctx := context.Background()
	c, err := dialogflow.NewFlowsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &cxpb.CreateFlowRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.CreateFlow(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleFlowsClient_DeleteFlow() {
	ctx := context.Background()
	c, err := dialogflow.NewFlowsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &cxpb.DeleteFlowRequest{
		// TODO: Fill request struct fields.
	}
	err = c.DeleteFlow(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleFlowsClient_ListFlows() {
	// import cxpb "google.golang.org/genproto/googleapis/cloud/dialogflow/cx/v3beta1"
	// import "google.golang.org/api/iterator"

	ctx := context.Background()
	c, err := dialogflow.NewFlowsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &cxpb.ListFlowsRequest{
		// TODO: Fill request struct fields.
	}
	it := c.ListFlows(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleFlowsClient_GetFlow() {
	// import cxpb "google.golang.org/genproto/googleapis/cloud/dialogflow/cx/v3beta1"

	ctx := context.Background()
	c, err := dialogflow.NewFlowsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &cxpb.GetFlowRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.GetFlow(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleFlowsClient_UpdateFlow() {
	// import cxpb "google.golang.org/genproto/googleapis/cloud/dialogflow/cx/v3beta1"

	ctx := context.Background()
	c, err := dialogflow.NewFlowsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &cxpb.UpdateFlowRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.UpdateFlow(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleFlowsClient_TrainFlow() {
	// import cxpb "google.golang.org/genproto/googleapis/cloud/dialogflow/cx/v3beta1"

	ctx := context.Background()
	c, err := dialogflow.NewFlowsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &cxpb.TrainFlowRequest{
		// TODO: Fill request struct fields.
	}
	op, err := c.TrainFlow(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	err = op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
}
