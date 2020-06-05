package service
import (
	"backend_template/models"
	"backend_template/repository"
)

type PersonService struct {
	BaseService
	repository.PersonRepository
}
// NewPersonService PersonService的构造函数
func NewPersonService(p repository.PersonRepository) PersonService {
	return PersonService{PersonRepository: p}
}


//FindAll Fetch all person data

func (s *PersonService) FindAll() []models.Person {
	var persons [] models.Person
	persons=s.PersonRepository.FindAll()

	return persons
}

func (p *PersonService) FindByID(id uint) models.Person {
	return p.PersonRepository.FindByID(id)
}
func (p *PersonService) Save(person models.Person) models.Person {
	p.PersonRepository.Save(person)

	return person
}

func (p *PersonService) Delete(person models.Person) {
	p.PersonRepository.Delete(person)
}