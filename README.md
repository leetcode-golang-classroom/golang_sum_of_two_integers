# golang_sum_of_two_integers

Given two integers `a` and `b`
, return *the sum of the two integers without using the operators*`+`*and*`-`

## Examples

**Example 1:**

```
Input: a = 1, b = 2
Output: 3

```

**Example 2:**

```
Input: a = 2, b = 3
Output: 5

```

**Constraints:**

- `1000 <= a, b <= 1000`

## 解析

給定兩個整數 a, b

要求不使用 + , - 運算來實作 兩數相加的演算法

對每個位元來說

相加分為兩個部份

1. 下一個位元的進位 
2. 這個位元的和除了進位之外的和

下一個位元的進位可以透過 把對應位元做 & 然後再做 left shift 

這個位元的合 可以發現除了因為不關心進位 所以是使用 XOR

進位 = 兩個位元同時為 1 需要進位

該位置的 bit = 除了進位之外的和 = 當 2個 bit 不同時 為 1 否則為 0

![](https://i.imgur.com/AWmoQkm.png)

## 程式碼
```go
package sol

func getSum(a int, b int) int {
	sum, carry := a, b
	for carry != 0 {
		temp := (sum & carry) << 1 // carry
		sum = sum ^ carry
		carry = temp
	}
	return sum
}

```
## 困難點

1. 要思考出 carry 與 sum 與位元運算的關係式

## Solve Point

- [x]  初始化 carry := b, sum := a
- [x]  當 carry ≠ 0, 更新 carry, sum = (carry&sum) << 1, (carry^sum)
- [x]  回傳 sum