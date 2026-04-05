import { theme } from 'antd'

export const defaultColor = {
  colorPrimary: '#1677ff',
  colorBgBase: '#ffffff',
  imgBgColor: '#f5f5f5'
}

export const darkColor = {
  colorPrimary: '#1677ff',
  colorBgBase: '#202020',
  imgBgColor: '#444444'
}

export const getAntThemeToken = (isLight = true) => {
  const colors = isLight ? defaultColor : darkColor
  return {
    token: {
      colorPrimary: colors.colorPrimary,
      colorBgBase: colors.colorBgBase,
      borderRadius: 6
    },
    algorithm: isLight ? theme.defaultAlgorithm : theme.darkAlgorithm
  }
}

export const setCssRoot = (isLight = true) => {
  const colors = isLight ? defaultColor : darkColor

  document.documentElement.style.setProperty('--color-primary', colors.colorPrimary)
  document.documentElement.style.setProperty('--color-bg', colors.colorBgBase)
  document.documentElement.style.setProperty('--img-bg-color', colors.imgBgColor)
}
