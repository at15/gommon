package errors_test

import (
	"os"
	"sync"
	"testing"

	asst "github.com/stretchr/testify/assert"

	"github.com/dyweb/gommon/errors"
)

func TestMultiErr_Error(t *testing.T) {
	errs := map[string]func() errors.MultiErr{
		"unsafe": errors.NewMultiErr,
		"safe":   errors.NewMultiErrSafe,
	}
	for name := range errs {
		t.Run(name, func(t *testing.T) {
			assert := asst.New(t)
			merr := errs[name]()
			// TODO: (at15) I am not sure if this is the intended behavior ...
			assert.Equal("0 errors", merr.Error())
		})
	}
}

func TestMultiErrSafe_Append(t *testing.T) {
	assert := asst.New(t)
	merr := errors.NewMultiErrSafe()
	nRoutine := 10
	errPerRoutine := 20
	var wg sync.WaitGroup
	wg.Add(nRoutine)
	for i := 0; i < nRoutine; i++ {
		go func() {
			for j := 0; j < errPerRoutine; j++ {
				merr.Append(os.ErrClosed)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	assert.Equal(nRoutine*errPerRoutine, len(merr.Errors()))
}

func TestMultiErr_AppendReturn(t *testing.T) {
	// NOTE: inspired by https://github.com/uber-go/multierr/issues/21
	errs := map[string]func() errors.MultiErr{
		"unsafe": errors.NewMultiErr,
		"safe":   errors.NewMultiErrSafe,
	}
	for name := range errs {
		t.Run(name, func(t *testing.T) {
			assert := asst.New(t)
			merr := errs[name]()
			assert.False(merr.Append(nil))
			assert.True(merr.Append(os.ErrPermission))
			assert.False(merr.Append(nil))
		})
	}
}

func TestMultiErr_Flatten(t *testing.T) {
	errs := map[string]func() errors.MultiErr{
		"unsafe": errors.NewMultiErr,
		"safe":   errors.NewMultiErrSafe,
	}
	for name := range errs {
		t.Run(name, func(t *testing.T) {
			assert := asst.New(t)
			merr := errs[name]()
			merr.Append(os.ErrPermission)
			merr.Append(os.ErrClosed)
			assert.Equal(2, len(merr.Errors()))

			merr2 := errs[name]()
			merr2.Append(merr)
			merr2.Append(os.ErrNotExist)
			merr2.Append(nil) // nil is not appended
			assert.Equal(3, len(merr2.Errors()))
			t.Log(merr2.Error())
		})
	}
}

func TestMultiErr_Errors(t *testing.T) {
	errs := map[string]func() errors.MultiErr{
		"unsafe": errors.NewMultiErr,
		"safe":   errors.NewMultiErrSafe,
	}
	for name := range errs {
		t.Run(name, func(t *testing.T) {
			assert := asst.New(t)
			merr := errs[name]()
			assert.Nil(merr.Errors())

			fakeErrs := []error{os.ErrClosed, os.ErrExist, os.ErrPermission, os.ErrNotExist}
			for i := 0; i < len(fakeErrs); i++ {
				assert.True(merr.Append(fakeErrs[i]))
			}
			errs := merr.Errors()
			assert.Equal(len(fakeErrs), merr.Len())
			assert.Equal(len(fakeErrs), len(errs))
			for i := 0; i < len(fakeErrs); i++ {
				assert.Equal(fakeErrs[i], errs[i])
			}
		})
	}
}

func TestMultiErr_ErrorOrNil(t *testing.T) {
	errs := map[string]func() errors.MultiErr{
		"unsafe": errors.NewMultiErr,
		"safe":   errors.NewMultiErrSafe,
	}
	for name := range errs {
		t.Run(name, func(t *testing.T) {
			assert := asst.New(t)
			merr := errs[name]()
			assert.Nil(merr.ErrorOrNil())
			assert.False(merr.HasError())

			merr.Append(os.ErrPermission)
			assert.NotNil(merr.ErrorOrNil())
			assert.True(merr.HasError())
		})
	}
}
