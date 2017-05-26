package string_concat

import (
	"testing"
	"text/template"
	"bytes"
	"fmt"
)

func BenchmarkSelfConcatOperator2000(b *testing.B) {
	for n := 0; n < b.N; n++ {
		SelfConcatOperator("test", 1000)
	}
}

func BenchmarkTemplateParallel(b *testing.B) {
	templ := template.Must(template.New("test").Parse("Hello, {{.}}!"))
	b.SetParallelism(1)
	b.RunParallel(func(pb *testing.PB) {
		var buf bytes.Buffer
		for pb.Next() {
			buf.Reset()
			templ.Execute(&buf, "World")
		}
	})
}

func ExampleSalutations() {
	fmt.Println("hello, and")
	fmt.Println("goodbye")
	// Output:
	// hello, and
	// goodbye
}