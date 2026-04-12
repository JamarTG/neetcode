func isValid(s string) bool {
    stack := []byte{}
    var ptr int

    for ptr < len(s) {
        isOpening := s[ptr] == '(' || s[ptr] == '{' || s[ptr] == '['

        if isOpening {
            stack = append(stack, s[ptr])
        } else {
            if len(stack) == 0 {
                return false
            }
            lastStackV := stack[len(stack)-1]
            if !(rune(s[ptr]) == '}' && rune(lastStackV) == '{' ||
                rune(s[ptr]) == ')' && rune(lastStackV) == '(' ||
                rune(s[ptr]) == ']' && rune(lastStackV) == '[') {
                return false
            }
            stack = stack[:len(stack)-1]
        }

        ptr += 1
    }

    return len(stack) == 0
}