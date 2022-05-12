/*
 * @Author: zwngkey
 * @Date: 2021-11-21 17:11:28
 * @LastEditTime: 2022-05-13 06:18:29
 * @Description:
 */
package main

import (
	"html/template"
	"log"
	"os"
	"zwngkey.cn/gopl/part_5/code2/github"
)

func main() {
	var issueList = template.Must(template.New("issuelist").Parse(`
		<h1>{{.TotalCount}} issues</h1>
		<table>
		<tr style='text-align: left'>
			<th>#</th>
			<th>State</th>
			<th>User</th>
			<th>Title</th>
		</tr>
		{{range .Items}}
		<tr>
			<td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
			<td>{{.State}}</td>
			<td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
			<td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
		</tr>
		{{end}}
		</table>
		`))
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)

	}

	if err := issueList.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}
