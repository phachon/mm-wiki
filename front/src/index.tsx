import { ConfigProvider } from 'antd'
import { BrowserRouter } from 'react-router-dom'
import zhCN from 'antd/es/locale/zh_CN'
import { createRoot } from 'react-dom/client'
import RouterView from '@/router'
import { getAntThemeToken } from '@/theme'
import './assets/styles/style.css'

const container = document.getElementById('root')
if (container != null) {
  const root = createRoot(container)
  root.render(
    <ConfigProvider locale={zhCN} theme={getAntThemeToken(true)}>
      <BrowserRouter>
        <RouterView />
      </BrowserRouter>
    </ConfigProvider>
  )
}
