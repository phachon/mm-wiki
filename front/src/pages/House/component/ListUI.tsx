import { Popconfirm, Space, Table, TablePaginationConfig, Tag } from 'antd'
import { HouseInfoType, HouseListItemType } from '@/types/houseType'
import { CloseSquareOutlined, FormOutlined, SelectOutlined } from '@ant-design/icons'
import {
  HouseAllowLeaseTag,
  HouseAllowSplitLeaseTag,
  HouseLeaseStatusTag,
  HouseShowWholeAddress
} from './ToolsUI'
import ActionButton from '@/components/Action/ActionButton'

interface HouseListUIProps {
  /**
   * 加载状态
   */
  listLoading: boolean

  /**
   * 房产列表
   */
  houseList: HouseListItemType[]

  /**
   * 分页信息
   */
  pagination: TablePaginationConfig

  /**
   * 房产列表分页方法
   * @param pageConfig 分页配置
   * @param filters 过滤器
   * @param sorter 排序
   */
  onHouseListChange: (pageConfig: TablePaginationConfig, filters: any, sorter: any) => void

  /**
   * 房产详情点击方法
   * @param houseInfo 房产信息
   */
  onHouseDetailClick: (houseInfo: HouseInfoType) => void

  /**
   * 房产修改点击方法
   * @param houseInfo 房产信息
   */
  onHouseEditClick: (houseInfo: HouseInfoType) => void

  /**
   * 房产删除确认方法
   * @param houseInfo 房产信息
   */
  onHouseDeleteConfim: (houseInfo: HouseInfoType) => void

  /**
   * 房产业主详情点击方法
   * @param houseInfo 房产信息
   */
  onHouseLandlordDetailClick: (houseInfo: HouseInfoType) => void

  /**
   * 房产托管点击方法
   * @param houseInfo 房产信息
   */
  onHouseTrustClick: (houseInfo: HouseInfoType) => void
}

const HouseListUI = (props: HouseListUIProps) => {
  return (
    <div className="panel-body">
      <Table
        rowKey={'house_id'}
        bordered={true}
        dataSource={props.houseList}
        loading={props.listLoading}
        pagination={props.pagination}
        onChange={props.onHouseListChange}
        footer={() => ''}
      >
        <Table.Column
          title={'房产ID'}
          dataIndex="house_id"
          width={80}
          key={'house_id'}
          align={'center'}
        />
        <Table.Column
          title={'详细地址'}
          key={'address'}
          render={(houseInfo: HouseInfoType) => (
            <a onClick={() => props.onHouseDetailClick(houseInfo)}>
              <Space>
                {HouseShowWholeAddress(houseInfo)}
                <SelectOutlined />
              </Space>
            </a>
          )}
        />
        <Table.Column
          title={'面积'}
          key={'area'}
          width={100}
          align={'center'}
          render={(houseInfo: HouseInfoType) => <span>{houseInfo.area} 平米</span>}
        />
        <Table.Column
          title={'月租金'}
          dataIndex={'mouth_rent'}
          key={'mouth_rent'}
          width={120}
          align={'center'}
          render={(mouthRent: number) => <span>{mouthRent}¥/月</span>}
        />
        <Table.Column
          title={'业主'}
          key={'landlord'}
          width={150}
          align={'center'}
          render={(houseInfo: HouseInfoType) => (
            <a onClick={() => props.onHouseLandlordDetailClick(houseInfo)}>
              <span className="button-text">
                业主详情 <SelectOutlined />
              </span>
            </a>
          )}
        />
        <Table.Column
          title={'允许整租'}
          dataIndex="allow_lease"
          width={100}
          key={'allow_lease'}
          align={'center'}
          render={HouseAllowLeaseTag}
        />
        <Table.Column
          title={'允许合租'}
          dataIndex="allow_split_lease"
          width={100}
          key={'allow_split_lease'}
          align={'center'}
          render={HouseAllowSplitLeaseTag}
        />
        <Table.Column
          title={'出租状态'}
          dataIndex="status"
          width={100}
          key={'lease_status'}
          align={'center'}
          render={HouseLeaseStatusTag}
        />
        <Table.Column
          title={'操作'}
          width={180}
          key={'action'}
          align={'center'}
          render={(houseListItem: HouseListItemType) => (
            <Space>
              <ActionButton
                text="修改"
                icon={<FormOutlined />}
                onClick={() => props.onHouseEditClick(houseListItem)}
                havePermission={houseListItem.action?.is_edit == 1}
              />
              <Popconfirm
                title="确定要删除吗?"
                onConfirm={() => {
                  props.onHouseDeleteConfim(houseListItem)
                }}
                okText="确定"
                cancelText="取消"
              >
                <ActionButton
                  text="删除"
                  icon={<CloseSquareOutlined />}
                  havePermission={houseListItem.action?.is_delete == 1}
                />
              </Popconfirm>
            </Space>
          )}
        />
      </Table>
    </div>
  )
}

export default HouseListUI
