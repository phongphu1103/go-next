package mappers

func MapSlice[T any, U any](in []*T, fn func(*T) *U) []*U {
    out := make([]*U, 0, len(in))
    for _, item := range in {
        out = append(out, fn(item))
    }
    return out
}
