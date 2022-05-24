package util

import (
	"go_blog/e"
	"strconv"
)

type SerID struct {
}

// 验证id是否合法, 并转化为uint类型
func (s *SerID) ConvertType(id string) (uint, error) {
	intId, err := strconv.Atoi(id)
	if err != nil {
		return 0, err
	}
	uintId := uint(intId)
	if uintId > 0 {
		return uintId, nil
	}
	return 0, e.ParamsInsError
}
