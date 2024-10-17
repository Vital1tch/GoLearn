package storageDevice

type SSD struct {
	Data string
}

func (s *SSD) ReadData() string {
	return s.Data
}

func (s *SSD) WriteData(text string) string {
	s.Data = text
	return text
}
