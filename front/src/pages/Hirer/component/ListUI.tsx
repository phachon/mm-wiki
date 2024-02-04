import { Button, Popconfirm, Space, Table, TablePaginationConfig, Tag } from 'antd'
import { HirerInfoType, HirerListItemType } from '@/types/hirerType'
import { CloseSquareOutlined, FormOutlined, SelectOutlined } from '@ant-design/icons'
import { HirerSexTag } from './ToolsUI'
import ActionButton from '@/components/Action/ActionButton'

interface HirerListUIProps {
  /**
   * 加载状态
   */
  listLoading: boolean

  /**
   * 租客列表
   */
  hirerList: HirerListItemType[]

  /**
   * 分页信息
   */
  pagination: TablePaginationConfig

  /**
   * 租客列表分页方法
   * @param pageConfig 分页配置
   * @param filters 过滤器
   * @param sorter 排序
   */
  onHirerListChange: (pageConfig: TablePaginationConfig, filters: any, sorter: any) => void

  /**
   * 租客详情点击方法
   * @param hirerInfo 租客信息
   */
  onHirerDetailClick: (hirerInfo: HirerInfoType) => void

  /**
   * 租客修改点击方法
   * @param hirerInfo 租客信息
   */
  onHirerEditClick: (hirerInfo: HirerInfoType) => void

  /**
   * 租客删除确认方法
   * @param hirerInfo 租客信息
   */
  onHirerDeleteConfim: (hirerInfo: HirerInfoType) => void
}

const HirerListUI = (props: HirerListUIProps) => {
  return (
    <div className="panel-body">
      <Table
        rowKey={'hirer_id'}
        bordered={true}
        dataSource={props.hirerList}
        loading={props.listLoading}
        pagination={props.pagination}
        onChange={props.onHirerListChange}
        footer={() => ''}
      >
        <Table.Column
          title={'租客ID'}
          dataIndex="hirer_id"
          width={100}
          key={'hirer_id'}
          align={'center'}
        />
        <Table.Column
          title={'租客姓名'}
          key={'name'}
          width={200}
          render={(hirerInfo: HirerInfoType) => (
            <a onClick={() => props.onHirerDetailClick(hirerInfo)}>
              <Space>
                {hirerInfo.name}
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
          render={(hirerInfo: HirerInfoType) => HirerSexTag(hirerInfo.sex)}
        />
        <Table.Column title={'电话'} key={'mobile'} dataIndex="mobile" />
        <Table.Column
          title={'紧急联系人'}
          key={'emergency_contact'}
          dataIndex="emergency_contact"
        />
        <Table.Column title={'联系人电话'} key={'emergency_mobile'} dataIndex="emergency_mobile" />
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
          render={(hirerListItem: HirerListItemType) => (
            <Space>
              <ActionButton
                text="修改"
                icon={<FormOutlined />}
                onClick={() => props.onHirerEditClick(hirerListItem)}
                havePermission={hirerListItem.action?.is_edit == 1}
              />
              <Popconfirm
                title="确定要删除吗?"
                onConfirm={() => {
                  props.onHirerDeleteConfim(hirerListItem)
                }}
                okText="确定"
                cancelText="取消"
              >
                <ActionButton
                  text="删除"
                  icon={<CloseSquareOutlined />}
                  havePermission={hirerListItem.action?.is_delete == 1}
                />
              </Popconfirm>
            </Space>
          )}
        />
      </Table>
    </div>
  )
}

export default HirerListUI
