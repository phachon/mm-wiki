import { PlusOutlined, EditOutlined, DeleteOutlined, SaveOutlined } from '@ant-design/icons'
import { Button, Divider, Typography, Table, Select, Space, Input, Form } from 'antd'
import { SpacePermissionSelectTagRender, SpacePermissionShowTagUI } from './ToolsUI'
import { useEffect, useState } from 'react'
import './space.css'
import { AccountInfoType } from '@/types/accountType'
import { SpacePermissionInfoType } from '@/types/spaceType'

const options: { label: string; value: string; color: string }[] = [
  { label: '查看', value: '1', color: 'green' },
  { label: '添加', value: '2', color: 'blue' },
  { label: '编辑', value: '3', color: 'orange' },
  { label: '删除', value: '4', color: 'red' },
  { label: '导出', value: '5', color: 'purple' }
]

const handleChange = (value: string[]) => {
  console.log(`selected ${value}`)
}

type SpacePermissionSettingUIProps = {
  adminPermissions?: SpacePermissionInfoType[]
  departmentPermissions?: SpacePermissionInfoType[]
  accountPermissions?: SpacePermissionInfoType[]
  allAccounts?: AccountInfoType[]
  allDepartments?: any[]
}

const SpacePermissionSettingUI = (props: SpacePermissionSettingUIProps) => {
  const [adminIsAdd, setAdminIsAdd] = useState(false)
  const [userGroupEditID, setUserGroupEditID] = useState(0)
  const [userEditID, setUserEditID] = useState(0)
  const form = Form.useFormInstance()

  useEffect(() => {
    if (adminIsAdd) {
      console.log('添加管理员add')
    }
  }, [adminIsAdd])

  const onAdminAdd = () => {
    console.log('添加管理员')
    setAdminIsAdd(true)
  }
  return (
    <div style={{ paddingLeft: 6 }} className="space-permission-setting">
      <div style={{ width: '80%' }}>
        <h3 style={{ marginTop: 8 }}>1、管理员</h3>
        <Table
          dataSource={props.adminPermissions}
          pagination={false}
          footer={() => (
            <Form form={form} name="admin_add" layout="inline">
              <Form.Item
                name="account_name"
                rules={[{ required: true, message: '请输入账号名！' }]}
                style={{ width: '29%' }}
              >
                <Select placeholder="请选择账号名" options={[]} />
              </Form.Item>
              <Form.Item>
                <Button size="small" type="link" style={{ gap: 3 }}>
                  <PlusOutlined />
                  添加
                </Button>
              </Form.Item>
            </Form>
          )}
          bordered={true}
        >
          <Table.Column title="账号名" dataIndex={'name'} width={'30%'} />
          <Table.Column
            title="权限"
            dataIndex={'permission'}
            width={'50%'}
            render={() => {
              return <Typography.Text>拥有全部权限</Typography.Text>
            }}
          />
          <Table.Column
            title="操作"
            dataIndex={'action'}
            align="center"
            render={() => {
              return (
                <Button type="link" size="small" icon={<DeleteOutlined />} style={{ gap: 3 }}>
                  删除
                </Button>
              )
            }}
          />
        </Table>
      </div>
      <Divider />
      <div style={{ width: '80%' }}>
        <h3 style={{ marginTop: 8 }}>2、账号组</h3>
        <Table
          dataSource={props.departmentPermissions}
          pagination={false}
          footer={() => (
            <Form form={form} name="admin_add" layout="inline">
              <Form.Item
                name="account_name"
                rules={[{ required: true, message: '请输入账号名！' }]}
                style={{ width: '29%' }}
              >
                <Select placeholder="请选择账号名" options={[]} />
              </Form.Item>
              <Form.Item
                name="account_name"
                rules={[{ required: true, message: '请输入账号名！' }]}
                style={{ width: '50%' }}
              >
                <Select
                  mode="multiple"
                  tagRender={SpacePermissionSelectTagRender}
                  placeholder="请选择权限"
                  defaultValue={['1', '2']}
                  onChange={handleChange}
                  options={options}
                />
              </Form.Item>
              <Form.Item>
                <Button size="small" type="link" style={{ gap: 3 }}>
                  <PlusOutlined />
                  添加
                </Button>
              </Form.Item>
            </Form>
          )}
          bordered={true}
        >
          <Table.Column title="组织名" dataIndex={'name'} width={'30%'} />
          <Table.Column
            title="权限"
            dataIndex={'permission'}
            width={'50%'}
            render={(value, record, index) => {
              return (
                <div>
                  {record.id === userGroupEditID ? (
                    <Select
                      mode="multiple"
                      tagRender={SpacePermissionSelectTagRender}
                      placeholder="请选择权限"
                      defaultValue={['1', '2']}
                      onChange={handleChange}
                      options={options}
                    />
                  ) : (
                    <SpacePermissionShowTagUI premissionTypes={[1, 2]} />
                  )}
                </div>
              )
            }}
          />
          <Table.Column
            title="操作"
            dataIndex={'action'}
            align="center"
            render={(value, record, index) => {
              return (
                <div>
                  {record.id === userGroupEditID ? (
                    <Button type="link" size="small" icon={<SaveOutlined />} style={{ gap: 3 }}>
                      保存
                    </Button>
                  ) : (
                    <Button
                      type="link"
                      size="small"
                      icon={<EditOutlined />}
                      style={{ gap: 3 }}
                      onClick={() => setUserGroupEditID(record.id)}
                    >
                      编辑
                    </Button>
                  )}
                  <Button type="link" size="small" icon={<DeleteOutlined />} style={{ gap: 3 }}>
                    删除
                  </Button>
                </div>
              )
            }}
          />
        </Table>
      </div>
      <Divider />
      <div style={{ width: '80%' }}>
        <h3 style={{ marginTop: 8 }}>3、普通账号</h3>
        <Table
          dataSource={props.accountPermissions}
          pagination={false}
          footer={() => (
            <Form form={form} name="admin_add" layout="inline">
              <Form.Item
                name="account_name"
                rules={[{ required: true, message: '请输入账号名！' }]}
                style={{ width: '29%' }}
              >
                <Select placeholder="请选择账号名" options={[]} />
              </Form.Item>
              <Form.Item
                name="account_name"
                rules={[{ required: true, message: '请输入账号名！' }]}
                style={{ width: '50%' }}
              >
                <Select
                  mode="multiple"
                  tagRender={SpacePermissionSelectTagRender}
                  placeholder="请选择权限"
                  defaultValue={['1', '2']}
                  onChange={handleChange}
                  options={options}
                />
              </Form.Item>
              <Form.Item>
                <Button size="small" type="link" style={{ gap: 3 }}>
                  <PlusOutlined />
                  添加
                </Button>
              </Form.Item>
            </Form>
          )}
          bordered={true}
        >
          <Table.Column title="账号名" dataIndex={'name'} width={'30%'} />
          <Table.Column
            title="权限"
            dataIndex={'permission'}
            width={'50%'}
            render={() => {
              return (
                <div>
                  {/* <SpacePermissionShowTagUI premissionTypes={[1, 2, 3, 4, 5]} /> */}
                  <Select
                    mode="multiple"
                    tagRender={SpacePermissionSelectTagRender}
                    placeholder="请选择权限"
                    defaultValue={['1', '2']}
                    onChange={handleChange}
                    options={options}
                  />
                </div>
              )
            }}
          />
          <Table.Column
            title="操作"
            dataIndex={'action'}
            align="center"
            render={() => {
              return (
                <div>
                  <Button type="link" size="small" icon={<EditOutlined />}>
                    编辑
                  </Button>
                  <Button type="link" size="small" icon={<SaveOutlined />}>
                    保存
                  </Button>
                  <Button type="link" size="small" icon={<DeleteOutlined />}>
                    删除
                  </Button>
                </div>
              )
            }}
          />
        </Table>
      </div>
    </div>
  )
}

export default SpacePermissionSettingUI
