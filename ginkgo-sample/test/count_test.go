package test

import (
	"ginkgo-sample/books"

	. "github.com/onsi/ginkgo/v2"

	. "github.com/onsi/gomega"
)

var _ = Describe("Books", func() {
	var foxInSocks, lesMis *books.Book

	BeforeEach(func() {
		lesMis = &books.Book{
			Title:  "Les Miserables",
			Author: "Victor Hugo",
			Pages:  2783,
		}

		foxInSocks = &books.Book{
			Title:  "Fox In Socks",
			Author: "Dr. Seuss",
			Pages:  24,
		}
	})

	Describe("Categorizing books", func() {
		Context("with more than 300 pages", func() {
			It("should be a novel", func() {
				Expect(lesMis.Category()).To(Equal(books.CategoryNovel))
			})
		})

		Context("with fewer than 300 pages", func() {
			It("should be a short story", func() {
				Expect(foxInSocks.Category()).To(Equal(books.CategoryShortStory))
			})
		})
	})
})
