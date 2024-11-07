package doubles

type MockSearchEngine struct {
	SearchByNameWasCalled  bool
	SearchByPhoneWasCalled bool
}

func (d *MockSearchEngine) SearchByName(name string) string {
	d.SearchByNameWasCalled = true
	return "12345678912"
}

func (d *MockSearchEngine) SearchByPhone(phone string) string {
	d.SearchByPhoneWasCalled = true
	return "fulano"
}

func (d *MockSearchEngine) AddEntry(name, phone string) error {
	return nil
}
