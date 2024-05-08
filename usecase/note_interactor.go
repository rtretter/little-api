package usecase

import (
    "little-api/model"
    "little-api/repo"
)

type NoteInteractor struct {
    NoteRepository repo.NoteRepository
}

func (interactor *NoteInteractor) Add(n model.Note) {
    interactor.NoteRepository.Store(n)
}

func (interactor *NoteInteractor) GetInfo() []model.Note {
    return interactor.NoteRepository.Select()
}

func (interactor *NoteInteractor) GetInfoById(id string) *model.Note {
    return interactor.NoteRepository.SelectById(id)
}

func (interactor *NoteInteractor) Update(n model.Note) {
    interactor.NoteRepository.Modify(n)
}

func (interactor *NoteInteractor) Delete(id string) int {
    return interactor.NoteRepository.Delete(id)
}
