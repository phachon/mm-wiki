import httpRequest from './http'
import Base from './Base'
import { CollectionType } from '@/types/collectionType'

const userInteractionUrl = {
  collection: '/user/interaction/collection',
  colectionCancel: '/user/interaction/collection_cancel',
  colectionStatus: '/user/interaction/collection_status'
}

/**
 * UserInteraction 用户 - 互动服务
 */
class UserInteraction extends Base {
  public constructor() {
    super()
  }

  /**
   * 收藏文档
   * @param docId 文档id
   */
  public collectionDoc(docId: number): Promise<any> {
    return httpRequest.post(
      this.getProxyUrl(userInteractionUrl.collection),
      {},
      {
        resource_id: docId,
        collection_type: CollectionType.CollectionDoc
      }
    )
  }

  /**
   * 收藏空间
   */
  public collectionSpace(spaceId: number): Promise<any> {
    return httpRequest.post(
      this.getProxyUrl(userInteractionUrl.collection),
      {},
      {
        resource_id: spaceId,
        collection_type: CollectionType.CollectionSpace
      }
    )
  }

  /**
   * 空间取消收藏
   */
  public collectionSpaceCancel(spaceId: number): Promise<any> {
    return httpRequest.post(
      this.getProxyUrl(userInteractionUrl.colectionCancel),
      {},
      {
        resource_id: spaceId,
        collection_type: CollectionType.CollectionSpace
      }
    )
  }

  /**
   * 文档取消收藏
   * @param docId 文档id
   */
  public collectionDocCancel(docId: number): Promise<any> {
    return httpRequest.post(
      this.getProxyUrl(userInteractionUrl.colectionCancel),
      {},
      {
        resource_id: docId,
        collection_type: CollectionType.CollectionDoc
      }
    )
  }

  /**
   * 查找空间收藏状态
   */
  public collectionSpaceStatus(spaceId: number): Promise<any> {
    return httpRequest.get(this.getProxyUrl(userInteractionUrl.colectionStatus), {
      resource_id: spaceId,
      collection_type: CollectionType.CollectionSpace
    })
  }

  /**
   * 查找文档收藏状态
   * @param docId 文档id
   */
  public collectionDocStatus(docId: number): Promise<any> {
    return httpRequest.get(this.getProxyUrl(userInteractionUrl.colectionStatus), {
      resource_id: docId,
      collection_type: CollectionType.CollectionDoc
    })
  }
}

export const UserInteractionService = new UserInteraction()
