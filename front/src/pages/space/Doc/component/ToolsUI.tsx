import DynamicIcon from '@/components/DynamicIcon/DynamicIcon'
import { DocTypes } from '@/types/docType'
import { CheckboxOptionType } from 'antd'

/**
 * 文档类型 Radio Options 组件
 */
export const DocTypeRadioOptions = () => {
  let options: CheckboxOptionType[] = []
  DocTypes.forEach((docType) => {
    options.push({
      label: (
        <span>
          <DynamicIcon name={docType.icon} /> {docType.name}
        </span>
      ),
      value: docType.type
    })
  })
  return options
}
