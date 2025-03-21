func maxDepthOfParentheses(s string) int {
    leftCount := 0
    rightCount := 0
    
    for _, char := range s {
        if char == '(' {
            leftCount++
        } else {
            rightCount++
            
            // If the current character is an opening parenthesis, we need to count its depth.
            if rightCount > 0 {
                rightCount--
            } else {
                // It's a closing parenthesis, so update the maximum depth and reset the right count for the next iteration.
                maxDepth := leftCount
                leftCount = rightCount
                rightCount = 0
            }
        }
    }
    
    return maxDepth
}
