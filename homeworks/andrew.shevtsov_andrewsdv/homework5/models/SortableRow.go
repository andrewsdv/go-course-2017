package models

type SortableRow []Figure

func (it SortableRow) Len() int {
	return len(it)
}
func (it SortableRow) Swap(i, j int) {
	it[i], it[j] = it[j], it[i]
}
func (it SortableRow) Less(i, j int) bool {
	return (it[i]).Weight() > (it[j]).Weight()
}
