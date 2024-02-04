import { HOME_ROOT_PATH } from '@/router/type'
import { Result, Button } from 'antd'
import { useNavigate } from 'react-router-dom'

const Error403 = () => {
  const navigate = useNavigate()

  function toHome() {
    navigate(HOME_ROOT_PATH)
  }

  return (
    <>
      <Result
        status="403"
        title="Forbidden"
        subTitle="对不起, 您没有权限访问该页面！"
        extra={
          <Button type="primary" onClick={toHome}>
            返回主页
          </Button>
        }
      />
    </>
  )
}

export default Error403
