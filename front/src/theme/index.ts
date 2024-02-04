import { theme } from 'antd'

export const defaultColor = {
  colorPrimary: '#1890ff',
  colorBgBase: '#ffffff',
  imgBgColor: '#f2f2f2'
}

export const darkColor = {
  colorPrimary: '#1890ff',
  colorBgBase: '#202020',
  imgBgColor: '#444444'
}

export const getAntThemeToken = (isLight = true) => {
  let colors = isLight ? defaultColor : darkColor
  return {
    token: {
      colorPrimary: colors.colorPrimary,
      colorBgBase: colors.colorBgBase
    },
    algorithm: isLight ? theme.defaultAlgorithm : theme.darkAlgorithm
  }
}

export const setCssRoot = (isLight = true) => {
  let colors = isLight ? defaultColor : darkColor

  document.documentElement.style.setProperty('--color-brimary', colors.colorPrimary)
  document.documentElement.style.setProperty('--color-bg', colors.colorBgBase)
  document.documentElement.style.setProperty('--img-bg-color', colors.imgBgColor)
}
