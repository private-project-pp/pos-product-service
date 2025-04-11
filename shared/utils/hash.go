package utils

import (
	"github.com/private-project-pp/pos-general-lib/shared/utils"
	"github.com/private-project-pp/pos-general-lib/stacktrace"
	"github.com/private-project-pp/product-rpc-service/shared/config"
)

func GenerateHashedPassword(value string) (hashed string, err error) {

	hashed, err = utils.GenerateSHA256Hash(value, config.Auth.PassHashCode)
	if err != nil {
		return hashed, stacktrace.Cascade(err, stacktrace.BAD_PROCESSING, err.Error())
	}
	return hashed, nil
}
