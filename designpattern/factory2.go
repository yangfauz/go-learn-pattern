package designpattern

// import "io"

// type Store interface {
// 	Open(string) (io.ReadWriteCloser, error)
// }

// type StorageType int

// const (
// 	DiskStorage StorageType = 1 << iota
// 	TempStorage
// 	MemoryStorage
// )

// func NewStore(t StorageType) Store {
// 	switch t {
// 	case MemoryStorage:
// 		return newMemoryStorage( /*...*/ )
// 	case DiskStorage:
// 		return newDiskStorage( /*...*/ )
// 	default:
// 		return newTempStorage( /*...*/ )
// 	}
// }

// s, _ := data.NewStore(data.MemoryStorage)
// f, _ := s.Open("file")

// n, _ := f.Write([]byte("data"))
// defer f.Close()
