package exceptionHandler

import (
	"github.com/hashicorp/go-multierror"
	"strings"
)

func Handle(result *multierror.Error) (code int, errStr string) {
	if result.ErrorOrNil() != nil {
		code := 500
		if strings.Contains(result.Error(), "不存在") {
			code = 404
		} else if strings.Contains(result.Error(), "有误") {
			code = 400
		}
		return code, result.Error()
	}
	return 200, ""

}
