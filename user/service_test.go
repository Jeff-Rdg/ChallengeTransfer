package user

import (
	"errors"
	"testing"
)

func TestService_Create(t *testing.T) {
	var service Service

	type testCase struct {
		testName    string
		request     Request
		expectedErr error
	}

	testCases := []testCase{
		{
			testName: "invalid name",
			request: Request{
				FullName:     "name with number 1",
				TaxNumber:    "933.203.710-88",
				Email:        "test@gmail.com",
				Password:     "123456",
				IsShopkeeper: false,
			},
			expectedErr: FullNameInvalidErr,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			_, err := service.Create(tc.request)
			if !errors.Is(err, tc.expectedErr) {
				t.Errorf("expected error %v, got %v", tc.expectedErr, err)
			}
		})
	}
}
