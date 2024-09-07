package classify

import "qingyu-wf/global"

func Create(c TableClassify) error {
	return global.MySql.Create(&c).Error
}

func SearchDb(key, val string) []TableClassify {
	var tc []TableClassify
	global.MySql.Where(key+" = ?", val).Find(&tc)
	return tc
}
