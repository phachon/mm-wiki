import { Descriptions } from 'antd'
import { HirerInfoType } from '@/types/hirerType'
import { HirerSexTag } from './ToolsUI'

interface HirerDetailUIProps {
  /**
   * 租客信息
   */
  hirerDetail: HirerInfoType | undefined
}

/**
 * 租客信息 UI 组件
 */
const HirerDetailUI = (props: HirerDetailUIProps) => {
  const hirerDetail = props?.hirerDetail
  return (
    <div key={hirerDetail?.hirer_id.toString()}>
      <Descriptions bordered size="small" column={12}>
        <Descriptions.Item label="租客ID" span={12} key="house_id">
          {hirerDetail?.hirer_id}
        </Descriptions.Item>
        <Descriptions.Item label="租客姓名" span={12} key="given_name">
          {hirerDetail?.name}
        </Descriptions.Item>
        <Descriptions.Item label="性别" span={12} key="sex">
          {HirerSexTag(hirerDetail?.sex)}
        </Descriptions.Item>
        <Descriptions.Item label="手机" span={12} key="mobile">
          {hirerDetail?.mobile}
        </Descriptions.Item>
        <Descriptions.Item label="身份证" span={12} key="id_crad_num">
          {/* {hirerDetail?.id_card_number} */}
          *********************
        </Descriptions.Item>
        <Descriptions.Item label="邮箱" span={12} key="email">
          {hirerDetail?.email}
        </Descriptions.Item>
        <Descriptions.Item label="创建时间" span={12} key="create_time">
          {hirerDetail?.create_time}
        </Descriptions.Item>
        <Descriptions.Item label="修改时间" span={12} key="update_time">
          {hirerDetail?.update_time}
        </Descriptions.Item>
      </Descriptions>
    </div>
  )
}

export default HirerDetailUI
