package repo

type Db interface {
    Create(object interface{})
    FindAll(object interface{})
    Find(object interface{}, id string) error
    DeleteById(object interface{}, id string)
    Update(object interface{})
}
