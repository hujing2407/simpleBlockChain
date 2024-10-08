package BLC

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"log"
)

func IntToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}

	return buff.Bytes()
}

// Standard JSON to String Array
func JSONToArray(jsonString string) []string {

	var sArr []string
	if err := json.Unmarshal([]byte(jsonString), &sArr); err != nil {
		log.Panic(err)
	}
	//fmt.Println("====json to []string  ====")
	//fmt.Println(sArr)
	return sArr
}
