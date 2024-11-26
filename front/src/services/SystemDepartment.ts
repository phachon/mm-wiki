import httpRequest from './http'
import Base from './Base'
import {
  DepartmentAddResp,
  DepartmentEditResp,
  DepartmentInfoType,
  DepartmentListResp
} from '../types/departmentType'

const systemDepartmentUrl = {
  departmentAdd: '/system/department/add',
  departmentSave: '/system/department/save',
  departmentList: '/system/department/list',
  departmentEdit: '/system/department/edit',
  departmentModify: '/system/department/modify',
  departmentDelete: '/system/department/delete'
}

/**
 * SystemDepartment 系统 - 部门服务
 */
class SystemDepartment extends Base {
  public constructor() {
    super()
  }

  /**
   * getAddDepartmentInfo 获取添加部门信息
   * @returns
   */
  public getAddDepartmentInfo(): Promise<DepartmentAddResp> {
    let departmentAddUrl = this.getProxyUrl(systemDepartmentUrl.departmentAdd)
    return httpRequest.get<DepartmentAddResp>(departmentAddUrl)
  }

  /**
   * saveDepartment 保存部门
   * @param departmentInfo 部门信息
   * @returns
   */
  public saveDepartment(departmentInfo: DepartmentInfoType): Promise<any> {
    let departmentSaveUrl = this.getProxyUrl(systemDepartmentUrl.departmentSave)
    return httpRequest.post<any>(departmentSaveUrl, {}, departmentInfo)
  }

  /**
   * 获取编辑部门信息
   * @returns
   */
  public getEditDepartmentInfo(departmentId: bigint): Promise<DepartmentEditResp> {
    let departmentEditUrl = this.getProxyUrl(systemDepartmentUrl.departmentEdit)
    return httpRequest.get<DepartmentEditResp>(departmentEditUrl, {
      department_id: departmentId
    })
  }

  /**
   * modifyDepartment 部门修改保存
   * @param departmentInfo 部门信息
   * @returns
   */
  public modifyDepartment(departmentInfo: DepartmentInfoType): Promise<any> {
    let departmentModifyUrl = this.getProxyUrl(systemDepartmentUrl.departmentModify)
    return httpRequest.post<any>(departmentModifyUrl, {}, departmentInfo)
  }

  /**
   * deleteDepartment 删除部门
   * @param departmentInfo 部门信息
   * @returns
   */
  public deleteDepartment(departmentId: string): Promise<any> {
    let departmentDeleteUrl = this.getProxyUrl(systemDepartmentUrl.departmentDelete)
    return httpRequest.post<any>(departmentDeleteUrl, {}, { department_id: departmentId })
  }

  /**
   * departmentList 部门列表
   */
  public departmentList(): Promise<DepartmentListResp> {
    let departmentListUrl = this.getProxyUrl(systemDepartmentUrl.departmentList)
    return httpRequest.get<DepartmentListResp>(departmentListUrl, {})
  }
}

export const SystemDepartmentService = new SystemDepartment()
