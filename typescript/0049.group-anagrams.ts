// Created by OrkWard at 2024/04/02 17:43
// leetgo: dev
// https://leetcode.com/problems/group-anagrams/

// @lc code=begin
const A_CODE = 'a'.charCodeAt(0)
function vertorize(str: string): string {
  const alphabetSeq: number[] = new Array(26).fill(0)
  for (const char of str) {
    const charInAlphabet = char.charCodeAt(0) - A_CODE
    alphabetSeq[charInAlphabet] += 1
  }
  return alphabetSeq.map((v) => String.fromCharCode(v)).join('')
}

function groupAnagrams(strs: string[]): string[][] {
  const strVectors = strs.map(vertorize)
  const vectorMap: Map<string, string[]> = new Map()
  strVectors.forEach((vec, i) => {
    if (vectorMap.has(vec)) {
      vectorMap.get(vec)!.push(strs[i])
    } else {
      vectorMap.set(vec, [strs[i]])
    }
  })

  return Array.from(vectorMap.values())
};

// @lc code=end
