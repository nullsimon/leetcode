#include <stdio.h>
#include <stdbool.h>
#include <string.h>

bool isPalindrome(int x) {
    if (x < 0) {
        return false;
    }
    if (x == 0) {
        return true;
    }

    char str[32];
    sprintf(str, "%d", x);
    int length = strlen(str);
    
    for (int i = 0; i < length/2 + 1; i++) {
        if (str[i] != str[length-i-1]) {
            return false;
        }
    }

    return true;
}

bool isPalindromeV2(int x) {
    if (x < 0) {
        return false;
    }
    if (x == 0) {
        return true;
    }

    int original = x;
    int reversed = 0;

    while (x > 0) {
        int number = x % 10;
        x = (x - number) / 10;
        if (reversed > original/10) {
            return false;
        }

        reversed = reversed * 10 + number;
    }

    if (original == reversed) {
        return true;
    }

    return false;
}

int main() {
    if (isPalindrome(121)) {
        printf("V1 Test passed\n");
    } else {
        printf("V1 Test failed\n");
    }

    if (isPalindromeV2(121)) {
        printf("V2 Test passed\n");
    } else {
        printf("V2 Test failed\n");
    }

    return 0;
}