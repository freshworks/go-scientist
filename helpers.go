package scientist

import (
	"fmt"
	"math/rand"
	"time"
)

func Bool(ok interface{}, err error) (bool, error) {
	if err != nil {
		return false, err
	}

	switch t := ok.(type) {
	case bool:
		return t, nil
	default:
		return false, fmt.Errorf("[scientist] bad result type: %v (%T)", ok, ok)
	}
}

// Shuffle randomizes the behavior access.
func (e *Experiment) Shuffle(behaviourName string, skip bool) []string {
	var behaviors []string
	for name := range e.behaviors {
		if skip && (behaviourName == name) {
			continue
		}
		behaviors = append(behaviors, name)
	}

	t := time.Now()
	rand.Seed(int64(t.Nanosecond()))

	arr := behaviors
	for i := len(arr) - 1; i > 0; i-- {
		j := rand.Intn(i)
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}
