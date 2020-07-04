func canFinish(numCourses int, prerequisites [][]int) bool {
    var courses map[int]map[int]bool = make(map[int]map[int]bool, numCourses)
    var coursesTaken map[int]bool = make(map[int]bool, numCourses)
    var queue []int
    
    
    for i := 0; i < numCourses; i++ {
        queue = append(queue, i)
        courses[i] = make(map[int]bool, 0)
        coursesTaken[i] = false
    }
    
    for i := 0; i < len(prerequisites); i++ {
        course, prereq := prerequisites[i][0], prerequisites[i][1]
        prereqs := courses[course]
        prereqs[prereq] = true
        courses[course] = prereqs
    }
    
    var iter int
    var maxIter int = numCourses * len(prerequisites)
    
    for len(queue) > 0 {
        if (iter > maxIter) {
            break
        }
        
        course := queue[0]
        queue = queue[1:]
        
        prereqs := courses[course]
        if len(prereqs) == 0 {
            coursesTaken[course] = true
            continue
        }
        
        allPreReqsTaken := true
        
        for prereq, _ := range prereqs {
            if !coursesTaken[prereq] {
                queue = append(queue, course)
                iter += 1
                allPreReqsTaken = false
                break
            }
        }
        
        if allPreReqsTaken {
            coursesTaken[course] = true
        }
    }
    
    for _, taken := range coursesTaken {
        if !taken {
            return false
        }
    }
    
    return true
}

/*
2
[[1,0]]
2
[[1,0],[0,1]]
1
[[0,0]]
5
[[1,0],[2,1],[3,2],[3,4],[4,0]]
10
[[5,8],[3,5],[1,9],[4,5],[0,2],[7,8],[4,9]]
7
[[1,0],[0,3],[0,2],[3,2],[2,5],[4,5],[5,6],[2,4]]
*/