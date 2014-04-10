package libvirt

import (
	"math"
	"math/rand"
	"testing"
)

func TestTypedParamsOpsBool(t *testing.T) {
	params := NewTypedParameters(1)

	if err := params.TypedParamsAddBool(true, "true"); err != nil {
		t.Errorf("Add boolean param failed: %#v", err)
	}

	if err := params.TypedParamsAddBool(true, "true"); err == nil {
		t.Errorf("Double add should fail")
	}

	if params.Len() != 1 {
		t.Errorf("Num params %d should be 1.", params.Len)
	}

	b, err := params.TypedParamsGetBool("true")
	if err != nil {
		t.Errorf("Get boolean param failed: %#v", err)
	}
	if !b {
		t.Errorf("Expected %v. Got %v", !b, b)
	}
}

func TestTypedParamsOpsFloat64(t *testing.T) {
	params := NewTypedParameters(1)

	if err := params.TypedParamsAddFloat64(math.Pi, "pi"); err != nil {
		t.Errorf("Add float64 param failed: %#v", err)
	}

	if err := params.TypedParamsAddFloat64(math.Pi, "pi"); err == nil {
		t.Errorf("Double add should fail")
	}

	if params.Len() != 1 {
		t.Errorf("Num params %d should be 1.", params.Len)
	}

	f, err := params.TypedParamsGetFloat64("pi")
	if err != nil {
		t.Errorf("Get float64 param failed: %#v", err)
	}
	if f != math.Pi {
		t.Errorf("Expected %v. Got %v", math.Pi, f)
	}

}

func TestTypedParamsOpsInt32(t *testing.T) {
	params := NewTypedParameters(1)

	pvalue := rand.Int31()
	if err := params.TypedParamsAddInt32(pvalue, "random"); err != nil {
		t.Errorf("Add int32 param failed: %#v", err)
	}

	if err := params.TypedParamsAddInt32(rand.Int31(), "random"); err == nil {
		t.Errorf("Double add should fail")
	}

	if params.Len() != 1 {
		t.Errorf("Num params %d should be 1.", params.Len)
	}
	v, err := params.TypedParamsGetInt32("random")
	if err != nil {
		t.Errorf("Get int32 param failed: %#v", err)
	}
	if v != pvalue {
		t.Errorf("Expected %v. Got %v", pvalue, v)
	}
}

func TestTypedParamsOpsInt64(t *testing.T) {
	params := NewTypedParameters(1)

	pvalue := rand.Int63()
	if err := params.TypedParamsAddInt64(pvalue, "random"); err != nil {
		t.Errorf("Add int64 param failed: %#v", err)
	}

	if err := params.TypedParamsAddInt64(rand.Int63(), "random"); err == nil {
		t.Errorf("Double add should fail")
	}

	if params.Len() != 1 {
		t.Errorf("Num params %d should be 1.", params.Len)
	}
	v, err := params.TypedParamsGetInt64("random")
	if err != nil {
		t.Errorf("Get int64 param failed: %#v", err)
	}
	if v != pvalue {
		t.Errorf("Expected %v. Got %v", pvalue, v)
	}
}

func TestTypedParamsOpsAddUInt32(t *testing.T) {
	params := NewTypedParameters(1)

	pvalue := uint32(rand.Int31())
	if err := params.TypedParamsAddUInt32(pvalue, "random"); err != nil {
		t.Errorf("Add uint32 param failed: %#v", err)
	}

	if err := params.TypedParamsAddUInt32(uint32(rand.Int31()), "random"); err == nil {
		t.Errorf("Double add should fail")
	}

	if params.Len() != 1 {
		t.Errorf("Num params %d should be 1.", params.Len)
	}
	v, err := params.TypedParamsGetUInt32("random")
	if err != nil {
		t.Errorf("Get uint32 param failed: %#v", err)
	}
	if v != pvalue {
		t.Errorf("Expected %v. Got %v", pvalue, v)
	}
}

func TestTypedParamsOpsUInt64(t *testing.T) {
	params := NewTypedParameters(1)

	pvalue := uint64(rand.Int63())
	if err := params.TypedParamsAddUInt64(pvalue, "random"); err != nil {
		t.Errorf("Add uint64 param failed: %#v", err)
	}

	if err := params.TypedParamsAddUInt64(uint64(rand.Int63()), "random"); err == nil {
		t.Errorf("Double add should fail")
	}

	if params.Len() != 1 {
		t.Errorf("Num params %d should be 1.", params.Len)
	}
	v, err := params.TypedParamsGetUInt64("random")
	if err != nil {
		t.Errorf("Get uint64 param failed: %#v", err)
	}
	if v != pvalue {
		t.Errorf("Expected %v. Got %v", pvalue, v)
	}
}
