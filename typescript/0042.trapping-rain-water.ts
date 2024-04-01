// Created by OrkWard at 2024/04/01 22:45
// leetgo: dev
// https://leetcode.com/problems/trapping-rain-water/

// @lc code=begin

function trap(height: number[]): number {
  let trappedRain = 0;
  // let leftHeight = height[0];
  // let rightHeight = Math.max(...height.slice(1));
  const leftHeightArr: number[] = [height[0]];
  height.slice(1).forEach((h, i) => {
    leftHeightArr.push(Math.max(h, leftHeightArr[i]))
  })
  const rightHeightArr: number[] = [height[height.length - 1]]
  height.slice(0, height.length - 1).reverse().forEach((h, i) => {
    rightHeightArr.push(Math.max(h, rightHeightArr[i]))
  })
  rightHeightArr.reverse()

  for (let i = 1; i < height.length - 1; i += 1) {
    const curHeight = height[i];
    // leftHeight = Math.max(leftHeight, height[i - 1]);
    // rightHeight = rightHeight === curHeight ? Math.max(...height.slice(i + 1)) : rightHeight;

    // if (leftHeight > curHeight && rightHeight > curHeight) {
    //   trappedRain += Math.min(leftHeight, rightHeight) - curHeight;
    // }
    if (leftHeightArr[i] > curHeight && rightHeightArr[i] > curHeight) {
      trappedRain += Math.min(leftHeightArr[i], rightHeightArr[i]) - curHeight;
    }
  }

  return trappedRain;
};

// @lc code=end
