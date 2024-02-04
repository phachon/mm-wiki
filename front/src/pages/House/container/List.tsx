import React, { useEffect, useState } from 'react'
import HouseListUI from '../component/ListUI'
import { HouseEditResp, HouseInfoType, HouseListResp, HouseTrustInfoType } from '@/types/houseType'
import { HouseService } from '@/services/House'
import { message, Modal, TablePaginationConfig } from 'antd'
import HouseSearchUI from '../component/SearchUI'
import HouseFormUI from '../component/FormUI'
import HouseDetailUI from '../component/DetailUI'
import { initPagination } from '@/types/adminType'
import { LandlordDetailResp, LandlordInfoType } from '@/types/landlordType'
import LandlordDetailUI from '@/pages/Landlord/component/DetailUI'
import HouseTrustUI from '../component/TrustUI'
import { LandlordService } from '@/services/Landlord'

let searchValues = {}

const HouseList: React.FC = () => {
  const [houseList, setHouseList] = useState<HouseInfoType[]>([])
  const [pagination, setPagination] = useState(initPagination)

  const [editModalOpen, setEditModalOpen] = useState<boolean>(false)
  const [editHouseInfo, setEditHouseInfo] = useState<HouseInfoType>()

  const [detailHouseInfo, setDetailHouseInfo] = useState<HouseInfoType>()
  const [detailModalOpen, setDetailModalOpen] = useState<boolean>(false)

  const [landlords, setLandlords] = useState<LandlordInfoType[]>([])
  const [defaultSelectLandlord, setDefaultSelectLandlord] = useState<string>('')

  const [landlordModalOpen, setLandlordModalOpen] = useState<boolean>(false)
  const [landlordInfo, setLandlordInfo] = useState<LandlordInfoType>()

  const [trustModalOpen, setTrustModalOpen] = useState<boolean>(false)
  const [houseTrustList, setHouseTrustList] = useState<HouseTrustInfoType[]>()

  useEffect(() => {
    getHouseList(initPagination, {})
  }, [])

  /**
   * 请求房产列表
   * @param pageConfig 翻页信息
   * @param searchKeywords 搜索信息
   */
  const getHouseList = (pageConfig: TablePaginationConfig, searchKeywords: {}) => {
    searchValues = searchKeywords
    HouseService.houseList(pageConfig.pageSize, pageConfig.current, searchKeywords)
      .then((res: HouseListResp) => {
        setHouseList(res.list)
        setPagination({
          ...initPagination,
          current: res.page_info?.page_num,
          pageSize: res.page_info?.page_size,
          total: res.page_info?.total_num
        })
      })
      .catch((e) => {
        console.log('获取房产列表失败：', e)
      })
  }

  /**
   * 搜索查询操作
   * @param values 搜索表单数据
   */
  const onHouseSearchChange = (values: {}) => {
    getHouseList(initPagination, values)
  }

  /**
   * 搜索重置操作
   */
  const onHouseSearchReset = () => {
    getHouseList(initPagination, {})
  }

  /**
   * 列表分页请求
   * @param pageConfig
   * @param filters
   * @param sorter
   */
  const onHouseListChange = (pageConfig: TablePaginationConfig, filters: any, sorter: any) => {
    getHouseList(pageConfig, searchValues)
  }

  /**
   * 修改点击操作
   * @param houseInfo 房产信息
   */
  const onHouseEditClick = (houseInfo: HouseInfoType) => {
    // 获取修改房产需要的信息
    HouseService.getEditHouseInfo(houseInfo.house_id)
      .then((editHouseInfo: HouseEditResp) => {
        setEditHouseInfo(editHouseInfo.house_info)
        setLandlords(editHouseInfo.landlords)
        setDefaultSelectLandlord(editHouseInfo.house_info.landlord_id.toString())
        setEditModalOpen(true)
      })
      .catch((e) => {
        console.log('获取修改信息失败err:', e)
      })
  }

  /**
   * 房产详情点击操作
   * @param houseInfo 房产信息
   */
  const onHouseDetailClick = (houseInfo: HouseInfoType) => {
    setDetailModalOpen(true)
    setDetailHouseInfo(houseInfo)
  }

  /**
   * 房产业主详情点击操作
   * @param houseInfo 房产信息
   */
  const onHouseLandlordDetailClick = (houseInfo: HouseInfoType) => {
    // 获取业主详细信息
    LandlordService.getLandlordDetail(houseInfo.landlord_id)
      .then((resp: LandlordDetailResp) => {
        setLandlordInfo(resp.landlord_info)
        setLandlordModalOpen(true)
      })
      .catch((e) => {
        console.log('获取房产业主详情失败, err:', e)
      })
  }

  /**
   * 删除房产操作
   * @param houseInfo
   * @param status
   */
  const onHouseDeleteConfim = (houseInfo: HouseInfoType) => {
    HouseService.deleteHouse(houseInfo.house_id)
      .then(() => {
        message.success('删除成功', 2, () => {
          getHouseList(pagination, searchValues)
        })
      })
      .catch((e) => {
        console.log('房产删除失败err:', e)
      })
  }

  /**
   * 修改弹框取消操作
   */
  const onEditModalCancel = () => {
    setEditModalOpen(false)
  }

  /**
   * 修改完成操作
   * @param houseInfo HouseInfoType
   */
  const onHouseInfoSave = (houseInfo: HouseInfoType) => {
    HouseService.modifyHouse(houseInfo)
      .then(() => {
        message.success('修改成功', 2, () => {
          setEditModalOpen(false)
          getHouseList(pagination, searchValues)
        })
      })
      .catch((e) => {
        console.log('修改保存失败:', e)
      })
  }

  /**
   * 房产托管点击操作
   * @param houseInfo 房产信息
   */
  const onHouseTrustClick = (houseInfo: HouseInfoType) => {
    setTrustModalOpen(true)
  }

  /**
   * 房产托管删除方法
   * @param houseTrustInfo 房产托管信息
   */
  const onHouseTrustDeleteConfim = (houseTrustInfo: HouseTrustInfoType) => {}

  /**
   * 房产托管列表翻页方法
   * @param pageConfig 分页信息
   * @param filters
   * @param sorter
   */
  const onHouseTrustListChange = (
    pageConfig: TablePaginationConfig,
    filters: any,
    sorter: any
  ) => {}

  /**
   * 房产托管保存方法
   * @param houseTrustInfo 房产托管信息
   */
  const onHouseTrustSave = (houseTrustInfo: any) => {
    console.log('onHouseTrustSave:', houseTrustInfo)
  }

  return (
    <div className="panel">
      <HouseSearchUI
        onHouseSearchChange={onHouseSearchChange}
        onHouseSearchReset={onHouseSearchReset}
      />
      <HouseListUI
        listLoading={false}
        houseList={houseList}
        pagination={pagination}
        onHouseListChange={onHouseListChange}
        onHouseEditClick={onHouseEditClick}
        onHouseDeleteConfim={onHouseDeleteConfim}
        onHouseDetailClick={onHouseDetailClick}
        onHouseLandlordDetailClick={onHouseLandlordDetailClick}
        onHouseTrustClick={onHouseTrustClick}
      />
      <Modal
        title="房产修改"
        width={1100}
        open={editModalOpen}
        onCancel={onEditModalCancel}
        footer={null}
      >
        <HouseFormUI
          houseInfo={editHouseInfo}
          selectLandlords={landlords}
          onHouseInfoSave={onHouseInfoSave}
          defaultSelectLandlordValue={defaultSelectLandlord}
        />
      </Modal>
      <Modal
        title="房产详情"
        width={570}
        open={detailModalOpen}
        onCancel={() => {
          setDetailModalOpen(false)
        }}
        footer={null}
      >
        <HouseDetailUI detailHouseInfo={detailHouseInfo} />
      </Modal>
      <Modal
        title="业主详情"
        width={570}
        open={landlordModalOpen}
        onCancel={() => {
          setLandlordModalOpen(false)
        }}
        footer={null}
      >
        <LandlordDetailUI landlordDetail={landlordInfo} />
      </Modal>
      <Modal
        title="托管详情"
        width={1070}
        open={trustModalOpen}
        onCancel={() => {
          setTrustModalOpen(false)
        }}
        footer={null}
      >
        <HouseTrustUI
          houseTrustList={houseTrustList}
          pagination={pagination}
          onHouseTrustListChange={onHouseListChange}
          onHouseTrustDeleteConfim={onHouseTrustDeleteConfim}
          onHouseTrustSave={onHouseTrustSave}
        />
      </Modal>
    </div>
  )
}

export default HouseList
