package snowflake

import (
	"log"
	"sync"
	"time"
)

const (
	epoch             = int64(1577808000000)                           // 设置起始时间(时间戳/毫秒)：2020-01-01 00:00:00，有效期 69 年
	timestampBits     = uint(41)                                       // 时间戳占用位数
	dataCenterIdBits  = uint(5)                                        // 数据中心 id 所占位数
	workerIdBits      = uint(5)                                        // 机器 id 所占位数
	sequenceBits      = uint(12)                                       // 序列所占的位数
	timestampMax      = int64(-1 ^ (-1 << timestampBits))              // 时间戳最大值
	dataCenterIdMax   = int64(-1 ^ (-1 << dataCenterIdBits))           // 支持的最大数据中心 id 数量
	workerIdMax       = int64(-1 ^ (-1 << workerIdBits))               // 支持的最大机器 id 数量
	sequenceMask      = int64(-1 ^ (-1 << sequenceBits))               // 支持的最大序列 id 数量
	workerIdShift     = sequenceBits                                   // 机器 id 左移位数
	dataCenterIdShift = sequenceBits + workerIdBits                    // 数据中心 id 左移位数
	timestampShift    = sequenceBits + workerIdBits + dataCenterIdBits // 时间戳左移位数
)

type Snowflake struct {
	sync.Mutex         // 锁
	timestamp    int64 // 时间戳 ，毫秒
	workerId     int64 // 工作节点
	dataCenterId int64 // 数据中心机房id
	sequence     int64 // 序列号
}

// NewSnowflake 创建一个雪花算法实例
// dataCenterId 数据中心 ID
// workerId 数据中心中机器 ID
func NewSnowflake(dataCenterId, workerId int64) *Snowflake {
	snowflake := &Snowflake{
		workerId:     workerId,
		dataCenterId: dataCenterId,
	}
	return snowflake
}

// NextVal 下一个 ID 值
func (s *Snowflake) NextVal() uint64 {
	s.Lock()                               // 上锁
	now := time.Now().UnixNano() / 1000000 // 转毫秒
	if s.timestamp == now {
		// 当同一时间戳（精度：毫秒）下多次生成id会增加序列号
		s.sequence = (s.sequence + 1) & sequenceMask
		if s.sequence == 0 {
			// 如果当前序列超出12bit长度，则需要等待下一毫秒
			// 下一毫秒将使用sequence:0
			for now <= s.timestamp {
				now = time.Now().UnixNano() / 1000000
			}
		}
	} else {
		// 不同时间戳（精度：毫秒）下直接使用序列号：0
		s.sequence = 0
	}
	t := now - epoch
	if t > timestampMax { // 如果当前时间戳大于所支持的最大时间戳
		s.Unlock()
		log.Printf("epoch must be between 0 and %d", timestampMax-1)
		return 0
	}
	s.timestamp = now
	r := uint64((t)<<timestampShift | (s.dataCenterId << dataCenterIdShift) | (s.workerId << workerIdShift) | (s.sequence))
	s.Unlock()
	return r
}
