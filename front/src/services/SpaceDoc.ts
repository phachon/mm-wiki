import httpRequest from './http'
import Base from './Base'
import { DocInfoResp, DocSaveReq, DocSaveResp } from '@/types/docType'

const spaceDocUrl = {
  docSave: '/space/doc/save',
  docInfo: '/space/doc/info'
}

/**
 * SpaceDoc 空间 - 文档服务
 */
class SpaceDoc extends Base {
  public constructor() {
    super()
  }

  /**
   * saveDoc 保存文档
   */
  public saveDoc(docInfo: DocSaveReq): Promise<DocSaveResp> {
    const addSpaceUrl = this.getProxyUrl(spaceDocUrl.docSave)
    return httpRequest.post<DocSaveResp>(addSpaceUrl, {}, docInfo)
  }

  /**
   * 获取文档信息
   * @param docId 文档id
   */
  public getDocInfo(docId: number): Promise<DocInfoResp> {
    return httpRequest.get<DocInfoResp>(this.getProxyUrl(spaceDocUrl.docInfo), {
      doc_id: docId
    })
  }
}

export const SpaceDocService = new SpaceDoc()
