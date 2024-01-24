package gonextcloud

import (
	"errors"
	"strconv"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserUpdateErrors(t *testing.T) {
	exp := map[string]error{}
	errs := make(chan *UpdateError)
	go func() {
		for i := 0; i < 10; i++ {
			f := strconv.Itoa(i)
			e := errors.New(f)
			err := UpdateError{
				Field: f,
				Error: e,
			}
			exp[f] = e
			errs <- &err
		}
		close(errs)
	}()
	uerrs := newUpdateError(errs)
	assert.Equal(t, exp, uerrs.Errors)
	assert.NotEmpty(t, uerrs.Error())
}

func TestUserUpdateErrorsNil(t *testing.T) {
	var wg sync.WaitGroup
	errs := make(chan *UpdateError)
	wg.Add(1)
	go func() {
		errs <- nil
		wg.Done()
	}()
	go func() {
		wg.Wait()
		close(errs)
	}()
	uerrs := newUpdateError(errs)
	assert.Nil(t, uerrs)
}
