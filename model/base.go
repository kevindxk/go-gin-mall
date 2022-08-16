/**
 * @Author: dexukong
 * @Description:
 * @File:  base
 * @Version: 1.0.0
 * @Date: 2022/08/16 9:00
 */

package model

type BasePage struct {
	PageNum  int `from:"pageNum"`
	PageSize int `from:"pageSize"`
}
