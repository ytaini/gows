/*
 * @Author: wzmiiiiii
 * @Date: 2022-07-25 15:03:50
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-11-21 11:53:26
 * @Description:
 */
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/axgle/mahonia"
)

func main() {
	htmlcontext := fetchURL("http://live.500.com")
	// fmt.Println(body)
	res := parseHTML(htmlcontext)
	// fmt.Println(res[0][0])
	// fmt.Println(res[1][0])

	parseOdds(htmlcontext)

	var gameInfos []*GameInfo

	for _, v := range res {
		gameInfo := parseGameInfo(v[0])
		gameInfos = append(gameInfos, gameInfo)
	}

	for i, gameInfo := range gameInfos {
		fmt.Println("比赛信息" + strconv.Itoa(i) + ":")
		fmt.Println("\t赛事信息:", gameInfo.GameType)
		fmt.Println("\t赛事id:", gameInfo.Fid)
		fmt.Println("\t比赛时间:", gameInfo.GameTime)
		fmt.Println("\t主队队伍:", gameInfo.RightClub)
		fmt.Println("\t客队队伍:", gameInfo.LeftClub)
		fmt.Println("\t赛事赔率:", gameInfo.Odds)
	}

}

type GameInfo struct {
	GameType  string
	GameTime  string
	RightClub string
	LeftClub  string
	Fid       string
	Odds      [3]string
}

func parseGameInfo(content string) *GameInfo {
	gameInfo := &GameInfo{}
	pat := `<td bgcolor=".*>(.+)</a></td>`
	typereg := regexp.MustCompile(pat)
	gametype := typereg.FindAllStringSubmatch(content, -1)
	// fmt.Println(gametype)
	// fmt.Println(gametype[0][1])
	gameInfo.GameType = gametype[0][1]

	//获取比赛时间 <td align="center">08-16 00:45</td>
	pat = `<td align="center">(.+)</td>`
	timereg := regexp.MustCompile(pat)
	gameTime := timereg.FindAllStringSubmatch(content, -1)
	// fmt.Println(gameTime)
	// fmt.Println(gameTime[1][1])
	gameInfo.GameTime = gameTime[1][1]

	//获取主队信息 <td align="right" class="p_lr01"><span class="gray">[12]</span><a target="_blank" href="//liansai.500.com/team/4492/">瓦尔贝里</a><span class="sp_sr">(+1)</span></td>
	pat = `<td align="right".*><a.*?>(.+)</a>.*</td>`
	rightreg := regexp.MustCompile(pat)
	gameRight := rightreg.FindAllStringSubmatch(content, -1)
	// fmt.Println(gameRight)
	gameInfo.RightClub = gameRight[0][1]

	//获取客队信息 <td align="left" class="p_lr01"><a target="_blank" href="//liansai.500.com/team/852/">哈马比</a><span class="gray">[05]</span></td>
	pat = `<td align="left".*><a.*?>(.+)</a>.*</td>`
	leftreg := regexp.MustCompile(pat)
	gameLeft := leftreg.FindAllStringSubmatch(content, -1)
	// fmt.Println(gameLeft)
	gameInfo.LeftClub = gameLeft[0][1]

	//获取fid 比赛唯一标识
	pat = `fid="(\d{7})"`
	fidreg := regexp.MustCompile(pat)
	gameFid := fidreg.FindAllStringSubmatch(content, -1)
	// fmt.Println(gameFid)
	gameInfo.Fid = gameFid[0][1]

	//获取赔率
	gameInfo.Odds = gameOdds[gameInfo.Fid]["sp"]
	return gameInfo
}

/*
{"808835":{"0":[1.35,5.05,7.38],"280":[1.33,4.95,6.7]},"809489":{"0":[2.67,2.9,2.79],"3":[2.75,2.87,2.87]}}
*/
//解析足彩赔率
var gameOdds map[string]map[string][3]string

func parseOdds(body string) {
	pat := `<script.*>\s?.*?({.*);\s?</script>`
	reg := regexp.MustCompile(pat)
	gameOdd := reg.FindAllStringSubmatch(body, -1)
	gloOdd := make(map[string]map[string][3]string)
	err := json.Unmarshal([]byte(gameOdd[0][1]), &gloOdd)
	if err != nil {
		log.Panic("failed to unmarshal json ", err)
	}
	gameOdds = gloOdd
}

//解析网页
/*
<tr id="a855640" order="4001" status="0" gy="欧罗巴,莫斯巴达,图恩" yy="欧罗巴,莫斯科斯巴達,图恩" lid="63" fid="855640" sid="5295" class="" infoid="132788" r="1">
    <td align="center" class=""><input type="checkbox" name="check_id[]" value="855640" />周四001</td>
    <td bgcolor="#6F00DD" class="ssbox_01"><a style="color:#fff" target="_blank" href="//liansai.500.com/zuqiu-5295/">欧罗巴</a></td>
    <td align="center">资格赛3</td>
    <td align="center">08-16 00:45</td>
    <td align="center">未</td>
    <td align="right" class="p_lr01"><span class="gray"></span><a target="_blank" href="//liansai.500.com/team/1114/">莫斯巴达</a><span class="sp_rq">(-1)</span></td>
    <td align="center"><div class="pk"><a href="./detail.php?fid=855640&r=1" target="_blank" class="clt1" ></a><a href="./detail.php?fid=855640&r=1" target="_blank" class="fgreen" data-ao="一球" data-pb="一球">一球</a><a href="./detail.php?fid=855640&r=1" target="_blank" class="clt3" ></a></div></td>
    <td align="left" class="p_lr01"><a target="_blank" href="//liansai.500.com/team/777/">图恩</a><span class="gray"></span></td>
    <td align="center" class="red"> - </td>
    <td align="center" class="bf_op">&nbsp;</td>
    <td align="center" class="red">&nbsp; <a href="./detail.php?fid=855640" class="live_animate" title="动画" target="_blank"></a></td>
    <td align="center" class="td_warn"><a target="_blank" href="//odds.500.com/fenxi/shuju-855640.shtml">析</a> <a target="_blank" href="//odds.500.com/fenxi/yazhi-855640.shtml">亚</a> <a target="_blank" href="//odds.500.com/fenxi/ouzhi-855640.shtml">欧</a> <a target="_blank" id="qing_855640" class="red hide"   href="//odds.500.com/fenxi/youliao-855640.shtml?channel=pc_score">情</a></td>
    <td align="center" class=""><a href="javascript:void(0)" class="icon_notop">置顶</a></td>
</tr>
*/
func parseHTML(html string) [][]string {
	pat := `<tr id="(?s).*?</tr>`
	reg := regexp.MustCompile(pat)
	res := reg.FindAllStringSubmatch(html, -1)
	return res
}

// 获取网页信息
func fetchURL(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		log.Panic("failed to request ", err)
	}

	if resp.StatusCode != 200 {
		log.Panic("err statuscode ")
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Panic("failed to ReadAll ", err)
	}

	decoder := mahonia.NewDecoder("GB18030")
	utf8body := decoder.ConvertString(string(body))
	return utf8body
}
