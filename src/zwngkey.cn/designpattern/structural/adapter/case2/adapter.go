/*
 * @Author: wzmiiiiii
 * @Date: 2022-12-09 02:50:24
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-12-09 03:31:06
 */
package adapter

import "fmt"

// 我们的接口（新接口）——音乐播放
type MusicPlayer interface {
	play(fileType string, fileName string)
}

// 在网上找的已实现好的库 音乐播放
// ( 旧接口）
type ExistPlayer struct {
}

func (*ExistPlayer) playMp3(fileName string) {
	fmt.Println("play mp3 :", fileName)
}
func (*ExistPlayer) playWma(fileName string) {
	fmt.Println("play wma :", fileName)
}

// 适配器
type PlayerAdaptor struct {
	// 持有一个旧接口
	existPlayer ExistPlayer
}

// 实现新接口
func (player *PlayerAdaptor) play(fileType string, fileName string) {
	switch fileType {
	case "mp3":
		player.existPlayer.playMp3(fileName)
	case "wma":
		player.existPlayer.playWma(fileName)
	default:
		fmt.Println("暂时不支持此类型文件播放")
	}
}
