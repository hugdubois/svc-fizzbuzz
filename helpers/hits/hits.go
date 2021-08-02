// Package hits provides a helper to nake some hit.
package hits

// this file is heavly inspired by "exprvar" from the standard go library

import (
	"fmt"
	"log"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
)

// hit is an abstract type for all exported variables.
type hit interface {
	// String returns a valid JSON value for the variable.
	// Types with String methods that do not return valid JSON
	// (such as time.Time) must not be used as a hit interface.
	String() string
}

// intVal is a 64-bit integer variable that satisfies the hit interface.
type intVal struct {
	i int64
}

func (v *intVal) String() string {
	return strconv.FormatInt(atomic.LoadInt64(&v.i), 10)
}

func (v *intVal) Add(delta int64) {
	atomic.AddInt64(&v.i, delta)
}

// Hits is a string-to-hit map variable that satisfies the hit interface.
type Hits struct {
	m      sync.Map // map[string]hit
	keysMu sync.RWMutex
	keys   []string // sorted
}

// keyValue represents a single entry in all hits.
type keyValue struct {
	Key   string
	Value hit
}

func (v *Hits) String() string {
	var b strings.Builder
	fmt.Fprintf(&b, "{")
	first := true
	v.Do(func(kv keyValue) {
		if !first {
			fmt.Fprintf(&b, ", ")
		}
		fmt.Fprintf(&b, "%q: %v", kv.Key, kv.Value)
		first = false
	})
	fmt.Fprintf(&b, "}")
	return b.String()
}

// Init removes all keys from the map.
func (v *Hits) Init() *Hits {
	v.keysMu.Lock()
	defer v.keysMu.Unlock()
	v.keys = v.keys[:0]
	v.m.Range(func(k, _ interface{}) bool {
		v.m.Delete(k)
		return true
	})
	return v
}

// addKey updates the sorted list of keys in v.keys.
func (v *Hits) addKey(key string) {
	v.keysMu.Lock()
	defer v.keysMu.Unlock()
	// Using insertion sort to place key into the already-sorted v.keys.
	if i := sort.SearchStrings(v.keys, key); i >= len(v.keys) {
		v.keys = append(v.keys, key)
	} else if v.keys[i] != key {
		v.keys = append(v.keys, "")
		copy(v.keys[i+1:], v.keys[i:])
		v.keys[i] = key
	}
}

// Add adds delta to the *intVal value stored under the given map key.
func (v *Hits) Add(key string, delta int64) {
	i, ok := v.m.Load(key)
	if !ok {
		var dup bool
		i, dup = v.m.LoadOrStore(key, new(intVal))
		if !dup {
			v.addKey(key)
		}
	}

	// Add to intVal; ignore otherwise.
	if iv, ok := i.(*intVal); ok {
		iv.Add(delta)
	}
}

// Delete deletes the given key from the map.
func (v *Hits) Delete(key string) {
	v.keysMu.Lock()
	defer v.keysMu.Unlock()
	i := sort.SearchStrings(v.keys, key)
	if i < len(v.keys) && key == v.keys[i] {
		v.keys = append(v.keys[:i], v.keys[i+1:]...)
		v.m.Delete(key)
	}
}

// Do calls f for each entry in the map.
// The map is locked during the iteration,
// but existing entries may be concurrently updated.
func (v *Hits) Do(f func(keyValue)) {
	v.keysMu.RLock()
	defer v.keysMu.RUnlock()
	for _, k := range v.keys {
		i, _ := v.m.Load(k)
		f(keyValue{k, i.(hit)})
	}
}

// Top returns the most hited key and its value
func (v *Hits) Top() (top string, count int64) {
	v.m.Range(func(key, value interface{}) bool {
		ival := value.(*intVal)
		i := ival.i
		if i > count {
			count = i
			top = key.(string)
		}

		return true
	})

	return
}

// All published variables.
var (
	vars      sync.Map // map[string]hit
	varKeysMu sync.RWMutex
	varKeys   []string // sorted
)

// Publish declares a named exported variable. This should be called from a
// package's init function when it creates its Vars. If the name is already
// registered then this will log.Panic.
func publish(name string, v hit) {
	if _, dup := vars.LoadOrStore(name, v); dup {
		log.Panicln("Reuse of exported var name:", name)
	}
	varKeysMu.Lock()
	defer varKeysMu.Unlock()
	varKeys = append(varKeys, name)
	sort.Strings(varKeys)
}

func NewHits(name string) *Hits {
	v := new(Hits).Init()
	publish(name, v)
	return v
}

// do calls f for each exported variable.
// The global variable map is locked during the iteration,
// but existing entries may be concurrently updated.
func do(f func(keyValue)) {
	varKeysMu.RLock()
	defer varKeysMu.RUnlock()
	for _, k := range varKeys {
		val, _ := vars.Load(k)
		f(keyValue{k, val.(hit)})
	}
}

// Handler is a http handler which returns the registered hits.
func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	fmt.Fprintf(w, "{\n")
	first := true
	do(func(kv keyValue) {
		if !first {
			fmt.Fprintf(w, ",\n")
		}
		first = false
		fmt.Fprintf(w, "%q: %s", kv.Key, kv.Value)
	})
	fmt.Fprintf(w, "\n}\n")
}
