import { create } from 'zustand'
import { createAccount, IAccount } from './account'
import { createFrame, IFrame } from './frame'
import { createDoc, IDoc } from './doc'

/**
 * 全局 store
 */
export const useGlobalStore = create<IAccount & IFrame & IDoc>()((...a) => ({
  ...createAccount(...a),
  ...createFrame(...a),
  ...createDoc(...a)
}))
