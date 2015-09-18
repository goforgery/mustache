package mustache

import (
	. "github.com/ricallinson/simplebdd"
	"testing"
)

func TestRender(t *testing.T) {

	Describe("Render()", func() {

		It("should return <h1>The Title/h1>", func() {
			m := Create()
			s, _ := m.Render("./fixtures/test.html", map[string]string{"title": "The Title"})
			AssertEqual(s, "<h1>The Title</h1>")
		})
	})

	Report(t)
}
