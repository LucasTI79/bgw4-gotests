package doubles

type SpySearchEngine struct {
	SearchByPhoneWasCalled bool
}

func (d *SpySearchEngine) SearchByName(name string) string {
	return ""
}

func (d *SpySearchEngine) SearchByPhone(phone string) string {
	d.SearchByPhoneWasCalled = true
	return "fulano"
}

func (d *SpySearchEngine) AddEntry(name, phone string) error {
	return nil
}
