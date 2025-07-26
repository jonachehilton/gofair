package streaming

import (
	"testing"

	"github.com/jonachehilton/gofair/streaming/models"
	"github.com/stretchr/testify/assert"
)

func TestUpdateMatchedLays(t *testing.T) {
	// Arrange
	testCases := []struct {
		selectionID   int64
		matchedLays   [][]float64
		existingState [][]float64
		expectedState [][]float64
	}{
		{
			selectionID:   1234,
			matchedLays:   [][]float64{{1.51, 2.0}},
			existingState: [][]float64{},
			expectedState: [][]float64{{1.51, 2.0}},
		},
		{
			selectionID:   2345,
			matchedLays:   [][]float64{{1.51, 7.0}},
			existingState: [][]float64{{1.51, 5.0}, {1.53, 2.0}, {1.85, 3.6}},
			expectedState: [][]float64{{1.51, 7.0}, {1.53, 2.0}, {1.85, 3.6}},
		},
	}

	// Act/Arrange
	for _, testCase := range testCases {
		cache := newOrderBookCache()
		cache.Runners[testCase.selectionID] = new(models.OrderRunnerChange)
		cache.Runners[testCase.selectionID].Ml = testCase.existingState
		cache.updateMatchedLays(testCase.selectionID, testCase.matchedLays)
		assert.Equal(t, cache.Runners[testCase.selectionID].Ml, testCase.expectedState)
	}
}

func TestUpdateMatchedBacks(t *testing.T) {
	// Arrange
	testCases := []struct {
		selectionID   int64
		matchedBacks  [][]float64
		existingState [][]float64
		expectedState [][]float64
	}{
		{
			selectionID:   1234,
			matchedBacks:  [][]float64{{1.51, 2.0}},
			existingState: [][]float64{},
			expectedState: [][]float64{{1.51, 2.0}},
		},
		{
			selectionID:   2345,
			matchedBacks:  [][]float64{{1.51, 7.0}},
			existingState: [][]float64{{1.51, 5.0}, {1.53, 2.0}, {1.85, 3.6}},
			expectedState: [][]float64{{1.51, 7.0}, {1.53, 2.0}, {1.85, 3.6}},
		},
	}

	// Act/Arrange
	for _, testCase := range testCases {
		cache := newOrderBookCache()
		cache.Runners[testCase.selectionID] = new(models.OrderRunnerChange)
		cache.Runners[testCase.selectionID].Mb = testCase.existingState
		cache.updateMatchedBacks(testCase.selectionID, testCase.matchedBacks)
		assert.Equal(t, cache.Runners[testCase.selectionID].Mb, testCase.expectedState)
	}
}
