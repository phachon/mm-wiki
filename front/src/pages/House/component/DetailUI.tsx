import { Descriptions, Tag } from 'antd'
import { HouseInfoType } from '@/types/houseType'
import {
  HouseAllowLeaseTag,
  HouseAllowSplitLeaseTag,
  HouseDecorationTypeText,
  HouseLeaseStatusTag,
  HouseSizeTypeText
} from './ToolsUI'

interface HouseDetailUIProps {
  /**
   * 房产详情信息
   */
  detailHouseInfo?: HouseInfoType
}

/**
 * 房产详情 UI 组件
 */
const HouseDetailUI = (props: HouseDetailUIProps) => {
  const houseInfo = props?.detailHouseInfo
  return (
    <div key={houseInfo?.house_id.toString()}>
      <Descriptions bordered size="small" column={12}>
        <Descriptions.Item label="房产ID" span={12} key="house_id">
          {houseInfo?.house_id}
        </Descriptions.Item>
        <Descriptions.Item label="房产编号" span={12} key="idx">
          {houseInfo?.idx}
        </Descriptions.Item>
        <Descriptions.Item label="所在区域" span={12} key="region">
          {houseInfo?.region}
        </Descriptions.Item>
        <Descriptions.Item label="详细地址" span={12} key="address">
          {houseInfo?.address}
        </Descriptions.Item>
        <Descriptions.Item label="门牌号" span={12} key="house_number">
          {houseInfo?.house_number}
        </Descriptions.Item>
        <Descriptions.Item label="装修类型" span={12} key="decoration_type">
          {HouseDecorationTypeText(houseInfo?.decoration_type)}
        </Descriptions.Item>
        <Descriptions.Item label="房产户型" span={12} key="size_type">
          {HouseSizeTypeText(houseInfo?.size_type)}
        </Descriptions.Item>
        <Descriptions.Item label="房产面积" span={12} key="area">
          {houseInfo?.area}
        </Descriptions.Item>
        <Descriptions.Item label="建成日期" span={12} key="build_time">
          {houseInfo?.build_time}
        </Descriptions.Item>
        <Descriptions.Item label="出租状态" span={12} key="lease_status">
          {HouseLeaseStatusTag(houseInfo?.lease_status)}
        </Descriptions.Item>
        <Descriptions.Item label="允许整租" span={12} key="allow_lease">
          {HouseAllowLeaseTag(houseInfo?.allow_lease)}
        </Descriptions.Item>
        <Descriptions.Item label="允许合租" span={12} key="allow_split_lease">
          {HouseAllowSplitLeaseTag(houseInfo?.allow_split_lease)}
        </Descriptions.Item>
        <Descriptions.Item label="月租金" span={12} key="mouth_rent">
          {houseInfo?.mouth_rent}¥/月
        </Descriptions.Item>
        <Descriptions.Item label="创建时间" span={12} key="create_time">
          {houseInfo?.create_time}
        </Descriptions.Item>
        <Descriptions.Item label="修改时间" span={12} key="update_time">
          {houseInfo?.update_time}
        </Descriptions.Item>
      </Descriptions>
    </div>
  )
}

export default HouseDetailUI
