package books

type Catalog int8

const (
	CategoryNovel      Catalog = iota
	CategoryShortStory Catalog = iota
)
const MaxShortStoryPages = 300

type Book struct {
	Title  string
	Author string
	Pages  int32
}

func (b Book) Category() Catalog {
	if b.Pages < MaxShortStoryPages {
		return CategoryShortStory
	} else {
		return CategoryNovel
	}
}
