package uniqidkey

import (
	"crypto/md5"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type Service struct {
}

func (s *Service) GetNewUniqidKey() (key string, err error) {
	i64 := time.Now().UnixNano() + rand.Int63()
	bl := []byte(strconv.FormatInt(i64, 10))

	b2 := md5.New().Sum(bl)
	key := fmt.Sprintf("%x", b2)

	bo := NewUniqidKeyBO(0)

	bo.UniqidKey = key
	err = bo.Save()
	if err != nil {
		err = fmt.Errorf("save fail, err=%w", err)
		return "", err
	}
	return key, nil
}

func (s *Service) SaveMap(key string, data map[string]interface{}) (err error) {

}
