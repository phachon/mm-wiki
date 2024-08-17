import { HOME_ROOT_PATH } from '@/router/type'
import { Result, Button } from 'antd'
import { useNavigate } from 'react-router-dom'

const Error404 = () => {
  const navigate = useNavigate()

  function toHome() {
    navigate(HOME_ROOT_PATH)
  }

  return (
    <>
      <Result
        status="404"
        title="Not Found"
        subTitle="对不起, 您访问的页面不存在！"
        extra={
          <Button type="primary" onClick={toHome}>
            返回主页
          </Button>
        }
      />
    </>
  )
}

export default Error404
