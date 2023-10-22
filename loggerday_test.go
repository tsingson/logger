package logger

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zapcore"
)

func BenchmarkBufferedWriteSyncer(b *testing.B) {
	b.Run("write file with buffer", func(b *testing.B) {
		file, err := os.CreateTemp("", "log")
		require.NoError(b, err)

		defer func() {
			assert.NoError(b, file.Close())
			assert.NoError(b, os.Remove(file.Name()))
		}()

		w := &zapcore.BufferedWriteSyncer{
			WS: zapcore.AddSync(file),
		}
		defer w.Stop()
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				_, _ = w.Write([]byte("foobarbazbabble"))
			}
		})
	})
}
