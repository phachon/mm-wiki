import { Tag } from 'antd'

export const NoticePublishStatusTag = (publishStatus: number) => {
  switch (publishStatus) {
    case 0:
      return <Tag color="error">未发布</Tag>
    case 1:
      return <Tag color="success">已发布</Tag>
    default:
      return <Tag color="cyan">Unkown</Tag>
  }
}
