import React, { useEffect, useState } from 'react'
import LandlordListUI from '../component/ListUI'
import { LandlordEditResp, LandlordInfoType, LandlordListResp } from '@/types/landlordType'
import { LandlordService } from '@/services/Landlord'
import { message, Modal, TablePaginationConfig } from 'antd'
import LandlordSearchUI from '../component/SearchUI'
import LandlordFormUI from '../component/FormUI'
import LandlordDetailUI from '../component/DetailUI'
import { initPagination } from '@/types/adminType'

let searchValues = {}

const LandlordList: React.FC = () => {
  const [landlordList, setLandlordList] = useState<LandlordInfoType[]>([])
  const [pagination, setPagination] = useState(initPagination)
  const [editModalOpen, setEditModalOpen] = useState<boolean>(false)
  const [editLandlordInfo, setEditLandlordInfo] = useState<LandlordInfoType>()
  const [detailLandlordInfo, setDetailLandlordInfo] = useState<LandlordInfoType>()
  const [detailModalOpen, setDetailModalOpen] = useState<boolean>(false)

  useEffect(() => {
    getLandlordList(initPagination, {})
  }, [])

  /**
   * 请求业主列表
   * @param pageConfig 翻页信息
   * @param searchKeywords 搜索信息
   */
  const getLandlordList = (pageConfig: TablePaginationConfig, searchKeywords: {}) => {
    searchValues = searchKeywords
    LandlordService.landlordList(pageConfig.pageSize, pageConfig.current, searchKeywords)
      .then((res: LandlordListResp) => {
        setLandlordList(res.list)
        setPagination({
          ...initPagination,
          current: res.page_info?.page_num,
          pageSize: res.page_info?.page_size,
          total: res.page_info?.total_num
        })
      })
      .catch((e) => {
        console.log('获取业主列表失败：', e)
      })
  }

  /**
   * 搜索查询操作
   * @param values 搜索表单数据
   */
  const onLandlordSearchChange = (values: {}) => {
    getLandlordList(initPagination, values)
  }

  /**
   * 搜索重置操作
   */
  const onLandlordSearchReset = () => {
    getLandlordList(initPagination, {})
  }

  /**
   * 列表分页请求
   * @param pageConfig
   * @param filters
   * @param sorter
   */
  const onLandlordListChange = (pageConfig: TablePaginationConfig, filters: any, sorter: any) => {
    getLandlordList(pageConfig, searchValues)
  }

  /**
   * 修改点击操作
   * @param landlordInfo 业主信息
   */
  const onLandlordEditClick = (landlordInfo: LandlordInfoType) => {
    // 获取修改业主需要的信息
    LandlordService.getEditLandlordInfo(landlordInfo.landlord_id)
      .then((editLandlordInfo: LandlordEditResp) => {
        setEditLandlordInfo(editLandlordInfo.landlord_info)
        setEditModalOpen(true)
      })
      .catch((e) => {
        console.log('获取修改信息失败err:', e)
      })
  }

  /**
   * 业主详情点击操作
   * @param landlordInfo 业主信息
   */
  const onLandlordDetailClick = (landlordInfo: LandlordInfoType) => {
    setDetailModalOpen(true)
    setDetailLandlordInfo(landlordInfo)
  }

  /**
   * 删除业主操作
   * @param landlordInfo
   * @param status
   */
  const onLandlordDeleteConfim = (landlordInfo: LandlordInfoType) => {
    LandlordService.deleteLandlord(landlordInfo.landlord_id)
      .then(() => {
        message.success('删除成功', 2, () => {
          getLandlordList(pagination, searchValues)
        })
      })
      .catch((e) => {
        console.log('业主删除失败err:', e)
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
   * @param landlordInfo LandlordInfoType
   */
  const onLandlordSave = (landlordInfo: LandlordInfoType) => {
    LandlordService.modifyLandlord(landlordInfo)
      .then(() => {
        message.success('修改成功', 2, () => {
          setEditModalOpen(false)
          getLandlordList(pagination, searchValues)
        })
      })
      .catch((e) => {
        console.log('修改保存失败:', e)
      })
  }

  return (
    <div className="panel">
      <LandlordSearchUI
        onLandlordSearchChange={onLandlordSearchChange}
        onLandlordSearchReset={onLandlordSearchReset}
      />
      <LandlordListUI
        listLoading={false}
        landlordList={landlordList}
        pagination={pagination}
        onLandlordListChange={onLandlordListChange}
        onLandlordEditClick={onLandlordEditClick}
        onLandlordDeleteConfim={onLandlordDeleteConfim}
        onLandlordDetailClick={onLandlordDetailClick}
      />
      <Modal
        title="业主修改"
        width={570}
        open={editModalOpen}
        onCancel={onEditModalCancel}
        footer={null}
      >
        <LandlordFormUI landlordInfo={editLandlordInfo} onLandlordSave={onLandlordSave} />
      </Modal>
      <Modal
        title="业主详情"
        width={570}
        open={detailModalOpen}
        onCancel={() => {
          setDetailModalOpen(false)
        }}
        footer={null}
      >
        <LandlordDetailUI landlordDetail={detailLandlordInfo} />
      </Modal>
    </div>
  )
}

export default LandlordList
