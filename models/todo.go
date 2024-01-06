package models

import (
	"FShare/dao"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"math/rand"
	"net"
	"net/http"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

type File struct {
	FileID      string `json:"id" gorm:"primary_key"`
	Name        string `json:"name"`
	FileOwner   string `json:"fileOwner"`
	Description string `json:"description"`
	Size        string `json:"size"`
	Time        string `json:"time"`
	Hash        string `json:"hash"`
	Status      int    `json:"status"` //1:没被申请；2：正在被申请中；3：申请被拒绝；4：可用不可转发；5：可用可转发
}

type Apply struct {
	ApplyOwner string `json:"applyOwner" gorm:"primary_key"`
	FileOwner  string `json:"fileOwner"`
	Time       string `json:"time"`
	FileID     string `json:"id" gorm:"primary_key"`
	FileName   string `json:"fileName"`
	Hash       string `json:"txHash"`
	Status     int    `json:"status"`
	IsHandled  bool   `json:"isHandled"`
}

type Applyrecord struct {
	FileID      string `json:"id"`
	Hash        string `json:"txHash" gorm:"primary_key" `
	FingerPrint string `json:"fingerprint"`
	Status      string `json:"status"`
}

type Hashdata struct {
	Result struct {
		Tx struct {
			Execer  string `json:"execer"`
			Payload struct {
				ContentStorage struct {
					Content interface{} `json:"content"`
					Key     string      `json:"key"`
					Op      int         `json:"op"`
					Value   string      `json:"value"`
				} `json:"contentStorage"`
				Ty int `json:"ty"`
			} `json:"payload"`
			RawPayload string `json:"rawPayload"`
			Signature  struct {
				Ty        int    `json:"ty"`
				Pubkey    string `json:"pubkey"`
				Signature string `json:"signature"`
			} `json:"signature"`
			Fee    int    `json:"fee"`
			Feefmt string `json:"feefmt"`
			Expire int    `json:"expire"`
			Nonce  int    `json:"nonce"`
			From   string `json:"from"`
			To     string `json:"to"`
			Hash   string `json:"hash"`
		} `json:"tx"`
	} `json:"result"`
}

var IP = gin.H{
	"A": "124.223.171.19", //王钺程
	"B": "101.43.94.172",  //李炳翰
	"C": "124.221.254.11", //金严
	"D": "124.223.210.53", //叶克炉
	"E": "124.222.196.78", //唐聪
	"F": "10.96.92.7",     //kxq
	"G": "10.96.208.18",   //wyc
	"Y": "10.0.4.14",      //云服务器
}

var Ip2Node = gin.H{
	"124.223.171.19": "A",
	"101.43.94.172":  "B",
	"124.221.254.11": "C",
	"124.223.210.53": "D", //叶克炉
	"124.222.196.78": "E", //唐聪
	"10.96.92.7":     "F", //kxq
	"10.96.208.18":   "G", //wyc
	"10.0.4.14":      "Y", //云服务器
}

var Node string //节点

/*
Todo这个model的增删改查放在这里
*/

func GetHostIp() string {
	conn, err := net.Dial("udp", "8.8.8.8:53")
	if err != nil {
		fmt.Println("get current host ip err: ", err)
		return ""
	}
	addr := conn.LocalAddr().(*net.UDPAddr)
	ip := strings.Split(addr.String(), ":")[0]
	return ip
}

func DownloadFile(context *gin.Context, node, fileName string) (err error) {
	str := fmt.Sprintf("%v", IP[node])
	address := "http://" + str + ":8080/download/" + fileName
	//fmt.Println(str)
	context.Redirect(http.StatusMovedPermanently, address)
	return

}

//func Download(context *gin.Context, fileName string) (err error) {
//	dst := fmt.Sprintf("./%s", fileName) //todo:这里修改文件路径
//	context.Header("Content-Type", "application/octet-stream")
//	context.Header("Content-Disposition", "attachment; filename="+fileName)
//	//context.Header("Content-Transfer-Encoding", "binary")
//	context.File(dst)
//	return
//}

func Download(context *gin.Context, fileName string) (err error) {
	// 构建文件路径
	var fN = strings.Split(fileName, ".")

	dst := fmt.Sprintf("./csvfile/%s_FP.csv", fN[0]) // 修改为正确的文件路径

	// 打开文件
	file, err := os.Open(dst)
	if err != nil {
		// 处理错误
		context.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
		return err
	}
	defer file.Close()

	// 设置响应头
	context.Header("Content-Disposition", "attachment; filename="+fN[0]+"_FP.csv")
	context.Header("Content-Type", "application/octet-stream")
	context.File(dst)
	// 将文件内容传递给客户端
	//_, err = io.Copy(context.Writer, file)
	//if err != nil {
	//	// 处理错误
	//	context.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
	//	return err
	//}

	return nil
}

func UploadFiles(context *gin.Context) (err error) {
	var file File

	//file.FileOwner = context.PostForm("fileOwner") //todo：后面改成Node，这里测试不同节点用
	file.FileOwner = Node
	file.Name = context.PostForm("name")
	file.Description = context.PostForm("description")
	file.Size = context.PostForm("size")
	file.Status, _ = strconv.Atoi(context.PostForm("status"))

	f, err := context.FormFile("f1")
	f.Filename = file.Name //+".xml" todo:这里更改文件名可以加xml后缀
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		//保存读取的文件到本地服务器
		dst := path.Join("./csvfile", f.Filename) //todo:这里修改文件路径
		_ = context.SaveUploadedFile(f, dst)
		context.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
		//生成文件ID
		t := time.Now()
		file.Time = t.Format("2006-01-02 15:04:05")
		timestamp := strconv.FormatInt(t.UTC().UnixNano(), 10)
		randnum := fmt.Sprintf("%04v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(10000))
		file.FileID = Node + timestamp + randnum
		//file.Status = 1
		fileproperties := file.FileID + "#" + file.Name + "#" + file.FileOwner + "#" + file.Description + "#" + file.Size + "#" + file.Time + "#" + strconv.Itoa(file.Status)
		fmt.Println(fileproperties)
		file.Hash = transfer("file", fileproperties)
		//file.Fingerprint = GenertaeFingerPrint(file)
		if err = dao.DB.Create(&file).Error; err != nil {
			return err
		}
	}
	return
}

func CreateApply(file *File) (err error) {
	var apply Apply
	apply.ApplyOwner = Node
	apply.FileOwner = file.FileOwner
	apply.FileID = file.FileID
	apply.FileName = file.Name
	apply.Status = 2
	apply.IsHandled = false
	t := time.Now()
	apply.Time = t.Format("2006-01-02 15:04:05")
	if err = dao.DB.Create(&apply).Error; err != nil {
		return err
	}
	return
}

func GetAllFile() (fileList []*File, err error) {
	if err = dao.DB.Where("file_owner <> ?", Node).Find(&fileList).Error; err != nil {
		return nil, err
	}
	return
}

func GetFileList() (fileList []*File, err error) {
	if err = dao.DB.Where("file_owner = ?", Node).Find(&fileList).Error; err != nil {
		return nil, err
	}
	return
}

func GetMyApply() (applyList []*Apply, err error) {
	if err = dao.DB.Where("apply_owner = ?", Node).Find(&applyList).Error; err != nil {
		return nil, err
	}
	return
}

func GetApplyList() (applyList []*Apply, err error) {
	if err = dao.DB.Where("file_owner = ?", Node).Find(&applyList).Error; err != nil {
		return nil, err
	}
	return
}

func GetFileByID(id string) (file *File, err error) {
	file = new(File)
	if err = dao.DB.Where("file_id = ?", id).First(&file).Error; err != nil {
		return nil, err
	}
	return
}

func GetApply(id, owner string) (apply *Apply, err error) {
	apply = new(Apply)
	if err = dao.DB.Where("file_id = ? and apply_owner = ?", id, owner).First(&apply).Error; err != nil {
		return nil, err
	}
	return
}

func UpdateFile(file *File) (err error) {
	err = dao.DB.Save(file).Error
	if err != nil {
		return err
	}
	return

}

func EmbedFingerprint(applyHash, fileName string) (string, error) {
	cmd := exec.Command("D:\\Reaserch\\System development\\project\\FShare\\venv\\Scripts\\python.exe", "python/embed.py", fileName, applyHash)

	output, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	result := string(output)
	if result == "false" {
		return "", errors.New("embed error")
	}

	return result, nil
}

func UpdateApply(apply *Apply, applyrecord *Applyrecord) (err error) {

	t := time.Now()
	apply.Time = t.Format("2006-01-02 15:04:05")
	applyproperties := apply.ApplyOwner + "#" + apply.FileOwner + "#" + apply.Time + "#" + apply.FileID + "#" + strconv.Itoa(apply.Status)
	apply.Hash = transfer("apply", applyproperties)
	apply.IsHandled = true
	err = dao.DB.Save(apply).Error
	if err != nil {
		return err
	}

	//todo:根据status=4的时候，可用不可转发
	FP, err := EmbedFingerprint(apply.Hash, apply.FileName)
	if err != nil {
		return err
	}
	applyrecord.FingerPrint = FP
	applyrecord.FileID = apply.FileID
	applyrecord.Hash = apply.Hash
	applyrecord.Status = strconv.Itoa(apply.Status)
	err = dao.DB.Save(applyrecord).Error
	if err != nil {
		return err
	}
	return

}

func FileIsExisted(filename string) bool {
	existed := true
	if _, err := os.Stat("csvfile/" + filename); os.IsNotExist(err) {
		existed = false
	}
	return existed
}

func DeleteAFileByID(id string) (err error) {
	var file File
	//删除数据库条目
	err = dao.DB.Where("file_id=?", id).Find(&file).Error
	if err != nil {
		return err
	}
	//删除本地源文件
	err = os.Remove("csvfile/" + file.Name)
	if err != nil {
		return err
	}
	//删除嵌入了水印的文件
	fileFP := strings.Split(file.Name, ".")
	fileFP_Name := fileFP[0] + "_FP.csv"
	if FileIsExisted(fileFP_Name) == true {
		fmt.Println("delete success")
		err = os.Remove("csvfile/" + fileFP_Name)
		if err != nil {
			return err
		}
	}

	err = dao.DB.Where("file_id=?", id).Delete(&file).Error
	if err != nil {
		return err
	}
	return
}

func DeleteApply(id, owner string) (err error) {
	err = dao.DB.Where("file_id=? and apply_owner=?", id, owner).Delete(&Apply{}).Error
	if err != nil {
		return err
	}
	return
}

// 将核验文件保存到核验缓存区
func SaveFilelocal(context *gin.Context) (Filetype string, err error) {
	// 将上传的文件取出来
	f, err := context.FormFile("FileVerify")
	if err != nil {
		context.JSON(http.StatusOK, gin.H{"message": err.Error()})
		return "", err
	}

	// 获取原始文件名和文件后缀
	originalFileName := f.Filename
	fileExtension := filepath.Ext(originalFileName)

	// 构建新的文件名
	newFileName := "verifyfile" + fileExtension

	// 设置文件名称为新的文件名
	f.Filename = newFileName

	// 将上传的文件保存到指定的本地地址，并返回响应
	log.Println(f.Filename)
	dst := fmt.Sprintf("./verifyfile/%s", f.Filename) // 设置核验文件保存的本地地址路径
	if err = context.SaveUploadedFile(f, dst); err != nil {
		return "", err
	}
	return fileExtension, nil
}

// 获取上传的核验文件的路径
func GetVerifyFile(filetype string) (FilePath string, err error) {
	dirPth := fmt.Sprintf("./verifyfile")
	fis, err := ioutil.ReadDir(filepath.Clean(filepath.ToSlash(dirPth)))
	if err != nil {
		return "", err
	}

	for _, f := range fis {
		_path := filepath.Join(dirPth, f.Name())
		// 指定格式
		if filepath.Ext(f.Name()) == filetype {
			FilePath = _path
			break // 一旦找到匹配的文件，就退出循环
		}
	}
	return FilePath, nil
}

func ExtractFingerPrint(filePath string) (string, string, error) {
	fmt.Println(filePath)
	cmd := exec.Command("D:\\Reaserch\\System development\\project\\FShare\\venv\\Scripts\\python.exe", "python/extract.py", filePath)

	output, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	result := string(output)
	if result == "false" {
		return "", "", errors.New("embed error")
	}
	fmt.Println(result)
	applyrecord := new(Applyrecord)
	if err = dao.DB.Where("finger_print=?", result).Find(applyrecord).Error; err != nil {
		return applyrecord.Hash, result, err
	}
	return applyrecord.Hash, result, nil
}

// 获取上传的核验文件的哈希值
//func FindtxHash(fingerprint string) (file *File, err error) {
//	file = new(File)
//	if err = dao.DB.Where("fingerprint=?", fingerprint).First(&file).Error; err != nil {
//		return nil, err
//	}
//	return
//}

func AnalyzeData(txHash string) (Hashdata, error) {
	encodeddata := queryTx(txHash)
	//定义结构体以匹配JSON数据
	var decodedata Hashdata
	err := json.Unmarshal(encodeddata, &decodedata)
	if err != nil {
		return decodedata, err
	}
	return decodedata, nil
}

func In(FileOwner string, checkNode []string) (err error) {
	for i := range checkNode {
		if FileOwner == checkNode[i] {
			return errors.New("FileOwner made invalid charge")
		}
	}
	return nil
}

func Verify(applydatalist []Hashdata, sourceNode string) ([]Hashdata, error) {
	//checkNode := sourceNode
	//NewcheckNode := ""
	//for init := 1; init <= len(applydatalist); init++ {
	//	for j := range applydatalist {
	//		applymessages := strings.Split(applydatalist[j].Result.Tx.Payload.ContentStorage.Value, "#")
	//		applyOwner := applymessages[0]
	//		if applyOwner == checkNode || applyOwner == NewcheckNode {
	//			if applyOwner == NewcheckNode {
	//				checkNode = NewcheckNode
	//			}
	//			switch applymessages[4] {
	//			case "4":
	//				applydatalist[j].Result.Tx.Payload.ContentStorage.Value += "#fail"
	//			case "5":
	//				applydatalist[j].Result.Tx.Payload.ContentStorage.Value += "#success"
	//			default:
	//				applydatalist[j].Result.Tx.Payload.ContentStorage.Value += "#fail"
	//				return applydatalist, errors.New("invalid param")
	//			}
	//			NewcheckNode = applymessages[1]
	//		}
	//	}
	//}
	var checkNode []string
	for j := range applydatalist {
		applymessages := strings.Split(applydatalist[j].Result.Tx.Payload.ContentStorage.Value, "#")
		if applymessages[4] == "4" {
			checkNode = append(checkNode, applymessages[0])
		}
	}
	for k := range applydatalist {
		applymessages := strings.Split(applydatalist[k].Result.Tx.Payload.ContentStorage.Value, "#")
		err := In(applymessages[1], checkNode)
		if err != nil {
			applydatalist[k].Result.Tx.Payload.ContentStorage.Value += "#fail"
		} else {
			applydatalist[k].Result.Tx.Payload.ContentStorage.Value += "#success"
		}
	}
	return applydatalist, nil
}

func TraceBackOnChain(txHash string, sourceNode string) ([]Hashdata, Hashdata, [][]string, error) {
	//todo: 传hash值，进行查询文件信息，进行错误验证。根据文件id查询申请哈希，得到申请hash，用一个数组存储循环查询申请记录
	//1.通过文件hash，查询文件信息,取出文件ID
	//定义结构体
	var filedata Hashdata
	var applydatalist []Hashdata

	fileRecord := new(Applyrecord)
	if err := dao.DB.Where("hash=?", txHash).Find(&fileRecord).Error; err != nil {
		return applydatalist, filedata, nil, err
	}
	//进行查询文件信息
	file := new(File)
	if err := dao.DB.Where("file_id=?", fileRecord.FileID).Find(&file).Error; err != nil {
		return applydatalist, filedata, nil, err
	}

	filedata, err := AnalyzeData(file.Hash)
	if err != nil {
		fmt.Println("JSON解析错误:", err)
		return applydatalist, filedata, nil, err
	}
	//2.取出文件ID,根据文件ID查询申请哈希，得到申请哈希，用一个数组存储循环查询申请记录
	applylist := new([]Applyrecord)
	if err = dao.DB.Where("file_id=?", file.FileID).Find(&applylist).Error; err != nil {
		return applydatalist, filedata, nil, err
	}
	//循环查询申请记录
	for i := range *applylist {
		apply := (*applylist)[i]
		applydata, _ := AnalyzeData(apply.Hash)
		applydatalist = append(applydatalist, applydata)
	}
	//对申请信息进行核验
	applydatalist_verified, _ := Verify(applydatalist, sourceNode)

	//简化追溯信息
	var checkdata [][]string
	for k := range applydatalist_verified {
		var data []string
		applydatalist_verified_message := applydatalist_verified[k].Result.Tx.Payload.ContentStorage.Value
		applydatalist_verified_messages := strings.Split(applydatalist_verified_message, "#")
		data = append(data, applydatalist_verified_messages[1])
		data = append(data, applydatalist_verified_messages[0])
		data = append(data, applydatalist_verified_messages[2])
		data = append(data, applydatalist_verified_messages[4])
		data = append(data, applydatalist_verified_messages[5])
		checkdata = append(checkdata, data)
	}
	return applydatalist_verified, filedata, checkdata, err
}
