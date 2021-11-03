package usecase_test

import (
	"errors"
	"id-maker/internal/entity"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

var errInternalServErr = errors.New("internal server error")

type test struct {
	name string
	mock func()
	res  interface{}
	err  error
}

func segment(t *testing.T) (*MockSegment, *MockSegmentRepo) {
	t.Helper()

	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()

	s := NewMockSegment(mockCtl)
	repo := NewMockSegmentRepo(mockCtl)

	return s, repo
}

func TestGetId(t *testing.T) {
	t.Parallel()

	segment, _ := segment(t)

	tests := []test{
		{
			name: "right result",
			mock: func() {
				segment.EXPECT().GetId("test").Return(int64(100), nil)
				segment.EXPECT().SnowFlakeGetId().Return(int64(100))
			},
			res: int64(100),
			err: nil,
		},
	}

	for _, tc := range tests {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			tc.mock()

			res, err := segment.GetId("test")

			require.Equal(t, res, tc.res)
			require.ErrorIs(t, err, tc.err)
		})
	}
}

func TestCreateTag(t *testing.T) {
	t.Parallel()

	segment, _ := segment(t)

	tests := []test{
		{
			name: "create tag",
			mock: func() {
				segment.EXPECT().CreateTag(&entity.Segments{}).Return(nil)
			},
			res: nil,
			err: nil,
		},
	}

	for _, tc := range tests {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			tc.mock()

			res := segment.CreateTag(&entity.Segments{})

			require.Equal(t, res, tc.res)
			require.ErrorIs(t, res, tc.err)
		})
	}
}

func TestGetList(t *testing.T) {
	t.Parallel()

	_, repo := segment(t)

	tests := []test{
		{
			name: "empty result",
			mock: func() {
				repo.EXPECT().GetList().Return(nil, nil)
			},
			res: []entity.Segments(nil),
			err: nil,
		},
		{
			name: "result with error",
			mock: func() {
				repo.EXPECT().GetList().Return(nil, errInternalServErr)
			},
			res: []entity.Segments(nil),
			err: errInternalServErr,
		},
	}

	for _, tc := range tests {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			tc.mock()

			res, err := repo.GetList()

			require.Equal(t, res, tc.res)
			require.ErrorIs(t, err, tc.err)
		})
	}
}
