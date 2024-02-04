import React, { useEffect, useState } from 'react'
import HirerListUI from '../component/ListUI'
import { HirerEditResp, HirerInfoType, HirerListResp } from '@/types/hirerType'
import { HirerService } from '@/services/Hirer'
import { message, Modal, TablePaginationConfig } from 'antd'
import HirerSearchUI from '../component/SearchUI'
import HirerFormUI from '../component/FormUI'
import HirerDetailUI from '../component/DetailUI'
import { initPagination } from '@/types/adminType'

let searchValues = {}

const HirerList: React.FC = () => {
  const [hirerList, setHirerList] = useState<HirerInfoType[]>([])
  const [pagination, setPagination] = useState(initPagination)
  const [editModalOpen, setEditModalOpen] = useState<boolean>(false)
  const [editHirerInfo, setEditHirerInfo] = useState<HirerInfoType>()
  const [detailHirerInfo, setDetailHirerInfo] = useState<HirerInfoType>()
  const [detailModalOpen, setDetailModalOpen] = useState<boolean>(false)

  useEffect(() => {
    getHirerList(initPagination, {})
  }, [])

  /**
   * 请求租客列表
   * @param pageConfig 翻页信息
   * @param searchKeywords 搜索信息
   */
  const getHirerList = (pageConfig: TablePaginationConfig, searchKeywords: {}) => {
    searchValues = searchKeywords
    HirerService.hirerList(pageConfig.pageSize, pageConfig.current, searchKeywords)
      .then((res: HirerListResp) => {
        setHirerList(res.list)
        setPagination({
          ...initPagination,
          current: res.page_info?.page_num,
          pageSize: res.page_info?.page_size,
          total: res.page_info?.total_num
        })
      })
      .catch((e) => {
        console.log('获取租客列表失败：', e)
      })
  }

  /**
   * 搜索查询操作
   * @param values 搜索表单数据
   */
  const onHirerSearchChange = (values: {}) => {
    getHirerList(initPagination, values)
  }

  /**
   * 搜索重置操作
   */
  const onHirerSearchReset = () => {
    getHirerList(initPagination, {})
  }

  /**
   * 列表分页请求
   * @param pageConfig
   * @param filters
   * @param sorter
   */
  const onHirerListChange = (pageConfig: TablePaginationConfig, filters: any, sorter: any) => {
    getHirerList(pageConfig, searchValues)
  }

  /**
   * 修改点击操作
   * @param hirerInfo 租客信息
   */
  const onHirerEditClick = (hirerInfo: HirerInfoType) => {
    // 获取修改租客需要的信息
    HirerService.getEditHirerInfo(hirerInfo.hirer_id)
      .then((editHirerInfo: HirerEditResp) => {
        setEditHirerInfo(editHirerInfo.hirer_info)
        setEditModalOpen(true)
      })
      .catch((e) => {
        console.log('获取修改信息失败err:', e)
      })
  }

  /**
   * 租客详情点击操作
   * @param hirerInfo 租客信息
   */
  const onHirerDetailClick = (hirerInfo: HirerInfoType) => {
    setDetailModalOpen(true)
    setDetailHirerInfo(hirerInfo)
  }

  /**
   * 删除租客操作
   * @param hirerInfo
   * @param status
   */
  const onHirerDeleteConfim = (hirerInfo: HirerInfoType) => {
    HirerService.deleteHirer(hirerInfo.hirer_id)
      .then(() => {
        message.success('删除成功', 2, () => {
          getHirerList(pagination, searchValues)
        })
      })
      .catch((e) => {
        console.log('租客删除失败err:', e)
      })
  }

  /**
   * 修改弹框取消操作
   */
  const onEditModalCancel = () => {
    setEditModalOpen(false)
  }

  /**
   * 修改保存操作
   * @param hirerInfo HirerInfoType
   */
  const onHirerSave = (hirerInfo: HirerInfoType) => {
    HirerService.modifyHirer(hirerInfo)
      .then(() => {
        message.success('修改成功', 2, () => {
          setEditModalOpen(false)
          getHirerList(pagination, searchValues)
        })
      })
      .catch((e) => {
        console.log('修改保存失败:', e)
      })
  }

  return (
    <div className="panel">
      <HirerSearchUI
        onHirerSearchChange={onHirerSearchChange}
        onHirerSearchReset={onHirerSearchReset}
      />
      <HirerListUI
        listLoading={false}
        hirerList={hirerList}
        pagination={pagination}
        onHirerListChange={onHirerListChange}
        onHirerEditClick={onHirerEditClick}
        onHirerDeleteConfim={onHirerDeleteConfim}
        onHirerDetailClick={onHirerDetailClick}
      />
      <Modal
        title="租客修改"
        width={570}
        open={editModalOpen}
        onCancel={onEditModalCancel}
        footer={null}
      >
        <HirerFormUI hirerInfo={editHirerInfo} onHirerSave={onHirerSave} />
      </Modal>
      <Modal
        title="租客详情"
        width={570}
        open={detailModalOpen}
        onCancel={() => {
          setDetailModalOpen(false)
        }}
        footer={null}
      >
        <HirerDetailUI hirerDetail={detailHirerInfo} />
      </Modal>
    </div>
  )
}

export default HirerList
