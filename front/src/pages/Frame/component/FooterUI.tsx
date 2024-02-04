import { Layout } from 'antd'

/**
 * FrameFooterUIProps 底部区块 props
 */
interface FrameFooterUIProps {
  text?: string
}

const FrameFooterUI = (props: FrameFooterUIProps) => {
  return <Layout.Footer className="admin-footer">{props.text}</Layout.Footer>
}

export default FrameFooterUI
