// test/inbound_service_test.go
package test

import (
	"time"
	"testing"
	"github.com/stretchr/testify/assert"
	"sistem-manajemen-gudang/model"
	"sistem-manajemen-gudang/service"
)

// 🔧 Mock Repository
type mockInboundRepository struct{}

func (m *mockInboundRepository) Create(p model.Inbound) (model.Inbound, error) {
	return  model.Inbound{}, nil
}
func (m *mockInboundRepository) FindAll() ([]model.Inbound, error) {
	return []model.Inbound{
		
		{
			ID:         1,
			ProductID:  101,
			Quantity:   20,
			ReceivedAt: time.Date(2025, 8, 1, 0, 0, 0, 0, time.UTC),
			Supplier:   "Supplier A",
			Product: model.Product{
				ID:   101,
				Name: "Monitor Samsung",
			},
		},
			{
			ID:         2,
			ProductID:  102,
			Quantity:   30,
			ReceivedAt: time.Date(2025, 8, 2, 0, 0, 0, 0, time.UTC),
			Supplier:   "Supplier B",
			Product: model.Product{
				ID:   102,
				Name: "Keyboard Logitech",
			},
		},
	}, nil
}
func (m *mockInboundRepository) FindByID(id uint) (model.Inbound, error) {
	return model.Inbound{
		ID:         id,
		ProductID:  10,
		Quantity:   50,
		ReceivedAt: time.Date(2025, 8, 5, 0, 0, 0, 0, time.UTC),
		Supplier:   "PT Sumber Makmur",
		Product: model.Product{
			ID:   10,
			Name: "Laptop ASUS",
		},
	}, nil
}
func (m *mockInboundRepository) Update(p model.Inbound) (model.Inbound, error) {
	return p, nil
}
func (m *mockInboundRepository) Delete(id uint) error {
	return nil
}
func (m *mockInboundRepository) GetInboundLaporan(start, end string) ([]model.Inbound, error) {
	return nil, nil
}

func TestInboundService_GetByID(t *testing.T) {
	mockRepo := &mockInboundRepository{}
	svc := service.NewInboundService(mockRepo)

	result, err := svc.GetByID(1)
	assert.NoError(t, err)
	assert.Equal(t, uint(1), result.ID)
	assert.Equal(t, uint(10), result.ProductID)
	assert.Equal(t, 50, result.Quantity)
	assert.Equal(t, "PT Sumber Makmur", result.Supplier)
	assert.Equal(t, "Laptop ASUS", result.Product.Name)
}
func TestInboundServiceGetAll(t *testing.T){
	mockRepo := &mockInboundRepository{}
	svc := service.NewInboundService(mockRepo)

	result, err := svc.GetAll() 

	assert.NoError(t, err)
	assert.Len(t, result, 2)

	assert.Equal(t, uint(1), result[0].ID)
	assert.Equal(t, "Monitor Samsung", result[0].Product.Name)

	assert.Equal(t, uint(2), result[1].ID)
	assert.Equal(t, "Keyboard Logitech", result[1].Product.Name)


}