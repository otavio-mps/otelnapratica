// Copyright Dose de Telemetria GmbH
// SPDX-License-Identifier: Apache-2.0

package grpc

import (
	"context"
	"testing"

	"github.com/dosedetelemetria/projeto-otel-na-pratica/api"
	"github.com/dosedetelemetria/projeto-otel-na-pratica/internal/pkg/model"
	"github.com/dosedetelemetria/projeto-otel-na-pratica/internal/pkg/store"
	"github.com/dosedetelemetria/projeto-otel-na-pratica/internal/pkg/store/memory"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var _ api.PlanServiceServer = (*planServer)(nil)

func TestPlanServer_Get(t *testing.T) {
	// prepare
	store := memory.NewPlanStore()
	createTestPlan(t, store)
	srv := NewPlanServer(store)

	// test
	req := &api.GetRequest{Id: "123"}
	resp, err := srv.Get(context.Background(), req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)

	// verify
	assert.Equal(t, req.Id, resp.Plan.Id)
}

func TestPlanServer_Create(t *testing.T) {
	// prepare
	store := memory.NewPlanStore()
	srv := NewPlanServer(store)

	// test
	req := &api.CreateRequest{
		Plan: &api.Plan{
			Id:          "456",
			Name:        "Another Test Plan",
			Description: "This is another test plan",
			Price:       20,
		},
	}
	resp, err := srv.Create(context.Background(), req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, req.Plan.Id, resp.Plan.Id)

	// verify
	plan, err := store.Get(context.Background(), resp.Plan.Id)
	assert.NoError(t, err)
	assert.NotNil(t, plan)
	assert.Equal(t, req.Plan.Id, plan.ID)
}

func TestPlanServer_Update(t *testing.T) {
	// prepare
	store := memory.NewPlanStore()
	createTestPlan(t, store)
	srv := NewPlanServer(store)

	// test
	req := &api.UpdateRequest{
		Plan: &api.Plan{
			Id:          "123",
			Name:        "Updated Test Plan",
			Description: "This is an updated test plan",
			Price:       15,
		},
	}
	resp, err := srv.Update(context.Background(), req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)

	// verify
	assert.Equal(t, req.Plan.Name, resp.Plan.Name)
	plan, err := store.Get(context.Background(), resp.Plan.Id)
	assert.NoError(t, err)
	assert.NotNil(t, plan)
	assert.Equal(t, req.Plan.Name, plan.Name)
}

func TestPlanServer_Delete(t *testing.T) {
	// prepare
	store := memory.NewPlanStore()
	createTestPlan(t, store)
	srv := NewPlanServer(store)

	// test
	req := &api.DeleteRequest{Id: "123"}
	_, err := srv.Delete(context.Background(), req)
	assert.NoError(t, err)

	// verify
	plan, err := store.Get(context.Background(), req.Id)
	assert.Nil(t, err)
	assert.Nil(t, plan)
}

func TestPlanServer_List(t *testing.T) {
	// prepare
	store := memory.NewPlanStore()
	createTestPlan(t, store)
	srv := NewPlanServer(store)

	// test
	req := &api.ListRequest{}
	resp, err := srv.List(context.Background(), req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)

	// verify
	assert.Len(t, resp.Plans, 1)
}

func createTestPlan(t *testing.T, store store.Plan) {
	_, err := store.Create(context.Background(), &model.Plan{
		ID:          "123",
		Name:        "Test Plan",
		Description: "This is a test plan",
		Price:       10,
	})
	require.NoError(t, err)
}
