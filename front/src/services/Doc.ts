import httpRequest from './http'
import Base from './Base'
import { DocSaveReq, DocSaveResp } from '@/types/docType'

const docUrl = {
  docSave: '/doc/save'
}

/**
 * Doc 文档服务
 */
class Doc extends Base {
  public constructor() {
    super()
  }

  /**
   * saveDoc 保存文档
   */
  public saveDoc(docInfo: DocSaveReq): Promise<DocSaveResp> {
    const addSpaceUrl = this.getProxyUrl(docUrl.docSave)
    return httpRequest.post<DocSaveResp>(addSpaceUrl, {}, docInfo)
  }
}

export const DocService = new Doc()
