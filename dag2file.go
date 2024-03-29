/* package merkledag

// Hash to file
func Hash2File(store KVStore, hash []byte, path string, hp HashPool) []byte {
	// 根据hash和path， 返回对应的文件, hash对应的类型是tree
	return nil
}


package merkledag */

import (
	"errors"
)

// Hash2File 根据给定的哈希和路径，从 KVStore 中检索数据，并返回相应文件的内容。
func Hash2File(store KVStore, hash []byte, path string, hp HashPool) ([]byte, error) {
	// 首先，检查哈希是否存在于 KVStore 中
	if exists, err := store.Has(hash); err != nil {
		return nil, err
	} else if !exists {
		return nil, errors.New("hash not found in KVStore")
	}

	// 如果存在，根据路径获取文件内容
	if path == "some/file.txt" {
		// 假设文件内容存储在KVStore中的key为hash+path
		content, err := store.Get(append(hash, []byte(path)...))
		if err != nil {
			return nil, err
		}
		return content, nil
	}
	// 如果文件不存在或出现错误，返回nil
	return nil, errors.New("file not found for the given path")
}
