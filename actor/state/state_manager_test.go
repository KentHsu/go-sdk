/*
Copyright 2021 The Dapr Authors
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package state

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/dapr/go-sdk/actor/mock"
	"github.com/dapr/go-sdk/actor/mock_client"
	"github.com/dapr/go-sdk/client"
	"github.com/golang/mock/gomock"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	testState = "state"
	testValue = "value"
	testTTL   = time.Second * 3600
)

func newMockStateManager(t *testing.T) *stateManagerCtx {
	ctrl := gomock.NewController(t)
	return &stateManagerCtx{
		actorTypeName: "test",
		actorID:       "fn",
		stateAsyncProvider: &DaprStateAsyncProvider{
			daprClient:      mock_client.NewMockClient(ctrl),
			stateSerializer: mock.NewMockCodec(ctrl),
		},
	}
}

func newGetActorStateRequest(sm *stateManagerCtx, key string) *client.GetActorStateRequest {
	return &client.GetActorStateRequest{
		ActorType: sm.actorTypeName,
		ActorID:   sm.actorID,
		KeyName:   key,
	}
}

func newGetActorStateResponse(data []byte) *client.GetActorStateResponse {
	return &client.GetActorStateResponse{Data: data}
}

func TestAdd_EmptyStateName(t *testing.T) {
	ctx := context.Background()
	sm := newMockStateManager(t)
	err := sm.Add(ctx, "", testValue)
	assert.Error(t, err)
}

func TestAdd_WithCachedStateChange(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name      string
		kind      ChangeKind
		shouldErr bool
	}{
		{"state change kind None", None, true},
		{"state change kind Add", Add, true},
		{"state change kind Update", Update, true},
		{"state change kind Remove", Remove, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sm := newMockStateManager(t)
			sm.stateChangeTracker.Store(testState, &ChangeMetadata{Kind: tt.kind, Value: testValue})
			mockClient := sm.stateAsyncProvider.daprClient.(*mock_client.MockClient)
			mockRequest := newGetActorStateRequest(sm, testState)
			mockResult := newGetActorStateResponse([]byte("result"))
			mockClient.EXPECT().GetActorState(ctx, mockRequest).Return(mockResult, nil)

			err := sm.Add(ctx, testState, testValue)
			if tt.shouldErr {
				assert.Error(t, err)
			} else {
				require.NoError(t, err)

				val, ok := sm.stateChangeTracker.Load(testState)
				require.True(t, ok)

				metadata := val.(*ChangeMetadata)
				assert.Equal(t, metadata.Kind, Update)
				assert.Equal(t, metadata.Value, testValue)
			}
		})
	}
}

func TestAdd_WithoutCachedStateChange(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name             string
		stateProviderErr bool
		duplicateState   bool
	}{
		{"state provider returns error", true, false},
		{"state provider returns data", false, true},
		{"successfully add new state", false, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sm := newMockStateManager(t)
			mockClient := sm.stateAsyncProvider.daprClient.(*mock_client.MockClient)
			mockRequest := newGetActorStateRequest(sm, testState)
			mockResult := newGetActorStateResponse([]byte("result"))
			if tt.stateProviderErr {
				mockClient.EXPECT().GetActorState(ctx, mockRequest).Return(nil, errors.New("mockErr"))
			} else {
				if tt.duplicateState {
					mockClient.EXPECT().GetActorState(ctx, mockRequest).Return(mockResult, nil)
				} else {
					mockClient.EXPECT().GetActorState(ctx, mockRequest).Return(nil, nil)
				}
			}

			err := sm.Add(ctx, testState, testValue)
			if tt.stateProviderErr || tt.duplicateState {
				assert.Error(t, err)
			} else {
				require.NoError(t, err)

				val, ok := sm.stateChangeTracker.Load(testState)
				require.True(t, ok)

				metadata := val.(*ChangeMetadata)
				assert.Equal(t, metadata.Kind, Add)
				assert.Equal(t, metadata.Value, testValue)
			}
		})
	}
}

func TestGet_EmptyStateName(t *testing.T) {
	ctx := context.Background()
	sm := newMockStateManager(t)
	err := sm.Get(ctx, "", testValue)
	assert.Error(t, err)
}

func TestGet_WithCachedStateChange(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name      string
		kind      ChangeKind
		shouldErr bool
	}{
		{"state change kind None", None, false},
		{"state change kind Add", Add, false},
		{"state change kind Update", Update, false},
		{"state change kind Remove", Remove, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sm := newMockStateManager(t)
			sm.stateChangeTracker.Store(testState, &ChangeMetadata{Kind: tt.kind, Value: testValue})

			var reply string
			err := sm.Get(ctx, testState, &reply)
			if tt.shouldErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, reply, testValue)
			}
		})
	}
}

func TestGet_WithoutCachedStateChange(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name      string
		shouldErr bool
	}{
		{"state provider returns error", true},
		{"state provider returns data", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sm := newMockStateManager(t)
			mockClient := sm.stateAsyncProvider.daprClient.(*mock_client.MockClient)
			mockCodec := sm.stateAsyncProvider.stateSerializer.(*mock.MockCodec)
			mockRequest := newGetActorStateRequest(sm, testState)
			mockResult := newGetActorStateResponse([]byte("result"))
			if tt.shouldErr {
				mockClient.EXPECT().GetActorState(ctx, mockRequest).Return(nil, errors.New("mockErr"))
			} else {
				mockClient.EXPECT().GetActorState(ctx, mockRequest).Return(mockResult, nil)
				mockCodec.EXPECT().Unmarshal(mockResult.Data, testValue)
			}

			err := sm.Get(ctx, testState, testValue)
			if tt.shouldErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}

			val, ok := sm.stateChangeTracker.Load(testState)
			require.True(t, ok)

			metadata := val.(*ChangeMetadata)
			assert.Equal(t, metadata.Kind, None)
			assert.Equal(t, metadata.Value, testValue)
		})
	}
}

func TestSet_EmptyStateName(t *testing.T) {
	ctx := context.Background()
	sm := newMockStateManager(t)
	err := sm.Set(ctx, "", testValue)
	assert.Error(t, err)
}

func TestSet_WithCachedStateChange(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name       string
		initKind   ChangeKind
		expectKind ChangeKind
	}{
		{"state change kind None", None, Update},
		{"state change kind Add", Add, Add},
		{"state change kind Update", Update, Update},
		{"state change kind Remove", Remove, Update},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sm := newMockStateManager(t)
			sm.stateChangeTracker.Store(testState, &ChangeMetadata{Kind: tt.initKind, Value: testValue})

			err := sm.Set(ctx, testState, testValue)
			require.NoError(t, err)

			val, ok := sm.stateChangeTracker.Load(testState)
			require.True(t, ok)

			metadata := val.(*ChangeMetadata)
			assert.Equal(t, metadata.Kind, tt.expectKind)
			assert.Equal(t, metadata.Value, testValue)
		})
	}
}

func TestSet_WithoutCachedStateChange(t *testing.T) {
	ctx := context.Background()
	sm := newMockStateManager(t)

	err := sm.Set(ctx, testState, testValue)
	require.NoError(t, err)

	val, ok := sm.stateChangeTracker.Load(testState)
	require.True(t, ok)

	metadata := val.(*ChangeMetadata)
	assert.Equal(t, metadata.Kind, Add)
	assert.Equal(t, metadata.Value, testValue)
}

func TestSetWithTTL_EmptyStateName(t *testing.T) {
	ctx := context.Background()
	sm := newMockStateManager(t)
	err := sm.SetWithTTL(ctx, "", testValue, testTTL)
	assert.Error(t, err)
}

func TestSetWithTTL_NegativeTTL(t *testing.T) {
	ctx := context.Background()
	sm := newMockStateManager(t)
	err := sm.SetWithTTL(ctx, testState, testValue, -testTTL)
	assert.Error(t, err)
}

func TestSetWithTTL_WithCachedStateChange(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name       string
		initKind   ChangeKind
		expectKind ChangeKind
	}{
		{"state change kind None", None, Update},
		{"state change kind Add", Add, Add},
		{"state change kind Update", Update, Update},
		{"state change kind Remove", Remove, Update},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sm := newMockStateManager(t)
			sm.stateChangeTracker.Store(testState, &ChangeMetadata{Kind: tt.initKind, Value: testValue})

			err := sm.SetWithTTL(ctx, testState, testValue, testTTL)
			require.NoError(t, err)

			val, ok := sm.stateChangeTracker.Load(testState)
			require.True(t, ok)

			metadata := val.(*ChangeMetadata)
			assert.Equal(t, metadata.Kind, tt.expectKind)
			assert.Equal(t, metadata.Value, testValue)
			assert.Equal(t, *metadata.TTL, testTTL)
		})
	}
}

func TestSetWithTTL_WithoutCachedStateChange(t *testing.T) {
	ctx := context.Background()
	sm := newMockStateManager(t)

	err := sm.SetWithTTL(ctx, testState, testValue, testTTL)
	require.NoError(t, err)

	val, ok := sm.stateChangeTracker.Load(testState)
	require.True(t, ok)

	metadata := val.(*ChangeMetadata)
	assert.Equal(t, metadata.Kind, Add)
	assert.Equal(t, metadata.Value, testValue)
	assert.Equal(t, *metadata.TTL, testTTL)
}

func TestRemove_EmptyStateName(t *testing.T) {
	ctx := context.Background()
	sm := newMockStateManager(t)
	err := sm.Remove(ctx, "")
	assert.Error(t, err)
}

func TestRemove_WithCachedStateChange(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name    string
		kind    ChangeKind
		inCache bool
	}{
		{"state change kind None", None, true},
		{"state change kind Add", Add, false},
		{"state change kind Update", Update, false},
		{"state change kind Remove", Remove, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sm := newMockStateManager(t)
			sm.stateChangeTracker.Store(testState, &ChangeMetadata{Kind: tt.kind, Value: testValue})

			err := sm.Remove(ctx, testState)
			require.NoError(t, err)

			val, ok := sm.stateChangeTracker.Load(testState)
			if tt.inCache {
				assert.Equal(t, val.(*ChangeMetadata).Kind, Remove)
				assert.True(t, ok)
			} else {
				assert.Nil(t, val)
				assert.False(t, ok)
			}
		})
	}
}

func TestRemove_WithoutCachedStateChange(t *testing.T) {
	ctx := context.Background()
	mockResult := newGetActorStateResponse([]byte("result"))
	mockErr := errors.New("mockErr")

	tests := []struct {
		name     string
		mockFunc func(sm *stateManagerCtx, mc *mock_client.MockClient)
	}{
		{"state provider returns error", func(sm *stateManagerCtx, mc *mock_client.MockClient) {
			mc.EXPECT().GetActorState(ctx, newGetActorStateRequest(sm, testState)).Return(nil, mockErr)
		}},
		{"state provider returns data", func(sm *stateManagerCtx, mc *mock_client.MockClient) {
			mc.EXPECT().GetActorState(ctx, newGetActorStateRequest(sm, testState)).Return(mockResult, nil)
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sm := newMockStateManager(t)
			mockClient := sm.stateAsyncProvider.daprClient.(*mock_client.MockClient)
			tt.mockFunc(sm, mockClient)
			err := sm.Remove(ctx, testState)
			assert.NoError(t, err)
		})
	}
}

func TestContains_EmptyStateName(t *testing.T) {
	ctx := context.Background()
	sm := newMockStateManager(t)
	res, err := sm.Contains(ctx, "")
	assert.False(t, res)
	assert.Error(t, err)
}

func TestContains_WithCachedStateChange(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name     string
		kind     ChangeKind
		expected bool
	}{
		{"state change kind None", None, true},
		{"state change kind Add", Add, true},
		{"state change kind Update", Update, true},
		{"state change kind Remove", Remove, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sm := newMockStateManager(t)
			sm.stateChangeTracker.Store(testState, &ChangeMetadata{Kind: tt.kind, Value: testValue})

			result, err := sm.Contains(ctx, testState)
			assert.Equal(t, result, tt.expected)
			assert.NoError(t, err)
		})
	}
}

func TestContains_WithoutCachedStateChange(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name      string
		shouldErr bool
	}{
		{"state provider returns error", true},
		{"state provider returns data", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sm := newMockStateManager(t)
			mockClient := sm.stateAsyncProvider.daprClient.(*mock_client.MockClient)
			mockRequest := newGetActorStateRequest(sm, testState)
			mockResult := newGetActorStateResponse([]byte("result"))
			if tt.shouldErr {
				mockClient.EXPECT().GetActorState(ctx, mockRequest).Return(nil, errors.New("mockErr"))
			} else {
				mockClient.EXPECT().GetActorState(ctx, mockRequest).Return(mockResult, nil)
			}

			result, err := sm.Contains(ctx, testState)
			if tt.shouldErr {
				assert.Error(t, err)
				assert.False(t, result)
			} else {
				assert.NoError(t, err)
				assert.True(t, result)
			}
		})
	}
}

func TestSave_SingleCachedStateChange(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name         string
		stateChanges *ChangeMetadata
	}{
		{"no state change", nil},
		{"state change kind None", &ChangeMetadata{Kind: None, Value: testValue}},
		{"state change kind Add", &ChangeMetadata{Kind: Add, Value: testValue}},
		{"state change kind Update", &ChangeMetadata{Kind: Update, Value: testValue}},
		{"state change kind Remove", &ChangeMetadata{Kind: Remove, Value: testValue}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sm := newMockStateManager(t)
			mockClient := sm.stateAsyncProvider.daprClient.(*mock_client.MockClient)
			mockCodec := sm.stateAsyncProvider.stateSerializer.(*mock.MockCodec)
			if tt.stateChanges != nil {
				sm.stateChangeTracker.Store(testState, tt.stateChanges)

				if tt.stateChanges.Kind == Remove {
					mockClient.EXPECT().SaveStateTransactionally(ctx, sm.actorTypeName, sm.actorID, gomock.Len(1))
				} else if tt.stateChanges.Kind == Add || tt.stateChanges.Kind == Update {
					mockClient.EXPECT().SaveStateTransactionally(ctx, sm.actorTypeName, sm.actorID, gomock.Len(1))
					mockCodec.EXPECT().Marshal(tt.stateChanges.Value)
				}
			}

			err := sm.Save(ctx)
			assert.Nil(t, err)
		})
	}
}

func TestSave_MultipleCachedStateChanges(t *testing.T) {
	ctx := context.Background()
	sm := newMockStateManager(t)
	mockClient := sm.stateAsyncProvider.daprClient.(*mock_client.MockClient)
	mockCodec := sm.stateAsyncProvider.stateSerializer.(*mock.MockCodec)

	stateChanges := []struct {
		stateName string
		value     string
		kind      ChangeKind
	}{
		{"stateNone", "valueNone", None},
		{"stateAdd", "valueAdd", Add},
		{"stateUpdate", "valueUpdate", Update},
		{"stateRemove", "valueRemove", Remove},
	}
	for _, sc := range stateChanges {
		sm.stateChangeTracker.Store(sc.stateName, &ChangeMetadata{Kind: sc.kind, Value: sc.value})
	}

	// 3 operations: 1 Add, 1 Update, 1 Remove
	mockClient.EXPECT().SaveStateTransactionally(ctx, sm.actorTypeName, sm.actorID, gomock.Len(3))
	// 2 times: 1 Add, 1 Update
	mockCodec.EXPECT().Marshal(gomock.Any()).Times(2)

	err := sm.Save(ctx)
	assert.Nil(t, err)
}

func TestFlush(t *testing.T) {
	ctx := context.Background()
	sm := newMockStateManager(t)
	stateChanges := []struct {
		stateName string
		value     string
		kind      ChangeKind
	}{
		{"stateNone", "valueNone", None},
		{"stateAdd", "valueAdd", Add},
		{"stateUpdate", "valueUpdate", Update},
		{"stateRemove", "valueRemove", Remove},
	}
	for _, sc := range stateChanges {
		sm.stateChangeTracker.Store(sc.stateName, &ChangeMetadata{Kind: sc.kind, Value: sc.value})
	}

	sm.Flush(ctx)

	for _, sc := range stateChanges {
		val, ok := sm.stateChangeTracker.Load(sc.stateName)
		if sc.kind == Remove {
			assert.Nil(t, val)
			assert.False(t, ok)
		} else {
			metadata := val.(*ChangeMetadata)
			assert.Equal(t, metadata.Kind, None)
			assert.Equal(t, metadata.Value, sc.value)
			assert.True(t, ok)
		}
	}
}

func TestNewActorStateManagerContext(t *testing.T) {
	type args struct {
		actorTypeName      string
		actorID            string
		stateAsyncProvider *DaprStateAsyncProvider
	}
	tests := []struct {
		name string
		args args
		want *stateManagerCtx
	}{
		{
			name: "init",
			args: args{
				actorTypeName:      "test",
				actorID:            "fn",
				stateAsyncProvider: &DaprStateAsyncProvider{},
			},
			want: &stateManagerCtx{
				actorTypeName:      "test",
				actorID:            "fn",
				stateAsyncProvider: &DaprStateAsyncProvider{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewActorStateManagerContext(tt.args.actorTypeName, tt.args.actorID, tt.args.stateAsyncProvider); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewActorStateManagerContext() = %v, want %v", got, tt.want)
			}
		})
	}
}
