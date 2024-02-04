function stringify(value: unknown): string {
  return JSON.stringify(value)
}

function parse<T>(value: string): T | null {
  try {
    return JSON.parse(value) as T
  } catch (error) {
    return null
  }
}

interface ILocalStorage {
  setValue(key: string, data: unknown): ILocalStorage
  getValue<T>(key: string, defaultValue?: T): T | null
  removeValue(key: string): ILocalStorage
}

const LocalStorage: ILocalStorage = {
  setValue(key: string, data: unknown): ILocalStorage {
    localStorage.setItem(key, stringify(data))
    return this
  },
  getValue<T>(key: string, defaultValue?: T): T | null {
    const value = localStorage.getItem(key)

    if (!value) return defaultValue || null
    const data = parse<T>(value)
    return data
  },
  removeValue(key: string): ILocalStorage {
    localStorage.removeItem(key)
    return this
  }
}

export default LocalStorage
