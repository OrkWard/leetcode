// Created by OrkWard at 2024/04/02 10:53
// leetgo: dev
// https://leetcode.com/problems/multiply-strings/

// @lc code=begin
function charAdd(char1: string, char2: string, carriage: boolean): [string, boolean] {
  const result = char1.charCodeAt(0) + char2.charCodeAt(0) - '0'.charCodeAt(0) * 2 + Number(carriage);
  return [String(result % 10), result >= 10]
}

function charMul(char1: string, char2: string): string {
  return String((char1.charCodeAt(0) - '0'.charCodeAt(0)) * (char2.charCodeAt(0) - '0'.charCodeAt(0)))
}

function add(num1: string, num2: string): string {
  let result = ''
  let carriage = false
  for (let i = 0; i < num1.length || i < num2.length; i += 1) {
    const charAddResult = charAdd(num1[num1.length - 1 - i] ?? '0', num2[num2.length - 1 - i] ?? '0', carriage)
    result = charAddResult[0] + result
    carriage = charAddResult[1]
  }
  return carriage ? '1' + result : result
}

function multiply(num1: string, num2: string): string {
  if (num1 === '0' || num2 === '0') {
    return '0'
  }

  let productArr: string[] = [];

  for (let i = 0; i < num1.length; i += 1) {
    for (let j = 0; j < num2.length; j += 1) {
      productArr.push(charMul(num1[num1.length - 1 - i], num2[num2.length - 1 - j]) + '0'.repeat(i + j))
    }
  }

  return productArr.reduce((prev, curr) => add(prev, curr), '0')
};

// @lc code=end

console.log(multiply('123456789', '987654321'))
