import { create } from 'zustand'

export const useTestStore = create((set) => ({
  accountInfo: null,
  setAccountInfo: () => set(() => ({ count: 31 }))
}))
