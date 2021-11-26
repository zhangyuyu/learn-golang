package adapter

import "testing"

func TestMediaAdapter(t *testing.T) {
	adapter := NewMediaAdapter("mp4")
	adapter.Play("mp4", "File")
}
