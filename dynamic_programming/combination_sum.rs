//TODO: implement leetcode based template (define in impl)

fn combination_sum_r(nums: &Vec<i32>, target: &i32) -> i32 {
    if target == &0 {
        return 1;
    }
    let mut out: i32 = 0;
    for num in nums {
        if num <= target {
            out += combination_sum_r(nums, &(target - num));
        }
    }
    return out;
}

fn main() {
    println!("{}", combination_sum_r(&(vec![1,2,3]), &4));
}
