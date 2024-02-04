import { Breadcrumb } from 'antd'
import { useGlobalStore } from '@/stores/index'

const FrameBreadcrumbUI = () => {
  const { iBreadcrumbItems } = useGlobalStore()
  return (
    <div
      className="admin-breadcrumb"
      style={{ marginTop: iBreadcrumbItems && iBreadcrumbItems.length > 0 ? 8 : 0 }}
    >
      <div className="admin-breadcrumb-nav">
        <Breadcrumb
          style={{ margin: '0 4px', lineHeight: '35px' }}
          separator="/"
          items={iBreadcrumbItems}
        />
      </div>
    </div>
  )
}

export default FrameBreadcrumbUI
