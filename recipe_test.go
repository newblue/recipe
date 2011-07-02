package recipe

import (
	"os"
	"testing"
	"gob"
)

func TestMe(t *testing.T) {
	if len(Recipe_list)==0 {
		t.Error(":(")
	}
	e := gob.NewEncoder(os.Stdout)
	//r := Recipe_list[0]
	e.Encode(Recipe_list)
}