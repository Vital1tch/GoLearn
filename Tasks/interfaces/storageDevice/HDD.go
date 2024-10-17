package storageDevice

type HDD struct {
	Data string
}

func (h *HDD) ReadData() string {
	return h.Data
}

func (h *HDD) WriteData(text string) string {
	h.Data = text
	return text
}
