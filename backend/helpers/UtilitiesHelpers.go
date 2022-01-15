package helpers

func StringLike(pointer *string) string {
	non_pointer := ""
	if pointer != nil {
		non_pointer = *pointer
	}
	return non_pointer
}
