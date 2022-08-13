package scripts

import (
	"ga_tuner/utils/tools"
	"math/rand"
	"os"
	"testing"
	"time"
)

func TestVectorFactory(t *testing.T) {
	rng1 := rand.New(rand.NewSource(time.Now().UnixNano()))
	vec1 := Vector(InitBinaryFloat64(50, 0, 2, rng1))
	if len(vec1) <= 0 {
		t.Error("VECTOR LENGTH MUST LARGER THAN ZERO")
	}
}

func TestMatchBinaryWithFlags(t *testing.T) {
	rng1 := rand.New(rand.NewSource(time.Now().UnixNano()))
	vec1 := Vector(InitBinaryFloat64(50, 0, 2, rng1))
	availableFlags := tools.ReturnAllFlags("gcc-11")
	cmd, flag_dict := MatchBinaryWithFlags(vec1, "O2", availableFlags)
	if len(cmd) <= 0 {
		t.Error("CMD must be longer than lenght zero. Got : ", cmd)
	}
	if len(flag_dict) != 50 {
		t.Error("FLAG DICTIONARY DOES NOT CONTAINS ALL FLAGS.")
	}

}

// EN SON BURADA KALDIM. HATA VERMESINI BEKLERDIM AMA VERMEDI.
// os.ARGS[1]'e erisemiyor olmasi gerekiyor.
func TestCollectBaseline(t *testing.T) {
	os.Args[1] = "2mm"
	O2_BASELINE := CollectBaseline("O2")
	O3_BASELINE := CollectBaseline("O3")
	if O2_BASELINE <= 0 {
		t.Error("RUN TIME OF CODE MUST BE LARGER THAN ZERO")
	}
	if O3_BASELINE <= 0 {
		t.Error("RUN TIME OF CODE MUST BE LARGER THAN ZERO")
	}
}
