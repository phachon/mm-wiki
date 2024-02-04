import httpRequest from './http'
import { NoticeEditResp, NoticeListResp } from '../types/noticeType'
import Base from './Base'

const noticeUrl = {
  noticeSave: '/system/notice/save',
  noticeEdit: '/system/notice/edit',
  noticeModify: '/system/notice/modify',
  noticeList: '/system/notice/list',
  noticePublishList: '/system/notice/publish_list',
  noticeDelete: '/system/notice/delete'
}

/**
 * Notice 公告服务
 */
class Notice extends Base {
  public constructor() {
    super()
  }

  /**
   * saveNotice 添加公告保存
   * @param noticeInfo 添加公告信息
   */
  public saveNotice(noticeInfo: {}): Promise<any> {
    const saveNoticeUrl = this.getProxyUrl(noticeUrl.noticeSave)
    return httpRequest.post<any>(saveNoticeUrl, {}, noticeInfo)
  }

  /**
   * getEditNoticeInfo 获取编辑公告信息
   * @param noticeId 公告id
   */
  public getEditNoticeInfo(noticeId: number): Promise<NoticeEditResp> {
    const noticeEditUrl = this.getProxyUrl(noticeUrl.noticeEdit)
    return httpRequest.get<NoticeEditResp>(noticeEditUrl, {
      notice_id: noticeId
    })
  }

  /**
   * modifyNotice 更新公告保存
   * @param editNoticeInfo 修改的信息
   * @returns
   */
  public modifyNotice(editNoticeInfo: {}): Promise<any> {
    const noticeModifyUrl = this.getProxyUrl(noticeUrl.noticeModify)
    return httpRequest.post<any>(noticeModifyUrl, {}, editNoticeInfo)
  }

  /**
   * getNoticeList 获取公告列表
   * @param pageSize 每一页条数
   * @param pageNum 页数
   * @param keywords 搜索值
   */
  public getNoticeList(
    pageSize?: number,
    pageNum?: number,
    keywords?: {}
  ): Promise<NoticeListResp> {
    const noticeListUrl = this.getProxyUrl(noticeUrl.noticeList)
    return httpRequest.get<NoticeListResp>(noticeListUrl, {
      page_size: pageSize,
      page_num: pageNum,
      keywords: JSON.stringify(keywords)
    })
  }

  /**
   * getPublishNoticeList 获取已发布公告列表
   * @param pageSize 每一页条数
   * @param pageNum 页数
   */
  public getPublishNoticeList(pageSize?: number, pageNum?: number): Promise<NoticeListResp> {
    const noticePublishListUrl = this.getProxyUrl(noticeUrl.noticePublishList)
    return httpRequest.get<NoticeListResp>(noticePublishListUrl, {
      page_size: pageSize,
      page_num: pageNum
    })
  }

  /**
   * deleteNotice 删除公告
   * @param noticeId 公告id
   */
  public deleteNotice(noticeId: number): Promise<any> {
    const deleteNoticeUrl = this.getProxyUrl(noticeUrl.noticeDelete)
    return httpRequest.post<any>(deleteNoticeUrl, {
      notice_id: noticeId
    })
  }
}

export const NoticeService = new Notice()
