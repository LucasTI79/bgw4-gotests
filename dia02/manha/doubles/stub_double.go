package doubles

type StubSearchEngine struct{}

func (d *StubSearchEngine) SearchByName(name string) string {
	return ""
}

func (d *StubSearchEngine) SearchByPhone(phone string) string {
	return "fulano"
}

func (d *StubSearchEngine) AddEntry(name, phone string) error {
	return nil
}
