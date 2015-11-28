package mustache

import (
	. "github.com/ricallinson/simplebdd"
	"testing"
)

func TestRender(t *testing.T) {

	Describe("Render()", func() {

		It("should use template from file", func() {
			m := Create()
			s, _ := m.Render("./fixtures/test.html", map[string]string{"title": "The Title"})
			AssertEqual(s, "<h1>The Title</h1>")
			AssertEqual(len(m.cache), 0)
		})

		It("should use template from cache", func() {
			m := Create(true)
			s, _ := m.Render("./fixtures/test.html", map[string]string{"title": "The Title"})
			AssertEqual(s, "<h1>The Title</h1>")
			AssertEqual(len(m.cache), 1)
		})

		It("should use template from file not cache", func() {
			m := Create(false)
			s, _ := m.Render("./fixtures/test.html", map[string]string{"title": "The Title"})
			AssertEqual(s, "<h1>The Title</h1>")
			AssertEqual(len(m.cache), 0)
		})

		It("should clean the cache", func() {
			m := Create(true)
			m.Render("./fixtures/test.html", map[string]string{"title": "The Title"})
			AssertEqual(len(m.cache), 1)
			m.Clean()
			AssertEqual(len(m.cache), 0)
		})
	})

	Report(t)
}
