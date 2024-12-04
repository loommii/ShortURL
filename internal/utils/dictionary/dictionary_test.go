package dictionary

import (
	"testing"
)

func Test_dictionary_Get(t *testing.T) {
	t.Run("测试场景", func(t *testing.T) {
		dictionary := NewDictionary()
		for i := 0; i < 100; i++ {
			t.Log(dictionary.Get())
		}
	})
}
