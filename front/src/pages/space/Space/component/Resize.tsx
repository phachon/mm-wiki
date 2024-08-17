import React, { useState } from 'react'
import { ResizableBox, ResizableBoxProps, ResizeCallbackData } from 'react-resizable'
import 'react-resizable/css/styles.css'

const ResizableComponent: React.FC = () => {
  const [width, setWidth] = useState(300)

  const handleResize = (e: React.SyntheticEvent, data: ResizeCallbackData) => {
    setWidth(data.size.width)
  }

  return (
    <ResizableBox
      width={width}
      height={0}
      onResize={handleResize}
      handleSize={[20, 20]}
      minConstraints={[200, Infinity]}
      maxConstraints={[800, Infinity]}
      axis="x"
    >
      <div
        style={{
          width: `${width}px`,
          border: '1px solid #ccc',
          padding: '10px',
          overflow: 'auto'
        }}
      >
        <h3>Resizable Content</h3>
        <p>
          Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nulla facilisi. Proin malesuada
          cursus sapien in elementum. Sed vel consequat enim, at rhoncus velit. Donec varius lectus
          id odio venenatis, eu tristique justo cursus. Duis in velit at lorem consequat bibendum.
          Pellentesque habitant morbi tristique senectus et netus et malesuada fames ac turpis
          egestas. Sed id augue consequat, pulvinar sem et, fringilla diam.
        </p>
      </div>
    </ResizableBox>
  )
}

export default ResizableComponent
