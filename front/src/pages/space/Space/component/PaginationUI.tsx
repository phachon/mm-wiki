import React from 'react'
import { Pagination } from 'antd'

interface SpacePaginationProps {
  total?: number
  pageSize?: number
  currentPage: number
  onChange?: (page: number) => void
}

const SpacePaginationUI: React.FC<SpacePaginationProps> = ({ total, pageSize, onChange }) => {
  return (
    <div style={{ textAlign: 'center', marginTop: 32 }}>
      <Pagination
        total={total}
        pageSize={pageSize}
        showQuickJumper
        showSizeChanger={false}
        showTotal={(total) => `总共 ${total} 条`}
        onChange={onChange}
      />
    </div>
  )
}

export default SpacePaginationUI
