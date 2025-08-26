# Go の map 基礎

## 1. map とは？

* **「キーと値」をペアで保存する箱**
* Python の dict、JavaScript の object に近い

例：

```go
ages := map[string]int{
    "Alice": 23,
    "Bob":   30,
}
```

* キーは `string` ("Alice" と "Bob")
* 値は `int` (23, 30)

---

## 2. 取り出す

```go
fmt.Println(ages["Alice"]) // 23
```

存在しないキーを指定したら？

```go
fmt.Println(ages["Carol"]) // 0 （int型のゼロ値）
```

---

## 3. 存在チェック

存在するかどうかを調べるときは **2値代入** を使う：

```go
value, ok := ages["Bob"]
fmt.Println(value, ok) // 30 true

_, ok = ages["Carol"]
fmt.Println(ok) // false
```

* `ok` が true なら存在する
* false なら存在しない

---

## 4. 追加と削除

```go
ages["Dave"] = 45    // 追加
delete(ages, "Alice") // 削除
```

---

# Bubble Tea の `map[int]struct{}`

ここで出てきたコードを思い出してください：

```go
selected map[int]struct{}
```

意味：

* キー → `int`（リストの番号を表す）
* 値 → `struct{}`（空の構造体）

---

## 5. なぜ `struct{}` なの？

* `struct{}` は **サイズ0** の特別な型（メモリを全く消費しない）
* つまり「値はどうでもいい、存在するかだけ管理したい」場合に使う

この場合 `selected` は「集合 (set)」のように使える。

---

## 6. 具体例

例えば

```go
selected := make(map[int]struct{})

// 2番目の要素を選択済みにする
selected[1] = struct{}{}

// チェックする
if _, ok := selected[1]; ok {
    fmt.Println("2番目は選択済み")
}
```

出力：

```
2番目は選択済み
```

---

# まとめ

* `map[キーの型]値の型` が基本形
* Bubble Tea では **「選択済みのインデックス集合」** を `map[int]struct{}` で表現している

  * `int` = リストの番号
  * `struct{}` = 値は不要（ただ存在だけ表せればいい）

---

課題：
map の理解を深めるために「自分で小さな **出席管理アプリ** を map で書いてみる」練習をしましょう。
例えば `"Alice" 出席済み` みたいに保存して表示する簡単なコード。

