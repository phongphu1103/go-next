package mappers

// MapSlice maps a slice of pointers, e.g. []*models.Tag -> []*domain.Tag.
func MapSlice[T any, U any](in []*T, fn func(*T) *U) []*U {
	out := make([]*U, 0, len(in))
	for _, item := range in {
		out = append(out, fn(item))
	}
	return out
}

// MapValues maps a slice of values to a slice of pointers,
// e.g. []models.Genre (as GORM loads it) -> []*domain.Genre.
// A nil input (relation not loaded) stays nil.
func MapValues[T any, U any](in []T, fn func(*T) *U) []*U {
	if in == nil {
		return nil
	}
	out := make([]*U, 0, len(in))
	for i := range in {
		out = append(out, fn(&in[i]))
	}
	return out
}

// MapToValues maps a slice of pointers to a slice of values, skipping nils,
// e.g. []*domain.Genre -> []models.Genre. A nil input stays nil.
func MapToValues[T any, U any](in []*T, fn func(*T) *U) []U {
	if in == nil {
		return nil
	}
	out := make([]U, 0, len(in))
	for _, item := range in {
		if u := fn(item); u != nil {
			out = append(out, *u)
		}
	}
	return out
}
