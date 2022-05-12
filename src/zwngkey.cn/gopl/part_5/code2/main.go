/*
 * @Author: zwngkey
 * @Date: 2021-11-21 15:21:56
 * @LastEditTime: 2022-05-13 06:17:56
 * @Description:
 */
package main

import (
	"log"
	"os"
	"text/template"
	"time"
	"zwngkey.cn/gopl/part_5/code2/github"
)

const templ = `{{.TotalCount}} issues:
{{range .Items}}-----------------------------
Number: {{.Number}}
User: 	{{.User.Login}}
Title: 	{{.Title | printf "%.64s"}}
Age:		{{.CreatedAt | daysAgo}} days
{{end}}`

func main() {
	report := template.Must(template.New("issueslist").
		Funcs(template.FuncMap{"daysAgo": daysAgo}).
		Parse(templ))

	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)

	}

	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}

	// fmt.Printf("%d issues:\n", result.TotalCount)
	// for _, item := range result.Items {
	// 	fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	// }
}

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}
