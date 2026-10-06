# columnar-scan

列式表的字典编码扫描：按谓词列筛选，再按投影列出结果。只用标准库。

```
go test ./...
go vet ./...
go run ./cmd/columnar --sample=project
go run ./cmd/columnar --sample=missingcol
go run ./cmd/columnar --sample=dupdict
go run ./cmd/columnar --sample=projorder
go run ./cmd/columnar --sample=work
```

## 口径

- **投影**：只输出请求的列，且按请求顺序；请求里没有的列不出现。
- **谓词列缺失**：谓词列不在表里时返回空结果（不是全表）。
- **字典**：字典里同一个值可能对应多个编码，凡是编码映射到该值的行都要命中。
- **筛选**：按行序输出命中的行。

## 不变量

- 结果里每行的列恰好是投影列，顺序一致。
- 命中行的判定只看谓词列的编码，与其它列无关。
- `scanned` 不随行数与字典大小相乘放大：两千行的 `scanned` 不超过 6000。

## 输出契约

场景打印一行 JSON：`cols` 是命中的列名列表（只含投影列），`rows` 是命中行数，`values` 是首行的值。
`--sample=work` 打印 `{"rows": N, "scanned": N}`。
