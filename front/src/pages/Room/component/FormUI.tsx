import {
  Button,
  Col,
  Divider,
  Form,
  Input,
  InputNumber,
  Radio,
  Row,
  Select,
  Space,
  Switch
} from 'antd'
import { EditLayoutForm, LayoutForm } from '../../../config/layout'
import { RoomInfoType } from '@/types/roomType'
import { HouseInfoType } from '@/types/houseType'
import { DefaultOptionType } from 'antd/lib/select'
import { useEffect } from 'react'
import { RoomDirectionTypeRadioOptions, RoomRoomTypeRadioOptions } from './ToolsUI'
import { HouseShowWholeAddress } from '@/pages/House/component/ToolsUI'
import { FormOutlined, PayCircleOutlined, SettingOutlined } from '@ant-design/icons'

interface RoomFormUIProps {
  /**
   * 房产列表
   */
  houseList?: HouseInfoType[]
  /**
   * 房间信息
   */
  roomInfo?: RoomInfoType
  /**
   * 表单布局
   */
  formLayout?: {
    labelCol: { span: number }
    wrapperCol: { span: number }
  }
  /**
   * 保存操作方法
   * @param values
   */
  onFinishCallback: (values: any) => void
}

/**
 * 获取房产 Option 数据
 * @param houseList 房产列表
 * @returns
 */
const getHouseOptions = (houseList: HouseInfoType[] | undefined): DefaultOptionType[] => {
  let houseOptions: DefaultOptionType[] = []
  if (!houseList) {
    return houseOptions
  }
  houseList.forEach((houseInfo) => {
    houseOptions.push({
      label: HouseShowWholeAddress(houseInfo),
      value: String(houseInfo.house_id)
    })
  })
  return houseOptions
}

/**
 * 房间表单 UI 组件
 * @param props 房间表单 UI 组件数据
 * @returns 房间表单 UI 组件
 */
const RoomFormUI = (props: RoomFormUIProps) => {
  const [form] = Form.useForm()
  const isEdit = props.roomInfo ? true : false

  const directionTypeOptions = RoomDirectionTypeRadioOptions()
  const roomTypeOptions = RoomRoomTypeRadioOptions()

  let layoutForm = isEdit ? EditLayoutForm : LayoutForm
  if (props.formLayout) {
    layoutForm = props.formLayout
  }

  useEffect(() => {
    if (props.roomInfo) {
      form.setFieldsValue(props.roomInfo)
    }
    if (props.roomInfo?.house_id) {
      form.setFieldValue('house_id', String(props.roomInfo.house_id))
    }
  }, [props.roomInfo])

  return (
    <div className="panel-body">
      <Form name="basic" onFinish={props.onFinishCallback} form={form} layout="vertical">
        <Divider orientation="left">
          <FormOutlined /> 基础信息
        </Divider>
        <Row gutter={24}>
          <Col span={6} key={'room_id'} offset={1}>
            <Form.Item label="房间id" name="room_id">
              <Input disabled placeholder="房间ID保存后自动生成" />
            </Form.Item>
          </Col>
          <Col span={6} key={'house_id'} offset={1}>
            <Form.Item
              label="所属房产"
              name="house_id"
              rules={[{ required: true, message: '请选择房产!' }]}
            >
              <Select
                style={{ width: '100%' }}
                placeholder="请选择房产"
                options={getHouseOptions(props.houseList)}
              />
            </Form.Item>
          </Col>
          <Col span={7} key={'name'} offset={1}>
            <Form.Item
              label="房间名称"
              name="name"
              rules={[{ required: true, message: '请输入房间名!' }]}
            >
              <Input placeholder="请输入房间名" />
            </Form.Item>
          </Col>

          <Col span={6} key={'area'} offset={1}>
            <Form.Item
              label="房间面积"
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
          <Col span={6} key={'room_type'} offset={1}>
            <Form.Item
              label="房间类型"
              name="room_type"
              rules={[{ required: true, message: '请选择房间类型!' }]}
              initialValue={0}
            >
              <Radio.Group options={roomTypeOptions} />
            </Form.Item>
          </Col>
          <Col span={7} key={'direction_type'} offset={1}>
            <Form.Item
              label="房间朝向"
              name="direction_type"
              rules={[{ required: true, message: '请选择房间朝向!' }]}
              initialValue={0}
            >
              <Radio.Group options={directionTypeOptions} />
            </Form.Item>
          </Col>
        </Row>
        <Divider orientation="left">
          <SettingOutlined /> 配置信息
        </Divider>
        <Row gutter={24}>
          <Col span={6} key={'has_toilet'} offset={1}>
            <Form.Item
              label="包含独卫"
              name="has_toilet"
              valuePropName="checked"
              initialValue={0}
              getValueProps={(value) => ({ checked: value === 1 ? true : false })}
              getValueFromEvent={(value) => {
                return value ? 1 : 0
              }}
              rules={[{ required: true, message: '请选择是否有独卫!' }]}
            >
              <Switch checkedChildren="是" unCheckedChildren="否" defaultChecked />
            </Form.Item>
          </Col>
          <Col span={6} key={'has_balcony'} offset={1}>
            <Form.Item
              label="包含阳台"
              name="has_balcony"
              valuePropName="checked"
              initialValue={0}
              getValueProps={(value) => ({ checked: value === 1 ? true : false })}
              getValueFromEvent={(value) => {
                return value ? 1 : 0
              }}
              rules={[{ required: true, message: '请选择是否有阳台!' }]}
            >
              <Switch checkedChildren="是" unCheckedChildren="否" defaultChecked />
            </Form.Item>
          </Col>
          <Col span={7} key={'allow_lease'} offset={1}>
            <Form.Item
              label="允许出租"
              name="allow_lease"
              valuePropName="checked"
              initialValue={0}
              getValueProps={(value) => ({ checked: value === 1 ? true : false })}
              getValueFromEvent={(value) => {
                return value ? 1 : 0
              }}
              rules={[{ required: true, message: '请选择是否允许出租!' }]}
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
            <Form.Item label="月租金" name="mouth_rent">
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

export default RoomFormUI
