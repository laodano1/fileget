package main
//
//import (
//	"fileget/util"
//	"github.com/howeyc/crc16"
//	redis "github.com/go-redis/redis/v8"
//)
//
//func NewRedisClient() {
//	rcl := redis.NewClusterClient()
//
//}
//
//func main() {
//	data := []byte("hello")
//	cksum := crc16.ChecksumCCITT(data)
//	util.Lg.Debugf("hello crc16 cksum: %v", cksum)
//
//	cksum = crc16.ChecksumCCITTFalse(data)
//	util.Lg.Debugf("hello crc16 cksum: %v", cksum)
//
//	cksum = crc16.ChecksumIBM(data)
//	util.Lg.Debugf("hello crc16 cksum: %v", cksum)
//
//	cksum = crc16.ChecksumMBus(data)
//	util.Lg.Debugf("hello crc16 cksum: %v", cksum)
//
//	cksum = crc16.ChecksumSCSI(data)
//	util.Lg.Debugf("hello crc16 cksum: %v", cksum)
//
//
//}
