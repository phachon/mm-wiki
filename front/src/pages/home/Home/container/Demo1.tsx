import React, { useState } from 'react'
import { Layout } from 'antd'
import Resizable from 'react-resizable'

const { Header, Content, Sider } = Layout

interface ResizableSiderProps {
  initialWidth?: number
  onResize?: (width: number) => void
}

const ResizableSider: React.FC<ResizableSiderProps> = ({ initialWidth = 200, onResize }) => {
  const [siderWidth, setSiderWidth] = useState(initialWidth)

  const handleResize = (e: React.MouseEvent, { size }: { size: { width: number } }) => {
    setSiderWidth(size.width)
    onResize?.(size.width)
  }

  return (
    <></>
    // <Resizable width={siderWidth} height={0} onResize={handleResize}>
    //   <Sider width={siderWidth} style={{ overflow: 'auto' }}>
    //     <div>Sider Content</div>
    //   </Sider>
    // </Resizable>
  )
}

export default ResizableSider
