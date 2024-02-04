import { Button, Form, Input, Radio, Select, Space, Tooltip, Typography } from 'antd'
import { EditLayoutForm, LayoutForm } from '../../../config/layout'
import { HirerInfoType } from '@/types/hirerType'
import { HirerSexRadioOptions } from './ToolsUI'

interface HirerFormUIProps {
  /**
   * 租客保存方法
   * @param values 保存的值
   * @returns
   */
  onHirerSave: (values: any) => void

  /**
   * 租客修改信息
   */
  hirerInfo?: HirerInfoType

  /**
   * 布局配置
   */
  formLayout?: {
    labelCol: { span: number }
    wrapperCol: { span: number }
  }
}

/**
 * 租客表单 UI 组件
 * @param props 租客表单 UI 组件数据
 * @returns 租客表单 UI 组件
 */
const HirerFormUI = (props: HirerFormUIProps) => {
  const [form] = Form.useForm()
  const isEdit = props.hirerInfo ? true : false
  let layoutForm = isEdit ? EditLayoutForm : LayoutForm
  if (props.formLayout) {
    layoutForm = props.formLayout
  }
  // 每次进来先清除一下原来的数据
  form.resetFields()
  if (props.hirerInfo) {
    form.setFieldsValue(props.hirerInfo)
  }

  const sexOptions = HirerSexRadioOptions()

  return (
    <div className="panel-body">
      <Form {...layoutForm} name="basic" onFinish={props.onHirerSave} form={form}>
        {isEdit && (
          <Form.Item label="租客id" name="hirer_id" rules={[{ required: true }]}>
            <Input disabled placeholder="请输入租客ID" />
          </Form.Item>
        )}

        <Form.Item label="姓名" name="name" rules={[{ required: true, message: '请输入姓名!' }]}>
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

        <Form.Item label="身份证" name="id_card_number">
          <Input placeholder="请输入身份证号" />
        </Form.Item>

        <Form.Item label="邮箱" name="email">
          <Input placeholder="请输入邮箱地址: xxx@xxx.com" />
        </Form.Item>

        <Form.Item label="紧急联系人" name="emergency_contact">
          <Input placeholder="请输入紧急联系人姓名" />
        </Form.Item>

        <Form.Item label="联系人电话" name="emergency_mobile">
          <Input placeholder="请输入紧急联系电话" />
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

export default HirerFormUI
