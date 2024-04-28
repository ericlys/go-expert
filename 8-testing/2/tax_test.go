package tax

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateTax(t *testing.T) {
	tax, err := CalculateTax(1000.0)
	assert.Nil(t, err)
	assert.Equal(t, 10.0, tax)

	tax, err = CalculateTax(0)
	assert.Error(t, err, "ammout must be greater than 0")
	assert.Equal(t, 0.0, tax)
}

func TestCalculateTaxAndSave(t *testing.T) {
	r := &TaxRepositoryMock{}
	r.On("SaveTax", 10.0).Return(nil)
	r.On("SaveTax", 0.0).Return(errors.New("error saving tax"))

	err := CalculateTaxAndSave(1000.0, r)
	assert.Nil(t, err)

	err = CalculateTaxAndSave(00.0, r)
	assert.Error(t, err, "error saving tax")

	r.AssertExpectations(t)
	r.AssertNumberOfCalls(t, "SaveTax", 2)
}
