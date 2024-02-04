import {
  Button,
  Col,
  DatePicker,
  Divider,
  Form,
  Input,
  InputNumber,
  Row,
  Select,
  Space,
  Switch
} from 'antd'
import { HouseInfoType } from '@/types/houseType'
import 'dayjs/locale/zh-cn'
import locale from 'antd/es/date-picker/locale/zh_CN'
import { LandlordInfoType } from '@/types/landlordType'
import { DefaultOptionType } from 'antd/lib/select'
import {
  HouseDecorationTypeSelectOptions,
  HouseRegionSelectOptions,
  HouseSizeTypeSelectOptions
} from './ToolsUI'
import dayjs from 'dayjs'
import { useEffect } from 'react'
import { FormOutlined, PayCircleOutlined, SettingOutlined } from '@ant-design/icons'

interface HouseFormUIProps {
  /**
   * 房产信息保存方法
   * @param values 保存结果
   */
  onHouseInfoSave: (values: HouseInfoType) => void

  /**
   * 新增业主点击方法
   * @param e 点击事件对象
   */
  onAddLandlordClick?: (e: any) => void

  /**
   * 选择业主列表
   */
  selectLandlords?: LandlordInfoType[]

  /**
   * 默认选中的业主
   */
  defaultSelectLandlordValue?: string

  /**
   * 房产信息（修改时用）
   */
  houseInfo?: HouseInfoType
}

const dateFormat = 'YYYY-MM'

/**
 * 获取业主 Option 数据
 * @param landlordList 业主列表
 * @returns
 */
const getLandlordOptions = (landlordList?: LandlordInfoType[]): DefaultOptionType[] => {
  let options: DefaultOptionType[] = []
  if (!landlordList) {
    return options
  }
  landlordList.forEach((landlordInfo) => {
    options.push({
      label: landlordInfo.nick_name + '(' + landlordInfo.given_name + ')', // 昵称（姓名）
      value: String(landlordInfo.landlord_id)
    })
  })
  return options
}

/**
 * 房产表单 UI 组件
 * @param props 房产表单 UI 组件数据
 * @returns 房产表单 UI 组件
 */
const HouseFormUI = (props: HouseFormUIProps) => {
  const [form] = Form.useForm()
  const isEdit = props.houseInfo ? true : false
  const sizeTypeOptions = HouseSizeTypeSelectOptions()
  const decorationTypeOptions = HouseDecorationTypeSelectOptions()
  const regionOptions = HouseRegionSelectOptions()

  useEffect(() => {
    if (props.houseInfo) {
      form.setFieldsValue({
        ...props.houseInfo,
        build_time: dayjs(props.houseInfo?.build_time, dateFormat),
        landlord_id: String(props.houseInfo.landlord_id)
      })
      return
    }
    form.setFieldsValue({})
  }, [props.houseInfo])

  useEffect(() => {
    if (props.defaultSelectLandlordValue) {
      form.setFieldValue('landlord_id', props.defaultSelectLandlordValue)
    }
  }, [props.defaultSelectLandlordValue])

  /**
   * 保存房产信息
   * @param houseInfo 房产信息
   */
  const onHouseSave = (houseInfo: any) => {
    // 日期格式转换
    houseInfo['build_time'] = houseInfo['build_time'].format('YYYY-MM')
    props.onHouseInfoSave(houseInfo)
  }

  return (
    <div className="panel-body">
      <Form name="basic" onFinish={onHouseSave} form={form} layout="vertical">
        <Divider orientation="left">
          <FormOutlined /> 基础信息
        </Divider>
        <Row gutter={24}>
          <Col span={6} key={'house_id'} offset={1}>
            <Form.Item label="房产id" name="house_id">
              <Input disabled placeholder="房产ID保存后自动生成" />
            </Form.Item>
          </Col>
          <Col span={6} key={'idx'} offset={1}>
            <Form.Item label="房产编号" name="idx">
              <Input disabled placeholder="房产编号保存后自动生成" />
            </Form.Item>
          </Col>
          <Col span={7} key={'region'} offset={1}>
            <Form.Item
              label="所在区域"
              name="region"
              rules={[{ required: true, message: '请选择区域!' }]}
            >
              <Select placeholder="选择房产所在区域" allowClear options={regionOptions} />
            </Form.Item>
          </Col>
          <Col span={6} key={'address'} offset={1}>
            <Form.Item label="小区地址" name="address" rules={[{ required: true }]}>
              <Input placeholder="请输入详细地址" />
            </Form.Item>
          </Col>
          <Col span={6} key={'house_number'} offset={1}>
            <Form.Item
              label="门牌号"
              name="house_number"
              rules={[{ required: true, message: '请输入门牌号!' }]}
            >
              <Input placeholder="请输入门牌号（5单元201）" />
            </Form.Item>
          </Col>
          <Col span={7} key={'decoration_type'} offset={1}>
            <Form.Item
              label="装修类型"
              name="decoration_type"
              rules={[{ required: true, message: '请选择装修类型!' }]}
            >
              <Select placeholder="请选择装修类型" allowClear options={decorationTypeOptions} />
            </Form.Item>
          </Col>
          <Col span={6} key={'size_type'} offset={1}>
            <Form.Item label="户型" name="size_type" rules={[{ required: true }]}>
              <Select placeholder="请选择户型" allowClear options={sizeTypeOptions} />
            </Form.Item>
          </Col>
          <Col span={6} key={'area'} offset={1}>
            <Form.Item
              label="面积"
              name="area"
              rules={[{ required: true, message: '请输入面积大小!' }]}
            >
              <InputNumber
                placeholder="请输入面积大小"
                style={{ width: '100%' }}
                addonAfter={'平米'}
              />
            </Form.Item>
          </Col>
          <Col span={7} key={'build_time'} offset={1}>
            <Form.Item label="建成日期" name="build_time">
              <DatePicker
                picker={'month'}
                locale={locale}
                placeholder="请选择建成日期"
                style={{ width: '100%' }}
                format={dateFormat}
              />
            </Form.Item>
          </Col>
          <Col span={6} key={'landlord_id'} offset={1}>
            <Form.Item
              label="选择业主"
              name="landlord_id"
              rules={[{ required: true, message: '请选择业主!' }]}
            >
              <Space.Compact style={{ width: '100%' }}>
                <Form.Item name="landlord_id" noStyle>
                  <Select
                    placeholder="请选择业主"
                    allowClear
                    showSearch
                    options={getLandlordOptions(props.selectLandlords)}
                  />
                </Form.Item>
                {!isEdit && <Button onClick={props.onAddLandlordClick}>新增业主</Button>}
              </Space.Compact>
            </Form.Item>
          </Col>
          <Col span={6} key={'allow_lease'} offset={1}>
            <Form.Item
              label="允许整租"
              name="allow_lease"
              valuePropName="checked"
              initialValue={0}
              getValueProps={(value) => ({ checked: value === 0 ? true : false })}
              getValueFromEvent={(value) => {
                return value ? 0 : -1
              }}
              rules={[{ required: true, message: '请选择是否允许整租!' }]}
            >
              <Switch checkedChildren="是" unCheckedChildren="否" defaultChecked />
            </Form.Item>
          </Col>
          <Col span={7} key={'allow_split_lease'} offset={1}>
            <Form.Item
              label="允许合租"
              name="allow_split_lease"
              valuePropName="checked"
              initialValue={0}
              getValueProps={(value) => ({ checked: value === 0 ? true : false })}
              getValueFromEvent={(value) => {
                return value ? 0 : -1
              }}
              rules={[{ required: true, message: '请选择是否允许合租!' }]}
            >
              <Switch checkedChildren="是" unCheckedChildren="否" defaultChecked />
            </Form.Item>
          </Col>
        </Row>
        <Divider orientation="left">
          <PayCircleOutlined /> 租金信息
        </Divider>
        <Row gutter={24}>
          <Col span={6} key={'mouth_rent'} offset={1}>
            <Form.Item label="月租金（整租）" name="mouth_rent">
              <InputNumber
                placeholder="请输入月租金"
                style={{ width: '100%' }}
                addonAfter={'CNY'}
              />
            </Form.Item>
          </Col>
        </Row>
        <Divider />
        <Row gutter={24} justify={'center'}>
          <Form.Item>
            <Space>
              <Button type="primary" htmlType="submit">
                保存
              </Button>
            </Space>
          </Form.Item>
        </Row>
      </Form>
    </div>
  )
}

export default HouseFormUI
