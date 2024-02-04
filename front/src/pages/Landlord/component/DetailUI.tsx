import { Descriptions } from 'antd'
import { LandlordInfoType } from '@/types/landlordType'
import { LandlordSexTag } from './ToolsUI'

interface LandlordDetailUIProps {
  /**
   * 业主信息
   */
  landlordDetail: LandlordInfoType | undefined
}

/**
 * 业主信息 UI 组件
 */
const LandlordDetailUI = (props: LandlordDetailUIProps) => {
  const landlordDetail = props?.landlordDetail
  return (
    <div key={landlordDetail?.landlord_id.toString()}>
      <Descriptions bordered size="small" column={12}>
        <Descriptions.Item label="业主ID" span={12} key="house_id">
          {landlordDetail?.landlord_id}
        </Descriptions.Item>
        <Descriptions.Item label="业主昵称" span={12} key="nick_name">
          {landlordDetail?.nick_name}
        </Descriptions.Item>
        <Descriptions.Item label="业主姓名" span={12} key="given_name">
          {landlordDetail?.given_name}
        </Descriptions.Item>
        <Descriptions.Item label="性别" span={12} key="sex">
          {LandlordSexTag(landlordDetail?.sex)}
        </Descriptions.Item>
        <Descriptions.Item label="手机" span={12} key="mobile">
          {landlordDetail?.mobile}
        </Descriptions.Item>
        <Descriptions.Item label="现住址" span={12} key="address">
          {landlordDetail?.address}
        </Descriptions.Item>
        <Descriptions.Item label="身份证" span={12} key="id_crad_num">
          {/* {landlordDetail?.id_card_number} */}
          *********************
        </Descriptions.Item>
        <Descriptions.Item label="邮箱" span={12} key="email">
          {landlordDetail?.email}
        </Descriptions.Item>
        <Descriptions.Item label="创建时间" span={12} key="create_time">
          {landlordDetail?.create_time}
        </Descriptions.Item>
        <Descriptions.Item label="修改时间" span={12} key="update_time">
          {landlordDetail?.update_time}
        </Descriptions.Item>
      </Descriptions>
    </div>
  )
}

export default LandlordDetailUI
