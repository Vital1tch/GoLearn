package storageDevice

type StorageDevice interface {
	ReadData() string
	WriteData(string) string
}
