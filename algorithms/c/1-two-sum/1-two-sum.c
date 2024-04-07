#include <stdio.h>
#include <stdlib.h>

/**
 * Note: The returned array must be malloced, assume caller calls free().
 */
int* twoSum(int* nums, int numsSize, int target, int* returnSize) {
    int *returnNums = malloc(numsSize * sizeof(int));
    *returnSize = 0;

    for (int i = 0; i < numsSize; i++) {
        for (int j = i + 1; j < numsSize; j++) {
            if (nums[i] + nums[j] == target) {
                returnNums[*returnSize]=i;
                (*returnSize)++;
                returnNums[*returnSize]=j;
                (*returnSize)++;
            }
        }
    }

    return returnNums;
}

int main() {
    int nums[] = {2, 7, 11, 15};
    int target = 9;
    int returnSize;
    int *returnNums = twoSum(nums, 4, target, &returnSize);

    for (int i = 0; i < returnSize; i++) {
        printf("%d\n", returnNums[i]);
    }

   if (returnSize == 2 && returnNums[0] == 0 && returnNums[1] == 1) {
        printf("Test passed\n");
    } else {
        printf("Test failed\n");
    }

    free(returnNums);

    return 0;
}