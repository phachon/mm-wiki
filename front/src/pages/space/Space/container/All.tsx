import React, { useState, useEffect } from 'react'
import { Col, Divider, Row, message } from 'antd'
import SpaceSearchUI from '../component/SearchUI'
import SpaceCardUI from '../component/CardUI'
import SpacePaginationUI from '../component/PaginationUI'
import { SpaceService } from '@/services/Space'
import type { SpaceInfoType } from '@/types/spaceType'
import type { PageInfoType } from '@/types/baseType'
import '../component/space.css'

/**
 * SpaceAll 空间列表组件
 * 展示所有公开的空间，支持搜索和分页
 */
const SpaceAll: React.FC = () => {
  const [currentPage, setCurrentPage] = useState(1) // 当前页码
  const [spaceList, setSpaceList] = useState<SpaceInfoType[]>([]) // 空间列表数据
  const [pageInfo, setPageInfo] = useState<PageInfoType>({
    total_num: 0,
    page_num: 1,
    page_size: 12,
    total_page: 0,
    has_next: 1
  }) // 分页信息
  const [searchKeywords, setSearchKeywords] = useState('') // 搜索关键词

  /**
   * 获取空间列表数据
   * @param page 页码
   * @param keywords 搜索关键词
   */
  const fetchSpaces = async (page: number, keywords: string = '') => {
    try {
      const response = await SpaceService.getSpaces(
        pageInfo.page_size,
        page,
        keywords ? { space_name: keywords } : undefined
      )
      setSpaceList(response.list)
      setPageInfo(response.page_info)
    } catch (error) {
      console.error('获取空间列表失败:', error)
    }
  }

  // 页面加载时获取数据
  useEffect(() => {
    fetchSpaces(currentPage)
  }, [currentPage])

  /**
   * 处理搜索事件
   * @param value 搜索关键词
   */
  const handleSearch = (value: string) => {
    setSearchKeywords(value)
    setCurrentPage(1) // 搜索时重置页码
    fetchSpaces(1, value)
  }

  /**
   * 处理分页变化
   * @param page 新的页码
   */
  const handlePageChange = (page: number) => {
    setCurrentPage(page)
  }

  /**
   * 处理空间收藏状态变化
   * @param spaceKey 空间标识
   * @param collected 是否收藏
   */
  const handleCollectionChange = async (spaceKey: string, collected: boolean) => {
    try {
      if (collected) {
        // 调用收藏接口
        await SpaceService.collectSpace(spaceKey)
        message.success('收藏成功')
      } else {
        // 调用取消收藏接口
        await SpaceService.uncollectSpace(spaceKey)
        message.success('已取消收藏')
      }
      // 刷新列表
      fetchSpaces(currentPage, searchKeywords)
    } catch (error) {
      message.error('操作失败，请稍后重试')
    }
  }

  return (
    <div>
      <SpaceSearchUI onSearch={handleSearch} />
      <Divider style={{ margin: '36px 0' }} />
      <Row gutter={24}>
        {spaceList.map((space) => (
          <Col key={space.space_id} span={6}>
            <SpaceCardUI
              title={space.name}
              description={space.description}
              creator={space.creator_name}
              createTime={space.create_time}
              spaceKey={space.space_key}
              isCollected={false}
              onCollectionChange={handleCollectionChange}
            />
          </Col>
        ))}
        <Col span={24}>
          <SpacePaginationUI
            total={pageInfo.total_num}
            pageSize={pageInfo.page_size}
            currentPage={currentPage}
            onChange={handlePageChange}
          />
        </Col>
      </Row>
    </div>
  )
}

export default SpaceAll
