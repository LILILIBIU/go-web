package live

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"
)

var client = &http.Client{
	Timeout: 10 * time.Second,
}
var strByte = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
var strByteLen = len(strByte)

func RandString(length int) string {

	bytes := make([]byte, length)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		bytes[i] = strByte[r.Intn(strByteLen)]
	}
	str := string(bytes)
	return str
}

// GetPassword 获取推流地址号
func GetPassword() (channel, room string) {
	room = RandString(64)
	str := String1{}
	url := "http://127.0.0.1:8090/control/get?room=" + room
	client := client
	resp, err := client.Get(url)
	if err != nil {
		log.Println("生成房间失败！")
		return "-1", "-1"
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	//解码
	err = json.Unmarshal(body, &str)
	if err != nil {
		return "-1", "-1"
	}
	channel = string(body[22 : len(body)-2])
	return
}
