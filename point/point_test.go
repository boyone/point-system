package point

import "testing"

func Test_ProductPrice99_00ShouldReceive0Point(t *testing.T) {
	// Arrange
	expectedResult := 0
	price := 99.00

	// Act
	actualResult := CalculatePoint(price)

	// Assert
	if expectedResult != actualResult {
		t.Errorf("expect %v got %v", expectedResult, actualResult)
	}
}

func Test_ProductPrice100_00ShouldReceive1Point(t *testing.T) {
	// Arrange
	expectedResult := 1
	price := 100.00

	// Act
	actualResult := CalculatePoint(price)

	// Assert
	if expectedResult != actualResult {
		t.Errorf("expect %v point got %v point", expectedResult, actualResult)
	}
}

func Test_ProductPrice499_00ShouldReceive4Point(t *testing.T) {
	// Arrange
	expectedResult := 4
	price := 499.00

	// Act
	actualResult := CalculatePoint(price)

	// Assert
	if expectedResult != actualResult {
		t.Errorf("expect %v point got %v point", expectedResult, actualResult)
	}
}

func Test_ProductPrice789_00ShouldReceive7Point(t *testing.T) {
	// Arrange
	expectedResult := 7
	price := 789.00

	// Act
	actualResult := CalculatePoint(price)

	// Assert
	if expectedResult != actualResult {
		t.Errorf("expect %v point got %v point", expectedResult, actualResult)
	}
}