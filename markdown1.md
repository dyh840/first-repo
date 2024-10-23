# 丁雨杭的题解

```go
func twoSum(nums []int, target int) []int {
    s:=make([]int ,0)
    for i,_:= range nums{
        a:=nums[i+1:]//取i后数字防止重复
        for j,_:=range a{
            if nums[i]+nums[j+i+1]==target&&len(s)==0{//条件判断以及只取一次
                s=append(s,i)
                s=append(s,j+i+1)
                break
            }
        }
    }
    return s
}
```

通过嵌套循环遍历暴力求解

