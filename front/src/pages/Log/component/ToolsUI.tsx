import { LogLevelTypes } from '@/types/logType'
import { Tag } from 'antd'
import { DefaultOptionType } from 'antd/es/select'

/**
 * 日志级别UI组件
 * @param level
 */
export const LogLevelTagUI = (level: number) => {
  let tagLebel = <Tag color="cyan">Unkown</Tag>
  LogLevelTypes.forEach((levelType) => {
    if (levelType.level === level) {
      tagLebel = <Tag color={levelType.color}>{levelType.name}</Tag>
      return
    }
  })
  return tagLebel
}

/**
 * 日志级别下拉选择 Options
 * @returns options
 */
export const LogLevelSelectOptions = (): DefaultOptionType[] => {
  let options: DefaultOptionType[] = []
  LogLevelTypes.forEach((levelType) => {
    options.push({
      label: levelType.name,
      value: levelType.level
    })
  })
  return options
}
