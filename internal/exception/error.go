package exception

type DuplicateEntryError struct {
	Message string
}

func (e *DuplicateEntryError) Error() string {
	return e.Message
}

type BadRequestError struct {
	Message string
}

func (b *BadRequestError) Error() string {
	return b.Message
}
