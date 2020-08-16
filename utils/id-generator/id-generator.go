package id_generator

import "github.com/google/uuid"

func New() uuid.UUID {
	return uuid.New()
}
