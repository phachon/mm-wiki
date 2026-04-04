import httpRequest from './http'
import { LinkEditResp, LinkListResp } from '../types/linkType'
import Base from './Base'

const systemLinkUrl = {
  linkSave: '/system/link/save',
  linkEdit: '/system/link/edit',
  linkModify: '/system/link/modify',
  linkList: '/system/link/list',
  linkDelete: '/system/link/delete'
}

class SystemLink extends Base {
  public constructor() {
    super()
  }

  public saveLink(linkInfo: {}): Promise<any> {
    const saveLinkUrl = this.getProxyUrl(systemLinkUrl.linkSave)
    return httpRequest.post<any>(saveLinkUrl, {}, linkInfo)
  }

  public getEditLinkInfo(linkId: number): Promise<LinkEditResp> {
    const linkEditUrl = this.getProxyUrl(systemLinkUrl.linkEdit)
    return httpRequest.get<LinkEditResp>(linkEditUrl, {
      link_id: linkId
    })
  }

  public modifyLink(editLinkInfo: {}): Promise<any> {
    const linkModifyUrl = this.getProxyUrl(systemLinkUrl.linkModify)
    return httpRequest.post<any>(linkModifyUrl, {}, editLinkInfo)
  }

  public getLinkList(
    pageSize?: number,
    pageNum?: number,
    keywords?: {}
  ): Promise<LinkListResp> {
    const linkListUrl = this.getProxyUrl(systemLinkUrl.linkList)
    return httpRequest.get<LinkListResp>(linkListUrl, {
      page_size: pageSize,
      page_num: pageNum,
      keywords: JSON.stringify(keywords)
    })
  }

  public deleteLink(linkId: number): Promise<any> {
    const deleteLinkUrl = this.getProxyUrl(systemLinkUrl.linkDelete)
    return httpRequest.post<any>(deleteLinkUrl, {}, { link_id: linkId })
  }
}

export const SystemLinkService = new SystemLink()
