package hashcash

import (
	"crypto/sha1"
	"fmt"
	"io"
	"testing"
	"time"
)

func TestGenerate(t *testing.T) {
	// valid payload
	// 1:20:MjAyMjA4Mjg=:a.maksadbek@gmail.com::1496475703:1010340

	timeFunc := func() time.Time {
		return time.Date(2022, time.August, 28, 0, 0, 0, 0, time.UTC)
	}

	randFunc := func() int32 {
		return 1496475703
	}

	hc := NewHashcash(timeFunc, randFunc)

	payload := hc.Generate("a.maksadbek@gmail.com", "")

	sha := sha1.New()
	io.WriteString(sha, payload)
	hash := fmt.Sprintf("%X", sha.Sum(nil))

	if hash[:5] != "00000" {
		t.Fatalf("invalid hash value %q for %q", hash, payload)
	}
}

func TestVerifySuccess(t *testing.T) {
	timeFunc := func() time.Time {
		return time.Date(2022, time.August, 28, 0, 0, 0, 0, time.UTC)
	}

	hc := NewHashcash(timeFunc, RandInt32)
	if !hc.Verify("X-Hashcash:1:20:MjAyMjA4Mjg=:a.maksadbek@gmail.com::1496475703:1010340") {
		t.Fatalf("invalid verification algorithm")
	}
}

func TestVerifyFailure(t *testing.T) {
	timeFunc := func() time.Time {
		return time.Date(2022, time.August, 28, 0, 0, 0, 0, time.UTC)
	}

	hc := NewHashcash(timeFunc, RandInt32)

	if hc.Verify("X-Hashcash:1:20:MjAyMjA4MjA=:a.maksadbek@gmail.com::1496475703:1713061") {
		t.Fatalf("expired token")
	}
}
