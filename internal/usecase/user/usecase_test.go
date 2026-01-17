package user

import (
	"context"
	"testing"

	"github.com/madkingxxx/backend-test/internal/core/skinport"
	"github.com/madkingxxx/backend-test/internal/core/user"
	"go.uber.org/mock/gomock"
)

type mocks struct {
	userService     *MockuserServiceI
	skinportService *MockskinportServiceI
}

func getUsecase(ctrl *gomock.Controller) (*UseCase, *mocks) {
	mocks := &mocks{
		userService:     NewMockuserServiceI(ctrl),
		skinportService: NewMockskinportServiceI(ctrl),
	}

	return &UseCase{
		userService:     mocks.userService,
		skinportService: mocks.skinportService,
	}, mocks
}

func TestNew(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	usecase, _ := getUsecase(ctrl)
	if usecase == nil {
		t.Fatal("expected usecase to be not nil")
	}
	if usecase.userService == nil {
		t.Fatal("expected userService to be not nil")
	}
	if usecase.skinportService == nil {
		t.Fatal("expected skinportService to be not nil")
	}
}

func TestPurchase(t *testing.T) {
	type input struct {
		userID   int
		hashName string
	}

	type testCase struct {
		name    string
		input   input
		output  user.User
		prepare func(m *mocks)
	}

	testCases := []testCase{
		{
			name: "Successful Purchase",
			input: input{
				userID:   1,
				hashName: "AK-47 | Redline (Field-Tested)",
			},

			output: user.User{
				ID:      1,
				Balance: 5.0, // Assuming initial balance was 100.0 and item price is 50.0
			},

			prepare: func(m *mocks) {
				m.skinportService.EXPECT().
					Get(gomock.Any(), "AK-47 | Redline (Field-Tested)").
					Return(skinport.Item{
						MarketHashName: "AK-47 | Redline (Field-Tested)",
						MinPrice:       50.0,
					}, nil)

				m.userService.EXPECT().
					Withdraw(gomock.Any(), 1, 50.0).
					Return(user.User{
						ID:      1,
						Balance: 5.0,
					}, nil)
			},
		},
	}

	for _, tc := range testCases {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		usecase, mocks := getUsecase(ctrl)
		tc.prepare(mocks)

		t.Run(tc.name, func(t *testing.T) {
			result, err := usecase.Purchase(context.Background(), tc.input.userID, tc.input.hashName)
			if err != nil {
				t.Fatalf("expected no error, got %v", err)
			}

			if result != tc.output {
				t.Fatalf("expected %v, got %v", tc.output, result)
			}
		})
	}
}
