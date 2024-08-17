import { Layout } from 'antd'

/**
 * SystemFooterUIProps 底部区块 props
 */
interface SystemFooterUIProps {
  text?: string
}

const SystemFooterUI = (props: SystemFooterUIProps) => {
  return <Layout.Footer className="admin-footer">{props.text}</Layout.Footer>
}

export default SystemFooterUI
