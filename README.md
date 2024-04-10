# 对Go操作mysql进行了非常简单的封装

## Go + sqlx + transaction + databasepool+batchoperation + crud

```
type User struct {
    UserId   int    `db:"user_id"`
    Username string `db:"username"`
    Sex      string `db:"sex"`
    Email    string `db:"email"`
}
```

### 使用方法，true为开启事务,false不开启

#### 插入：*Controller.Insert(true, username, sex,email)*

#### 更新：*Controller.Update(true, username, user_id)*

#### 删除：*Controller.Delete(true, user_id)*

#### 批量删除：*Controller.BatchDelete(true, user_id0,user_id1, user_id2, user_id3)*

```go
Controller.BatchDelete(true, 10, 11, 12, 0, 1, 2)
```

```go
slice := []int{10, 11, 12, 0, 1, 2}
Controller.BatchSelect(true, slice...)
```

#### 查找：*result := Controller.Select(True,user_id)*

#### 批量查找：Controller.BatchSelect(true, user_id0,user_id1, user_id2, user_id3)

```go
result := Controller.BatchSelect(true, 17,18,19,20)
fmt.Printf("找到%d条\n", len(result))
for _, val := range result {
    fmt.Println(val)
}
```

```go
slice := []int{17, 18, 19, 20}
result := Controller.BatchSelect(true, slice...)
fmt.Printf("找到%d条\n", len(result))
for _, val := range result {
    fmt.Println(val)
}
```