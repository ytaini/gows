// @author: wzmiiiiii
// @since: 2022/12/23 21:14:16
// @desc: TODO

package controller

import (
	"log"
	"regexp"
	"testing"
)

func TestRegExp(t *testing.T) {
	url := `/companies/edit/0b727af5-840e-4025-bd89-9611406b31cc`
	re := regexp.MustCompile(`/companies/edit/([a-zA-Z0-9-]*$)`)
	matched := re.FindStringSubmatch(url)
	log.Println(matched)
}
