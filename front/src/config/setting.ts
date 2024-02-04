// @ts-ignore
import logoImg from '../assets/images/logo_2.png'
/**
 * ISetting 系统设置
 */
export interface ISetting {
  /**
   * @name string 登录页面标题
   */
  loginPageTitle?: string
  /**
   * @name string 登录页面 icon 地址
   */
  loginPageIconUrl?: string
  /**
   * @name string 登录页面副标题
   */
  loginPageSubTitle?: string
  /**
   * @name boolean 登录是否使用域账号登录
   */
  loginUseDomainLogin?: boolean
  /**
   * @name string 版权信息 Copyright © 2023 Kitter Incorporated. All rights reserved.
   */
  copyright?: string
  /**
   * @name string 后台框架头部标题
   */
  frameHeaderTitle?: string
  /**
   * @name string 后台框架头部icon地址
   */
  frameHeaderIconUrl?: string
  /**
   * @name string 页角显示文案
   */
  footerShowText?: string
}

/**
 * 默认的设置配置
 */
export const SettingConfig: ISetting = {
  loginPageTitle: 'MM-Wiki 系统登录',
  loginPageIconUrl: logoImg,
  loginPageSubTitle: '一个轻量级的企业知识分享与团队协同软件',
  loginUseDomainLogin: false,
  copyright: 'Copyright © 2023 @Kitter Incorporated. All rights reserved.',
  frameHeaderTitle: ' MM-Wiki ',
  frameHeaderIconUrl: logoImg,
  footerShowText: 'MM-Wiki © 2023 Created by @kitter'
}
