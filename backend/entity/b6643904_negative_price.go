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
		Price: 55,
		Code:  "BK1234568785",
	}

	ok, err := govalidtor.Struct(books)
	g.Expect(ok).To(BeFalse())
	g.Expect(err).ToNot(BeNil())
	g.Expect(err.Error().ContainSubString("Code"))

}
