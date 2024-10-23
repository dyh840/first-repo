func twoSum(nums []int, target int) []int {

  s:=make([]int ,0)//<u>新建切片</u>

  for i,_:= range nums{//<u>遍历nums</u>

​    a:=nums[i+1:]

​    for j,_:=range a{<u>//遍历</u>i<u>后的数字防止重复</u>

​      if nums[i]+nums[j+i+1]==target&&len(s)==0{//<u>判断条件以及只取一次</u>

​        s=append(s,i)

​        s=append(s,j+i+1)//<u>取值</u>

​        break//<u>终止</u>

​      }

​    }

  }

  return s

}