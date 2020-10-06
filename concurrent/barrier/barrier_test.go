package barrier

import (
	"fmt"
	"strings"
	"testing"
)

func TestBarrier(t *testing.T) {
	t.Run("Correct endpoints", func(t *testing.T) {
		endpoints := []string{
			"http://httpbin.org/headers",
			"http://httpbin.org/user-agent",
		}

		result := captureBarrierOutput(endpoints...)
		fmt.Println(!strings.Contains(result, "Accept-Encoding"))
		fmt.Println(!strings.Contains(result, "User-Agent"))
		if !strings.Contains(result, "Accept-Encoding") || !strings.Contains(result, "User-Agent") {
			t.Fail()
		}
		println(result)
	})

	t.Run("One incorrect endpoint", func(t *testing.T) {
		endpoints := []string{
			"http://httpbin.org/user-agent",
			"http://malformed-url",
		}

		result := captureBarrierOutput(endpoints...)
		if !strings.Contains(result, "ERROR") {
			t.Fail()
		}
		println(result)
	})

	t.Run("Very short timeout", func(t *testing.T) {
		endpoints := []string{
			"http://httpbin.org/headers",
			"http://httpbin.org/user-agent",
		}
		timeoutMilliseconds = 1
		result := captureBarrierOutput(endpoints...)
		if !strings.Contains(result, "Timeout") {
			t.Fail()
		}
		println(result)
	})
}
