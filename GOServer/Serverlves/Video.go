package Serverlves

import (
	"GOServer/DaoModle"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type tmpVideoData struct {
	StrSize string
	Size    int64
	seek    int64
	tmpSize int64
	data    []byte
}

func pathSafe(path string) bool {
	var count int
	for _, v := range path {
		if v == '\\' {
			count++
		}
	}

	if count != 2 {
		return false
	}
	fmt.Println("yes")
	prefix := strings.HasPrefix(path, "..\\ShopVideo\\")
	if !prefix {
		return false
	}
	return true
}

//第一版 缓冲视频流
//设计一个带缓冲的文件流
//但有bug无法解决
//

// var tmpMap = sync.Map{}
//
//	func readVideo(remote, path string, seek int64) (*tmpVideoData, error) {
//		if !pathSafe(path) {
//			return nil, errors.New("路径错误")
//		}
//		open, err := os.Open(path)
//		if err != nil {
//			return nil, err
//		}
//		defer open.Close()
//		stat, err := open.Stat()
//		if err != nil {
//			return nil, err
//		}
//		tmpData := tmpVideoData{
//			StrSize: strconv.FormatInt(stat.Size()-1, 10),
//			Size:    stat.Size(),
//			seek:    seek,
//			tmpSize: 0,
//			data:    nil,
//		}
//		//_, err = open.Seek(seek, 0)
//		if err != nil {
//			return nil, err
//		}
//		if tmpData.Size-seek > (1<<20)*8 {
//			tmpData.tmpSize = (1 << 20) * 8
//			tmpData.data = make([]byte, (1<<20)*8)
//			_, err1 := open.Read(tmpData.data)
//			if err1 != nil {
//				return nil, err1
//			}
//			tmpMap.Store(remote+path, &tmpData)
//		} else {
//			tmpData.tmpSize = tmpData.Size - seek
//			tmpData.data = make([]byte, tmpData.tmpSize)
//			file, err := os.ReadFile(path)
//			if err != nil {
//				panic(err)
//			}
//			_, err1 := open.Read(tmpData.data)
//			if err1 != nil {
//				return nil, err1
//			}
//			if bytes.Compare(file, tmpData.data) == 0 {
//				fmt.Println("equal")
//			} else {
//				fmt.Println(bytes.Compare(file, tmpData.data))
//				fmt.Println("FUNK!!!!!!")
//			}
//			tmpMap.Store(path, &tmpData)
//		}
//		return &tmpData, nil
//	}
func findVideo(context *gin.Context) {
	value := context.Query("ShopId")
	shopId, _ := strconv.ParseUint(value, 10, 64)
	video := DaoModle.ShopVideo{
		ShopId: shopId,
	}
	err := DB.Where("shop_id = ?", video.ShopId).Find(&video).Error
	if err != nil {
		context.JSON(http.StatusOK, response{
			Request: false,
			Err:     NOTFIND,
		})
		return
	}
	context.JSON(http.StatusOK, video)
}

//
//func getVideo(context *gin.Context) {
//	path := context.Query("path")
//	RangeStr := context.GetHeader("Range")
//	if len(RangeStr) == 1 { // 用普通Get请求
//		RangeStr = "bytes=0-"
//	}
//	RangeNum := strings.Split(RangeStr, "=")[1]
//	RangeIdx := strings.Split(RangeNum, "-")[0]
//	seek, _ := strconv.ParseInt(RangeIdx, 10, 64)
//	var fileData *tmpVideoData
//	value, ok := tmpMap.Load(context.Request.RemoteAddr + path)
//	if ok {
//		fileData = value.(*tmpVideoData)
//		if fileData.Size-seek < (1<<20)*4 {
//			context.Writer.Header().Add("Content-Range", "bytes "+strconv.FormatInt(seek, 10)+"-"+strconv.FormatInt(fileData.Size-1, 10)+"/"+strconv.FormatInt(fileData.Size, 10))
//			context.Data(206, "video/mp4", fileData.data)
//			return
//		}
//		if seek+(1<<20) > fileData.seek+fileData.tmpSize || seek < fileData.seek {
//			_fileData, err1 := readVideo(context.Request.RemoteAddr, path, seek)
//			if err1 != nil {
//				context.JSON(http.StatusOK, response{
//					Request: false,
//					Err:     err1.Error(),
//				})
//				return
//			}
//			fileData = _fileData
//		}
//	} else {
//		_fileData, err1 := readVideo(context.Request.RemoteAddr, path, seek)
//		if err1 != nil {
//			context.JSON(http.StatusOK, response{
//				Request: false,
//				Err:     err1.Error(),
//			})
//			return
//		}
//		if _fileData.tmpSize < (1<<20)*4 {
//			context.Writer.Header().Add("Content-Range", "bytes "+strconv.FormatInt(seek, 10)+"-"+strconv.FormatInt(_fileData.Size-1, 10)+"/"+strconv.FormatInt(_fileData.Size, 10))
//			context.Data(206, "video/mp4", _fileData.data)
//			return
//		}
//		fileData = _fileData
//	}
//	var sendLen int64
//	var rangeRes = "bytes " + RangeIdx + "-"
//	if fileData.tmpSize >= 1<<20 {
//		rangeRes += strconv.FormatInt(seek+(1<<20-1), 10)
//		sendLen = 1 << 20
//	} else {
//		rangeRes += strconv.FormatInt(seek+fileData.tmpSize, 10)
//		sendLen = fileData.tmpSize - 1
//	}
//	fmt.Println(rangeRes, "  ", 1<<20, "  ", seek-fileData.seek+sendLen, "  ", len(fileData.data[seek-fileData.seek:seek-fileData.seek+sendLen]))
//	if fileData == nil {
//		panic("nil")
//	}
//	rangeRes += "/" + fileData.StrSize
//	context.Writer.Header().Add("Content-Range", rangeRes)
//	fmt.Println("walk")
//	context.Data(206, "video/mp4", fileData.data[seek-fileData.seek:seek-fileData.seek+sendLen])
//}

func getVideo(context *gin.Context) {
	path := context.Query("path")
	safe := pathSafe(path)
	if !safe {
		context.JSON(http.StatusOK, response{
			Request: false,
			Err:     VIDEOPATHERR,
		})
		return
	}
	file, err := os.ReadFile(path)
	if err != nil {
		context.JSON(http.StatusOK, response{
			Request: false,
			Err:     err.Error(),
		})
		return
	}
	RangeStr := context.GetHeader("Range")
	if len(RangeStr) == 1 { // 用普通Get请求
		RangeStr = "bytes=0-"
	}
	RangeNum := strings.Split(RangeStr, "=")[1]
	RangeIdx := strings.Split(RangeNum, "-")[0]
	seek, _ := strconv.ParseInt(RangeIdx, 10, 64)
	rep := "bytes "
	rep += strconv.FormatInt(seek, 10) + "-" + strconv.FormatInt(int64(len(file)-1), 10) + "/" + strconv.FormatInt(int64(len(file)), 10)
	context.Writer.Header().Add("Content-Range", rep)
	context.Data(206, "video/mp4", file[seek:])
}

func addVideo(context *gin.Context) {
	shopId := context.PostForm("ShopId")
	fmt.Println("ShopId: ", shopId)
	file, err := context.FormFile("video")
	if err != nil {
		context.JSON(http.StatusOK, response{
			Request: false,
			Err:     MODLEERR,
		})
		return
	}
	file.Filename = shopId + ".mp4"
	parseUint, err := strconv.ParseUint(shopId, 10, 64)
	if err != nil {
		context.JSON(http.StatusOK, response{
			Request: false,
			Err:     MODLEERR,
		})
		return
	}
	video := DaoModle.ShopVideo{
		ShopId: parseUint,
		Path:   "..\\ShopVideo\\" + file.Filename,
	}
	DB.Create(&video)
	err = context.SaveUploadedFile(file, "..\\ShopVideo\\"+file.Filename)
	if err != nil {
		context.JSON(http.StatusOK, response{
			Request: false,
			Err:     VIDEOPATHERR,
		})
	}
}

func init() {
	Engine.GET("/video/:opt", func(context *gin.Context) {
		context.Writer.Header().Add("Access-Control-Allow-Origin", CORS)
		param := context.Param("opt")
		if param == "findVideo" {
			findVideo(context)
			return
		} else if param == "getVideo" {
			getVideo(context)
			return
		}
	})
	Engine.POST("/video/addVideo", func(context *gin.Context) {
		context.Writer.Header().Add("Access-Control-Allow-Origin", CORS)
		addVideo(context)
	})
	accessOrigin("/video/:opt")
}
