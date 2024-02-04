import { Button, Form, Input, Radio, Select, Space, Tooltip, Typography } from 'antd'
import { EditLayoutForm, LayoutForm } from '../../../config/layout'
import { LandlordInfoType } from '@/types/landlordType'
import { LandlordSexRadioOptions } from './ToolsUI'

interface LandlordFormUIProps {
  /**
   * 业主保存方法
   * @param values 保存的值
   * @returns
   */
  onLandlordSave: (values: any) => void

  /**
   * 业主修改信息
   */
  landlordInfo?: LandlordInfoType

  /**
   * 布局配置
   */
  formLayout?: {
    labelCol: { span: number }
    wrapperCol: { span: number }
  }
}

/**
 * 业主表单 UI 组件
 * @param props 业主表单 UI 组件数据
 * @returns 业主表单 UI 组件
 */
const LandlordFormUI = (props: LandlordFormUIProps) => {
  const [form] = Form.useForm()
  const isEdit = props.landlordInfo ? true : false
  let layoutForm = isEdit ? EditLayoutForm : LayoutForm
  if (props.formLayout) {
    layoutForm = props.formLayout
  }
  // 每次进来先清除一下原来的数据
  form.resetFields()
  if (props.landlordInfo) {
    form.setFieldsValue(props.landlordInfo)
  }

  const sexOptions = LandlordSexRadioOptions()

  return (
    <div className="panel-body">
      <Form {...layoutForm} name="basic" onFinish={props.onLandlordSave} form={form}>
        {isEdit && (
          <Form.Item label="业主id" name="landlord_id" rules={[{ required: true }]}>
            <Input disabled placeholder="请输入业主ID" />
          </Form.Item>
        )}

        <Form.Item
          name="nick_name"
          label="昵称"
          rules={[{ required: true, message: '请输入昵称!' }]}
        >
          <Space>
            <Form.Item name="nick_name" noStyle>
              <Input placeholder="请输入昵称（昵称必须唯一）" style={{ width: 220 }} />
            </Form.Item>
            <Typography.Text type="danger">*昵称必须唯一</Typography.Text>
          </Space>
        </Form.Item>

        <Form.Item
          label="姓名"
          name="given_name"
          rules={[{ required: true, message: '请输入姓名!' }]}
        >
          <Input placeholder="请输入姓名" />
        </Form.Item>

        <Form.Item
          label="性别"
          name="sex"
          rules={[{ required: true, message: '请选择性别!' }]}
          initialValue={0}
        >
          <Radio.Group options={sexOptions} />
        </Form.Item>

        <Form.Item
          label="手机"
          name="mobile"
          rules={[{ required: true, message: '请输入手机号码!' }]}
        >
          <Input placeholder="请输入手机号码" />
        </Form.Item>

        <Form.Item
          label="住址"
          name="address"
          rules={[{ required: true, message: '请输入现住址!' }]}
        >
          <Input placeholder="请输入现住址" />
        </Form.Item>

        <Form.Item label="身份证" name="id_card_number">
          <Input placeholder="请输入身份证号" />
        </Form.Item>

        <Form.Item label="邮箱" name="email">
          <Input placeholder="请输入邮箱地址: xxx@xxx.com" />
        </Form.Item>

        <Form.Item wrapperCol={{ offset: layoutForm.labelCol.span }}>
          <Button type="primary" htmlType="submit">
            保存
          </Button>
        </Form.Item>
      </Form>
    </div>
  )
}

export default LandlordFormUI
