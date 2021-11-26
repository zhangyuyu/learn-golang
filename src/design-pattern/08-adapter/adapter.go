package adapter

import "fmt"

// MediaPlayer 多媒体播放接口
type MediaPlayer interface {
	Play(audioType, fileName string)
}

// AdvanceMediaPlayer 新多媒体播放接口
type AdvanceMediaPlayer interface {
	PlayRmvb(fileName string)
	PlayMP4(fileName string)
}

// RmvbPlayer rmvb播放器类
type RmvbPlayer struct {
}

// PlayRmvb 播放rmvb格式视频方法
func (r *RmvbPlayer) PlayRmvb(fileName string) {
	fmt.Println("Playing rmbv file. Name: ", fileName)
}

//PlayMP4 为了实现接口所写，不使用这个方法
func (r *RmvbPlayer) PlayMP4(fileName string) {
	//do nothing
}

// Mp4Player mp4播放器类
type Mp4Player struct {
}

// PlayRmvb 为了实现接口所写，不使用这个方法
func (r *Mp4Player) PlayRmvb(fileName string) {
	//do nothing
}

// PlayMP4 播放mp4格式视频方法
func (r *Mp4Player) PlayMP4(fileName string) {
	fmt.Println("Playing mp4 file. Name: ", fileName)
}

//MediaAdapter 多媒体播放器适配器
type MediaAdapter struct {
	MusicPlayer AdvanceMediaPlayer
}

//NewMediaAdapter 初始化适配器
func NewMediaAdapter(audioType string) *MediaAdapter {
	switch audioType {
	case "rmvb":
		return &MediaAdapter{MusicPlayer: &RmvbPlayer{}}
	case "mp4":
		return &MediaAdapter{MusicPlayer: &Mp4Player{}}
	}
	return nil
}

//Play 根据文件类型，调用相应的适配器进行播放
func (mAdapter *MediaAdapter) Play(audioType, fileName string) {
	switch audioType {
	case "rmvb":
		mAdapter.MusicPlayer.PlayRmvb(fileName)
		break
	case "mp4":
		mAdapter.MusicPlayer.PlayMP4(fileName)
		break
	}
}
