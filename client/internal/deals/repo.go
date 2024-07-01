package deals

import (
	"errors"
	"hakaton2024/client/pkg/models"
	"sync"
)

var (
	ErrNoDeals = errors.New("no deals found")
	mu         = &sync.RWMutex{}
)

type DealMemoryRepository struct {
	data map[string][]models.Deal
}

func NewMemoryRepo() *DealMemoryRepository {
	return &DealMemoryRepository{
		data: map[string][]models.Deal{
			"RTS": {
				{
					TimeClosed: models.Time{Hour: 12, Minute: 1},
					OpenVal:    1,
					ClosedVal:  3,
					HighVal:    5,
					LowVal:     1,
					VolumeVal:  10,
				},
				{
					TimeClosed: models.Time{Hour: 12, Minute: 2},
					OpenVal:    3,
					ClosedVal:  7,
					HighVal:    8,
					LowVal:     2,
					VolumeVal:  18,
				},
			},
		},
	}
}

func (repo *DealMemoryRepository) GetTicker(ticket string) ([]models.Deal, error) {
	mu.RLock()
	deals, ok := repo.data[ticket]
	mu.RUnlock()
	if !ok {
		return nil, ErrNoDeals
	}
	return deals, nil
}

func (repo *DealMemoryRepository) AddTickerDeals(ticket string, deals []models.Deal) (bool, error) {
	mu.Lock()
	repo.data[ticket] = deals
	mu.Unlock()
	return true, nil
}
