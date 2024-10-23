func twoSum(nums []int, target int) []int {
    s:=make([]int ,0)
    for i,_:= range nums{
        a:=nums[i+1:]
        for j,_:=range a{
            if nums[i]+nums[j+i+1]==target&&len(s)==0{
                s=append(s,i)
                s=append(s,j+i+1)
                break
            }
        }
    }
    return s
}