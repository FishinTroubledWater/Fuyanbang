package exceptionHandler

import (
	"fmt"
	"github.com/hashicorp/go-multierror"
	"strings"
)

func Handle(result *multierror.Error) (code int, errStr string) {
	if result.ErrorOrNil() != nil {
		errs := result.Errors
		for i := 0; i < len(errs); i++ {
			err := errs[i]
			errStr := err.Error()
			fmt.Println(errStr)
			if strings.Contains(errStr, "查询的记录不存在") {
				return 404, "查询的记录不存在"
			}
			if strings.Contains(errStr, "用户已存在") {
				return 404, "用户已存在"
			}
		}
		return 500, result.Error()
	}
	return 200, ""

}
