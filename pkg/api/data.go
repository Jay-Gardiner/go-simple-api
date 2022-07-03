package api

type DataService interface{}

type DataRepository interface{}

type dataService struct {
	storage DataRepository
}

func NewDataService(dataRepo DataRepository) DataService {
	return &dataService{
		storage: dataRepo,
	}
}
