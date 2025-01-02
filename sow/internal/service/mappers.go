package service

import (
	"errors"
	"fmt"

	"github.com/calamity-m/reap/pkg/errs"
	"github.com/calamity-m/reap/proto/sow/v1"
	"github.com/calamity-m/reap/sow/internal/persistence"
	"github.com/google/uuid"
)

func MapEntryToRecord(entry persistence.FoodRecordEntry) *sow.Record {
	record := &sow.Record{
		UserId:      entry.UserId.String(),
		Name:        entry.Name,
		Description: entry.Description,
		Kj:          entry.KJ,
		Grams:       entry.Grams,
		Ml:          entry.ML,
		Calories:    kjToCals(entry.KJ),
		Oz:          gramsToOz(entry.Grams),
		FlOz:        mlToFLOz(entry.ML),
	}

	return record
}

func MapRecordToEntry(record *sow.Record) (persistence.FoodRecordEntry, error) {
	userId, err := uuid.Parse(record.GetUserId())
	if err != nil {
		return persistence.FoodRecordEntry{}, errors.Join(fmt.Errorf("user id is not valid"), errs.ErrInvalidRequest)
	}

	entry := persistence.FoodRecordEntry{
		UserId:      userId,
		Name:        record.Name,
		Description: record.Description,
		KJ:          calsToKJ(record.Calories),
		ML:          flOzToML(record.FlOz),
		Grams:       ozToGrams(record.Oz),
		Created:     record.Time.AsTime(),
	}

	// Yucky imperial system
	if record.Kj != 0 {
		entry.KJ = record.Kj
	}
	if record.Grams != 0 {
		entry.Grams = record.Grams
	}
	if record.Ml != 0 {
		entry.ML = record.Ml
	}

	return entry, nil
}

func calsToKJ(cals float32) float32 {
	return cals * 4.184
}

func ozToGrams(oz float32) float32 {
	return oz * 28.35
}

func flOzToML(flOz float32) float32 {
	return flOz * 29.574
}

func kjToCals(kj float32) float32 {
	return kj / 4.184
}

func gramsToOz(grams float32) float32 {
	return grams / 28.35
}

func mlToFLOz(ml float32) float32 {
	return ml / 29.574
}
