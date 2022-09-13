package binding

import (
	"sync"

	"github.com/go-playground/validator/v10"
)

type any interface{}

type ValidationError []error

type DefaultValidator struct {
	once     sync.Once
	validate *validator.Validate
}

func (err ValidationError) Error() string {
	
}

func (v *DefaultValidator) lazyInit() {
	v.once.Do(func() {
		v.validate = validator.New()
		v.validate.SetTagName("binding")
	})
}

func (v *DefaultValidator) Engine() any {
	v.lazyInit()
	return v.validate
}
