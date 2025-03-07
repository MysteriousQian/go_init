package ip

import (
	"encoding/csv"
	"fmt"
	"io"
	"strconv"

	"os"

	"github.com/kentik/patricia"
	"github.com/kentik/patricia/string_tree"
)

/*
Geo地理位置信息

	ID: 地理位置ID
	Country: 国家名称 中国
	Province: 国家一级行政区 例如河南省
	City: 城市名称 例如 郑州
*/
type LocationInfo struct {
	Country  string
	Province string
	City     string
}

/*
全局IP前缀树
*/
var ipTree string_tree.TreeV4

/*
全局地理位置信息
*/
var locationMap map[int64]LocationInfo

/*
初始化自动加载IP库到内存
*/
func init() {
	ipTree = *string_tree.NewTreeV4()
	locationMap = make(map[int64]LocationInfo)
	fmt.Print("开始加载IP前缀树\n")
	err := loadIpTree()
	if err != nil {
		panic(err)
	}
	fmt.Print("IP前缀树加载成功\n")
	fmt.Print("开始加载地理位置信息...\n")
	err = loadLocation()
	if err != nil {
		panic(err)
	}
	fmt.Print("地理位置加载成功\n")
}

/*
加载ipv4.csv文件到内存
*/
func loadIpTree() (err error) {
	reader, file, err := loadCsvFile("./locales/ipv4.csv")
	defer file.Close()
	if err != nil {
		return
	}
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		if record[1] == "" {
			continue
		}
		ip, _, err := patricia.ParseIPFromString(record[0])
		if err != nil {
			return err
		}
		ok, _ := ipTree.Add(*ip, record[1], nil)
		if !ok {
			return fmt.Errorf("failed to add data to the tree ip:[%v]", record[0])
		}
	}
	return
}

/*
加载地理位置信息csv文件到内存
*/
func loadLocation() (err error) {
	reader, file, err := loadCsvFile("./locales/ipLocation-en.csv")
	defer file.Close()
	if err != nil {
		return
	}
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		Geoid, err := strconv.ParseInt(record[0], 10, 32)
		if err != nil {
			return err
		}
		locationMap[Geoid] = LocationInfo{
			Country:  record[4],
			Province: record[7],
			City:     record[10],
		}
	}
	return
}

/*
FindIpLoction 根据IP查找IP归属地

	param:
		ip: IP地址
	return:
		ipinfo: IP归属地信息 LocationInfo
		err: 错误信息
*/
func FindIpLoction(ip string) (ipIdInfo LocationInfo, err error) {
	var ipinfo LocationInfo
	a, _, err := patricia.ParseIPFromString(ip)
	if err != nil {
		return ipIdInfo, err
	}
	if a == nil {
		return ipIdInfo, fmt.Errorf("invalid IP address: %s", ip)
	}
	ok, v := ipTree.FindDeepestTag(*a)
	if ok {
		value, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return ipIdInfo, err
		}
		ipinfo = locationMap[value]
		return ipinfo, nil
	}

	return
}

/*
读取CSV文件

	param:
		filepath: 文件路径
	return:
		reader: 读取器
		err: 错误信息
*/
func loadCsvFile(filepath string) (reader *csv.Reader, file *os.File, err error) {
	file, err = os.Open(filepath)
	if err != nil {
		return
	}
	reader = csv.NewReader(file) // 读取CSV文件
	reader.Comma = ','           // 设置分隔符为制表符
	_, _ = reader.Read()         // 跳过第一行
	return
}

func main() {
	ipinfo, err := FindIpLoction("104.194.79.86")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("IP归属地信息: 国家ID:%s,一级行政区ID:%s,城市ID:%s", ipinfo.Country, ipinfo.Province, ipinfo.City)
}
