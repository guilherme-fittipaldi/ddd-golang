package aggregate_test

import (
	"testing"

	"github.com/guilherme-fittipaldi/ddd-golang/aggregate"
)

func TestCustomer_NewCustomer(t *testing.T) {
	type testCase struct {
		test        string
		name        string
		expectedErr error
	}

	testCases := []testCase{
		{
			test:        "Empty name validation",
			name:        "",
			expectedErr: aggregate.ErrInvalidPerson,
		}, {
			test:        "Valid name",
			name:        "Guilherme Fittipaldi",
			expectedErr: nil,
		},
	}

	for _,tc := range testCases {
		t.Run(tc.test, func(t *testing.T){
			_, err := aggregate.NewCustomer(tc.name)
			if err != tc.expectedErr{
				t.Errorf("Exprcted error %v, got %v", tc.expectedErr, err)
			}
		})
	}
}
