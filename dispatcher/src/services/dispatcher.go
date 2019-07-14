package services

import (
	"database/sql"
	"sync"

	log "github.com/sirupsen/logrus"
	"pic-collage.com/dispatcher/models"
)

const gap = 100000

var dispatchDao models.DispatchDao
var start, end uint64
var lock = sync.Mutex{}

func InitDispatchService(db *sql.DB) (err error) {
	dispatchDao = models.NewDispatchSQLDao(db)

	counter, err := dispatchDao.GetCounter()
	if err != nil {
		log.WithError(err).Error("Fail to get current counter")
		return err
	}

	log.WithField("counter", counter).Info("Current counter")

	start = counter
	end = start + gap

	return nil
}

func DispatchCounter() (uint64, uint64, error) {
	lock.Lock()
	defer lock.Unlock()

	_start := start
	_end := end
	start += gap
	end += gap

	err := dispatchDao.UpdateCounter(start)
	if err != nil {
		// Reset the counter.
		start = _start
		end = _end
		return 0, 0, err
	}

	return _start, _end, nil
}
