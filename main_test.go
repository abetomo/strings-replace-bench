package main

import (
	"strings"
	"testing"
)

const str = "ABCDEFGHIJKLMNOPQRSTUVWXYZ abcdefghijklmnopqrstuvwxyz ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func BenchmarkStringsReplace(b *testing.B) {
	b.Run("2 chars", func(b *testing.B) {
		b.Run("Replace all", func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				s := strings.Replace(str, "A", "1", -1)
				strings.Replace(s, "C", "3", -1)
			}
		})

		b.Run("Replace some", func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				s := strings.Replace(str, "A", "1", -1)
				strings.Replace(s, "1", "3", -1)
			}
		})
	})

	b.Run("2 strings", func(b *testing.B) {
		b.Run("Replace all", func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				s := strings.Replace(str, "ABC", "123", -1)
				strings.Replace(s, "DEF", "456", -1)
			}
		})

		b.Run("Replace some", func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				s := strings.Replace(str, "ABC", "123", -1)
				strings.Replace(s, "456", "___", -1)
			}
		})
	})

	b.Run("5 chars", func(b *testing.B) {
		b.Run("Replace all", func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				s := strings.Replace(str, "A", "1", -1)
				s = strings.Replace(s, "C", "3", -1)
				s = strings.Replace(s, "H", "8", -1)
				s = strings.Replace(s, "N", "14", -1)
				strings.Replace(s, "Z", "26", -1)
			}
		})

		b.Run("Replace some", func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				s := strings.Replace(str, "A", "1", -1)
				s = strings.Replace(s, "1", "3", -1)
				s = strings.Replace(s, "H", "8", -1)
				s = strings.Replace(s, "2", "14", -1)
				strings.Replace(s, "Z", "26", -1)
			}
		})
	})

	b.Run("5 strings", func(b *testing.B) {
		b.Run("Replace all", func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				s := strings.Replace(str, "ABC", "123", -1)
				s = strings.Replace(s, "DEF", "456", -1)
				s = strings.Replace(s, "HIJ", "8910", -1)
				s = strings.Replace(s, "NOP", "141516", -1)
				strings.Replace(s, "XYZ", "242526", -1)
			}
		})

		b.Run("Replace some", func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				s := strings.Replace(str, "ABC", "123", -1)
				s = strings.Replace(s, "456", "___", -1)
				s = strings.Replace(s, "HIJ", "8910", -1)
				s = strings.Replace(s, "141516", "___", -1)
				strings.Replace(s, "XYZ", "242526", -1)
			}
		})
	})
}

func BenchmarkStringsReplacer(b *testing.B) {
	b.Run("2 chars", func(b *testing.B) {
		b.Run("Replace all.", func(b *testing.B) {
			oldnew := []string{
				"A", "1",
				"C", "3",
			}
			r := strings.NewReplacer(oldnew...)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				r.Replace(str)
			}
		})

		b.Run("Replace some.", func(b *testing.B) {
			oldnew := []string{
				"A", "1",
				"1", "3",
			}
			r := strings.NewReplacer(oldnew...)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				r.Replace(str)
			}
		})
	})

	b.Run("2 strings", func(b *testing.B) {
		b.Run("Replace all.", func(b *testing.B) {
			oldnew := []string{
				"ABC", "123",
				"DEF", "456",
			}
			r := strings.NewReplacer(oldnew...)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				r.Replace(str)
			}
		})

		b.Run("Replace some.", func(b *testing.B) {
			oldnew := []string{
				"ABC", "123",
				"456", "___",
			}
			r := strings.NewReplacer(oldnew...)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				r.Replace(str)
			}
		})
	})

	b.Run("5 chars", func(b *testing.B) {
		b.Run("Replace all.", func(b *testing.B) {
			oldnew := []string{
				"A", "1",
				"C", "3",
				"H", "8",
				"N", "14",
				"Z", "26",
			}
			r := strings.NewReplacer(oldnew...)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				r.Replace(str)
			}
		})

		b.Run("Replace some.", func(b *testing.B) {
			oldnew := []string{
				"A", "1",
				"1", "3",
				"H", "8",
				"2", "14",
				"Z", "26",
			}
			r := strings.NewReplacer(oldnew...)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				r.Replace(str)
			}
		})
	})

	b.Run("5 strings", func(b *testing.B) {
		b.Run("Replace all.", func(b *testing.B) {
			oldnew := []string{
				"ABC", "123",
				"DEF", "456",
				"HIJ", "8910",
				"NOP", "141516",
				"XYZ", "242526",
			}
			r := strings.NewReplacer(oldnew...)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				r.Replace(str)
			}
		})

		b.Run("Replace some.", func(b *testing.B) {
			oldnew := []string{
				"ABC", "123",
				"456", "___",
				"HIJ", "8910",
				"141516", "___",
				"XYZ", "242526",
			}
			r := strings.NewReplacer(oldnew...)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				r.Replace(str)
			}
		})
	})
}
