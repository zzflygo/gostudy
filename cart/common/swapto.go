package common

import (
	"encoding/json"
	"go-micro.dev/v4/util/log"
)

func SwapTo(in, out interface{}) error {
	byteStr, err := json.Marshal(in)
	if err != nil {
		log.Fatal("json marshal failed err:", err)
		return err
	}
	err = json.Unmarshal(byteStr, out)
	if err != nil {
		log.Fatal("json unmarshal failed err:", err)
		return err
	}
	return nil
}
