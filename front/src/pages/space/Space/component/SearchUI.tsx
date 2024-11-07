import React from 'react'
import { Col, Input, Row } from 'antd'

const { Search } = Input

interface SpaceSearchUIProps {
  onSearch?: (value: string) => void
}

const SpaceSearchUI: React.FC<SpaceSearchUIProps> = ({ onSearch }) => {
  return (
    <Row justify={'center'}>
      <Col span={12}>
        <Search
          className="space-search"
          placeholder="请输入空间名"
          allowClear
          enterButton="搜索"
          size="large"
          onSearch={onSearch}
        />
      </Col>
    </Row>
  )
}

export default SpaceSearchUI
