package doubles

type FakeSearchEngine struct {
	db map[string]string
}

func (d *FakeSearchEngine) SearchByName(name string) string {
	return d.db[name]
}

func (d *FakeSearchEngine) SearchByPhone(phone string) string {
	for name, value := range d.db {
		if value == phone {
			return name
		}
	}
	return ""
}

func (d *FakeSearchEngine) AddEntry(name, phone string) error {
	d.db[name] = phone
	return nil
}
