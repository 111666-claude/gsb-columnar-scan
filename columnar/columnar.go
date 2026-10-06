// Package columnar 做列式字典编码扫描。
package columnar

// Counter 记录比较次数。
type Counter struct {
	Scanned int
}

// Column 是一个字典编码列。
type Column struct {
	Name  string
	Dict  []string
	Codes []int
}

// Table 是一张列式表。
type Table struct {
	Columns []Column
	Rows    int
}

// KV 是结果里的一格。
type KV struct {
	Name  string
	Value string
}

func findColumn(t *Table, name string, c *Counter) *Column {
	for i := range t.Columns {
		c.Scanned++
		if t.Columns[i].Name == name {
			return &t.Columns[i]
		}
	}
	return nil
}

// Scan 按谓词筛选并返回投影列。
func Scan(t *Table, project []string, predCol, predVal string, c *Counter) [][]KV {
	col := findColumn(t, predCol, c)
	if col == nil {
		return [][]KV{}
	}
	predCodes := map[int]bool{}
	for code, text := range col.Dict {
		c.Scanned++
		if text == predVal {
			predCodes[code] = true
		}
	}
	projCols := make([]*Column, 0, len(project))
	for _, name := range project {
		if pc := findColumn(t, name, c); pc != nil {
			projCols = append(projCols, pc)
		}
	}
	out := [][]KV{}
	for i := 0; i < t.Rows; i++ {
		c.Scanned++
		if !predCodes[col.Codes[i]] {
			continue
		}
		row := make([]KV, 0, len(projCols))
		for _, pc := range projCols {
			row = append(row, KV{Name: pc.Name, Value: pc.Dict[pc.Codes[i]]})
		}
		out = append(out, row)
	}
	return out
}
