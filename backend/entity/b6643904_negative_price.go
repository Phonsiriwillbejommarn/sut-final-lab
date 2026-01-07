package entity

import (
	"testing"

	"github.com/asaskevich/govalidtor"
	. "github.com/onsi/gomega"
)

func TestnegativePrice(t *testing.T) {
	g := NewGomegaWithT(t)
	books := Books{
		Title: "Jommarn",
		Price: 0,
		Code:  "BK123456",
	}

	ok, err := govalidtor.Struct(books)
	g.Expect(ok).To(BeFalse())
	g.Expect(err).ToNot(BeNil())
	g.Expect(err.Error().ContainSubString())

}
