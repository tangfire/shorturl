package connect

import (
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestGet(t *testing.T) {
	convey.Convey("基础用例", t, func() {
		url := "https://cs.gdut.edu.cn/jsdw/fg/azmpx/Y.htm"
		got := Get(url)
		// 断言
		convey.So(got, convey.ShouldEqual, true)
		convey.ShouldBeTrue(got)
	})

	convey.Convey("url请求不通过的示例", t, func() {
		url := "posts/Go/unit-test-5/"
		got := Get(url)
		fmt.Println(got)
		// 断言
		convey.ShouldBeFalse(got)
	})

}
