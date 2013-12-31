package libvirt

import (
    "math"
    "math/rand"
    "testing"
)

func TestTypedParamsAddBool(t *testing.T) {
    params := new(TypedParameters)

    if err := params.TypedParamsAddBool(true, "true"); err != nil {
        t.Errorf("Add boolean param failed: %#v", err)
    }

    if err := params.TypedParamsAddBool(true, "true"); err == nil {
        t.Errorf("Double add should fail")
    }
    if params.Len() != 1 {
        t.Errorf("Num params %d should be 1.", params.Len)
    }
}

func TestTypedParamsAddFloat64(t *testing.T) {
    params := new(TypedParameters)

    if err := params.TypedParamsAddFloat64(math.Pi, "pi"); err != nil {
        t.Errorf("Add boolean param failed: %#v", err)
    }

    if err := params.TypedParamsAddFloat64(math.Pi, "pi"); err == nil {
        t.Errorf("Double add should fail")
    }
    if params.Len() != 1 {
        t.Errorf("Num params %d should be 1.", params.Len)
    }
}

func TestTypedParamsAddInt32(t *testing.T) {
    params := new(TypedParameters)

    if err := params.TypedParamsAddInt32(rand.Int31(), "random"); err != nil {
        t.Errorf("Add boolean param failed: %#v", err)
    }

    if err := params.TypedParamsAddInt32(rand.Int31(), "random"); err == nil {
        t.Errorf("Double add should fail")
    }
    if params.Len() != 1 {
        t.Errorf("Num params %d should be 1.", params.Len)
    }
}

func TestTypedParamsAddInt64(t *testing.T) {
    params := new(TypedParameters)

    if err := params.TypedParamsAddInt64(rand.Int63(), "random"); err != nil {
        t.Errorf("Add boolean param failed: %#v", err)
    }

    if err := params.TypedParamsAddInt64(rand.Int63(), "random"); err == nil {
        t.Errorf("Double add should fail")
    }
    if params.Len() != 1 {
        t.Errorf("Num params %d should be 1.", params.Len)
    }
}

func TestTypedParamsAddUInt32(t *testing.T) {
    params := new(TypedParameters)

    if err := params.TypedParamsAddUInt32(uint32(rand.Int31()), "random"); err != nil {
        t.Errorf("Add boolean param failed: %#v", err)
    }

    if err := params.TypedParamsAddUInt32(uint32(rand.Int31()), "random"); err == nil {
        t.Errorf("Double add should fail")
    }
    if params.Len() != 1 {
        t.Errorf("Num params %d should be 1.", params.Len)
    }
}

func TestTypedParamsAddUInt64(t *testing.T) {
    params := new(TypedParameters)

    if err := params.TypedParamsAddUInt64(uint64(rand.Int63()), "random"); err != nil {
        t.Errorf("Add boolean param failed: %#v", err)
    }

    if err := params.TypedParamsAddUInt64(uint64(rand.Int63()), "random"); err == nil {
        t.Errorf("Double add should fail")
    }
    if params.Len() != 1 {
        t.Errorf("Num params %d should be 1.", params.Len)
    }
}


func TestTypedParamsGetBool(t *testing.T) {}
func TestTypedParamsGetFloat64(t *testing.T) {}
func TestTypedParamsGetInt32(t *testing.T) {}
func TestTypedParamsGetInt64(t *testing.T) {}
func TestTypedParamsGetUInt32(t *testing.T) {}
func TestTypedParamsGetUInt64(t *testing.T) {}
