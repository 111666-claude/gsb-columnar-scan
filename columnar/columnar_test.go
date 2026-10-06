package columnar

import "testing"

func TestScanSingleMatch(t *testing.T) {
	tbl := &Table{Rows: 1, Columns: []Column{{Name: "a", Dict: []string{"x"}, Codes: []int{0}}}}
	rows := Scan(tbl, []string{"a"}, "a", "x", &Counter{})
	if len(rows) != 1 || rows[0][0].Value != "x" {
		t.Fatalf("应命中一行 a=x：%+v", rows)
	}
}

func TestScanNoMatch(t *testing.T) {
	tbl := &Table{Rows: 1, Columns: []Column{{Name: "a", Dict: []string{"x"}, Codes: []int{0}}}}
	if rows := Scan(tbl, []string{"a"}, "a", "z", &Counter{}); len(rows) != 0 {
		t.Fatalf("不该命中：%+v", rows)
	}
}

func TestRowCount(t *testing.T) {
	if table2().Rows != 2 {
		t.Fatal("表应有两行")
	}
}

func TestProjectedValue(t *testing.T) {
	rows := Scan(table2(), []string{"a"}, "a", "x", &Counter{})
	if len(rows) != 1 || rows[0][0].Name != "a" {
		t.Fatalf("投影首列应是 a：%+v", rows)
	}
}

func TestCounterCounts(t *testing.T) {
	c := &Counter{}
	Scan(table2(), []string{"a"}, "a", "x", c)
	if c.Scanned == 0 {
		t.Fatal("应统计比较次数")
	}
}

func TestEmptyTable(t *testing.T) {
	tbl := &Table{Rows: 0}
	if rows := Scan(tbl, []string{"a"}, "a", "x", &Counter{}); len(rows) != 0 {
		t.Fatalf("空表不该有结果：%+v", rows)
	}
}
