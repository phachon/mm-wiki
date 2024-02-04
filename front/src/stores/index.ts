import { create } from 'zustand'
import { createAccount, IAccount } from './account'
import { createFrame, IFrame } from './frame'
import { devtools } from 'zustand/middleware'

/**
 * 全局 store
 */
export const useGlobalStore = create<IAccount & IFrame>()((...a) => ({
  ...createAccount(...a),
  ...createFrame(...a)
}))
