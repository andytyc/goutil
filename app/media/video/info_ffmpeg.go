package video

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/andytyc/goutil/op/opos/opex"
)

// MediaFFprobe MediaFFprobe
type MediaFFprobe struct {
	Streams []*MediaStream
	Format  *MediaFormat
}

// MediaStream MediaStream
type MediaStream struct {
	Index     int    `json:"index"`
	CodecType string `json:"codec_type"`
	StartTime string `json:"start_time"`
	Duration  string `json:"duration"`
	BitRate   string `json:"bit_rate"`

	// video
	Width         int    `json:"width"`
	Height        int    `json:"height"`
	CodedWidth    int    `json:"coded_width"`
	CodedHeight   int    `json:"coded_height"`
	RealFrameRate string `json:"r_frame_rate"`
	AvgFrameRate  string `json:"avg_frame_rate"`
	TotalFrame    string `json:"nb_frames"`
}

// MediaFormat MediaFormat
type MediaFormat struct {
	FileName   string `json:"filename"`
	StreamNum  int    `json:"nb_streams"`
	FormatName string `json:"format_name"`
	StartTime  string `json:"start_time"`
	Duration   string `json:"duration"`
	Size       string `json:"size"`
	BitRate    string `json:"bit_rate"`
}

// GetMediaFileInfo 获取视频元数据
func GetMediaFileInfo(filename string) (ret *opex.Result, mediaInfo *MediaFFprobe, err error) {
	cmder := new(opex.Cmder)
	cmdline := "ffprobe -v quiet -print_format json -show_format -show_streams " + filename
	ret = cmder.Run(cmdline, 10)

	if ret.Err() != nil {
		err = fmt.Errorf("exec cmd failed :%s:%s", filename, ret.Err())
		return
	}

	mediaInfo = new(MediaFFprobe)
	buf := ret.Buf()
	err = json.Unmarshal(buf.Bytes(), mediaInfo)
	if err != nil {
		err = fmt.Errorf("exec cmd success but Unmarshal failed :%s:%s:%s", filename, ret, err)
		return
	}

	return
}

// GetVideoTotalFrame 获取视频总帧数
func GetVideoTotalFrame(mediaInfo *MediaFFprobe) (totalFrame int64, err error) {
	filename := mediaInfo.Format.FileName

	var videoFrameRateStr string
	for _, stream := range mediaInfo.Streams {
		// if stream.CodecType == "video" && stream.TotalFrame != "" {
		// 	return stream.TotalFrame
		// }
		if stream.CodecType == "video" {
			if stream.RealFrameRate == "" {
				err = fmt.Errorf("stream.RealFrameRate is empty :%s:%s", filename, stream.RealFrameRate)
				return
			}
			rateSlice := strings.Split(stream.RealFrameRate, "/")
			if len(rateSlice) != 2 {
				err = fmt.Errorf("stream.RealFrameRate is not format like 25/1 :%s:%s:%s", filename, stream.RealFrameRate, rateSlice)
				return
			}
			videoFrameRateStr = rateSlice[0]
			break
		}
	}
	if videoFrameRateStr == "" {
		err = fmt.Errorf("videoFrameRateStr is empty :%s", filename)
		return
	}
	if mediaInfo.Format.Duration == "" {
		err = fmt.Errorf("mediaInfo.Format.Duration is empty :%s", filename)
		return
	}

	videoDuration, err := strconv.ParseFloat(mediaInfo.Format.Duration, 64)
	if err != nil {
		err = fmt.Errorf("mediaInfo.Format.Duration strconv.Atoi failed :%s:%s", filename, err)
		return
	}
	videoFrameRate, err := strconv.ParseInt(videoFrameRateStr, 10, 64)
	if err != nil {
		err = fmt.Errorf("videoFrameRateStr strconv.Atoi failed :%s:%s", filename, err)
		return
	}

	totalFrame = int64(videoDuration * float64(videoFrameRate))
	return
}
