import { Layout } from 'antd'

/**
 * SpaceFooterUIProps 底部区块 props
 */
interface SpaceFooterUIProps {
  text?: string
}

const SpaceFooterUI = (props: SpaceFooterUIProps) => {
  return <Layout.Footer className="admin-footer">{props.text}</Layout.Footer>
}

export default SpaceFooterUI
