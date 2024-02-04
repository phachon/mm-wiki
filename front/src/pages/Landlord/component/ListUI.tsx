import { Popconfirm, Space, Table, TablePaginationConfig, Tag } from 'antd'
import { LandlordInfoType, LandlordListItemType } from '@/types/landlordType'
import { CloseSquareOutlined, FormOutlined, SelectOutlined } from '@ant-design/icons'
import { LandlordSexTag } from './ToolsUI'
import ActionButton from '@/components/Action/ActionButton'

interface LandlordListUIProps {
  /**
   * 加载状态
   */
  listLoading: boolean

  /**
   * 业主列表
   */
  landlordList: LandlordListItemType[]

  /**
   * 分页信息
   */
  pagination: TablePaginationConfig

  /**
   * 业主列表分页方法
   * @param pageConfig 分页配置
   * @param filters 过滤器
   * @param sorter 排序
   */
  onLandlordListChange: (pageConfig: TablePaginationConfig, filters: any, sorter: any) => void

  /**
   * 业主详情点击方法
   * @param landlordInfo 业主信息
   */
  onLandlordDetailClick: (landlordInfo: LandlordInfoType) => void

  /**
   * 业主修改点击方法
   * @param landlordInfo 业主信息
   */
  onLandlordEditClick: (landlordInfo: LandlordInfoType) => void

  /**
   * 业主删除确认方法
   * @param landlordInfo 业主信息
   */
  onLandlordDeleteConfim: (landlordInfo: LandlordInfoType) => void
}

const LandlordListUI = (props: LandlordListUIProps) => {
  return (
    <div className="panel-body">
      <Table
        rowKey={'landlord_id'}
        bordered={true}
        dataSource={props.landlordList}
        loading={props.listLoading}
        pagination={props.pagination}
        onChange={props.onLandlordListChange}
        footer={() => ''}
      >
        <Table.Column
          title={'业主ID'}
          dataIndex="landlord_id"
          width={100}
          key={'landlord_id'}
          align={'center'}
        />
        <Table.Column title={'业主昵称'} dataIndex="nick_name" key={'nick_name'} width={150} />
        <Table.Column
          title={'业主姓名'}
          key={'given_name'}
          width={150}
          render={(landlordInfo: LandlordInfoType) => (
            <a onClick={() => props.onLandlordDetailClick(landlordInfo)}>
              <Space>
                {landlordInfo.given_name}
                <SelectOutlined />
              </Space>
            </a>
          )}
        />
        <Table.Column
          title={'性别'}
          key={'sex'}
          width={100}
          align={'center'}
          render={(landlordInfo: LandlordInfoType) => LandlordSexTag(landlordInfo.sex)}
        />
        <Table.Column title={'现住址'} key={'address'} dataIndex="address" />
        <Table.Column
          title={'创建时间'}
          dataIndex={'create_time'}
          key={'create_time'}
          width={200}
          align={'center'}
        />
        <Table.Column
          title={'操作'}
          width={180}
          key={'action'}
          align={'center'}
          render={(landlordListItem: LandlordListItemType) => (
            <Space>
              <ActionButton
                text="修改"
                icon={<FormOutlined />}
                onClick={() => props.onLandlordEditClick(landlordListItem)}
                havePermission={landlordListItem.action?.is_edit == 1}
              />
              <Popconfirm
                title="确定要删除吗?"
                onConfirm={() => {
                  props.onLandlordDeleteConfim(landlordListItem)
                }}
                okText="确定"
                cancelText="取消"
              >
                <ActionButton
                  text="删除"
                  icon={<CloseSquareOutlined />}
                  havePermission={landlordListItem.action?.is_delete == 1}
                />
              </Popconfirm>
            </Space>
          )}
        />
      </Table>
    </div>
  )
}

export default LandlordListUI
