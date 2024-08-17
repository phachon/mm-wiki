import { Row, Col } from 'antd'

const LoginFooterUI = () => {
  return (
    <footer className="login-footer">
      <Row className="login-footer-bar">
        <Col lg={24} sm={24}>
          <span style={{ marginRight: 12 }}>ICP 证x xx-x-xxx</span>
          <span style={{}}>2023 Copyright © Kitter 科技</span>
        </Col>
      </Row>
    </footer>
  )
}

export default LoginFooterUI
