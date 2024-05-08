package impl

import (
    "little-api/model"
    "little-api/repo"
)

type NoteRepository struct {
    repo.Db
}

func (db *NoteRepository) Store(n model.Note) {
    db.Create(&n)
}

func (db *NoteRepository) Select() []model.Note {
    notes := []model.Note{}
    db.FindAll(&notes)
    return notes
}

func (db *NoteRepository) SelectById(id string) *model.Note {
    note := model.Note{}
    err := db.Find(&note, id)
    if err != nil {
        return nil
    }
    return &note
}

func (db *NoteRepository) Modify(n model.Note) {
    db.Update(n)
}

func (db *NoteRepository) Delete(id string) int {
    user := []model.Note{}
    db.DeleteById(&user, id)
    return len(user)
}
