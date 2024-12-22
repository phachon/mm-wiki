import httpRequest from './http'
import Base from './Base'
import {
  DocAddSaveReq,
  DocAddSaveResp,
  DocInfoResp,
  DocContentSaveReq,
  DocContentSaveResp
} from '@/types/docType'

const spaceDocUrl = {
  docSave: '/space/doc/save',
  docInfo: '/space/doc/info',
  docContentSave: '/space/doc/content_save',
  docUploadFile: '/space/doc/upload_file'
}

/**
 * SpaceDoc 空间 - 文档服务
 */
class SpaceDoc extends Base {
  public constructor() {
    super()
  }

  /**
   * addSaveDoc 添加文档保存
   */
  public addSaveDoc(docInfo: DocAddSaveReq): Promise<DocAddSaveResp> {
    const addSpaceUrl = this.getProxyUrl(spaceDocUrl.docSave)
    return httpRequest.post<DocAddSaveResp>(addSpaceUrl, {}, docInfo)
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

  /**
   * 保存文档内容
   * @param docContent 文档内容
   */
  public saveDocContent(docContent: DocContentSaveReq): Promise<DocContentSaveResp> {
    return httpRequest.post<DocContentSaveResp>(
      this.getProxyUrl(spaceDocUrl.docContentSave),
      {},
      docContent
    )
  }

  /**
   * 上传文件
   * @param file 文件
   * @param params 参数
   */
  public uploadFile(file: File, params: { [key: string]: any } = {}): Promise<any> {
    return httpRequest.uploadFile(this.getProxyUrl(spaceDocUrl.docUploadFile), file, params)
  }
}

export const SpaceDocService = new SpaceDoc()
