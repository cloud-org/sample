package handler

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"gomock-sample/store"
)

func TestHandler_GetValue(t *testing.T) {
	t.Parallel()
	type TestCase struct {
		name        string
		key         string
		errorReason string
		mockFn      func() store.Store
	}
	cases := []TestCase{
		{
			name:        "error",
			key:         "panda",
			errorReason: "get error",
			mockFn: func() store.Store {
				mockStore := store.NewMockStore(gomock.NewController(t))
				mockStore.EXPECT().Get("panda").Return("", fmt.Errorf("get error"))

				return mockStore
			},
		},
		{
			name: "get key - panda",
			key:  "panda",
			mockFn: func() store.Store {
				mockStore := store.NewMockStore(gomock.NewController(t))
				mockStore.EXPECT().Get("panda").Return("panda", nil)

				return mockStore
			},
		},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			var h *Handler
			if tc.mockFn != nil {
				h = NewHandler(tc.mockFn())
			}
			got, err := h.GetValue(tc.key)
			if err != nil {
				assert.Contains(t, err.Error(), tc.errorReason)
				return
			}
			assert.Nil(t, err)
			assert.Equal(t, tc.key, got)
		})
	}
}

func TestHandler_SetValue(t *testing.T) {
	t.Parallel()
	type TestCase struct {
		name        string
		key         string
		errorReason string
		mockFn      func() store.Store
	}
	cases := []TestCase{
		{
			name:        "error",
			key:         "panda",
			errorReason: "set error",
			mockFn: func() store.Store {
				mockStore := store.NewMockStore(gomock.NewController(t))
				mockStore.EXPECT().Set("panda", "panda").Return(fmt.Errorf("set error"))

				return mockStore
			},
		},
		{
			name: "set key - panda",
			key:  "panda",
			mockFn: func() store.Store {
				mockStore := store.NewMockStore(gomock.NewController(t))
				mockStore.EXPECT().Set("panda", "panda").Return(nil)

				return mockStore
			},
		},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			var h *Handler
			if tc.mockFn != nil {
				h = NewHandler(tc.mockFn())
			}
			err := h.SetValue(tc.key, tc.key)
			if err != nil {
				assert.Contains(t, err.Error(), tc.errorReason)
				return
			}
			assert.Nil(t, err)
		})
	}
}
