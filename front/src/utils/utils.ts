/**
 * 获取 uuid
 */
export const getUUId = () => {
  return Number(Math.random().toString().substr(2, 10) + Date.now()).toString(36)
}

export const arrayToMapBool = <T>(array: T[]): Map<T, boolean> => {
  let map = new Map<T, boolean>()
  for (const item of array) {
    map.set(item, true)
  }
  return map
}

export const arrayToString = <T>(array: T[]): string[] => {
  let res: string[] = []
  for (const item of array) {
    res.push(String(item))
  }
  return res
}

// 获取当前0点时间
export const getNowStartTimestamp = () => {
  const date = new Date()
  const timestamp = Date.UTC(date.getFullYear(), date.getMonth(), date.getDate(), 0, 0, 0, 0)
  return timestamp
}
