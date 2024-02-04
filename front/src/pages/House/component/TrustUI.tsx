import { HouseInfoType, HouseTrustInfoType } from '@/types/houseType'
import {
  Button,
  Card,
  Col,
  DatePicker,
  Form,
  Input,
  InputNumber,
  Popconfirm,
  Row,
  Select,
  Space,
  Table,
  TablePaginationConfig
} from 'antd'
import { CloseSquareOutlined } from '@ant-design/icons'
import locale from 'antd/es/date-picker/locale/zh_CN'
import moment from 'moment'
import dayjs from 'dayjs'
import { useEffect } from 'react'

const dateFormat = 'YYYY-MM-DD'
const { Option } = Select
// 设置默认的起始日期
const disabledDate = (current: any) => {
  return current <= moment().startOf('day')
}

interface HouseTrustUIProps {
  /**
   * 当前房产托管信息
   */
  currentHouseTrustInfo?: HouseTrustInfoType

  /**
   * 房产托管历史列表
   */
  houseTrustList?: HouseTrustInfoType[]

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
  onHouseTrustListChange: (pageConfig: TablePaginationConfig, filters: any, sorter: any) => void

  /**
   * 房产托管保存方法
   * @param values 保存结果
   */
  onHouseTrustSave: (values: HouseTrustInfoType) => void

  /**
   * 房产托管删除方法
   */
  onHouseTrustDeleteConfim?: (values: HouseTrustInfoType) => void
}

/**
 * 房产托管 UI 组件
 * @param props 房产托管 UI 组件数据
 * @returns 房产托管 UI 组件
 */
const HouseTrustUI = (props: HouseTrustUIProps) => {
  const [form] = Form.useForm()
  const isEdit = props.currentHouseTrustInfo ? true : false

  useEffect(() => {
    if (props.currentHouseTrustInfo) {
      form.setFieldsValue({
        ...props.currentHouseTrustInfo,
        start_time: dayjs(props.currentHouseTrustInfo?.start_time, dateFormat),
        end_time: dayjs(props.currentHouseTrustInfo?.end_time, dateFormat)
      })
      return
    }
    form.setFieldsValue({})
  }, [props.currentHouseTrustInfo])

  const suffixSelector = (
    <Form.Item name="rent_type" noStyle>
      <Select style={{ width: 90 }} defaultValue={'1'}>
        <Option value="1">¥/月</Option>
        <Option value="2">¥/季度</Option>
        <Option value="3">¥/半年</Option>
        <Option value="4">¥/一年</Option>
      </Select>
    </Form.Item>
  )

  /**
   * 房产托管保存方法
   * @param values 保存的信息
   */
  const houseTrustSave = (values: any) => {
    // 日期格式转换
    const startTime = values['start_time'].format('YYYY-MM-DD')
    const endTime = values['end_time'].format('YYYY-MM-DD')
    props.onHouseTrustSave({
      ...values,
      start_time: startTime,
      end_time: endTime
    })
  }

  return (
    <div style={{ marginTop: 18 }}>
      <Card title="托管中" bodyStyle={{ paddingBottom: 12 }}>
        <Form name="basic" onFinish={houseTrustSave} form={form} layout="vertical">
          <Row gutter={24}>
            <Col span={7} key={'start_time'}>
              <Form.Item label="开始日期" name="start_time" rules={[{ required: true }]}>
                <DatePicker
                  picker={'date'}
                  locale={locale}
                  placeholder="请选择托管开始日期"
                  style={{ width: '100%' }}
                  format={dateFormat}
                  disabledDate={disabledDate}
                />
              </Form.Item>
            </Col>
            <Col span={7} key={'end_time'}>
              <Form.Item label="结束日期" name="end_time" rules={[{ required: true }]}>
                <DatePicker
                  picker={'date'}
                  locale={locale}
                  placeholder="请选择托管结束日期"
                  style={{ width: '100%' }}
                  format={dateFormat}
                  disabledDate={disabledDate}
                />
              </Form.Item>
            </Col>
            <Col span={8} key={'rent'}>
              <Form.Item
                name="rent"
                label="租金"
                rules={[{ required: true, message: '请输入租金' }]}
              >
                <InputNumber
                  addonAfter={suffixSelector}
                  style={{ width: '100%' }}
                  placeholder="请输入租金"
                />
              </Form.Item>
            </Col>
            <Col span={2} key={'submit'}>
              <Form.Item label=" ">
                <Button type="primary" htmlType="submit">
                  保存
                </Button>
              </Form.Item>
            </Col>
          </Row>
        </Form>
      </Card>
      <div style={{ marginTop: 18 }}>
        <Table
          title={() => (
            <span style={{ fontWeight: 600, fontSize: 16, marginLeft: 8 }}>托管历史</span>
          )}
          rowKey={'house_id'}
          bordered={true}
          dataSource={props.houseTrustList}
          pagination={props.pagination}
          onChange={props.onHouseTrustListChange}
          footer={() => ''}
        >
          <Table.Column
            title={'ID'}
            dataIndex="house_id"
            width={80}
            key={'house_id'}
            align={'center'}
          />
          <Table.Column
            title={'托管时间'}
            dataIndex={'build_time'}
            key={'build_time'}
            align={'center'}
          />
          <Table.Column
            title={'托管租金'}
            dataIndex={'build_time'}
            key={'build_time'}
            align={'center'}
          />
          <Table.Column
            title={'创建时间'}
            dataIndex={'create_time'}
            key={'create_time'}
            align={'center'}
          />
          <Table.Column title={'操作人'} dataIndex="status" width={100} align={'center'} />
          <Table.Column
            title={'操作'}
            width={180}
            key={'action'}
            align={'center'}
            render={(houseTrustInfo: HouseTrustInfoType) => (
              <>
                <Space>
                  <Popconfirm
                    title="确定要删除吗?"
                    onConfirm={() => {
                      props.onHouseTrustDeleteConfim
                        ? props.onHouseTrustDeleteConfim(houseTrustInfo)
                        : null
                    }}
                    okText="确定"
                    cancelText="取消"
                  >
                    <a>
                      <CloseSquareOutlined />
                      <span className="button-text">删除</span>
                    </a>
                  </Popconfirm>
                </Space>
              </>
            )}
          />
        </Table>
      </div>
    </div>
  )
}

export default HouseTrustUI
