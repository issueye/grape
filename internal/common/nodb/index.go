package nodb

import (
	"context"
	"fmt"
	"path/filepath"
	"sync"
	"time"

	"github.com/issueye/grape/internal/common/model"
	"github.com/issueye/grape/internal/global"
	"github.com/issueye/grape/pkg/db/nosql/leveldb"
	"github.com/issueye/grape/pkg/utils"
)

// 单例
var (
	IndexDB *leveldb.Index[model.TrafficStatistics]
)

var once sync.Once

func InitDB(ctx context.Context) {
	GetDB()

	Listen(ctx)
}

func GetDB() *leveldb.Index[model.TrafficStatistics] {
	if IndexDB == nil {
		once.Do(func() {
			// 初始化数据库
			IndexDB = NewDB()
		})
	}

	return IndexDB
}

// 获取数据库
func GetIndexDB() *leveldb.Index[model.TrafficStatistics] {
	if IndexDB == nil {
		GetDB()
	}

	return IndexDB
}

// 创建一个 leveldb 数据库
func NewDB() *leveldb.Index[model.TrafficStatistics] {
	httpMessagesPath := filepath.Join("runtime", "data", "http")
	return leveldb.NewIndex[model.TrafficStatistics](httpMessagesPath)
}

// 监听 global.IndexDB 管道传入的数据
func Listen(ctx context.Context) {
	// 监听全局变量 global.IndexDB 管道传入的数据
	go func() {
		for {
			select {
			case data := <-global.IndexDB:
				{
					if data == nil {
						continue
					}

					// int64 转 字节数组
					id := utils.GenID()
					// 使用 messages-时间戳—id 作为数据库的 key
					// 当前时间的毫秒级时间戳
					// 例如：messages-1678756000000-123456789

					byteID := fmt.Sprintf("messages-%d-%d", time.Now().UnixNano(), id)
					data.ID = byteID
					// 写入数据库
					if err := GetIndexDB().Put(byteID, data); err != nil {
						// 处理写入数据库的错误
						global.Log.Errorf("数据写入失败 %v", err)
						continue
					}
				}
			case <-ctx.Done():
				// 取消监听
				return
			}
		}
	}()
}
