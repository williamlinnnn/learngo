## 以下敘述的解答?
```go
"\"Hello\\"" + ` \"World\"`
```

1. "Hello" "World"
2. "Hello" \"World\" *CORRECT*
3. "Hello" `"World"`
4. "\"Hello\" `\"World\"`"

> **1:** Go 不會解讀原始字串（raw string literal）中的跳脫字元
>
> **2:** 沒錯，Go 會將 \" 解讀為 "，但對於 \"World\" 則不會這麼處理
>


## 在程式碼中表示以下文字的最佳方式是什麼？
```xml
<xml>
  <items>
    <item>"Teddy Bear"</item>
  </items>
</xml>
```

1. *CORRECT*
```go
`<xml>
    <items>
        <item>"Teddy Bear"</item>
    </items>
</xml>`
```

2. 
```go
"<xml>
    <items>
        <item>"Teddy Bear"</item>
    </items>
</xml>"
```

3. 
```go
"<xml>
    <items>
        <item>\"Teddy Bear\"</item>
    </items>
</xml>"
```

4. 
```go
`<xml>
    <items>
        <item>\"Teddy Bear\"</item>
    </items>
</xml>`
```

> **2-3:** 你不能這樣寫字串字面值，它不能跨多行
>
> **4:** 在原始字串（raw string literal）中，不需要使用跳脫字元
>


## 以下敘述的解答?
```go
len("lovely")
```

1. 7
2. 8
3. 6 *CORRECT*
4. 0

> **2:** 記住! "a" 是 1 char. `a` 也是 1 char.
>


## 以下敘述的解答?
```go
len("very") + len(`\"cool\"`)
```

1. 8
2. 12 *CORRECT*
3. 16
4. 10

> **1:** 還有雙引號，也要計算它們
>
> **2:** 沒錯，Go 不會解讀原始字串（raw string literal）中的 \"
>
> **3:** 記住！ "very" 是 4 個字元，`very` 也是 4 個字元
>
> **4:** 記住！Go 不會解讀原始字串（raw string literal）中的 \"
>


## 以下敘述的解答?
```go
len("very") + len("\"cool\"")
```

1. 8
2. 12
3. 16
4. 10 *CORRECT*

> **1:** 還有雙引號，也要計算它們
>
> **2:** 記住！Go 會解讀字串字面值中的跳脫字元
>
> **4:** 沒錯，Go 會解讀字串字面值中的 \"。因此，"\"" 表示 "，這是一個字元
>


## 以下敘述的解答?
```go
len("péripatéticien")
```

**提示:** é 佔 2 個位元組，而 len 函數計算的是位元組數，而不是字母數

**無用的資訊:** "péripatéticien" 意思是 "wanderer".

1. 14
2. 16 *CORRECT*
3. 18
4. 20

> **1:** 記住！é 佔 2 個位元組
>
> **2:** 英文字母佔 1 個位元組，但 é 佔 2 個位元組，所以總共是 16 個位元組。很酷吧。
>
> **3:** 你沒有算上雙引號，對吧？
>


## 如何取得這個字串字面值中實際字元的正確長度？
```go
"péripatéticien"
```

1. `len(péripatéticien)`
2. `len("péripatéticien")`
3. `utf8.RuneCountInString("péripatéticien")` *CORRECT*
4. `unicode/utf8.RuneCountInString("péripatéticien")`

> **1:** 雙引號在哪裡？
>
> **2:** 這只能找到字串值中的位元組數
>
> **4:** 你很接近了，但套件名稱是 utf8，不是 unicode/utf8
>


## 以下敘述的解答?
```go
utf8.RuneCountInString("péripatéticien")
```

1. 16
2. 14 *CORRECT*
3. 18
4. 20

> **1:** 這是它的位元組數。RuneCountInString 計算的是 rune（碼位），不是位元組
>
> **2:** 沒錯。RuneCountInString 回傳字串值中 rune（碼位）的數量
>


## 哪個套件包含字串操作的函數？
1. string
2. unicode/utf8
3. strings *CORRECT*
4. unicode/strings


## 以下敘述的解答?
```go
strings.Repeat("*x", 3) + "*"
```

**提示:** Repeat 函數會重複給定的字串

1. `*x*x*x`
2. `x*x*x`
3. `*x3`
4. `*x*x*x*` *CORRECT*

> **1:** 你很接近了，但最後的字串連接漏掉了
>
> **2:** 仔細看看
>
> **3:** 哇！你應該真的再看一次課程。抱歉
>
> **4:** 沒錯。Repeat 函數會重複給定的字串，而連接運算子則會將字串合併
>


## 以下敘述的解答?
```go
strings.ToUpper("bye bye ") + "see you!"
```

1. `bye bye see you!`
2. `BYE BYE SEE YOU!`
3. `bye bye + see you!`
4. `BYE BYE see you!` *CORRECT*

> **1:** 你漏掉了 ToUpper?
>
> **2:** 你很接近，但仔細看看。ToUpper 只改變了字串的前半部分
>
> **3:** 完全不接近。抱歉
>
> **4:** 完美！抓得很好！ToUpper 只改變了字串的前半部分
>