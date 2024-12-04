package doubao

import "testing"

func TestGenChart(t *testing.T) {
	GenChart("网站用户增长分析", "日期,用户数\n1号,10\n2号,20\n3号,30")
}
