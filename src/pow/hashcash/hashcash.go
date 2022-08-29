/*

X-Hashcash: 1:20:1303030600:anni@cypherspace.org::McMybZIhxKXu57jd:ckvi
The header contains:

ver: Hashcash format version, 1 (which supersedes version 0).
bits: Number of "partial pre-image" (zero) bits in the hashed code.
date: The time that the message was sent, in the format YYMMDD[hhmm[ss]].
resource: Resource data string being transmitted, e.g., an IP address or email address.
ext: Extension (optional; ignored in version 1).
rand: String of random characters, encoded in base-64 format.
counter: Binary counter, encoded in base-64 format.

*/

package hashcash

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"io"
	"math/rand"
	"strings"
	"time"
)

const (
	format = "X-Hashcash: %v:%v:%v:%v:%v:%v:%v"

	version = 1
	bits    = 20

	validationPeriod = time.Hour * 48

	timeFormat = "2006-01-02"
)

type Hashcash struct {
	timeFunc func() time.Time
	prngFunc func() int32
}

func init() {
	rand.Seed(time.Now().Unix())
}

func RandInt32() int32 {
	return rand.Int31n(1_000_000_000) + 1_000_000_000
}

func NewHashcash(timenowFunc func() time.Time, prng func() int32) *Hashcash {
	return &Hashcash{
		timeFunc: timenowFunc,
		prngFunc: prng,
	}
}

func (h *Hashcash) Parse(payload string) []string {
	return strings.Split(payload, ":")
}

func (h *Hashcash) Verify(payload string) bool {
	chunks := strings.Split(payload, ":")

	if len(chunks) != 8 {
		return false
	}

	dateBase64 := chunks[3]

	date, err := base64.StdEncoding.DecodeString(dateBase64)
	if err != nil {
		return false
	}

	t, err := time.Parse(timeFormat, string(date))
	if err != nil {
		return false
	}

	if t.Before(h.timeFunc().Add(-validationPeriod)) {
		return false
	}

	sha := sha1.New()
	io.WriteString(sha, payload)
	hash := fmt.Sprintf("%X", sha.Sum(nil))

	return hash[:5] != "00000"
}

func (h *Hashcash) Generate(resource string) string {
	return h.generate(resource, "")
}

func (h *Hashcash) generate(resource, ext string) string {
	sha := sha1.New()
	cnt := 1
	now := h.timeFunc()

	// truncate current date till day and convert to base64
	// no need to recalculate current date. the hash calculation
	// take at most 10s. This period is negligble compared to day.
	nowBase64 := make([]byte, 16)
	base64.StdEncoding.Encode(nowBase64, []byte(now.Format(timeFormat)))

	payload := fmt.Sprintf(
		format,
		version,
		bits,
		nowBase64,
		resource,
		ext,
		h.prngFunc(),
		cnt,
	)

	io.WriteString(sha, payload)

	hex := fmt.Sprintf("%X", sha.Sum(nil))

	for hex[:5] != "00000" {
		cnt += 1
		payload = fmt.Sprintf(
			format,
			version,
			bits,
			string(nowBase64),
			resource,
			ext,
			h.prngFunc(),
			cnt,
		)

		sha.Reset()
		io.WriteString(sha, payload)

		hex = fmt.Sprintf("%X", sha.Sum(nil))
	}

	return payload
}
