package widgets

import (
	"testing"

	"fyne.io/fyne/v2/test"
	"github.com/stretchr/testify/assert"
)

func TestProgressBuilder_Default(t *testing.T) {
	bar := NewProgressBuilder().Build()
	assert.Equal(t, float32(0.5), bar.progress)
	assert.Equal(t, float32(3.0), bar.borderRadius)
	assert.Equal(t, float32(100), bar.size.Width)
	assert.Equal(t, float32(5), bar.size.Height)
}

func TestProgressBar_SetProgress(t *testing.T) {
	bar := NewProgressBuilder().Progress(0).Build()
	assert.Equal(t, float32(0.0), bar.progress)

	bar.SetProgress(0.5)
	assert.Equal(t, float32(0.5), bar.progress)
}

func TestProgressBar_CreateRenderer(t *testing.T) {
	bar := NewProgressBuilder().Progress(0).Build()
	r := bar.CreateRenderer()
	test.WidgetRenderer(bar)
	assert.NotNil(t, r)
}
