# Path Coverage by Exhaustion

for pdq

explain

pdq interface
pdq [-t think][-z sleep][-s service][-vd] from to by

where 

- t is positive number, upper bound unknown
- z sleep is synonymous with t
- s is service time, positibe number, upper bound unknown
- v is a verbose flag
- d is a debug flag
- h is a usage flag, causing an exit
- from is a positive number, defaulting to 1 
- to is a positive number, invalid if less than from
- by is a positive number, smaller than to-from

The combination of the last three has a limit


## Go range of an array
    nums := []int{2, 3, 4}
    sum := 0
    for _, num := range nums {
        sum += num
    }
