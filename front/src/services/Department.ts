import httpRequest from './http'
import Base from './Base'
import {
  DepartmentAddResp,
  DepartmentEditResp,
  DepartmentInfoType,
  DepartmentListResp
} from '../types/departmentType'

const departmentUrl = {
  departmentAdd: '/system/department/add',
  departmentSave: '/system/department/save',
  departmentList: '/system/department/list',
  departmentEdit: '/system/department/edit',
  departmentModify: '/system/department/modify',
  departmentDelete: '/system/department/delete'
}

/**
 * Department 部门服务
 */
class Department extends Base {
  public constructor() {
    super()
  }

  /**
   * getAddDepartmentInfo 获取添加部门信息
   * @returns
   */
  public getAddDepartmentInfo(): Promise<DepartmentAddResp> {
    let departmentAddUrl = this.getProxyUrl(departmentUrl.departmentAdd)
    return httpRequest.get<DepartmentAddResp>(departmentAddUrl)
  }

  /**
   * saveDepartment 保存部门
   * @param departmentInfo 部门信息
   * @returns
   */
  public saveDepartment(departmentInfo: DepartmentInfoType): Promise<any> {
    let departmentSaveUrl = this.getProxyUrl(departmentUrl.departmentSave)
    return httpRequest.post<any>(departmentSaveUrl, {}, departmentInfo)
  }

  /**
   * 获取编辑部门信息
   * @returns
   */
  public getEditDepartmentInfo(departmentId: bigint): Promise<DepartmentEditResp> {
    let departmentEditUrl = this.getProxyUrl(departmentUrl.departmentEdit)
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
    let departmentModifyUrl = this.getProxyUrl(departmentUrl.departmentModify)
    return httpRequest.post<any>(departmentModifyUrl, {}, departmentInfo)
  }

  /**
   * deleteDepartment 删除部门
   * @param departmentInfo 部门信息
   * @returns
   */
  public deleteDepartment(departmentInfo: DepartmentInfoType): Promise<any> {
    let departmentDeleteUrl = this.getProxyUrl(departmentUrl.departmentDelete)
    return httpRequest.post<any>(departmentDeleteUrl, {}, departmentInfo)
  }

  /**
   * departmentList 部门列表
   */
  public departmentList(): Promise<DepartmentListResp> {
    let departmentListUrl = this.getProxyUrl(departmentUrl.departmentList)
    return httpRequest.get<DepartmentListResp>(departmentListUrl, {})
  }
}

export const DepartmentService = new Department()
