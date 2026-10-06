package columnar

import "fmt"

// Sample 是一个固定场景，用来复现题面里的锚点。
type Sample struct {
	Name string
	Run  func() map[string]any
}

func table2() *Table {
	return &Table{
		Rows: 2,
		Columns: []Column{
			{Name: "a", Dict: []string{"x", "y"}, Codes: []int{0, 1}},
			{Name: "b", Dict: []string{"p", "q"}, Codes: []int{0, 1}},
		},
	}
}

func names(row []KV) []string {
	out := []string{}
	for _, kv := range row {
		out = append(out, kv.Name)
	}
	return out
}

// Samples 返回全部场景。
func Samples() []Sample {
	return []Sample{
		{Name: "project", Run: func() map[string]any {
			rows := Scan(table2(), []string{"a"}, "a", "x", &Counter{})
			return map[string]any{"rows": len(rows), "cols": names(rows[0])}
		}},
		{Name: "missingcol", Run: func() map[string]any {
			rows := Scan(table2(), []string{"a"}, "zz", "x", &Counter{})
			return map[string]any{"rows": len(rows)}
		}},
		{Name: "dupdict", Run: func() map[string]any {
			t := &Table{Rows: 3, Columns: []Column{
				{Name: "a", Dict: []string{"x", "y", "x"}, Codes: []int{0, 1, 2}},
			}}
			rows := Scan(t, []string{"a"}, "a", "x", &Counter{})
			return map[string]any{"rows": len(rows)}
		}},
		{Name: "projorder", Run: func() map[string]any {
			rows := Scan(table2(), []string{"b", "a"}, "a", "x", &Counter{})
			return map[string]any{"cols": names(rows[0])}
		}},
	}
}

// Work 跑规模线场景。
func Work(n int) map[string]any {
	dict := make([]string, n)
	codes := make([]int, n)
	for i := 0; i < n; i++ {
		dict[i] = fmt.Sprintf("v%04d", i)
		codes[i] = i
	}
	t := &Table{Rows: n, Columns: []Column{{Name: "a", Dict: dict, Codes: codes}}}
	c := &Counter{}
	Scan(t, []string{"a"}, "a", fmt.Sprintf("v%04d", n-1), c)
	return map[string]any{"rows": n, "scanned": c.Scanned}
}
