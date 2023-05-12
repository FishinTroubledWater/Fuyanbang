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
			if strings.Contains(errStr, "要查询的记录不存在") {
				return 404, "要查询的记录不存在"
			}
			if strings.Contains(errStr, "记录已存在") {
				return 404, "记录已存在"
			}
			if strings.Contains(errStr, "要删除的记录不存在") {
				return 400, "要删除的记录不存在"
			}
			if strings.Contains(errStr, "要修改的记录不存在") {
				return 400, "要修改的记录不存在"
			}
			if strings.Contains(errStr, "要插入的记录有误") {
				return 400, "要插入的记录有误"
			}
		}
		return 500, result.Error()
	}
	return 200, ""

}
