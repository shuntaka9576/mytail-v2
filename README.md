# Usage
 * 標準的な使い方  
オプション-nで出力したい行数を指定可能(デフォルト:10)
```bash
> mytail -n=3 filename
```
* 空行を無視するモード  
オプション「-v=true」で空行を無視して結果を出力できます(デフォルト:false)  
※結果の改行コードは、LFで出力されます
```bash
> mytail -n=3 -v=true filename
```
# Build
```go
> go build
```
# Test
```go
> go test
```
