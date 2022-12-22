package m3u

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

// 播放列表
type PlayList []Track

// 单个频道
// #EXTINF:-1 tvg-id="CCTVPlus1.cn" tvg-logo="https://i.imgur.com/AvYpKxL.png" group-title="News",CCTV+ 1 (600p) [Not 24/7]
type Track struct {
	Path  string
	Title string
	Time  int64
	// tvg-logo
	Logo string
	// tvg-id
	Id string
	// tvg-group
	Group string
}

// tvg-id="CCTVPlus1.cn"
func parseTvgAttr(s string) (string, string) {
	kv := strings.Split(s, "=")
	if len(kv) == 2 {
		return kv[0], strings.ReplaceAll(kv[1], "\"", "")
	}
	return "", ""
}

// 校验是否可用
func checker(allPl *PlayList) PlayList {
	pl := PlayList{}
	for _, t := range *allPl {
		if checkerPath(t.Path) {
			pl = append(pl, t)
		}
	}
	return pl
}

func checkerPath(path string) bool {
	return true
}

func Parse(r io.Reader) (PlayList, error) {
	pl := PlayList{}
	br := bufio.NewReader(r)

	for {
		line, err := br.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				return pl, nil
			}
			return pl, err
		}
		line = line[:len(line)-1]
		if len(line) > 8 && line[:8] == "#EXTINF:" {
			strs := strings.Split(line[8:], ",")
			if len(strs) == 2 {
				exts := strings.Split(strs[0], " ")
				ext_len := len(exts)
				if ext_len > 1 {
					id := ""
					logo := ""
					group := ""
					for i := 1; i < ext_len; i++ {
						k, v := parseTvgAttr(exts[i])
						if k == "tvg-id" {
							id = v
						}
						if k == "tvg-group" {
							group = v
						}
						if k == "tvg-logo" {
							logo = v
						}
					}
					ftime, err := strconv.ParseFloat(exts[0], 64)
					if err != nil {
						continue
					}
					time := int64(ftime)
					path, err := br.ReadString('\n')
					if err != nil {
						continue
					}
					pl = append(pl, Track{
						Path:  path,
						Time:  time,
						Id:    id,
						Group: group,
						Logo:  logo,
					})
				}
			}
		}
	}
}

// 保存可用列表到 csv 文件
func ToCsv(allPl *PlayList, f string) {
	pl := checker(allPl)
	for _, t := range pl {
		fmt.Printf("%s,%s,%s,%s,%s,%d\n", t.Id, t.Title, t.Group, t.Path, t.Logo, t.Time)
	}
}

func ToM3u(allPl *PlayList, f string) {}
