// NOTE(lxf): Override the 'cm' package
//
//go:generate go tool wit-bindgen-go generate --world example --cm github.com/cosmonic-labs/cm-override/cm --out gen ./wit

package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/cosmonic-labs/cm-override/gen/example/http-server/encoding"
	"go.wasmcloud.dev/component/net/wasihttp"

	// NOTE(lxf): Note we use the custom 'cm' package
	"github.com/cosmonic-labs/cm-override/cm"
)

func init() {
	wasihttp.HandleFunc(indexHandler)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	vals := &encoding.TestRecord{
		StringValue: cm.Some("hello"),
		IntValue:    cm.Some(int32(42)),
		BoolValue:   cm.Some(true),
		ListValue:   cm.Some(cm.ToList([]string{"val1", "val2"})),
	}

	fmt.Fprintf(w, "marshal\n")
	json.NewEncoder(w).Encode(vals)

	enc, _ := json.Marshal(vals)
	decodedVals := &encoding.TestRecord{}
	_ = json.Unmarshal(enc, decodedVals)
	fmt.Fprintf(w, "\nunmarshal\n")
	json.NewEncoder(w).Encode(decodedVals)

	recnils := &encoding.TestRecord{
		StringValue: cm.None[string](),
		IntValue:    cm.None[int32](),
		BoolValue:   cm.None[bool](),
		ListValue:   cm.None[cm.List[string]](),
	}
	fmt.Fprintf(w, "\nmarshal with nulls\n")
	json.NewEncoder(w).Encode(recnils)

	enc, _ = json.Marshal(recnils)
	decodedRecnils := &encoding.TestRecord{}
	_ = json.Unmarshal(enc, decodedRecnils)
	fmt.Fprintf(w, "\nunmarshal with nulls\n")
	json.NewEncoder(w).Encode(decodedRecnils)
}

func main() {}
