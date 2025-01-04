import DynamicIcon from '@/components/DynamicIcon/DynamicIcon'
import { SettingConfig } from '@/config/setting'
import { DocTypes } from '@/types/docType'
import { CheckboxOptionType } from 'antd'
import { NavigateFunction } from 'react-router-dom'

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

/**
 * 文档 url 处理函数
 */
export const DocUrlProcessor = (url: string, type: string) => {
  if (type == 'image' || type == 'video' || type == 'audio' || type == 'file') {
    return SettingConfig.fileDomain + '/' + url
  }

  return url
}

// navigateDocView 导航到 doc 页面
export const navigateDocView = (navigate: NavigateFunction, docId?: number) => {
  if (!docId) {
    return
  }
  navigate(`/doc/` + docId)
}
