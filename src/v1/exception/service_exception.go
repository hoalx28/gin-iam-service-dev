package exception

import (
	"errors"
	"iam/src/v1/constant"

	"github.com/samber/lo"
)

type serviceException struct {
	Root   error
	Failed constant.Failed
}

type ServiceException interface {
	Error() string
	GetRootCause() string
	GetFailed() constant.Failed
}

func NewServiceException(root error, failed constant.Failed) *serviceException {
	if lo.IsNil(root) {
		return &serviceException{Root: errors.New(failed.Message), Failed: failed}
	}
	return &serviceException{Root: root, Failed: failed}
}

func (e serviceException) Error() string {
	return e.Root.Error()
}

func (e serviceException) GetRootCause() string {
	if err, ok := e.Root.(serviceException); ok {
		return err.GetRootCause()
	}
	return e.Root.Error()
}

func (e serviceException) GetFailed() constant.Failed {
	return e.Failed
}
