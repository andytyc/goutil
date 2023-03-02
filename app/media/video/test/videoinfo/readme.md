# 例子

```bash
# 背景
# 运行环境已经安装好ffmpeg, 包中需要用到:ffprobe

# 编译
# mac
go build -ldflags "-s -w"
# linux
GOOS=linux GOARCH=amd64 go build -ldflags "-s -w"

# 运行
./videoinfo -f ./test.mp4

'''
2023/03/02 19:43:14 FlagFileName : ./test.mp4
2023/03/02 19:43:14 file : ./test.mp4 run command ret : {
    "streams": [
        {
            "index": 0,
            "codec_name": "h264",
            "codec_long_name": "H.264 / AVC / MPEG-4 AVC / MPEG-4 part 10",
            "profile": "High",
            "codec_type": "video",
            "codec_time_base": "1/50",
            "codec_tag_string": "avc1",
            "codec_tag": "0x31637661",
            "width": 1920,
            "height": 1080,
            "coded_width": 1920,
            "coded_height": 1088,
            "has_b_frames": 0,
            "pix_fmt": "yuvj420p",
            "level": 42,
            "color_range": "pc",
            "color_space": "bt709",
            "color_transfer": "bt709",
            "color_primaries": "bt709",
            "chroma_location": "left",
            "refs": 1,
            "is_avc": "true",
            "nal_length_size": "4",
            "r_frame_rate": "25/1",
            "avg_frame_rate": "25/1",
            "time_base": "1/1000",
            "start_pts": 0,
            "start_time": "0.000000",
            "duration_ts": 60040,
            "duration": "60.040000",
            "bit_rate": "4010131",
            "bits_per_raw_sample": "8",
            "nb_frames": "1501",
            "disposition": {
                "default": 1,
                "dub": 0,
                "original": 0,
                "comment": 0,
                "lyrics": 0,
                "karaoke": 0,
                "forced": 0,
                "hearing_impaired": 0,
                "visual_impaired": 0,
                "clean_effects": 0,
                "attached_pic": 0,
                "timed_thumbnails": 0
            },
            "tags": {
                "creation_time": "2021-12-13T01:01:40.000000Z",
                "language": "und",
                "handler_name": "VideoHandler"
            }
        }
    ],
    "format": {
        "filename": "./test.mp4",
        "nb_streams": 1,
        "nb_programs": 0,
        "format_name": "mov,mp4,m4a,3gp,3g2,mj2",
        "format_long_name": "QuickTime / MOV",
        "start_time": "0.000000",
        "duration": "60.040000",
        "size": "30108861",
        "bit_rate": "4011840",
        "probe_score": 100,
        "tags": {
            "major_brand": "isom",
            "minor_version": "512",
            "compatible_brands": "isomiso2avc1mp41",
            "creation_time": "2021-12-13T01:01:40.000000Z"
        }
    }
}
2023/03/02 19:43:14 funcGetVideoInfo 文件大小 30108861
2023/03/02 19:43:14 funcGetVideoInfo 文件总帧数 1501
'''
```
