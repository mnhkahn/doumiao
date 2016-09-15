/*
 * @Author: mnhkahn
 * @Date: 2016-09-15 09:54:51
 * @Last Modified by: mnhkahn
 * @Last Modified time: 2016-09-15 10:29:41
 */
// http://www.01happy.com/golang-reverse-proxy/
package main

import (
	"github.com/astaxie/beego"
	"github.com/mnhkahn/doumiao/proxy"
)

func main() {
	beego.Handler("/", proxy.DEFAULT_PROXY_HANDER)
	beego.Run()
}
