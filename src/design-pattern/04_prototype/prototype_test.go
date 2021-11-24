package prototype

import (
	"testing"
	"time"

	. "github.com/stretchr/testify/assert"
)

func TestKeywords_Clone(t *testing.T) {
	updateAt, _ := time.Parse("2006", "2020")
	keyword := &Keyword{
		Word:      "testA",
		Visit:     1,
		UpdatedAt: &updateAt,
	}
	newKeyWord := keyword.Clone()

	Equal(t, newKeyWord, keyword)

}
